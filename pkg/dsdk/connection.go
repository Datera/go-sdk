package dsdk

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	udc "github.com/Datera/go-udc/pkg/udc"
	uuid "github.com/google/uuid"
	greq "github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
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
)

type ApiConnection struct {
	username   string
	password   string
	apiVersion string
	tenant     string
	secure     bool
	baseUrl    *url.URL
	apikey     string
}

type ApiErrorResponse struct {
	Name         string            `json:"name"`
	Code         int               `json:"code"`
	Http         int               `json:"http"`
	Message      string            `json:"message"`
	Ts           string            `json:"ts"`
	Version      string            `json:"version"`
	Op           string            `json:"op"`
	Tenant       string            `json:"tenant"`
	Path         string            `json:"path"`
	Params       map[string]string `json:"params"`
	ConnInfo     map[string]string `json:"connInfo"`
	ClientId     string            `json:"client_id"`
	ClientType   string            `json:"client_type"`
	Id           string            `json:"api_req_id"`
	TenancyClass string            `json:"tenancy_class"`
}

type ApiLogin struct {
	Key     string `json:"key"`
	Version string `json:"version"`
	ReqTime int    `json:"request_time"`
}

type ApiVersions struct {
	ApiVersions []string `json:"api_versions"`
}

type ApiListOuter struct {
	Data     []interface{}          `json:"data"`
	Version  string                 `json:"version"`
	Metadata map[string]interface{} `json:"metadata"`
	ReqTime  int                    `json:"request_time"`
	Tenant   string                 `json:"tenant"`
	Path     string                 `json:"path"`
}

type ApiOuter struct {
	Data     map[string]interface{} `json:"data"`
	Version  string                 `json:"version"`
	Metadata map[string]interface{} `json:"metadata"`
	ReqTime  int                    `json:"request_time"`
	Tenant   string                 `json:"tenant"`
	Path     string                 `json:"path"`
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
		log.Error(err)
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
		log.Errorf("Couldn't stringify data, %s", ro.JSON)
	}
	if sensitive {
		sdata = []byte("********")
	}
	sheaders, err := json.Marshal(ro.Headers)
	if err != nil {
		log.Errorf("Couldn't stringify headers, %s", ro.Headers)
	}
	tid, ok := ctxt.Value("tid").(string)
	if !ok {
		tid = "nil"
	}
	log.Debugf(strings.Join([]string{"\nDatera Trace ID: %s",
		"Datera Request ID: %s",
		"Datera Request URL: %s",
		"Datera Request Method: %s",
		"Datera Request Payload: %s",
		"Datera Request Headers: %s\n"}, "\n"),
		tid, reqId, gurl.String(), method, string(sdata), sheaders)
	t1 := time.Now()
	resp, err := greq.DoRegularRequest(method, gurl.String(), ro)
	t2 := time.Now()
	tDelta := t2.Sub(t1)
	log.Debugf(strings.Join([]string{"\nDatera Trace ID: %s",
		"Datera Response ID: %s",
		"Datera Response TimeDelta: %fs",
		"Datera Response URL: %s",
		"Datera Response Payload: %s",
		"Datera Response Object: %s\n"}, "\n"),
		tid, reqId, tDelta.Seconds(), gurl.String(), resp.String(), "nil")
	eresp, err := checkResponse(resp, err, retry)
	if err == badStatus[RetryRequestAfterLogin] {
		if apiresp, err2 := c.Login(ctxt); err2 != nil {
			log.Errorf("%s", err)
			log.Errorf("%s", err2)
			return apiresp, err2
		}
		return c.do(ctxt, method, url, ro, rs, false, sensitive)
	}
	if err != nil {
		log.Errorf("Error during checkResponse: %s\n", err)
		return nil, err
	}
	if eresp != nil {
		log.Errorf("Recieved API Error %s\n", Pretty(eresp))
		return eresp, nil
	}
	err = resp.JSON(rs)
	if err != nil {
		log.Errorf("Could not unpack response, %s\n", err)
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
			log.Errorf("Login failure: %s, %s", Pretty(apierr), err)
			return apierr, err
		}
	}
	ro.Headers = map[string]string{"tenant": c.tenant, "Auth-Token": c.apikey}
	return c.do(ctxt, method, url, ro, rs, true, false)
}

func NewApiConnection(c *udc.UDC, secure bool) *ApiConnection {
	url, err := makeBaseUrl(c.MgmtIp, c.ApiVersion, secure)
	if err != nil {
		log.Fatalf("%s", err)
	}
	return &ApiConnection{
		username:   c.Username,
		password:   c.Password,
		apiVersion: c.ApiVersion,
		tenant:     c.Tenant,
		secure:     secure,
		baseUrl:    url,
	}
}

func (c *ApiConnection) Get(ctxt context.Context, url string, ro *greq.RequestOptions) (*ApiOuter, *ApiErrorResponse, error) {
	rs := &ApiOuter{}
	apiresp, err := c.doWithAuth(ctxt, "GET", url, ro, rs)
	return rs, apiresp, err
}

func (c *ApiConnection) GetList(ctxt context.Context, url string, ro *greq.RequestOptions) (*ApiListOuter, *ApiErrorResponse, error) {
	rs := &ApiListOuter{}
	// TODO:(_alastor_) handle pulling paged entries
	apiresp, err := c.doWithAuth(ctxt, "GET", url, ro, rs)
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
		Data: map[string]string{"name": "admin", "password": "password"},
	}
	apiresp, err := c.do(ctxt, "PUT", "login", ro, login, false, true)
	c.apikey = login.Key
	return apiresp, err
}
