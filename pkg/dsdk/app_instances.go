package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type AppInstance struct {
	AccessControlMode       string            `json:"access_control_mode,omitempty"`
	AdminState              string            `json:"admin_state,omitempty"`
	AppTemplate             string            `json:"app_template,omitempty"`
	Causes                  []string          `json:"causes,omitempty"`
	CloneSrc                map[string]string `json:"clone_src,omitempty"`
	CreateMode              string            `json:"create_mode,omitempty"`
	DeploymentState         string            `json:"deployment_state,omitempty"`
	Descr                   string            `json:"descr,omitempty"`
	Health                  string            `json:"health,omitempty"`
	Id                      string            `json:"id,omitempty"`
	Name                    string            `json:"name,omitempty"`
	OpState                 string            `json:"op_state,omitempty"`
	Path                    string            `json:"path,omitempty"`
	RemoteRestorePercentage int               `json:"remote_restore_percentage,omitempty"`
	RemoteRestoreProgress   string            `json:"remote_restore_progress,omitempty"`
	RepairPriority          string            `json:"repair_priority,omitempty"`
	RestorePoint            string            `json:"restore_point,omitempty"`
	RestoreProgress         string            `json:"restore_progress,omitempty"`
	SnapshotPolicies        []SnapshotPolicy  `json:"snapshot_policies,omitempty"`
	Snapshots               []Snapshot        `json:"snapshots,omitempty"`
	StorageInstances        []StorageInstance `json:"storage_instances,omitempty"`
	StoragePool             []StoragePool     `json:"storage_pool,omitempty"`
	Uuid                    string            `json:"uuid,omitempty"`
	StorageInstancesEp      *StorageInstances
	ctxt                    context.Context
	conn                    *ApiConnection
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
	resp.StorageInstancesEp = newStorageInstances(e.ctxt, e.conn, e.Path)
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
		r.StorageInstancesEp = newStorageInstances(e.ctxt, e.conn, e.Path)
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
	resp.StorageInstancesEp = newStorageInstances(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type AppInstanceSetRequest struct {
	AdminState         string            `json:"admin_state,omitempty"`
	Descr              string            `json:"descr,omitempty"`
	Force              bool              `json:"force,omitempty"`
	Name               string            `json:"name,omitempty"`
	Provisioned        string            `json:"provisioned,omitempty"`
	RemoteProvider     string            `json:"remote_provider,omitempty"`
	RemoteRestorePoint string            `json:"remote_restore_point,omitempty"`
	RepairPriority     string            `json:"repair_priority,omitempty"`
	RestorePoint       string            `json:"restore_point,omitempty"`
	SnapshotPolicies   []SnapshotPolicy  `json:"snapshot_policies,omitempty"`
	StorageInstances   []StorageInstance `json:"storage_instances,omitempty"`
	StoragePool        []StoragePool     `json:"storage_pool,omitempty"`
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
	resp.StorageInstancesEp = newStorageInstances(e.ctxt, e.conn, e.Path)
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
	resp.StorageInstancesEp = newStorageInstances(e.ctxt, e.conn, e.Path)
	return resp, nil
}
