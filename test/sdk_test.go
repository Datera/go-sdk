package dsdk_test

import (
	"fmt"
	"testing"

	dsdk "github.com/Datera/go-sdk/pkg/dsdk"
)

func TestSDKInsecure(t *testing.T) {
	sdk, err := dsdk.NewSDK(nil, false)
	if err != nil {
		t.Error(err)
	}
	sdk.HealthCheck()
}

func TestSDKSecure(t *testing.T) {
	sdk, err := dsdk.NewSDK(nil, true)
	if err != nil {
		t.Error(err)
	}
	sdk.HealthCheck()
}

func TestSDKInitiatorCreate(t *testing.T) {
	sdk, err := dsdk.NewSDK(nil, true)
	if err != nil {
		t.Error(err)
	}
	ro := &dsdk.InitiatorsCreateRequest{
		Id:   fmt.Sprintf("iqn.1993-08.org.debian:01:%s", dsdk.RandString(12)),
		Name: dsdk.RandString(12),
	}
	var init *dsdk.Initiator
	if init, err = sdk.Initiators.Create(ro); err != nil {
		t.Errorf("%s", err)
	}
	if _, err = init.Delete(nil); err != nil {
		t.Errorf("%s", err)
	}
}
