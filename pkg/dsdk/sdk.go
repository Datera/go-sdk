package dsdk

import (
	"context"

	udc "github.com/Datera/go-udc/pkg/udc"
	log "github.com/sirupsen/logrus"
)

const (
	VERSION         = "2.0.0"
	VERSION_HISTORY = `
		2.0.0 -- Revamped SDK to new directory structure, switched to using grequests and added UDC support
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
	StorageNodes         *StorageNodes
	StoragePools         *StoragePools
	Tenants              *Tenants
}

func NewSDK(c *udc.UDC, secure bool) (*SDK, error) {
	var err error
	if c == nil {
		c, err = udc.GetConfig()
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}
	ctxt := context.Background()
	conn := NewApiConnection(ctxt, c, secure)
	return &SDK{
		conf:                 c,
		Ctxt:                 ctxt,
		Conn:                 conn,
		AccessNetworkIpPools: newAccessNetworkIpPools("/"),
		AppInstances:         newAppInstances("/"),
		AppTemplates:         newAppTemplates("/"),
		Initiators:           newInitiators("/"),
		InitiatorGroups:      newInitiatorGroups("/"),
		StorageNodes:         newStorageNodes("/"),
		StoragePools:         newStoragePools("/"),
		Tenants:              newTenants("/"),
	}, nil
}

func (c SDK) Context() context.Context {
	ctxt := context.WithValue(c.Ctxt, "conn", c.Conn)
	ctxt = context.WithValue(ctxt, "tid", RandString(34))
	return ctxt
}

// Cleans AppInstances, AppTemplates, StorageInstances, Initiators and InitiatorGroups under
// the currently configured tenant
func (c SDK) HealthCheck() error {
	resp, err := c.StorageNodes.List(&StorageNodesListRequest{})
	if err != nil {
		return err
	}
	log.Debugf("Connected to cluster: %s with tenant %s.\n", c.conf.MgmtIp, c.conf.Tenant)
	for _, r := range *resp {
		sn := StorageNode(r)
		log.Debugf("Found Storage Node: %s\n", sn.Uuid)
	}
	return nil
}
