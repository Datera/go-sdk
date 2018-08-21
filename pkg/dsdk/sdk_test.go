package dsdk

import (
	"fmt"
	"testing"
)

func TestSDKInsecure(t *testing.T) {
	sdk, err := NewSDK(nil, false)
	if err != nil {
		t.Error(err)
	}
	sdk.conn.Login()
	fmt.Printf("%s\n", sdk.conn.apikey)
}

func TestSDKSecure(t *testing.T) {
	sdk, err := NewSDK(nil, true)
	if err != nil {
		t.Error(err)
	}
	sdk.conn.Login()
	fmt.Printf("%s\n", sdk.conn.apikey)
}

func TestSDKInitiatorCreate(t *testing.T) {
	sdk, err := NewSDK(nil, true)
	if err != nil {
		t.Error(err)
	}
	ro := &InitiatorsCreateRequest{
		Id:   fmt.Sprintf("iqn.1993-08.org.debian:01:%s", RandString(12)),
		Name: RandString(12),
	}
	var resp *InitiatorsCreateResponse
	if resp, err = sdk.Initiators.Create(ro); err != nil {
		t.Errorf("%s", err)
	}
	init := Initiator(*resp)
	if _, err = init.Delete(nil); err != nil {
		t.Errorf("%s", err)
	}
}
