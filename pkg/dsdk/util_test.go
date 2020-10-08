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

func Test_canonicalizeRoute(t *testing.T) {
	type args struct {
		route string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{route: "/v2.2/app_instances/123/storage_instances/345"},
			want: "/v2.2/app_instances/:id/storage_instances/:id",
		},
		{
			name: "",
			args: args{route: "/v2.2/metrics/hw/cpu"},
			want: "/v2.2/metrics/hw/:id",
		},
		{
			name: "",
			args: args{route: "/v2.2/system"},
			want: "/v2.2/system",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canonicalizeRoute(tt.args.route, "2.2"); got != tt.want {
				t.Errorf("canonicalizeRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
