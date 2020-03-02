package dsdk_test

import (
	"fmt"
	"testing"

	"github.com/Datera/go-udc/pkg/udc"
	dsdk "github.com/tjcelaya/go-datera/pkg/dsdk"
	"gopkg.in/h2non/gock.v1"
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
	ctxt := sdk.NewContext()
	ro := &dsdk.InitiatorsCreateRequest{
		Ctxt: ctxt,
		Id:   fmt.Sprintf("iqn.1993-08.org.debian:01:%s", dsdk.RandString(12)),
		Name: dsdk.RandString(12),
	}
	var init *dsdk.Initiator
	if init, _, err = sdk.Initiators.Create(ro); err != nil {
		t.Errorf("%s", err)
	}
	if _, _, err = init.Delete(&dsdk.InitiatorDeleteRequest{
		Ctxt: ctxt,
	}); err != nil {
		t.Errorf("%s", err)
	}
}

// TestRetry ensures that a request that gets a 503 will retry and if the
// next response is a 200 will return the result
func TestRetry(t *testing.T) {
	defer gock.Off()

	gock.New("http://127.0.0.1:7717").
		Put("/v1/login").
		Reply(200).
		JSON(&dsdk.ApiLogin{Key: "thekey"})

	// mock a 503 followed by success
	gock.New("http://127.0.0.1:7717").
		Get("/v1/system").
		Reply(503)

	gock.New("http://127.0.0.1:7717").
		Get("/v1/system").
		Reply(200).
		JSON(dsdk.ApiOuter{Data: map[string]interface{}{"name": "the system"}})

	sdk, err := dsdk.NewSDK(&udc.UDC{
		MgmtIp:     "127.0.0.1",
		Username:   "foo",
		Password:   "bar",
		ApiVersion: "1",
	}, false)
	if err != nil {
		t.Error(err)
	}
	ctxt := sdk.NewContext()
	s, aer, err := sdk.System.Get(&dsdk.SystemGetRequest{
		Ctxt: ctxt,
	})

	if gock.HasUnmatchedRequest() {
		for _, un := range gock.GetUnmatchedRequests() {
			t.Errorf("unmatched request: %+v", un)
		}
		t.Fatal("received unexpected requests")
	}

	if aer != nil || err != nil {
		t.Fatalf("errors should have been nil: %v %v", aer, err)
	}

	if s.Name != "the system" {
		t.Fatalf("did not get expected result: %+v", s)
	}
}
