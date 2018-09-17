package dsdk_test

import (
	"context"
	"fmt"
	"testing"

	dsdk "github.com/Datera/go-sdk/pkg/dsdk"
	udc "github.com/Datera/go-udc/pkg/udc"
	greq "github.com/levigross/grequests"
)

func TestApiVersions(t *testing.T) {
	c, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	conn := dsdk.NewApiConnection(c, false)
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
	conn := dsdk.NewApiConnection(c, false)
	_, err = conn.Login(context.Background())
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestConnReAuth(t *testing.T) {
	c, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	conn := dsdk.NewApiConnection(c, false)
	_, _, err = conn.GetList(context.Background(), "app_instances", &greq.RequestOptions{})
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestCreateInitiator(t *testing.T) {
	c, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	conn := dsdk.NewApiConnection(c, false)
	id := fmt.Sprintf("iqn.1993-08.org.debian:01:%s", dsdk.RandString(12))
	ro := &greq.RequestOptions{
		Data: map[string]string{
			"id":    id,
			"name":  "my-go-test",
			"force": "true",
		},
	}
	_, _, err = conn.Post(context.Background(), "initiators", ro)
	if err != nil {
		t.Errorf("%s", err)
	}
	ro = &greq.RequestOptions{}
	_, _, err = conn.Delete(context.Background(), fmt.Sprintf("initiators/%s", id), ro)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestCreateAi(t *testing.T) {
	c, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	conn := dsdk.NewApiConnection(c, false)
	ro := &greq.RequestOptions{
		Data: map[string]string{
			"id":    fmt.Sprintf("iqn.1993-08.org.debian:01:%s", dsdk.RandString(12)),
			"name":  "my-go-test",
			"force": "true",
		},
	}
	_, _, err = conn.Post(context.Background(), "initiators", ro)
	if err != nil {
		t.Errorf("%s", err)
	}
}
