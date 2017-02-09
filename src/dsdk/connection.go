package dsdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
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
	ConnTemplate    = "http://{{.hostname}}:{{.port}}/v{{.version}}/{{.endpoint}}"
	SecConnTemplate = "https://{{.hostname}}:{{.port}}/v{{.version}}/{{.endpoint}}"
)

var Errors = map[int]bool{
	400: true,
	401: true,
	422: true,
	500: true}

type ApiConnection struct {
	Mutex      *sync.Mutex
	Method     string
	Endpoint   string
	Headers    map[string]string
	QParams    []string
	Hostname   string
	ApiVersion string
	Port       string
	Username   string
	Password   string
	Secure     bool
	Client     *http.Client
	ApiToken   string
	Tenant     string
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
	Http                int      `json:"http`
	Message             string   `json:"message"`
	Debug               string   `json:"debug"`
	Ts                  string   `json:"ts"`
	ApiReqId            int      `json:"api_req_id"`
	StorageNodeUuid     string   `json:"storage_node_uuid"`
	StorageNodeHostname string   `json:"storage_node_hostname"`
	Schema              string   `json:"schema,omitempty"`
	Errors              []string `json:"errors,omitempty"`
}

type Ep interface {
	SetEp(string, *ApiConnection)
}

// Changing tenant should require changing the API connection object maybe?
func NewApiConnection(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*ApiConnection, error) {
	InitLog(true, "")
	t, err := time.ParseDuration(timeout)
	if err != nil {
		return nil, err
	}
	h := map[string]string{"Content-Type": "application/json"}
	for p, v := range headers {
		h[p] = v
	}
	c := &ApiConnection{
		Mutex:      &sync.Mutex{},
		Hostname:   hostname,
		Port:       port,
		Username:   username,
		Password:   password,
		Tenant:     tenant,
		Headers:    h,
		ApiVersion: apiVersion,
		Secure:     secure,
		Client:     &http.Client{Timeout: t},
	}
	c.UpdateHeaders(fmt.Sprintf("tenant=%s", tenant))
	log.Debugf("New API connection: %#v", c)
	return c, nil
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
	for k, _ := range argm {
		if !strings.Contains(fstring, "{{."+k+"}}") {
			err := fmt.Errorf("Could not find arg: '%s' in template: '%s'", fstring)
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
func (r *ApiConnection) UpdateHeaders(headers ...string) error {
	for _, h := range headers {
		h := strings.Split(h, "=")
		r.Headers[h[0]] = h[1]
	}
	return nil
}

func (r *ApiConnection) prepConn() (string, error) {
	var fstring string
	if r.Secure {
		fstring = SecConnTemplate
	} else {
		fstring = ConnTemplate
	}
	m := map[string]string{
		"hostname": r.Hostname,
		"port":     r.Port,
		"endpoint": r.Endpoint,
		"version":  r.ApiVersion,
	}
	conn, err := parseTemplate(fstring, m)
	if err != nil {
		return "", err
	}
	if r.ApiToken != "" {
		r.UpdateHeaders(fmt.Sprintf("Auth-Token=%s", r.ApiToken))
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

func (r *ApiConnection) doRequest(method, endpoint string, body []byte, qparams []string, sensitive bool) ([]byte, error) {
	r.Mutex.Lock()
	// Handle method
	var m string
	switch strings.ToLower(method) {
	default:
		panic(fmt.Sprintf("Did not understand method request %s"))
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
	reqUuid, err := NewUUID()
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
		"Datera Request URL: /v%s/%s",
		"Datera Request Method: %s",
		"Datera Request Payload: %s",
		"Datera Request Headers: %s"}, "\n"),
		nil,
		reqUuid,
		r.ApiVersion,
		conn,
		r.Method,
		logb,
		r.Headers)
	resp, err := r.Client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}
	log.Debugf(strings.Join([]string{
		"\nDatera Trace ID: %s",
		"Datera Response ID: %s",
		"Datera Response Status: %s",
		"Datera Response Payload: %s",
		"Datera Response Headers: %s"}, "\n"),
		nil,
		reqUuid,
		resp.Status,
		rbody,
		resp.Header)
	err = handleBadResponse(resp)
	r.Mutex.Unlock()
	return rbody, err
}

// qparams have form "param=value"
func (r *ApiConnection) Get(endpoint string, qparams ...string) ([]byte, error) {
	return r.doRequest("get", endpoint, nil, qparams, false)
}

func (r *ApiConnection) Put(endpoint string, sensitive bool, bodyp ...string) ([]byte, error) {
	params := make(map[string]interface{})
	for _, b := range bodyp {
		p := strings.Split(b, "=")
		var v interface{}
		v = p[1]
		if p[1] == "true" || p[1] == "false" {
			v, _ = strconv.ParseBool(p[1])
		}
		params[p[0]] = v
	}
	body, err := json.Marshal(params)
	if err != nil {
		return []byte(""), err
	}
	return r.doRequest("put", endpoint, body, nil, sensitive)
}

func (r *ApiConnection) Post(endpoint string, bodyp ...string) ([]byte, error) {
	params := make(map[string]interface{})
	for _, b := range bodyp {
		p := strings.Split(b, "=")
		var v interface{}
		v = p[1]
		if p[1] == "true" || p[1] == "false" {
			v, _ = strconv.ParseBool(p[1])
		}
		params[p[0]] = v
	}
	body, err := json.Marshal(params)
	if err != nil {
		return []byte(""), err
	}
	return r.doRequest("post", endpoint, body, nil, false)
}

// qparams have form "param=value"
func (r *ApiConnection) Delete(endpoint string, bodyp ...string) ([]byte, error) {
	params := make(map[string]interface{})
	for _, b := range bodyp {
		p := strings.Split(b, "=")
		var v interface{}
		v = p[1]
		if p[1] == "true" || p[1] == "false" {
			v, _ = strconv.ParseBool(p[1])
		}
		params[p[0]] = v
	}
	body, err := json.Marshal(params)
	if err != nil {
		return []byte(""), err
	}
	return r.doRequest("delete", endpoint, body, nil, false)
}

func (r *ApiConnection) Login() error {
	p1 := fmt.Sprintf("name=%s", r.Username)
	p2 := fmt.Sprintf("password=%s", r.Password)
	var l ReturnLogin
	var e ErrResponse21
	resp, err := r.Put("login", true, p1, p2)
	if err != nil {
		serr := json.Unmarshal(resp, &e)
		if serr != nil {
			return err
		}
		return errors.New(e.Message)
	}
	err = json.Unmarshal(resp, &l)
	if err != nil {
		return err
	}
	if l.Key == "" {
		return errors.New(
			fmt.Sprintf("No Api Token In Response: %s", resp))
	}
	r.ApiToken = l.Key
	return nil
}

func getData(resp []byte) (json.RawMessage, *ErrResponse21, error) {
	var r Response21
	var e ErrResponse21
	err := json.Unmarshal(resp, &r)
	if err != nil {
		return []byte{}, nil, err
	}
	err = json.Unmarshal(resp, &e)
	return r.DataRaw, &e, nil
}

func handleBadResponse(resp *http.Response) error {
	_, ok := Errors[resp.StatusCode]
	if ok {
		return errors.New(fmt.Sprintf("%s", resp.Status))
	}
	return nil
}
