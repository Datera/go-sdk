package dapi

import (
	"bytes"
	"fmt"
	// "encoding/json"
	// "errors"
	"io/ioutil"
	"net/http"
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
	Hostname   string
	ApiVersion string
	Port       string
	Username   string
	Password   string
	Secure     bool
	Client     *http.Client
	LastResult string
	ApiToken   string
	Tenant     string
}

// Changing tenant should require changing the API connection object maybe?
func NewApiConnection(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*ApiConnection, error) {
	t, err := time.ParseDuration(timeout)
	if err != nil {
		return nil, err
	}
	return &ApiConnection{
		Hostname:   hostname,
		Port:       port,
		Username:   username,
		Password:   password,
		Tenant:     tenant,
		Headers:    headers,
		ApiVersion: apiVersion,
		Secure:     secure,
		Client:     &http.Client{Timeout: t},
	}, nil
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
	fmt.Println(buf.String())
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
	return conn, err
}

func (r *ApiConnection) Get(endpoint string) (string, error) {
	r.Method = http.MethodGet
	r.Endpoint = endpoint
	conn, err := r.prepConn()
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(r.Method, conn, nil)
	for h, v := range r.Headers {
		req.Header.Set(h, v)
	}
	if err != nil {
		return "", err
	}
	resp, err := r.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(rbody), err
}

func (r *ApiConnection) Put(endpoint string, body []byte) (string, error) {
	r.Method = http.MethodPut
	r.Endpoint = endpoint
	conn, err := r.prepConn()
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(r.Method, conn, bytes.NewReader(body))
	for h, v := range r.Headers {
		req.Header.Set(h, v)
	}
	if err != nil {
		return "", err
	}
	resp, err := r.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(rbody), err
}
