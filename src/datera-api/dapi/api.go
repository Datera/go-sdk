package dapi

import (
	"fmt"
)

type DateraApi struct {
	Addr       string
	ApiVersion string
	Username   string
	Password   string
}

func NewApi(addr, apiVersion, username, password string) *DateraApi {

	d := DateraApi{addr, apiVersion, username, password}
	return &d
}

func (d *DateraApi) ConnString(endpoint string) string {

	return fmt.Sprintf(
		"https://%s:7718/v%s/%s", d.Addr, d.ApiVersion, endpoint)

}
