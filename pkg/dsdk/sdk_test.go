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
