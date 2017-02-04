package dsdk

import (
	"encoding/json"
	"fmt"
)

type (
	RootEp struct {
		Path         string
		conn         *ApiConnection
		AppInstances AppInstancesEp
	}

	AppInstancesEp struct {
		Path string
		conn *ApiConnection
	}

	ApiEp struct {
		Path string
		conn *ApiConnection
	}

	AppInstanceEntity struct {
		Path string
		conn *ApiConnection
		Name string `json:"name"`
	}

	ApiEntity struct {
		Path string
		conn *ApiConnection
	}

	MetricsEntityEp struct {
		Path string
		conn *ApiConnection
	}
)

//Endpoint Functions:
// - Create
// - Get
// - List

func NewRootEp(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*RootEp, error) {
	conn, err := NewApiConnection(hostname, port, username, password, apiVersion, tenant, timeout, headers, secure)
	if err != nil {
		return nil, err
	}
	err = conn.Login()
	if err != nil {
		return nil, err
	}
	return &RootEp{
		Path:         "",
		conn:         conn,
		AppInstances: NewAppInstanceEp(conn),
	}, nil

}

func NewAppInstanceEp(conn *ApiConnection) AppInstancesEp {
	return AppInstancesEp{
		Path: "app_instances",
		conn: conn,
	}
}

func (ep AppInstancesEp) List(queryp ...string) ([]AppInstanceEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ais []AppInstanceEntity
	err = json.Unmarshal(d, &ais)
	if err != nil {
		panic(err)
	}
	for _, ai := range ais {
		ai.conn = ep.conn
	}
	fmt.Printf("Ais: %#v", ais)
	return ais, nil
}

// func (ep AppInstancesEp) Get(queryp ...string) (AppInstanceEntity, error) {

// }

// // Body Args have this form: "param=value", must have at least one
// func (ep AppInstancesEp) Create(bodyp ...string) (AppInstanceEntity, error) {

// }
