package dsdk

import (
	"strings"
	"time"
)

const (
	VERSION = "1.0.2"
)

var (
	Cpool *ConnectionPool
)

type Client struct {
}

func NewClient(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*Client, error) {
	var err error
	//Initialize global connection object
	Cpool, err = NewConnPool(hostname, port, username, password, apiVersion, tenant, timeout, headers, secure)
	if err != nil {
		return nil, err
	}
	conn := Cpool.GetConn()
	defer Cpool.ReleaseConn(conn)
	return &Client{}, nil
}

func (c Client) GetEp(path string) IEndpoint {
	return NewEp("", path)
}

// Cleans AppInstances, AppTemplates, StorageInstances, Initiators and InitiatorGroups under
// the currently configured tenant
func (c Client) ForceClean() {
	f := func(lc chan int, en IEntity) {
		if strings.Contains(en.GetPath(), "app_instances") {
			en.Set("admin_state=offline", "force=true")
		}
		if strings.Contains(en.GetPath(), "app_templates") {
			for {
				err := en.Delete("force=true")
				if err != nil {
					if strings.Contains(err.Error(), "read-only") {
						break
					} else {
						time.Sleep(2 * time.Second)
					}
				} else {
					break
				}
			}
		}
		en.Delete("force=true")
		lc <- 1
	}

	var dones []chan int
	chi := 0
	for _, epStr := range []string{"app_instances", "app_templates", "initiators", "initiator_groups"} {
		items, _ := c.GetEp(epStr).List()
		numItems := len(items)
		for i := 0; i < numItems; i++ {
			dones = append(dones, make(chan int))
		}
		for _, item := range items {
			go f(dones[chi], item)
			chi += 1
		}
	}
	// Check channels for completion
	for _, ch := range dones {
		<-ch
	}
}
