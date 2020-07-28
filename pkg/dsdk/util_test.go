package dsdk

import (
	"context"
	"reflect"
	"testing"

	greq "github.com/levigross/grequests"
)

func TestUtil_FormatQuery(test *testing.T) {
	ctx := context.Background()
	{
		ro := &SnapshotDeleteRequest{
			Ctxt: ctx,
		}

		v := reflect.ValueOf(*ro)
		t := reflect.TypeOf(*ro)
		gro := &greq.RequestOptions{
			JSON: ro,
		}
		formatQueryParams(gro, v, t)

		for k, _ := range gro.Params {
			if k == "remote_provider_uuid" {
				test.Error("unexpected empty param")
			}
		}
	}
	{
		ro := &SnapshotDeleteRequest{
			Ctxt:               ctx,
			RemoteProviderUuid: "test",
		}

		v := reflect.ValueOf(*ro)
		t := reflect.TypeOf(*ro)
		gro := &greq.RequestOptions{
			JSON: ro,
		}
		formatQueryParams(gro, v, t)

		for k, v := range gro.Params {
			if k == "remote_provider_uuid" && v != "test" {
				test.Error("missing param value")
			}
		}
	}
}
