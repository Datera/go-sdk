package dsdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	structs "github.com/fatih/structs"
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

type apiError string

const (
	connTemplate = "{{.schema}}://{{.hostname}}:{{.port}}/v{{.version}}/{{.endpoint}}"
	USetToken    = ""
	MaxPoolConn  = 5
)

const (
	permDeniedError apiError = "PermissionDeniedError"
	authFailedError          = "AuthFailedError"
)

var (
	httpErrors = map[int]bool{
		400: true,
		401: true,
		422: true,
		500: true}

	retryError = fmt.Errorf("Retry")
)

type IAPIConnection interface {
	post(string, ...interface{}) ([]byte, error)
	get(string, ...string) ([]byte, error)
	put(string, bool, ...interface{}) ([]byte, error)
	delete(string, ...interface{}) ([]byte, error)
	login() error
	updateHeaders(...string) error
}

type IHTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Connection Pool to allow for concurrent connections to the backend
type connectionPool struct {
	Conns chan IAPIConnection
}

func newConnPool(hostname, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*connectionPool, error) {
	c := &connectionPool{}
	c.Conns = make(chan IAPIConnection, MaxPoolConn)
	auth := newLogAuth(username, password)
	for i := 0; i < MaxPoolConn; i++ {
		api, err := newAPIConnection(hostname, apiVersion, tenant, timeout, headers, secure, auth)
		if err != nil {
			return nil, err
		}
		c.Conns <- api
	}
	return c, nil
}

func (c *connectionPool) getConn() IAPIConnection {
	return <-c.Conns
}

func (c *connectionPool) releaseConn(api IAPIConnection) {
	c.Conns <- api
}

type apiConnection struct {
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
	Auth       *logAuth
	Schema     string
	id         string
}

// Unrelated to Auth object in entity.go
type logAuth struct {
	APIToken string
	Username string
	Password string
	Mutex    *sync.Mutex
}

func newLogAuth(username, password string) *logAuth {
	return &logAuth{
		Username: username,
		Password: password,
		APIToken: "",
		Mutex:    &sync.Mutex{},
	}
}

func (a *logAuth) setToken(t string) {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	a.APIToken = t
}

func (a *logAuth) getToken() string {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	return a.APIToken
}

type returnLogin struct {
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
func newAPIConnection(hostname, apiVersion, tenant, timeout string, headers map[string]string, secure bool, auth *logAuth) (IAPIConnection, error) {
	InitLog(true, "")
	t, err := time.ParseDuration(timeout)
	if err != nil {
		return nil, err
	}
	h := map[string]string{"Content-Type": "application/json"}
	for p, v := range headers {
		h[p] = v
	}
	port := "7717"
	schema := "http"
	client := &http.Client{Timeout: t}
	if secure {
		port = "7718"
		schema = "https"
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}

	}
	apiUUID, err := NewUUID()
	c := apiConnection{
		Hostname:   hostname,
		Port:       port,
		Tenant:     tenant,
		Headers:    h,
		APIVersion: apiVersion,
		Secure:     secure,
		Client:     client,
		Auth:       auth,
		Schema:     schema,
		id:         apiUUID,
	}
	c.updateHeaders(fmt.Sprintf("tenant=%s", tenant))
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
func (r *apiConnection) updateHeaders(headers ...string) error {
	for _, h := range headers {
		h := strings.Split(h, "=")
		r.Headers[h[0]] = h[1]
	}
	return nil
}

func (r *apiConnection) prepConn() (string, error) {
	m := map[string]string{
		"hostname": r.Hostname,
		"port":     r.Port,
		"endpoint": r.Endpoint,
		"version":  r.APIVersion,
		"schema":   r.Schema,
	}
	conn, err := parseTemplate(connTemplate, m)
	if err != nil {
		return "", err
	}
	if r.Auth.getToken() != USetToken {
		r.updateHeaders(fmt.Sprintf("Auth-Token=%s", r.Auth.getToken()))
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

func (r *apiConnection) doRequest(method, endpoint string, body []byte, qparams []string, sensitive bool, retry bool) ([]byte, error) {
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
		"\nDatera Connector ID: %s",
		"Datera Request ID: %s",
		"Datera Request URL: %s",
		"Datera Request Method: %s",
		"Datera Request Payload: %s",
		"Datera Request Headers: %s"}, "\n"),
		r.id,
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
		"\nDatera Connector ID: %s",
		"Datera Response ID: %s",
		"Datera Response Status: %s",
		"Datera Response Payload: %s",
		"Datera Response Headers: %s"}, "\n"),
		r.id,
		reqUUID,
		resp.Status,
		rbody,
		resp.Header)
	log.Debugf("\nRequest %s Duration Response: %.2fs", reqUUID, dur)
	log.Debugf("\nRequest %s Duration Read: %.2fs", reqUUID, dur2)
	err = handleBadResponse(resp, rbody)
	// Retry if we need to login, but only once
	if err == retryError && !retry {
		r.Auth.setToken(USetToken)
		r.Headers["Auth-Token"] = USetToken
		r.login()
		return r.doRequest(method, endpoint, body, qparams, sensitive, true)
	}
	return rbody, err
}

func (r *apiConnection) get(endpoint string, qparams ...string) ([]byte, error) {
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
func (r *apiConnection) put(endpoint string, sensitive bool, bodyp ...interface{}) ([]byte, error) {
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
func (r *apiConnection) post(endpoint string, bodyp ...interface{}) ([]byte, error) {
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
func (r *apiConnection) delete(endpoint string, bodyp ...interface{}) ([]byte, error) {
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

// After successful login the API token is saved in the apiConnection object
func (r *apiConnection) login() error {
	p1 := fmt.Sprintf("name=%s", r.Auth.Username)
	p2 := fmt.Sprintf("password=%s", r.Auth.Password)
	var l returnLogin
	var e ErrResponse21
	// Only login if we need to
	if r.Auth.getToken() == USetToken {
		resp, err := r.put("login", true, p1, p2)
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
		r.Auth.setToken(l.Key)
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
		apierr := apiError(e.Name)
		if apierr == permDeniedError || apierr == authFailedError {
			return retryError
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
