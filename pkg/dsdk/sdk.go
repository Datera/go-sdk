package dsdk

import (
	"context"

	udc "github.com/Datera/go-udc/pkg/udc"
	log "github.com/sirupsen/logrus"
)

const (
	VERSION         = "1.1.0"
	VERSION_HISTORY = `
		1.1.0 -- Revamped SDK to new directory structure, switched to using grequests and added UDC support
	`
)

type SDK struct {
	conf *udc.UDC
	conn *ApiConnection
	ctxt context.Context
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
	return &SDK{
		conf: c,
		ctxt: ctxt,
		conn: NewApiConnection(ctxt, c, secure),
	}, nil
}

// Cleans AppInstances, AppTemplates, StorageInstances, Initiators and InitiatorGroups under
// the currently configured tenant
func (c SDK) ForceClean() {
}
