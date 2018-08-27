package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type AppInstance struct {
	AccessControlMode       string             `json:"access_control_mode,omitempty" mapstructure:"access_control_mode"`
	AdminState              string             `json:"admin_state,omitempty" mapstructure:"admin_state"`
	AppTemplate             *AppTemplate       `json:"app_template,omitempty" mapstructure:"app_template"`
	Causes                  []string           `json:"causes,omitempty" mapstructure:"causes"`
	CloneSrc                *AppInstance       `json:"clone_src,omitempty" mapstructure:"clone_src"`
	CreateMode              string             `json:"create_mode,omitempty" mapstructure:"create_mode"`
	DeploymentState         string             `json:"deployment_state,omitempty" mapstructure:"deployment_state"`
	Descr                   string             `json:"descr,omitempty" mapstructure:"descr"`
	Health                  string             `json:"health,omitempty" mapstructure:"health"`
	Id                      string             `json:"id,omitempty" mapstructure:"id"`
	Name                    string             `json:"name,omitempty" mapstructure:"name"`
	OpState                 string             `json:"op_state,omitempty" mapstructure:"op_state"`
	Path                    string             `json:"path,omitempty" mapstructure:"path"`
	RemoteRestorePercentage int                `json:"remote_restore_percentage,omitempty" mapstructure:"remote_restore_percentage"`
	RemoteRestoreProgress   string             `json:"remote_restore_progress,omitempty" mapstructure:"remote_restore_progress"`
	RepairPriority          string             `json:"repair_priority,omitempty" mapstructure:"repair_priority"`
	RestorePoint            string             `json:"restore_point,omitempty" mapstructure:"restore_point"`
	RestoreProgress         string             `json:"restore_progress,omitempty" mapstructure:"restore_progress"`
	SnapshotPolicies        []*SnapshotPolicy  `json:"snapshot_policies,omitempty" mapstructure:"snapshot_policies"`
	Snapshots               []*Snapshot        `json:"snapshots,omitempty" mapstructure:"snapshots"`
	StorageInstances        []*StorageInstance `json:"storage_instances,omitempty" mapstructure:"storage_instances"`
	StoragePool             []*StoragePool     `json:"storage_pool,omitempty" mapstructure:"storage_pool"`
	Uuid                    string             `json:"uuid,omitempty" mapstructure:"uuid"`
	StorageInstancesEp      *StorageInstances  `json:"-"`
	ctxt                    context.Context
	conn                    *ApiConnection
}

type AppInstances struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type AppInstancesCreateRequest struct {
	AppTemplate      *AppTemplate           `json:"app_template,omitempty" mapstructure:"app_template"`
	CloneSnapshotSrc *Snapshot              `json:"clone_snapshot_src,omitempty" mapstructure:"clone_snapshot_src"`
	CloneVolumeSrc   *Volume                `json:"clone_volume_src,omitempty" mapstructure:"clone_volume_src"`
	CloneSrc         *AppInstance           `json:"clone_src,omitempty" mapstructure:"clone_src"`
	CreateMode       string                 `json:"create_mode,omitempty" mapstructure:"create_mode"`
	Descr            string                 `json:"descr,omitempty" mapstructure:"descr"`
	Name             string                 `json:"name,omitempty" mapstructure:"name"`
	RepairPriority   string                 `json:"repair_priority,omitempty" mapstructure:"repair_priority"`
	SnapshotPolicies []*SnapshotPolicy      `json:"snapshot_policies,omitempty" mapstructure:"snapshot_policies"`
	StorageInstances []*StorageInstance     `json:"storage_instances,omitempty" mapstructure:"storage_instances"`
	StoragePool      []*StoragePool         `json:"storage_pool,omitempty" mapstructure:"storage_pool"`
	TemplateOverride map[string]interface{} `json:"template_override,omitempty" mapstructure:"template_override"`
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
	AdminState         string             `json:"admin_state,omitempty" mapstructure:"admin_state"`
	Descr              string             `json:"descr,omitempty" mapstructure:"descr"`
	Force              bool               `json:"force,omitempty" mapstructure:"force"`
	Name               string             `json:"name,omitempty" mapstructure:"name"`
	Provisioned        string             `json:"provisioned,omitempty" mapstructure:"provisioned"`
	RemoteProvider     string             `json:"remote_provider,omitempty" mapstructure:"remote_provider"`
	RemoteRestorePoint string             `json:"remote_restore_point,omitempty" mapstructure:"remote_restore_point"`
	RepairPriority     string             `json:"repair_priority,omitempty" mapstructure:"repair_priority"`
	RestorePoint       string             `json:"restore_point,omitempty" mapstructure:"restore_point"`
	SnapshotPolicies   []*SnapshotPolicy  `json:"snapshot_policies,omitempty" mapstructure:"snapshot_policies"`
	StorageInstances   []*StorageInstance `json:"storage_instances,omitempty" mapstructure:"storage_instances"`
	StoragePool        []*StoragePool     `json:"storage_pool,omitempty" mapstructure:"storage_pool"`
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
	Force bool `json:"force,omitempty" mapstructure:"force"`
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
