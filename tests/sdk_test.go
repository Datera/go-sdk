package dsdk_test

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/Datera/go-udc/pkg/udc"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/tjcelaya/go-datera/pkg/dsdk"
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

func TestRetryScenarios(t *testing.T) {
	originalTO := dsdk.RetryTimeout
	dsdk.RetryTimeout = int64(5) // lower the retry timeout so any test failures that result in a retry loop don't take 5 minutes
	defer func() { dsdk.RetryTimeout = originalTO }()
	testApiResponse := dsdk.ApiOuter{Data: map[string]interface{}{"name": "the system"}}
	testSystem := &dsdk.System{}
	if err := dsdk.FillStruct(testApiResponse.Data, testSystem); err != nil {
		t.Fatal(err)
	}

	apiErr401 := &dsdk.ApiErrorResponse{Name: "AuthFailedError", Http: 401}

	type expected struct {
		ApiErr *dsdk.ApiErrorResponse
		Err    error
		Data   *dsdk.System
	}
	testCases := []struct {
		desc     string
		setup    func()
		expected expected
	}{
		{
			desc: "returns success when a 503 is followed by a 200",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				// mock a 503 followed by success
				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(503).
					// On 503 errors the api may not return all fields of the ApiErrorResponse
					JSON(&dsdk.ApiErrorResponse{Message: "overloaded"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(200).
					JSON(testApiResponse)
			},
			expected: expected{
				Data: testSystem,
			},
		},
		{
			desc: "returns success when multiple 503s are followed by a 200",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				// mock 2 503s followed by success
				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(503).
					// On 503 errors the api may not return all fields of the ApiErrorResponse
					JSON(&dsdk.ApiErrorResponse{Message: "overloaded"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(503).
					// On 503 errors the api may not return all fields of the ApiErrorResponse
					JSON(&dsdk.ApiErrorResponse{Message: "overloaded"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(200).
					JSON(testApiResponse)
			},
			expected: expected{
				Data: testSystem,
			},
		},
		{
			desc: "returns success on connection error followed by 200",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				// mock 2 503s followed by success
				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					ReplyError(errors.New("connect: connection refused"))

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(200).
					JSON(testApiResponse)
			},
			expected: expected{
				Data: testSystem,
			},
		},
		{
			desc: "returns success on connection error followed by 200 during login",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					ReplyError(errors.New("connect: connection refused"))

				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				// mock 2 503s followed by success
				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					ReplyError(errors.New("connect: connection refused"))

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(200).
					JSON(testApiResponse)
			},
			expected: expected{
				Data: testSystem,
			},
		},
		{
			desc: "returns success on a 401 followed by successful re-authentication",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				// mock 401 followed by relogin and success
				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(dsdk.PermissionDenied).
					JSON(&dsdk.ApiErrorResponse{Name: "AuthFailedError", Http: 401})

				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(200).
					JSON(testApiResponse)
			},
			expected: expected{
				Data: testSystem,
			},
		},
		{
			desc: "returns success on a 503 -> 401 -> 200",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				// mock 503 followed by 401 then relogin and success
				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(503).
					// On 503 errors the api may not return all fields of the ApiErrorResponse
					JSON(&dsdk.ApiErrorResponse{Message: "overloaded"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(dsdk.PermissionDenied).
					JSON(apiErr401)

				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(200).
					JSON(testApiResponse)
			},
			expected: expected{
				Data: testSystem,
			},
		},
		{
			desc: "returns success on a 503 -> 401 -> 503 -> 200",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				// mock 503 followed by 401 then relogin and success
				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(503).
					// On 503 errors the api may not return all fields of the ApiErrorResponse
					JSON(&dsdk.ApiErrorResponse{Message: "overloaded"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(dsdk.PermissionDenied).
					JSON(apiErr401)

				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(503).
					JSON(&dsdk.ApiErrorResponse{Message: "overloaded"})

				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(200).
					JSON(testApiResponse)
			},
			expected: expected{
				Data: testSystem,
			},
		},
		{
			desc: "retries stop after a time limit",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				// mock multiple 503s followed by 200
				for i := 0; i < 3; i++ {
					gock.New("http://127.0.0.1:7717").
						Get("/v1/system").
						Reply(503).
						// On 503 errors the api may not return all fields of the ApiErrorResponse
						JSON(&dsdk.ApiErrorResponse{Message: "overloaded"})
				}

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(200).
					JSON(testApiResponse)
			},
			expected: expected{
				Err: dsdk.ErrRetryTimeout,
			},
		},
		{
			desc: "does not retry on an initial login auth error",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(dsdk.PermissionDenied).
					JSON(apiErr401)
			},
			expected: expected{
				ApiErr: apiErr401,
			},
		},
		{
			desc: "does not retry after a 401 during a re-login",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(dsdk.PermissionDenied).
					JSON(apiErr401)

				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Persist().
					Reply(dsdk.PermissionDenied).
					JSON(apiErr401)
			},
			expected: expected{
				ApiErr: apiErr401,
			},
		},
		{
			desc: "does not retry on a 400",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(400).
					JSON(&dsdk.ApiErrorResponse{Message: "invalid", Http: 400})
			},
			expected: expected{
				ApiErr: &dsdk.ApiErrorResponse{Message: "invalid", Http: 400},
			},
		},
		{
			desc: "does not retry on a 400 encountered during a retry",
			setup: func() {
				gock.New("http://127.0.0.1:7717").
					Put("/v1/login").
					Reply(200).
					JSON(&dsdk.ApiLogin{Key: "thekey"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(503).
					// On 503 errors the api may not return all fields of the ApiErrorResponse
					JSON(&dsdk.ApiErrorResponse{Message: "overloaded"})

				gock.New("http://127.0.0.1:7717").
					Get("/v1/system").
					Reply(400).
					JSON(&dsdk.ApiErrorResponse{Message: "invalid", Http: 400})
			},
			expected: expected{
				ApiErr: &dsdk.ApiErrorResponse{Message: "invalid", Http: 400},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			defer gock.OffAll()
			tC.setup()

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

			actual := expected{
				ApiErr: aer,
				Err:    err,
				Data:   s,
			}

			if diff := cmp.Diff(tC.expected, actual, cmpopts.EquateErrors()); diff != "" {
				t.Fatalf("did not get expected result: %s", diff)
			}
		})
	}
}

func TestConcurrentUsage(t *testing.T) {
	originalTO := dsdk.RetryTimeout
	dsdk.RetryTimeout = int64(5) // lower the retry timeout so any test failures that result in a retry loop don't take 5 minutes
	defer func() { dsdk.RetryTimeout = originalTO }()
	gock.New("http://127.0.0.1:7717").
		Put("/v1/login").
		Persist().
		Reply(200).
		JSON(&dsdk.ApiLogin{Key: "thekey"})
	gock.New("http://127.0.0.1:7717").
		Get("/v1/system").
		Persist().
		Reply(200).
		JSON(dsdk.ApiOuter{Data: map[string]interface{}{"name": "the system"}})

	// work around https://github.com/levigross/grequests/issues/78
	http.DefaultClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return nil
	}

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
	doneCh := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			select {
			case <-doneCh:
				wg.Done()
				return
			default:
			}
			_, _, _ = sdk.System.Get(&dsdk.SystemGetRequest{
				Ctxt: ctxt,
			})
			dsdk.GetConn(ctxt).Logout()
			time.Sleep(20 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case <-doneCh:
				wg.Done()
				return
			default:
			}
			_, _, _ = sdk.System.Get(&dsdk.SystemGetRequest{
				Ctxt: ctxt,
			})
			dsdk.GetConn(ctxt).Logout()
			time.Sleep(19 * time.Millisecond)
		}
	}()

	time.Sleep(3 * time.Second)
	close(doneCh)
	wg.Wait()

	gock.OffAll()
}
