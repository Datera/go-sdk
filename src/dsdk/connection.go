package dsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	structs "github.com/fatih/structs"
	// snaker "github.com/serenize/snaker"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

const (
	connTemplate    = "http://{{.hostname}}:{{.port}}/v{{.version}}/{{.endpoint}}"
	secConnTemplate = "https://{{.hostname}}:{{.port}}/v{{.version}}/{{.endpoint}}"
	permDeniedError = "PermissionDeniedError"
	USetToken       = ""
	MaxPoolConn     = 5
)

var (
	httpErrors = map[int]bool{
		400: true,
		401: true,
		422: true,
		500: true}

	Retry = fmt.Errorf("Retry")
)

type IAPIConnection interface {
	Post(string, ...interface{}) ([]byte, error)
	Get(string, ...string) ([]byte, error)
	Put(string, bool, ...interface{}) ([]byte, error)
	Delete(string, ...interface{}) ([]byte, error)
	Login() error
	UpdateHeaders(...string) error
}

type IHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Connection Pool to allow for concurrent connections to the backend
type ConnectionPool struct {
	Conns chan IAPIConnection
}

func NewConnPool(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*ConnectionPool, error) {
	c := &ConnectionPool{}
	c.Conns = make(chan IAPIConnection, MaxPoolConn)
	auth := NewLogAuth(username, password)
	for i := 0; i < MaxPoolConn; i++ {
		api, err := NewAPIConnection(hostname, port, apiVersion, tenant, timeout, headers, secure, auth)
		if err != nil {
			return nil, err
		}
		c.Conns <- api
	}
	return c, nil
}

func (c *ConnectionPool) GetConn() IAPIConnection {
	return <-c.Conns
}

func (c *ConnectionPool) ReleaseConn(api IAPIConnection) {
	c.Conns <- api
}

type APIConnection struct {
	Method     string
	Endpoint   string
	Headers    map[string]string
	QParams    []string
	Hostname   string
	APIVersion string
	Port       string
	Secure     bool
	Client     IHTTPClient
	Tenant     string
	Auth       *LogAuth
}

// Unrelated to Auth object in entity.go
type LogAuth struct {
	APIToken string
	Username string
	Password string
	Mutex    *sync.Mutex
}

func NewLogAuth(username, password string) *LogAuth {
	return &LogAuth{
		Username: username,
		Password: password,
		APIToken: "",
		Mutex:    &sync.Mutex{},
	}
}

func (a *LogAuth) SetToken(t string) {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	a.APIToken = t
}

func (a *LogAuth) GetToken() string {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	return a.APIToken
}

type ReturnLogin struct {
	Key     string `json:"key"`
	Version string `json:"version"`
}

type Response21 struct {
	Tenant  string          `json:"tenant"`
	Path    string          `json:"path"`
	Version string          `json:"version"`
	DataRaw json.RawMessage `json:"data"`
}

type ErrResponse21 struct {
	Name                string   `json:"name"`
	Code                int      `json:"code"`
	HTTP                int      `json:"http"`
	Message             string   `json:"message"`
	Debug               string   `json:"debug"`
	Ts                  string   `json:"ts"`
	APIReqId            int      `json:"api_req_id"`
	StorageNodeUUID     string   `json:"storage_node_uuid"`
	StorageNodeHostname string   `json:"storage_node_hostname"`
	Schema              string   `json:"schema,omitempty"`
	Errors              []string `json:"errors,omitempty"`
}

// Changing tenant should require changing the API connection object maybe?
func NewAPIConnection(hostname, port, apiVersion, tenant, timeout string, headers map[string]string, secure bool, auth *LogAuth) (IAPIConnection, error) {
	InitLog(true, "")
	t, err := time.ParseDuration(timeout)
	if err != nil {
		return nil, err
	}
	h := map[string]string{"Content-Type": "application/json"}
	for p, v := range headers {
		h[p] = v
	}
	c := APIConnection{
		Hostname:   hostname,
		Port:       port,
		Tenant:     tenant,
		Headers:    h,
		APIVersion: apiVersion,
		Secure:     secure,
		Client:     &http.Client{Timeout: t},
		Auth:       auth,
	}
	c.UpdateHeaders(fmt.Sprintf("tenant=%s", tenant))
	log.Debugf("New API connection: %#v", c)
	return &c, nil
}

// Args have the form "name=value"
func parseTemplate(fstring string, args ...interface{}) (string, error) {
	tpl, err := template.New("format").Parse(fstring)
	if err != nil {
		return "", err
	}
	argm := make(map[string]string)
	switch t := args[0].(type) {
	default:
		fmt.Println("Error")
	case string:
		for _, i := range args {
			arg := i.(string)
			x := strings.Split(arg, "=")
			if len(x) == 2 {
				argm[x[0]] = x[1]
			}
		}
	case map[string]string:
		argm = t
	}
	for k := range argm {
		if !strings.Contains(fstring, "{{."+k+"}}") {
			err := fmt.Errorf("Could not find arg: '%s' in template: '%s'", fstring, k)
			return "", err
		}
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, argm)
	if err != nil {
		return "", err
	}
	// fmt.Println(buf.String())
	return buf.String(), nil
}

// Headers: "header=value"
func (r *APIConnection) UpdateHeaders(headers ...string) error {
	for _, h := range headers {
		h := strings.Split(h, "=")
		r.Headers[h[0]] = h[1]
	}
	return nil
}

func (r *APIConnection) prepConn() (string, error) {
	var fstring string
	if r.Secure {
		fstring = secConnTemplate
	} else {
		fstring = connTemplate
	}
	m := map[string]string{
		"hostname": r.Hostname,
		"port":     r.Port,
		"endpoint": r.Endpoint,
		"version":  r.APIVersion,
	}
	conn, err := parseTemplate(fstring, m)
	if err != nil {
		return "", err
	}
	if r.Auth.GetToken() != USetToken {
		r.UpdateHeaders(fmt.Sprintf("Auth-Token=%s", r.Auth.GetToken()))
	}
	for i, p := range r.QParams {
		r.QParams[i] = url.QueryEscape(p)
	}
	qparams := strings.Join(r.QParams, "&")
	if len(qparams) > 0 {
		conn = strings.Join([]string{conn, qparams}, "?")
	}
	return conn, err
}

func (r *APIConnection) doRequest(method, endpoint string, body []byte, qparams []string, sensitive bool, retry bool) ([]byte, error) {
	// Handle method
	var m string
	switch strings.ToLower(method) {
	default:
		panic(fmt.Sprintf("Did not understand method request %s", method))
	case "get":
		m = http.MethodGet
	case "put":
		m = http.MethodPut
	case "post":
		m = http.MethodPost
	case "delete":
		m = http.MethodDelete
	}
	r.Method = m
	// Handle Endpoint
	r.Endpoint = strings.Trim(endpoint, "/")
	// Set query parameters
	r.QParams = qparams
	// prepConn handles header addition, url construction and query params
	conn, err := r.prepConn()
	if err != nil {
		return []byte(""), err
	}
	var b io.Reader
	if body == nil {
		b = nil
	} else {
		b = bytes.NewReader(body)
	}
	req, err := http.NewRequest(r.Method, conn, b)
	for h, v := range r.Headers {
		req.Header.Set(h, v)
	}
	if err != nil {
		return []byte(""), err
	}
	reqUUID, err := NewUUID()
	if err != nil {
		return []byte(""), err
	}
	// Obscure sensitive information
	var logb io.Reader
	if sensitive {
		logb = bytes.NewReader([]byte("************"))
	} else {
		logb = b
	}
	log.Debugf(strings.Join([]string{
		"\nDatera Trace ID: %s",
		"Datera Request ID: %s",
		"Datera Request URL: %s",
		"Datera Request Method: %s",
		"Datera Request Payload: %s",
		"Datera Request Headers: %s"}, "\n"),
		nil,
		reqUUID,
		conn,
		r.Method,
		logb,
		r.Headers)
	start1 := time.Now()
	resp, err := r.Client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()
	dur := time.Since(start1).Seconds()
	start2 := time.Now()
	rbody, err := ioutil.ReadAll(resp.Body)
	dur2 := time.Since(start2).Seconds()
	if err != nil {
		return []byte(""), err
	}
	log.Debugf(strings.Join([]string{
		// "\nDatera Trace ID: %s",
		"Datera Response ID: %s",
		"Datera Response Status: %s",
		"Datera Response Payload: %s",
		"Datera Response Headers: %s"}, "\n"),
		// nil,
		reqUUID,
		resp.Status,
		rbody,
		resp.Header)
	log.Debugf("\nRequest %s Duration Response: %.2fs", reqUUID, dur)
	log.Debugf("\nRequest %s Duration Read: %.2fs", reqUUID, dur2)
	err = handleBadResponse(resp, rbody)
	// Retry if we need to login, but only once
	if err == Retry && !retry {
		r.Auth.SetToken(USetToken)
		r.Login()
		return r.doRequest(method, endpoint, body, qparams, sensitive, true)
	}
	return rbody, err
}

func (r *APIConnection) Get(endpoint string, qparams ...string) ([]byte, error) {
	return r.doRequest("get", endpoint, nil, qparams, false, false)
}

// bodyp arguments can be in one of two forms
//
// 1. Vararg strings follwing this pattern: "key=value"
//    These strings have a limitation in that they cannot be arbitrarily nested
//    JSON values, instead they must be simple strings
//    Eg.  "key=value" is fine, but `key=["some", "list"]` will fail
//    the arbitrary JSON usecase is handled by #2
//
// 2. A single map[string]interface{} argument.  This handles the case where
//    we need to send arbitrarily nested JSON as an argument
//
// Function arguments are setup this way to provide an easy way to handle 90%
// of the use cases (where we're just passing key, value string pairs) but that
// remaining 10% we need to pass something more complex
func (r *APIConnection) Put(endpoint string, sensitive bool, bodyp ...interface{}) ([]byte, error) {
	var body []byte
	var params map[string]interface{}
	var p interface{}
	if len(bodyp) > 0 {
		p = bodyp[0]
		b, err := parseStruct(p)
		if err == nil {
			body = b
		} else {
			params, err = parseParams(bodyp...)
			body, err = json.Marshal(params)
			if err != nil {
				return []byte(""), err
			}
		}
	}
	return r.doRequest("put", endpoint, body, nil, sensitive, false)
}

// bodyp arguments can be in one of two forms
//
// 1. Vararg strings follwing this pattern: "key=value"
//    These strings have a limitation in that they cannot be arbitrarily nested
//    JSON values, instead they must be simple strings
//    Eg.  "key=value" is fine, but `key=["some", "list"]` will fail
//    the arbitrary JSON usecase is handled by #2
//
// 2. A single map[string]interface{} argument.  This handles the case where
//    we need to send arbitrarily nested JSON as an argument
//
// Function arguments are setup this way to provide an easy way to handle 90%
// of the use cases (where we're just passing key, value string pairs) but that
// remaining 10% we need to pass something more complex
func (r *APIConnection) Post(endpoint string, bodyp ...interface{}) ([]byte, error) {
	var body []byte
	var params map[string]interface{}
	var p interface{}
	if len(bodyp) > 0 {
		p = bodyp[0]
		b, err := parseStruct(p)
		if err == nil {
			body = b
		} else {
			params, err = parseParams(bodyp...)
			body, err = json.Marshal(params)
			if err != nil {
				return []byte(""), err
			}
		}
	}
	return r.doRequest("post", endpoint, body, nil, false, false)
}

// bodyp arguments can be in one of two forms
//
// 1. Vararg strings follwing this pattern: "key=value"
//    These strings have a limitation in that they cannot be arbitrarily nested
//    JSON values, instead they must be simple strings
//    Eg.  "key=value" is fine, but `key=["some", "list"]` will fail
//    the arbitrary JSON usecase is handled by #2
//
// 2. A single map[string]interface{} argument.  This handles the case where
//    we need to send arbitrarily nested JSON as an argument
//
// Function arguments are setup this way to provide an easy way to handle 90%
// of the use cases (where we're just passing key, value string pairs) but that
// remaining 10% we need to pass something more complex
func (r *APIConnection) Delete(endpoint string, bodyp ...interface{}) ([]byte, error) {
	var body []byte
	var params map[string]interface{}
	var p interface{}
	if len(bodyp) > 0 {
		p = bodyp[0]
		b, err := parseStruct(p)
		if err == nil {
			body = b
		} else {
			params, err = parseParams(bodyp...)
			body, err = json.Marshal(params)
			if err != nil {
				return []byte(""), err
			}
		}
	}
	return r.doRequest("delete", endpoint, body, nil, false, false)
}

// After successful login the API token is saved in the APIConnection object
func (r *APIConnection) Login() error {
	p1 := fmt.Sprintf("name=%s", r.Auth.Username)
	p2 := fmt.Sprintf("password=%s", r.Auth.Password)
	var l ReturnLogin
	var e ErrResponse21
	// Only login if we need to
	if r.Auth.GetToken() == USetToken {
		resp, err := r.Put("login", true, p1, p2)
		if err != nil {
			serr := json.Unmarshal(resp, &e)
			if serr != nil {
				return err
			}
			return fmt.Errorf("%s", e.Message)
		}
		err = json.Unmarshal(resp, &l)
		if err != nil {
			return err
		}
		if l.Key == "" {
			return fmt.Errorf("No Api Token In Response: %s", resp)
		}
		r.Auth.SetToken(l.Key)
	}
	return nil
}

func getData(resp []byte) (json.RawMessage, *Response21, *ErrResponse21, error) {
	var r Response21
	var e ErrResponse21
	err := json.Unmarshal(resp, &r)
	if err != nil {
		return []byte{}, nil, nil, err
	}
	err = json.Unmarshal(resp, &e)
	return r.DataRaw, &r, &e, nil
}

func handleBadResponse(resp *http.Response, rbody []byte) error {
	_, ok := httpErrors[resp.StatusCode]
	if resp.StatusCode == 401 {
		var e ErrResponse21
		err := json.Unmarshal(rbody, &e)
		if err != nil {
			log.Errorf("Bad Response: %#v", resp)
			panic(err)
		}
		if e.Name == permDeniedError {
			return Retry
		}
	}
	if ok {
		return fmt.Errorf("%s", resp.Status)
	}
	return nil
}

func parseParams(params ...interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	if len(params) == 0 {
		return result, nil
	}
	fparam := params[0]
	switch fparam.(type) {
	case map[string]interface{}:
		r := fparam.(map[string]interface{})
		return r, nil
	case interface{}:
		for _, param := range params {
			s := param.(string)
			p := strings.Split(s, "=")
			var v interface{}
			v = p[1]
			if p[1] == "true" || p[1] == "false" {
				v, _ = strconv.ParseBool(p[1])
			}
			result[p[0]] = v
		}
		return result, nil
	default:
		return result, fmt.Errorf("Couldn't Parse Params: %s", params)
	}

}

func parseStruct(s interface{}) ([]byte, error) {
	if structs.IsStruct(s) {
		b, err := json.Marshal(s)
		return b, err
	}
	return []byte(""), fmt.Errorf("Not a struct")
}
