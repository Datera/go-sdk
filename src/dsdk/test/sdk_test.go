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

func getClient(t *testing.T) *dsdk.RootEndpoint {
	headers := make(map[string]string)
	client, err := dsdk.NewRootEndpoint("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		t.Fatalf("%s", err)
	}
	return client
}

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

}

func TestEndpoint(t *testing.T) {
	headers := make(map[string]string)
	client, err := dsdk.NewRootEndpoint("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		t.Fatalf("%s", err)
	}
	_, err = client.GetEndpoint("app_instances").List()
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestSubendpoint(t *testing.T) {
	headers := make(map[string]string)
	client, err := dsdk.NewRootEndpoint("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		t.Fatalf("%s", err)
	}
	name, _ := dsdk.NewUUID()
	ai, err := client.GetEndpoint("app_instances").Create(
		fmt.Sprintf("name=%s", name))
	ai.GetEndpoint("storage_instances").Create()
	ais, err := client.GetEndpoint("app_instances").List()
	if err != nil {
		t.Fatalf("%s", err)
	}
	ai = ais[0]
	ai.GetEndpoint("storage_instances").Create("name=storage-1")
	ai, _ = ai.Reload()
	si := ai.GetEntities("storage_instances")[0]
	si, err = si.Reload()
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCreate(t *testing.T) {
	headers := make(map[string]string)
	client, err := dsdk.NewRootEndpoint("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		t.Fatalf("%s", err)
	}
	name, _ := dsdk.NewUUID()
	ai, err := client.GetEndpoint("app_instances").Create(
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

// func TestACL(t *testing.T) {
// client := getClient(t)
// fmt.Println(client)
// }

func TestClean(t *testing.T) {
	headers := make(map[string]string)
	client, err := dsdk.NewRootEndpoint("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	if err != nil {
		t.Fatalf("%s", err)
	}

	ais, err := client.GetEndpoint("app_instances").List()
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
