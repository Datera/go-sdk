package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type StorageInstance struct {
	Path                 string              `json:"path,omitempty"`
	Access               Access              `json:"access,omitempty"`
	AccessControlMode    string              `json:"access_control_mode,omitempty"`
	AclPolicy            AclPolicy           `json:"acl_policy,omitempty"`
	ActiveInitiators     []Initiator         `json:"active_initiators,omitempty"`
	ActiveStorageNodes   []StorageNode       `json:"active_storage_nodes,omitempty"`
	AdminState           string              `json:"admin_state,omitempty"`
	Auth                 Auth                `json:"auth,omitempty"`
	Causes               string              `json:"causes,omitempty"`
	DeploymentState      string              `json:"deployment_state,omitempty"`
	Health               string              `json:"health,omitempty"`
	IpPool               AccessNetworkIpPool `json:"ip_pool,omitempty"`
	Name                 string              `json:"name,omitempty"`
	OpState              string              `json:"op_state,omitempty"`
	ServiceConfiguration string              `json:"service_configuration,omitempty"`
	Uuid                 string              `json:"uuid,omitempty"`
	Volumes              []Volume            `json:"volumes,omitempty"`
	VolumesEp            *Volumes
	ctxt                 context.Context
	conn                 *ApiConnection
}

type StorageInstances struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type StorageInstancesCreateRequest struct {
	Name            string `json:"name,omitempty"`
	ReplicaCount    int    `json:"replica_count,omitempty"`
	Size            int    `json:"size,omitempty"`
	PlacementMode   string `json:"placement_mode,omitempty"`
	PlacementPolicy string `json:"placement_policy,omitempty"`
	Force           bool   `json:"force,omitempty"`
}

type StorageInstancesCreateResponse StorageInstance

func newStorageInstances(ctxt context.Context, conn *ApiConnection, path string) *StorageInstances {
	return &StorageInstances{
		Path: _path.Join(path, "storage_instances"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *StorageInstances) Create(ro *StorageInstancesCreateRequest) (*StorageInstancesCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageInstancesCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumesEp = newVolumes(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type StorageInstancesListRequest struct {
	Params map[string]string
}

type StorageInstancesListResponse []StorageInstance

func (e *StorageInstances) List(ro *StorageInstancesListRequest) (*StorageInstancesListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := StorageInstancesListResponse{}
	for _, data := range rs.Data {
		elem := &StorageInstance{}
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

type StorageInstancesGetRequest struct {
	Id string
}

type StorageInstancesGetResponse StorageInstance

func (e *StorageInstances) Get(ro *StorageInstancesGetRequest) (*StorageInstancesGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageInstancesGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumesEp = newVolumes(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type StorageInstanceSetRequest struct {
	AccessControlMode string              `json:"access_control_mode,omitempty"`
	AclPolicy         AclPolicy           `json:"acl_policy,omitempty"`
	AdminState        string              `json:"admin_state,omitempty"`
	Auth              Auth                `json:"auth,omitempty"`
	Force             bool                `json:"force,omitempty"`
	IpPool            AccessNetworkIpPool `json:"ip_pool,omitempty"`
	Volumes           []Volume            `json:"volumes,omitempty"`
}

type StorageInstanceSetResponse StorageInstance

func (e *StorageInstance) Set(ro *StorageInstanceSetRequest) (*StorageInstanceSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageInstanceSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumesEp = newVolumes(e.ctxt, e.conn, e.Path)
	return resp, nil

}

type StorageInstanceDeleteRequest struct {
	Force bool `json:"force,omitempty"`
}

type StorageInstanceDeleteResponse StorageInstance

func (e *StorageInstance) Delete(ro *StorageInstanceDeleteRequest) (*StorageInstanceDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &StorageInstanceDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumesEp = newVolumes(e.ctxt, e.conn, e.Path)
	return resp, nil
}
