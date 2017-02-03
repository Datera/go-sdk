package dapi_test

import (
	// "encoding/json"
	"fmt"
	// "strings"
	"datera-api/dapi"
	"testing"
	// "net/http"
	// "net/http/httptest"
	// "github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

const (
	ADDR     = "192.168.1.1"
	APIVER   = "2.1"
	USERNAME = "testuser"
	PASSWORD = "testpass"
)

func TestApiBasic(t *testing.T) {

}

func TestConnection(t *testing.T) {
	headers := make(map[string]string)
	// test := map[string]string{"name": "admin", "password": "password"}
	// j, _ := json.Marshal(test)
	conn, err := dapi.NewApiConnection("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		panic(err)
	}
	conn.UpdateHeaders("Content-Type=application/json")
	_, err = conn.Get("api", "myparam=test")
	if err != nil {
		panic(err)
	}
	err = conn.Login()
	if err != nil {
		panic(err)
	}
	fmt.Println(conn.ApiToken)

	// ts := httptest.NewServer(http.HandlerFunc(
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		fmt.Fprintln(w, "Hello mudda, Hello fadda")
	// 	}))
	// defer ts.Close()
	// s := strings.Split(ts.URL, ":")
	// port := s[2]
	// hostname = strings.Strip(s[1], "/")
	// conn, _ := dapi.NewApiConnection(hostname, port, "admin", "password", "2.1", "/root", "30s", headers, false)

}
