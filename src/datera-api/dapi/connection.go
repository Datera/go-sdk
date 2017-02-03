package dapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	// "errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"
)

const (
	ConnTemplate    = "http://{{.hostname}}:{{.port}}/v{{.version}}/{{.endpoint}}/"
	SecConnTemplate = "https://{{.hostname}}:{{.port}}/v{{.version}}/{{.endpoint}}/"
)

type ApiConnection struct {
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

// Changing tenant should require changing the API connection object maybe?
func NewApiConnection(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*ApiConnection, error) {
	InitLog(true, "")
	t, err := time.ParseDuration(timeout)
	if err != nil {
		return nil, err
	}
	c := &ApiConnection{
		Hostname:   hostname,
		Port:       port,
		Username:   username,
		Password:   password,
		Tenant:     tenant,
		Headers:    headers,
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
	r.Endpoint = endpoint
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
	reqUuid, err := newUUID()
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
	return rbody, err
}

// qparams have form "param=value"
func (r *ApiConnection) Get(endpoint string, qparams ...string) ([]byte, error) {
	return r.doRequest("get", endpoint, nil, qparams, false)
}

func (r *ApiConnection) Put(endpoint string, body []byte, sensitive bool) ([]byte, error) {
	return r.doRequest("put", endpoint, body, nil, sensitive)
}

func (r *ApiConnection) Post(endpoint string, body []byte) ([]byte, error) {
	return r.doRequest("post", endpoint, body, nil, false)
}

// qparams have form "param=value"
func (r *ApiConnection) Delete(endpoint string, qparams ...string) ([]byte, error) {
	return r.doRequest("delete", endpoint, nil, qparams, false)
}

func (r *ApiConnection) Login() error {
	j, err := json.Marshal(map[string]string{
		"name":     r.Username,
		"password": r.Password,
	})
	if err != nil {
		return err
	}
	var l ReturnLogin
	resp, err := r.Put("login", j, true)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, &l)
	if err != nil {
		return err
	}
	r.ApiToken = l.Key
	return nil
}
