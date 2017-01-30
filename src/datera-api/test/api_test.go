package dapi_test

import (
	// "fmt"
	"testing"

	"datera-api/dapi"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

const (
	ADDR     = "192.168.1.1"
	APIVER   = "2.1"
	USERNAME = "testuser"
	PASSWORD = "testpass"
)

func TestApiBasic(t *testing.T) {
	assert := assert.New(t)

	d := dapi.NewApi(ADDR, APIVER, USERNAME, PASSWORD)

	assert.True(d.Addr == ADDR)
	assert.True(d.ApiVersion == APIVER)
	assert.True(d.Username == USERNAME)
	assert.True(d.Password == PASSWORD)
}

func TestConnString(t *testing.T) {
	assert := assert.New(t)
	d := dapi.NewApi(ADDR, APIVER, USERNAME, PASSWORD)

	endpoint := "test_endpoint/"

	assert.True("https://192.168.1.1:7718/v2.1/test_endpoint/" == d.ConnString(endpoint))
}
