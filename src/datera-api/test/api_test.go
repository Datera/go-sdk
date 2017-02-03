package dapi_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"datera-api/dapi"
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
	conn, _ := dapi.NewApiConnection("172.19.1.41", "7717", "admin", "password", "2.1", "/root", "30s", headers, false)
	test := map[string]string{"name": "admin", "password": "password"}
	j, _ := json.Marshal(test)
	conn.UpdateHeaders("Content-Type=application/json")
	fmt.Println(conn.Put("login", j))
}
