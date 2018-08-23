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
	conn                 *ApiConnection
	ctxt                 context.Context
	AccessNetworkIpPools *AccessNetworkIpPools
	AppInstances         *AppInstances
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
		ctxt:                 ctxt,
		conn:                 conn,
		AccessNetworkIpPools: newAccessNetworkIpPools(ctxt, conn, "/"),
		AppInstances:         newAppInstances(ctxt, conn, "/"),
		Initiators:           newInitiators(ctxt, conn, "/"),
		InitiatorGroups:      newInitiatorGroups(ctxt, conn, "/"),
		StorageNodes:         newStorageNodes(ctxt, conn, "/"),
		StoragePools:         newStoragePools(ctxt, conn, "/"),
		Tenants:              newTenants(ctxt, conn, "/"),
	}, nil
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
