package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type AppInstance struct {
	Path                 string        `json:"path,omitempty"`
	Access               Access        `json:"access,omitempty"`
	AccessControlMode    string        `json:"access_control_mode,omitempty"`
	AclPolicy            AclPolicy     `json:"acl_policy,omitempty"`
	ActiveInitiators     []Initiator   `json:"active_initiators,omitempty"`
	ActiveStorageNodes   []StorageNode `json:"active_storage_nodes,omitempty"`
	AdminState           string        `json:"admin_state,omitempty"`
	Auth                 Auth          `json:"auth,omitempty"`
	Causes               string        `json:"causes,omitempty"`
	DeploymentState      string        `json:"deployment_state,omitempty"`
	Health               string        `json:"health,omitempty"`
	IpPool               IpPool        `json:"ip_pool,omitempty"`
	Name                 string        `json:"name,omitempty"`
	OpState              string        `json:"op_state,omitempty"`
	ServiceConfiguration string        `json:"service_configuration,omitempty"`
	Uuid                 string        `json:"uuid,omitempty"`
	Volumes              []Volume      `json:"volumes,omitempty"`
	VolumesEp            *Volumes
	ctxt                 context.Context
	conn                 *ApiConnection
}

type AppInstances struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type AppInstancesCreateRequest struct {
	Name            string `json:"name,omitempty"`
	ReplicaCount    int    `json:"replica_count,omitempty"`
	Size            int    `json:"size,omitempty"`
	PlacementMode   string `json:"placement_mode,omitempty"`
	PlacementPolicy string `json:"placement_policy,omitempty"`
	Force           bool   `json:"force,omitempty"`
}

type AppInstancesCreateResponse AppInstance

func newAppInstances(ctxt context.Context, conn *ApiConnection, path string) *AppInstances {
	return &AppInstances{
		Path: _path.Join(path, "app_instances"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *AppInstances) Create(ro *AppInstancesCreateRequest) (*AppInstancesCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AppInstancesCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumesEp = newVolumes(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type AppInstancesListRequest struct {
	Params map[string]string
}

type AppInstancesListResponse []AppInstance

func (e *AppInstances) List(ro *AppInstancesListRequest) (*AppInstancesListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := AppInstancesListResponse{}
	for _, data := range rs.Data {
		elem := &AppInstance{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, *elem)
	}
	for _, r := range resp {
		r.conn = e.conn
		r.ctxt = e.ctxt
		r.VolumesEp = newVolumes(e.ctxt, e.conn, e.Path)
	}
	return &resp, nil
}

type AppInstancesGetRequest struct {
	Id string
}

type AppInstancesGetResponse AppInstance

func (e *AppInstances) Get(ro *AppInstancesGetRequest) (*AppInstancesGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &AppInstancesGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumesEp = newVolumes(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type AppInstanceSetRequest struct {
	AccessControlMode string    `json:"access_control_mode,omitempty"`
	AclPolicy         AclPolicy `json:"acl_policy,omitempty"`
	AdminState        string    `json:"admin_state,omitempty"`
	Auth              Auth      `json:"auth,omitempty"`
	Force             bool      `json:"force,omitempty"`
	IpPool            IpPool    `json:"ip_pool,omitempty"`
	Volumes           []Volume  `json:"volumes,omitempty"`
}

type AppInstanceSetResponse AppInstance

func (e *AppInstance) Set(ro *AppInstanceSetRequest) (*AppInstanceSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AppInstanceSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumesEp = newVolumes(e.ctxt, e.conn, e.Path)
	return resp, nil

}

type AppInstanceDeleteRequest struct {
	Force bool `json:"force,omitempty"`
}

type AppInstanceDeleteResponse AppInstance

func (e *AppInstance) Delete(ro *AppInstanceDeleteRequest) (*AppInstanceDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &AppInstanceDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumesEp = newVolumes(e.ctxt, e.conn, e.Path)
	return resp, nil
}
