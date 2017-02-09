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
		t.Fatalf("%s", err)
	}
	conn.UpdateHeaders("Content-Type=application/json")
	err = conn.Login()
	if err != nil {
		t.Fatalf("%s", err)
	}
	_, err = conn.Get("users")
	if err != nil {
		t.Fatalf("%s", err)
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
		t.Fatalf("%s", err)
	}
	_, err = client.AppInstances.List()
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestSubendpoint(t *testing.T) {
	headers := make(map[string]string)
	client, err := dsdk.NewRootEp("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		t.Fatalf("%s", err)
	}
	ais, err := client.AppInstances.List()
	if err != nil {
		t.Fatalf("%s", err)
	}
	ai := ais[0]
	si := ai.StorageInstances[0]
	si, err = si.Reload()
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCreate(t *testing.T) {
	headers := make(map[string]string)
	client, err := dsdk.NewRootEp("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		t.Fatalf("%s", err)
	}
	name, _ := dsdk.NewUUID()
	ai, err := client.AppInstances.Create(
		fmt.Sprintf("name=%s", name))
	if err != nil {
		t.Fatalf("%s", err)
	}
	ai, err = ai.Reload()
	if err != nil {
		t.Fatalf("%s", err)
	}
	ai, err = ai.Set("admin_state=offline")
	if err != nil {
		t.Fatalf("%s", err)
	}
	err = ai.Delete("force=true")
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestClean(t *testing.T) {
	headers := make(map[string]string)
	client, err := dsdk.NewRootEp("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		t.Fatalf("%s", err)
	}

	ais, err := client.AppInstances.List()
	for _, ai := range ais {
		_, err = ai.Set("admin_state=offline", "force=true")
		if err != nil {
			t.Fatal(err)
		}
		err = ai.Delete("force=true")
		if err != nil {
			t.Fatal(err)
		}
	}
}
