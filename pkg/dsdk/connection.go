package dsdk

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
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
	log "github.com/sirupsen/logrus"
)

var (
	RetryTimeout           = int64(300)
	ErrRetryTimeout        = errors.New("timeout reached before request completed successfully during retries")
	InvalidRequest         = 400
	PermissionDenied       = 401
	Retry503               = 503
	ConnectionError        = 9998
	RetryRequestAfterLogin = 9999
	badStatus              = map[int]error{
		InvalidRequest:         fmt.Errorf("InvalidRequest"),
		PermissionDenied:       fmt.Errorf("PermissionDenied"),
		Retry503:               fmt.Errorf("Retry503"),
		ConnectionError:        fmt.Errorf("ConnectionError"),
		RetryRequestAfterLogin: fmt.Errorf("RetryRequestAfterLogin"),
	}
	DateraDriver = fmt.Sprintf("Golang-SDK-%s", VERSION)
	logTraceID   = "trace_id"
)

type ApiConnection struct {
	m          *sync.RWMutex
	username   string
	password   string
	apiVersion string
	tenant     string
	secure     bool
	ldap       string
	apikey     string
	baseUrl    *url.URL
	httpClient *http.Client
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

type ListRangeParams struct {
	Since  string `json:"since,omitempty" mapstructure:"since"`
	From   string `json:"from,omitempty" mapstructure:"from"`
	To     string `json:"to,omitempty" mapstructure:"to"`
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

func (s ListRangeParams) ToMap() map[string]string {
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
	if s.Since != "" {
		r["since"] = s.Since
	}
	if s.From != "" {
		r["from"] = s.From
	}
	if s.To != "" {
		r["to"] = s.To
	}
	return r
}

func ListRangeParamsFromMap(m map[string]string) *ListRangeParams {
	lp := &ListRangeParams{}
	lp.Filter = m["filter"]
	lp.Sort = m["sort"]
	lp.Since = m["since"]
	lp.From = m["from"]
	lp.To = m["to"]
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

func translateErrors(resp *greq.Response, err error) (*ApiErrorResponse, error) {
	if err != nil {
		Log().Error(err)
		if strings.Contains(err.Error(), "connect: connection refused") {
			return nil, badStatus[ConnectionError]
		}
		return nil, err
	}

	if !resp.Ok {
		eresp := &ApiErrorResponse{}
		err := resp.JSON(eresp)
		if err != nil {
			eresp.Message = fmt.Sprintf("failed to unmarshal ApiErrorResponse: %v", err)
		}

		// in some cases (like 503s) the response JSON doesn't contain
		// all the fields of the ApiErrorResponse and we want to always
		// be able to rely on at least having the status code
		if eresp.Http == 0 {
			eresp.Http = resp.StatusCode
		}
		return eresp, badStatus[resp.StatusCode]
	}
	return nil, nil
}

// hasLoggedIn reports whether the ApiConnection has successfully authenticated once
func (c *ApiConnection) hasLoggedIn() bool {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.apikey != ""
}

func (c *ApiConnection) retry(ctxt context.Context, method, url string, ro *greq.RequestOptions, rs interface{}, sensitive bool) (*ApiErrorResponse, error) {
	t1 := time.Now().Unix()
	backoff := 1
	var apiresp *ApiErrorResponse
	for time.Now().Unix()-t1 < RetryTimeout {
		apiresp, err := c.do(ctxt, method, url, ro, rs, false, sensitive)
		if apiresp == nil && err == nil {
			return nil, nil
		}

		// Retry on 503 and ConnectionErrors only
		if apiresp != nil && apiresp.Http != 503 {
			return apiresp, nil
		} else if err != nil && !strings.Contains(err.Error(), "connect: connection refused") {
			return nil, err
		}

		time.Sleep(time.Second * time.Duration(backoff*backoff))
		backoff += 1
	}
	return apiresp, ErrRetryTimeout
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
	if ro.HTTPClient == nil && c.httpClient != nil {
		ro.HTTPClient = c.httpClient
	}
	if ro.Context == nil {
		ro.Context = ctxt
	}
	if ro.Headers == nil {
		ro.Headers = make(map[string]string, 1)
	}
	ro.Headers["Datera-Driver"] = DateraDriver
	tid, ok := ctxt.Value("tid").(string)
	if !ok {
		tid = "nil"
	}
	if _, ok := ctxt.Value("quiet").(bool); ok {
		sdata = []byte("<muted>")
	}
	t1 := time.Now()
	// This will be run before each request.  It's needed so we can get access
	// to the headers/body passed with the request instead of just our custom ones
	ro.BeforeRequest = func(h *http.Request) error {
		sheaders, err := json.Marshal(h.Header)
		if err != nil {
			Log().Errorf("Couldn't stringify headers, %s", h.Header)
		}
		Log().WithFields(log.Fields{
			logTraceID:        tid,
			"request_id":      reqId,
			"request_method":  method,
			"request_url":     gurl.String(),
			"request_headers": sheaders,
			"request_payload": string(sdata),
		}).Debugf("Datera SDK making request")
		return nil
	}

	// The actual request happens here
	// Context is passed through ro.Context
	resp, err := greq.DoRegularRequest(method, gurl.String(), ro)

	t2 := time.Now()
	tDelta := t2.Sub(t1)
	rdata := resp.String()
	if _, ok := ctxt.Value("quiet").(bool); ok {
		rdata = "<muted>"
	}
	Log().WithFields(log.Fields{
		logTraceID:           tid,
		"request_id":         reqId,
		"response_timedelta": tDelta.Seconds(),
		"request_url":        gurl.String(),
		"response_payload":   rdata,
		"response_code":      resp.StatusCode,
	}).Debugf("Datera SDK response received")

	eresp, err := translateErrors(resp, err)

	if err == badStatus[PermissionDenied] {
		// if we have logged in successfully before we may just need to refresh the apikey
		// and retry the original request
		if c.hasLoggedIn() {
			c.Logout()
			if apiresp, err2 := c.Login(ctxt); apiresp != nil || err2 != nil {
				Log().Errorf("%s", err)
				Log().Errorf("%s", err2)
				return apiresp, err2
			}
			c.m.RLock()
			ro.Headers["Auth-Token"] = c.apikey
			c.m.RUnlock()
			return c.do(ctxt, method, url, ro, rs, false, sensitive)
		}

		// but if we get here while logged out then then credentials may no longer be valid and we shouldn't
		// retry the login again.  Just return the permission denied error
		return eresp, nil

	}
	if retry && (err == badStatus[Retry503] || err == badStatus[ConnectionError]) {
		return c.retry(ctxt, method, url, ro, rs, sensitive)
	}
	if eresp != nil {
		Log().Errorf("Recieved API Error %s", Pretty(eresp))
		return eresp, nil
	}
	if err != nil {
		Log().Errorf("Error during checkResponse: %s", err)
		return nil, err
	}
	err = resp.JSON(rs)
	if err != nil {
		Log().Errorf("Could not unpack response, %s", err)
		Log().Errorf("Response, %s", resp.String())
		return nil, err
	}
	return nil, nil
}

func (c *ApiConnection) doWithAuth(ctxt context.Context, method, url string, ro *greq.RequestOptions, rs interface{}) (*ApiErrorResponse, error) {
	if ro == nil {
		ro = &greq.RequestOptions{}
	}
	if !c.hasLoggedIn() {
		if apierr, err := c.Login(ctxt); apierr != nil || err != nil {
			Log().Errorf("Login failure: %s, %s", Pretty(apierr), err)
			return apierr, err
		}
	}
	c.m.RLock()
	ro.Headers = map[string]string{"tenant": c.tenant, "Auth-Token": c.apikey}
	c.m.RUnlock()
	return c.do(ctxt, method, url, ro, rs, true, false)
}

func NewApiConnection(c *udc.UDC, secure bool) *ApiConnection {
	return NewApiConnectionWithHTTPClient(c, secure, nil)
}

func NewApiConnectionWithHTTPClient(c *udc.UDC, secure bool, client *http.Client) *ApiConnection {
	u, err := makeBaseUrl(c.MgmtIp, c.ApiVersion, secure)
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
		baseUrl:    u,
		httpClient: client,
		m:          &sync.RWMutex{},
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
			// there are api endpoints that handle lists with more fields than
			// ListParams (but still have offset/limit in common)
			// just update offset directly here to preserve those extra fields
			ro.Params["offset"] = strconv.FormatInt(int64(offset), 10)
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
	apiresp, err := c.do(ctxt, "PUT", "login", ro, login, true, true)
	c.m.Lock()
	if (apiresp != nil && apiresp.Http == PermissionDenied) || (err != nil && err == badStatus[PermissionDenied]) {
		c.apikey = ""
	} else {
		c.apikey = login.Key
	}

	c.m.Unlock()
	return apiresp, err
}

func (c *ApiConnection) Logout() {
	c.m.Lock()
	defer c.m.Unlock()
	c.apikey = ""
}
