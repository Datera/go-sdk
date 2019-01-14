Datera Golang SDK
=================

Building
--------

.. _here: http://golang.org/dl/

Requires Go 1.8+
You can download the latest version of Go here_

::

    $ make

Running Tests
-------------

::

    $ make test

Getting Started
---------------

.. code:: go

    import (
        "fmt"
        udc "github.com/Datera/go-udc/pkg/udc"
        dsdk "github.com/Datera/go-sdk/pkg/dsdk"
    )

    func main() {

        // Get Universal Datera Config (UDC).  See http://github.com/Datera/go-udc
        c, err = udc.GetConfig()
        if err != nil {
            panic(err)
        }

        // Instantiate SDK with UDC
        sdk, err := dsdk.NewSdk(udc, true)
        if err != nil {
            panic(err)
        }

        // Run HealthCheck
        if err = sdk.HealthCheck(); err != nil {
            panic(err)
        }

        // Get Context for future requests
        ctxt := sdk.NewContext()

        // You can also use your own context by providing one with
        // a "tid" key
        ctxt := context.Background()
        ctxt = context.WithValue(ctxt, "tid", "C8DF241A-FF24-4939-B8CE-987B2344FF23")
        ctxt = sdk.WithContext(ctxt)

        // NOTE: You MUST provide a valid ctxt object with each request to the
        // SDK.  Not doing so will result in a panic.  A valid ctxt object
        // contains the following keys:
        // "tid" -- A uuid or other string indicating the current transaction
        //          for tracing purposes
        // "conn" -- An ApiConnection object reference.  This is obtained via
        //           the sdk.WithContext(ctxt) function

        // List AppInstances
        params := dsdk.ListParams{
            Limit:  maxEntries,
            Offset: startToken,
        }
        ais, apierr, err := r.sdk.AppInstances.List(&dsdk.AppInstancesListRequest{
            Ctxt:   ctxt,     // This is required, see note above
            Params: params,   // These can be omitted if uneeded
        })
        if err != nil {
            panic(err)
            return nil, err
        } else if apierr != nil {
            panic(fmt.Errorf("%#v", apierr))
        }
        for _, ai := range ais {
            fmt.Println(ai.Name)
        }

        // Get System Attributes
        sys, apierr, err := sdk.System.Get(&dsdk.SystemGetRequest{Ctxt: ctxt})
        if err != nil {
            panic(err)
            return nil, err
        } else if apierr != nil {
            panic(fmt.Errorf("%#v", apierr))
        }
        fmt.Printf("System: %s\n", dsdk.Pretty(sys))
    }

All requests made by the Datera Golang SDK are within the same tenant specified
at instantiation time.  If multiple tenants are desired, multiple SDK objects
must be used, each with a different tenant.  You can accomplish this with
the following code

.. code:: go

    import (
        udc "github.com/Datera/go-udc/pkg/udc"
    )

    c1 := &udc.UDC{
        Username: "my-user"
        Password: "my-pass"
        MgmtIp: "1.1.1.1"
        ApiVersion: "2.2"
        Tenant: "tenant-A"
    }
    c2 := &udc.UDC{
        Username: "my-user"
        Password: "my-pass"
        MgmtIp: "1.1.1.1"
        ApiVersion: "2.2"
        Tenant: "tenant-B"
    }

    sdkA, err := dsdk.NewSdk(c1, true)
    if err != nil {
        panic(err)
    }

    sdkB, err := dsdk.NewSdk(c2, true)
    if err != nil {
        panic(err)
    }

Now all requests made with sdkA will go to "tenant-A", all requests with sdkB
will be routed to "tenant-B".  Changing the tenant for an existing SDK object
is currently unsupported.

Please consult the test files for more in depth API usage
