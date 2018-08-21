package dsdk

import (
	"context"
	"fmt"
	"testing"

	udc "github.com/Datera/go-udc/pkg/udc"
	// log "github.com/sirupsen/logrus"
	greq "github.com/levigross/grequests"
)

func TestApiVersions(t *testing.T) {
	c, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	conn := NewApiConnection(context.Background(), c, false)
	apiv := conn.ApiVersions()
	if len(apiv) != 3 {
		t.Errorf("%d", len(apiv))
	}
}

func TestConnAuth(t *testing.T) {
	c, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	conn := NewApiConnection(context.Background(), c, false)
	err = conn.Login()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestConnReAuth(t *testing.T) {
	c, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	conn := NewApiConnection(context.Background(), c, false)
	_, err = conn.GetList("app_instances", nil)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestCreateInitiator(t *testing.T) {
	c, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	conn := NewApiConnection(context.Background(), c, false)
	id := fmt.Sprintf("iqn.1993-08.org.debian:01:%s", RandString(12))
	ro := &greq.RequestOptions{
		Data: map[string]string{
			"id":    id,
			"name":  "my-go-test",
			"force": "true",
		},
	}
	_, err = conn.Post("initiators", ro)
	if err != nil {
		t.Errorf("%s", err)
	}
	ro = &greq.RequestOptions{}
	_, err = conn.Delete(fmt.Sprintf("initiators/%s", id), ro)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestCreateAi(t *testing.T) {
	c, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	conn := NewApiConnection(context.Background(), c, false)
	ro := &greq.RequestOptions{
		Data: map[string]string{
			"id":    fmt.Sprintf("iqn.1993-08.org.debian:01:%s", RandString(12)),
			"name":  "my-go-test",
			"force": "true",
		},
	}
	_, err = conn.Post("initiators", ro)
	if err != nil {
		t.Errorf("%s", err)
	}
}
