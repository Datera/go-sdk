package dsdk_test

import (
	"encoding/json"
	"fmt"
	// "strings"
	"dsdk"
	"testing"
	// "net/http"
	// "net/http/httptest"
	// "github.com/stretchr/testify/assert"
	// "github.com/pkg/profile"
)

const (
	ADDR     = "172.19.1.41"
	PORT     = "7717"
	APIVER   = "2.1"
	USERNAME = "admin"
	PASSWORD = "password"
	TENANT   = "/root"
	TIMEOUT  = "30s"
)

func getClient(t *testing.T) *dsdk.RootEp {
	headers := make(map[string]string)
	client, err := dsdk.NewRootEp(
		ADDR, PORT, USERNAME, PASSWORD, APIVER, TENANT, TIMEOUT, headers, false)
	if err != nil {
		t.Fatalf("%s", err)
	}
	return client
}

func TestApiBasic(t *testing.T) {

}

func TestConnection(t *testing.T) {
	headers := make(map[string]string)
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
	client := getClient(t)
	_, err := client.GetEp("app_instances").List()
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestSubendpoint(t *testing.T) {
	client := getClient(t)
	name, _ := dsdk.NewUUID()
	ai, err := client.GetEp("app_instances").Create(
		fmt.Sprintf("name=%s", name))
	ai.GetEp("storage_instances").Create()
	ais, err := client.GetEp("app_instances").List()
	if err != nil {
		t.Fatalf("%s", err)
	}
	ai = ais[0]
	ai.GetEp("storage_instances").Create("name=storage-1")
	ai, _ = ai.Reload()
	si := ai.GetEn("storage_instances")[0]
	si, err = si.Reload()
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestCreate(t *testing.T) {
	client := getClient(t)
	name, _ := dsdk.NewUUID()
	ai, err := client.GetEp("app_instances").Create(
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

func TestACL(t *testing.T) {
	client := getClient(t)
	name, _ := dsdk.NewUUID()
	ai, err := client.GetEp("app_instances").Create(
		fmt.Sprintf("name=%s", name))
	if err != nil {
		t.Fatalf("%s", err)
	}
	si, _ := ai.GetEp("storage_instances").Create("name=storage-1")
	initep := client.GetEp("initiators")
	_, err = initep.Create(
		"name=test-initiator",
		"id=iqn.1993-08.org.debian:01:71be38c985a")
	if err != nil {
		t.Fatalf("%s", err)
	}
	aclep := si.GetEp("acl_policy")
	var args map[string]interface{}
	err = json.Unmarshal([]byte(`{"initiators":[{"path": "/initiators/iqn.1993-08.org.debian:01:71be38c985a"}]}`), &args)
	if err != nil {
		t.Fatalf("%s", err)
	}
	_, err = aclep.Set(args)
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestClean(t *testing.T) {
	client := getClient(t)
	ais, err := client.GetEp("app_instances").List()
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
	inits, _ := client.GetEp("initiators").List()
	for _, init := range inits {
		err = init.Delete()
		if err != nil {
			t.Fatal(err)
		}
	}
}
