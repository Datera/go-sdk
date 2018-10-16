package dsdk

import (
	"context"
	"fmt"

	udc "github.com/Datera/go-udc/pkg/udc"
	uuid "github.com/google/uuid"
)

const (
	VERSION         = "2.1.0"
	VERSION_HISTORY = `
		2.0.0 -- Revamped SDK to new directory structure, switched to using grequests and added UDC support
		2.1.0 -- Added LDAP server support
	`
)

type SDK struct {
	conf                 *udc.UDC
	Conn                 *ApiConnection
	Ctxt                 context.Context
	AccessNetworkIpPools *AccessNetworkIpPools
	AppInstances         *AppInstances
	AppTemplates         *AppTemplates
	Initiators           *Initiators
	InitiatorGroups      *InitiatorGroups
	LogsUpload           *LogsUpload
	StorageNodes         *StorageNodes
	StoragePools         *StoragePools
	System               *System
	Tenants              *Tenants
}

func NewSDK(c *udc.UDC, secure bool) (*SDK, error) {
	var err error
	if c == nil {
		c, err = udc.GetConfig()
		if err != nil {
			Log().Error(err)
			return nil, err
		}
	}
	conn := NewApiConnection(c, secure)
	return &SDK{
		conf:                 c,
		Conn:                 conn,
		AccessNetworkIpPools: newAccessNetworkIpPools("/"),
		AppInstances:         newAppInstances("/"),
		AppTemplates:         newAppTemplates("/"),
		Initiators:           newInitiators("/"),
		InitiatorGroups:      newInitiatorGroups("/"),
		LogsUpload:           newLogsUpload("/"),
		StorageNodes:         newStorageNodes("/"),
		StoragePools:         newStoragePools("/"),
		System:               newSystem("/"),
		Tenants:              newTenants("/"),
	}, nil
}

func (c SDK) SetDriver(d string) {
	DateraDriver = d
}

func (c SDK) WithContext(ctxt context.Context) context.Context {
	return context.WithValue(ctxt, "conn", c.Conn)
}

func (c SDK) NewContext() context.Context {
	ctxt := context.WithValue(context.Background(), "conn", c.Conn)
	ctxt = context.WithValue(ctxt, "tid", uuid.Must(uuid.NewRandom()).String())
	return ctxt
}

// Cleans AppInstances, AppTemplates, StorageInstances, Initiators and InitiatorGroups under
// the currently configured tenant
func (c SDK) HealthCheck() error {
	sns, apierr, err := c.StorageNodes.List(&StorageNodesListRequest{
		Ctxt: context.WithValue(c.NewContext(), "quiet", true),
	})
	if err != nil {
		return err
	}
	if apierr != nil {
		return fmt.Errorf("ApiError: %s", Pretty(apierr))
	}
	Log().Debugf("Connected to cluster: %s with tenant %s.\n", c.conf.MgmtIp, c.conf.Tenant)
	for _, sn := range sns {
		Log().Debugf("Found Storage Node: %s\n", sn.Uuid)
	}
	return nil
}
