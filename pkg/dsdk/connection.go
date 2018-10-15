package dsdk

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	udc "github.com/Datera/go-udc/pkg/udc"
	uuid "github.com/google/uuid"
	greq "github.com/levigross/grequests"
)

var (
	InvalidRequest         = 400
	PermissionDenied       = 401
	RetryRequestAfterLogin = 9999
	badStatus              = map[int]error{
		InvalidRequest:         fmt.Errorf("InvalidRequest"),
		PermissionDenied:       fmt.Errorf("PermissionDenied"),
		RetryRequestAfterLogin: fmt.Errorf("RetryRequestAfterLogin"),
	}
	DateraDriver = fmt.Sprintf("Golang-SDK-%s", VERSION)
)

type ApiConnection struct {
	m          *sync.Mutex
	username   string
	password   string
	apiVersion string
	tenant     string
	secure     bool
	ldap       string
	apikey     string
	baseUrl    *url.URL
}

type ApiErrorResponse struct {
	Name         string            `json:"name,omitempty"`
	Code         int               `json:"code,omitempty"`
	Http         int               `json:"http,omitempty"`
	Message      string            `json:"message,omitempty"`
	Ts           string            `json:"ts,omitempty"`
	Version      string            `json:"version,omitempty"`
	Op           string            `json:"op,omitempty"`
	Tenant       string            `json:"tenant,omitempty"`
	Path         string            `json:"path,omitempty"`
	Params       map[string]string `json:"params,omitempty"`
	ConnInfo     map[string]string `json:"connInfo,omitempty"`
	ClientId     string            `json:"client_id,omitempty"`
	ClientType   string            `json:"client_type,omitempty"`
	Id           string            `json:"api_req_id,omitempty"`
	TenancyClass string            `json:"tenancy_class,omitempty"`
	Errors       []string          `json:"errors,omitempty"`
}

type ApiLogin struct {
	Key     string `json:"key,omitempty,omitempty"`
	Version string `json:"version,omitempty,omitempty"`
	ReqTime int    `json:"request_time,omitempty,omitempty"`
}

type ApiVersions struct {
	ApiVersions []string `json:"api_versions"`
}

type ApiListOuter struct {
	Data     []interface{}          `json:"data,omitempty"`
	Version  string                 `json:"version,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	ReqTime  int                    `json:"request_time,omitempty"`
	Tenant   string                 `json:"tenant,omitempty"`
	Path     string                 `json:"path,omitempty"`
}

type ApiOuter struct {
	Data     map[string]interface{} `json:"data,omitempty"`
	Version  string                 `json:"version,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	ReqTime  int                    `json:"request_time,omitempty"`
	Tenant   string                 `json:"tenant,omitempty"`
	Path     string                 `json:"path,omitempty"`
}

type ListParams struct {
	Filter string `json:"filter,omitempty" mapstructure:"filter"`
	Limit  int    `json:"limit,omitempty" mapstructure:"limit"`
	Sort   string `json:"sort,omitempty" mapstructure:"sort"`
	Offset int    `json:"offset,omitempty" mapstructure:"offset"`
}

func (s ListParams) ToMap() map[string]string {
	r := map[string]string{}
	if s.Filter != "" {
		r["filter"] = s.Filter
	}
	if s.Limit != 0 {
		r["limit"] = strconv.FormatInt(int64(s.Limit), 10)
	}
	if s.Sort != "" {
		r["sort"] = s.Sort
	}
	if s.Offset != 0 {
		r["offset"] = strconv.FormatInt(int64(s.Offset), 10)
	}
	return r
}

func ListParamsFromMap(m map[string]string) *ListParams {
	lp := &ListParams{}
	lp.Filter = m["filter"]
	lp.Sort = m["sort"]
	if m["offset"] != "" {
		o, err := strconv.ParseInt(m["offset"], 0, 0)
		if err != nil {
			panic(err)
		}
		lp.Offset = int(o)
	} else {
		lp.Offset = 0
	}
	if m["limit"] != "" {
		o, err := strconv.ParseInt(m["limit"], 0, 0)
		if err != nil {
			panic(err)
		}
		lp.Limit = int(o)
	} else {
		lp.Limit = 0
	}
	return lp
}

func init() {
	// TODO(_alastor_): Disable this and do real certificate verification
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func makeBaseUrl(h, apiv string, secure bool) (*url.URL, error) {
	h = strings.Trim(h, "/")
	if secure {
		return url.Parse(fmt.Sprintf("https://%s:7718/v%s", h, apiv))
	}
	return url.Parse(fmt.Sprintf("http://%s:7717/v%s", h, apiv))
}

func checkResponse(resp *greq.Response, err error, retry bool) (*ApiErrorResponse, error) {
	if err != nil {
		Log().Error(err)
		return nil, err
	}
	if resp.StatusCode == PermissionDenied && retry {
		return nil, badStatus[RetryRequestAfterLogin]
	}
	if !resp.Ok {
		eresp := &ApiErrorResponse{}
		resp.JSON(eresp)
		return eresp, badStatus[resp.StatusCode]
	}
	return nil, nil
}

func (c *ApiConnection) do(ctxt context.Context, method, url string, ro *greq.RequestOptions, rs interface{}, retry, sensitive bool) (*ApiErrorResponse, error) {
	gurl := *c.baseUrl
	gurl.Path = path.Join(gurl.Path, url)
	reqId := uuid.Must(uuid.NewRandom()).String()
	sdata, err := json.Marshal(ro.JSON)
	if err != nil {
		Log().Errorf("Couldn't stringify data, %s", ro.JSON)
	}
	if sensitive {
		sdata = []byte("********")
	}
	if ro.Headers == nil {
		ro.Headers = make(map[string]string, 1)
	}
	ro.Headers["Datera-Driver"] = DateraDriver
	tid, ok := ctxt.Value("tid").(string)
	if !ok {
		tid = "nil"
	}
	t1 := time.Now()
	// This will be run before each request.  It's needed so we can get access
	// to the headers/body passed with the request instead of just our custom ones
	ro.BeforeRequest = func(h *http.Request) error {
		sheaders, err := json.Marshal(h.Header)
		if err != nil {
			Log().Errorf("Couldn't stringify headers, %s", h.Header)
		}
		Log().Debugf(strings.Join([]string{"\nDatera Trace ID: %s",
			"Datera Request ID: %s",
			"Datera Request URL: %s",
			"Datera Request Method: %s",
			"Datera Request Payload: %s",
			"Datera Request Headers: %s\n"}, "\n"),
			tid, reqId, gurl.String(), method, string(sdata), sheaders)
		return nil
	}

	// The actual request happens here
	resp, err := greq.DoRegularRequest(method, gurl.String(), ro)

	t2 := time.Now()
	tDelta := t2.Sub(t1)
	Log().Debugf(strings.Join([]string{"\nDatera Trace ID: %s",
		"Datera Response ID: %s",
		"Datera Response TimeDelta: %fs",
		"Datera Response URL: %s",
		"Datera Response Payload: %s",
		"Datera Response Object: %s\n"}, "\n"),
		tid, reqId, tDelta.Seconds(), gurl.String(), resp.String(), "nil")
	eresp, err := checkResponse(resp, err, retry)
	if err == badStatus[RetryRequestAfterLogin] {
		if apiresp, err2 := c.Login(ctxt); err2 != nil {
			Log().Errorf("%s", err)
			Log().Errorf("%s", err2)
			return apiresp, err2
		}
		return c.do(ctxt, method, url, ro, rs, false, sensitive)
	}
	if eresp != nil {
		Log().Errorf("Recieved API Error %s\n", Pretty(eresp))
		return eresp, nil
	}
	if err != nil {
		Log().Errorf("Error during checkResponse: %s\n", err)
		return nil, err
	}
	err = resp.JSON(rs)
	if err != nil {
		Log().Errorf("Could not unpack response, %s\n", err)
		Log().Errorf("Response, %s\n", resp.String())
		return nil, err
	}
	return nil, nil
}

func (c *ApiConnection) doWithAuth(ctxt context.Context, method, url string, ro *greq.RequestOptions, rs interface{}) (*ApiErrorResponse, error) {
	if ro == nil {
		ro = &greq.RequestOptions{}
	}
	if c.apikey == "" {
		if apierr, err := c.Login(ctxt); err != nil {
			Log().Errorf("Login failure: %s, %s", Pretty(apierr), err)
			return apierr, err
		}
	}
	ro.Headers = map[string]string{"tenant": c.tenant, "Auth-Token": c.apikey}
	return c.do(ctxt, method, url, ro, rs, true, false)
}

func NewApiConnection(c *udc.UDC, secure bool) *ApiConnection {
	url, err := makeBaseUrl(c.MgmtIp, c.ApiVersion, secure)
	if err != nil {
		Log().Fatalf("%s", err)
	}
	return &ApiConnection{
		username:   c.Username,
		password:   c.Password,
		apiVersion: c.ApiVersion,
		tenant:     c.Tenant,
		ldap:       c.Ldap,
		secure:     secure,
		baseUrl:    url,
		m:          &sync.Mutex{},
	}
}

func (c *ApiConnection) Get(ctxt context.Context, url string, ro *greq.RequestOptions) (*ApiOuter, *ApiErrorResponse, error) {
	rs := &ApiOuter{}
	apiresp, err := c.doWithAuth(ctxt, "GET", url, ro, rs)
	return rs, apiresp, err
}

func (c *ApiConnection) GetList(ctxt context.Context, url string, ro *greq.RequestOptions) (*ApiListOuter, *ApiErrorResponse, error) {
	rs := &ApiListOuter{}
	apiresp, err := c.doWithAuth(ctxt, "GET", url, ro, rs)
	// TODO:(_alastor_) handle pulling paged entries
	if apiresp == nil && len(rs.Metadata) > 0 {
		lp := ListParamsFromMap(ro.Params)
		if lp.Limit != 0 || lp.Offset != 0 {
			return rs, apiresp, err
		}
		data := rs.Data
		offset := 0
		tcnt := 0
		for ldata := len(data); ldata != tcnt; {
			tcnt := int(rs.Metadata["total_count"].(float64))
			offset += len(rs.Data)
			if offset >= tcnt {
				break
			}
			ro.Params = ListParams{
				Offset: offset,
			}.ToMap()
			rs.Data = []interface{}{}
			apiresp, err := c.doWithAuth(ctxt, "GET", url, ro, rs)
			if apiresp != nil || err != nil {
				rs.Data = data
				return rs, apiresp, err
			}
			data = append(data, rs.Data...)
		}
		rs.Data = data
	}
	return rs, apiresp, err
}

func (c *ApiConnection) Put(ctxt context.Context, url string, ro *greq.RequestOptions) (*ApiOuter, *ApiErrorResponse, error) {
	rs := &ApiOuter{}
	apiresp, err := c.doWithAuth(ctxt, "PUT", url, ro, rs)
	return rs, apiresp, err
}

func (c *ApiConnection) Post(ctxt context.Context, url string, ro *greq.RequestOptions) (*ApiOuter, *ApiErrorResponse, error) {
	rs := &ApiOuter{}
	apiresp, err := c.doWithAuth(ctxt, "POST", url, ro, rs)
	return rs, apiresp, err
}

func (c *ApiConnection) Delete(ctxt context.Context, url string, ro *greq.RequestOptions) (*ApiOuter, *ApiErrorResponse, error) {
	rs := &ApiOuter{}
	apiresp, err := c.doWithAuth(ctxt, "DELETE", url, ro, rs)
	return rs, apiresp, err
}

func (c *ApiConnection) ApiVersions() []string {
	gurl := *c.baseUrl
	gurl.Path = "api_versions"
	resp, err := greq.DoRegularRequest("GET", gurl.String(), nil)
	if err != nil {
		return []string{}
	}
	apiv := &ApiVersions{}
	resp.JSON(apiv)
	return apiv.ApiVersions
}

func (c *ApiConnection) Login(ctxt context.Context) (*ApiErrorResponse, error) {
	c.m.Lock()
	defer c.m.Unlock()
	login := &ApiLogin{}
	ro := &greq.RequestOptions{
		Data: map[string]string{
			"name":     c.username,
			"password": c.password,
		},
	}
	if c.ldap != "" {
		ro.Data["remote_server"] = c.ldap
	}
	apiresp, err := c.do(ctxt, "PUT", "login", ro, login, false, true)
	c.apikey = login.Key
	return apiresp, err
}
