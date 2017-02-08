package dsdk_test

import (
	// "encoding/json"
	"fmt"
	// "strings"
	"dsdk"
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
	conn, err := dsdk.NewApiConnection("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	conn.UpdateHeaders("Content-Type=application/json")
	err = conn.Login()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(conn.ApiToken)
	_, err = conn.Get("users")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// ts := httptest.NewServer(http.HandlerFunc(
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		fmt.Fprintln(w, "Hello mudda, Hello fadda")
	// 	}))
	// defer ts.Close()
	// s := strings.Split(ts.URL, ":")
	// port := s[2]
	// hostname = strings.Strip(s[1], "/")
	// conn, _ := dsdk.NewApiConnection(hostname, port, "admin", "password", "2.1", "/root", "30s", headers, false)
}

func TestEndpoint(t *testing.T) {
	headers := make(map[string]string)
	client, err := dsdk.NewRootEp("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	_, err = client.AppInstances.List()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestSubendpoint(t *testing.T) {
	headers := make(map[string]string)
	client, err := dsdk.NewRootEp("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	ais, err := client.AppInstances.List()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	ai := ais[0]
	si := ai.StorageInstances[0]
	fmt.Printf("%s", si.Path)
}
