package dsdk

import (
	"bytes"
	"encoding/json"
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
	connTemplate    = "http://{{.hostname}}:{{.port}}/v{{.version}}/{{.endpoint}}"
	secConnTemplate = "https://{{.hostname}}:{{.port}}/v{{.version}}/{{.endpoint}}"
)

var httpErrors = map[int]bool{
	400: true,
	401: true,
	422: true,
	500: true}

type IAPIConnection interface {
	Post(string, ...interface{}) ([]byte, error)
	Get(string, ...string) ([]byte, error)
	Put(string, bool, ...interface{}) ([]byte, error)
	Delete(string, ...interface{}) ([]byte, error)
	Login() error
	UpdateHeaders(...string) error
}

type APIConnection struct {
	Mutex      *sync.Mutex
	Method     string
	Endpoint   string
	Headers    map[string]string
	QParams    []string
	Hostname   string
	APIVersion string
	Port       string
	Username   string
	Password   string
	Secure     bool
	Client     *http.Client
	APIToken   string
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
func NewAPIConnection(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (IAPIConnection, error) {
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
		Mutex:      &sync.Mutex{},
		Hostname:   hostname,
		Port:       port,
		Username:   username,
		Password:   password,
		Tenant:     tenant,
		Headers:    h,
		APIVersion: apiVersion,
		Secure:     secure,
		Client:     &http.Client{Timeout: t},
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
	if r.APIToken != "" {
		r.UpdateHeaders(fmt.Sprintf("Auth-Token=%s", r.APIToken))
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

func (r *APIConnection) doRequest(method, endpoint string, body []byte, qparams []string, sensitive bool) ([]byte, error) {
	r.Mutex.Lock()
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
		reqUUID,
		resp.Status,
		rbody,
		resp.Header)
	err = handleBadResponse(resp)
	r.Mutex.Unlock()
	return rbody, err
}

// qparams have form "param=value"
func (r *APIConnection) Get(endpoint string, qparams ...string) ([]byte, error) {
	return r.doRequest("get", endpoint, nil, qparams, false)
}

func (r *APIConnection) Put(endpoint string, sensitive bool, bodyp ...interface{}) ([]byte, error) {
	params, err := parseParams(bodyp...)
	if err != nil {
		return []byte(""), err
	}
	body, err := json.Marshal(params)
	if err != nil {
		return []byte(""), err
	}
	return r.doRequest("put", endpoint, body, nil, sensitive)
}

func (r *APIConnection) Post(endpoint string, bodyp ...interface{}) ([]byte, error) {
	params, err := parseParams(bodyp...)
	if err != nil {
		return []byte(""), err
	}
	body, err := json.Marshal(params)
	if err != nil {
		return []byte(""), err
	}
	return r.doRequest("post", endpoint, body, nil, false)
}

// qparams have form "param=value"
func (r *APIConnection) Delete(endpoint string, bodyp ...interface{}) ([]byte, error) {
	params, err := parseParams(bodyp...)
	if err != nil {
		return []byte(""), err
	}
	body, err := json.Marshal(params)
	if err != nil {
		return []byte(""), err
	}
	return r.doRequest("delete", endpoint, body, nil, false)
}

func (r *APIConnection) Login() error {
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
		return fmt.Errorf("%s", e.Message)
	}
	err = json.Unmarshal(resp, &l)
	if err != nil {
		return err
	}
	if l.Key == "" {
		return fmt.Errorf("No Api Token In Response: %s", resp)
	}
	r.APIToken = l.Key
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
	_, ok := httpErrors[resp.StatusCode]
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
