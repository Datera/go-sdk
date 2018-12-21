package dsdk

import (
	"context"
	"fmt"
	_path "path"
	"strconv"

	greq "github.com/levigross/grequests"
)

type AppInstance struct {
	AccessControlMode       string                 `json:"access_control_mode,omitempty" mapstructure:"access_control_mode"`
	AdminState              string                 `json:"admin_state,omitempty" mapstructure:"admin_state"`
	AppTemplate             *AppTemplate           `json:"app_template,omitempty" mapstructure:"app_template"`
	Causes                  []string               `json:"causes,omitempty" mapstructure:"causes"`
	CloneSrc                *AppInstance           `json:"clone_src,omitempty" mapstructure:"clone_src"`
	CreateMode              string                 `json:"create_mode,omitempty" mapstructure:"create_mode"`
	DeploymentState         string                 `json:"deployment_state,omitempty" mapstructure:"deployment_state"`
	Descr                   string                 `json:"descr,omitempty" mapstructure:"descr"`
	Health                  string                 `json:"health,omitempty" mapstructure:"health"`
	Id                      string                 `json:"id,omitempty" mapstructure:"id"`
	Name                    string                 `json:"name,omitempty" mapstructure:"name"`
	OpState                 string                 `json:"op_state,omitempty" mapstructure:"op_state"`
	Path                    string                 `json:"path,omitempty" mapstructure:"path"`
	RemoteRestorePercentage int                    `json:"remote_restore_percentage,omitempty" mapstructure:"remote_restore_percentage"`
	RemoteRestoreProgress   string                 `json:"remote_restore_progress,omitempty" mapstructure:"remote_restore_progress"`
	RepairPriority          string                 `json:"repair_priority,omitempty" mapstructure:"repair_priority"`
	RestorePoint            string                 `json:"restore_point,omitempty" mapstructure:"restore_point"`
	RestoreProgress         string                 `json:"restore_progress,omitempty" mapstructure:"restore_progress"`
	SnapshotPolicies        []*SnapshotPolicy      `json:"snapshot_policies,omitempty" mapstructure:"snapshot_policies"`
	Snapshots               []*Snapshot            `json:"snapshots,omitempty" mapstructure:"snapshots"`
	StorageInstances        []*StorageInstance     `json:"storage_instances,omitempty" mapstructure:"storage_instances"`
	StoragePool             []*StoragePool         `json:"storage_pool,omitempty" mapstructure:"storage_pool"`
	TemplateOverride        map[string]interface{} `json:"template_override,omitempty" mapstructure:"template_override"`
	Uuid                    string                 `json:"uuid,omitempty" mapstructure:"uuid"`
	StorageInstancesEp      *StorageInstances      `json:"-"`
}

func RegisterAppInstanceEndpoints(a *AppInstance) {
	a.StorageInstancesEp = newStorageInstances(a.Path)
	for _, si := range a.StorageInstances {
		RegisterStorageInstanceEndpoints(si)
	}
}

type AppInstances struct {
	Path string
}

type AppInstancesCreateRequest struct {
	Ctxt             context.Context        `json:"-"`
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

func newAppInstances(path string) *AppInstances {
	return &AppInstances{
		Path: _path.Join(path, "app_instances"),
	}
}

func (e *AppInstances) Create(ro *AppInstancesCreateRequest) (*AppInstance, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AppInstance{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterAppInstanceEndpoints(resp)
	return resp, nil, nil
}

type AppInstancesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListParams      `json:"params,omitempty"`
}

func (e *AppInstances) List(ro *AppInstancesListRequest) ([]*AppInstance, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params.ToMap()}
	rs, apierr, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := []*AppInstance{}
	for _, data := range rs.Data {
		elem := &AppInstance{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		RegisterAppInstanceEndpoints(elem)
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

type AppInstancesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"-"`
}

func (e *AppInstances) Get(ro *AppInstancesGetRequest) (*AppInstance, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Id), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AppInstance{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterAppInstanceEndpoints(resp)
	return resp, nil, nil
}

type AppInstanceSetRequest struct {
	Ctxt               context.Context    `json:"-"`
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

func (e *AppInstance) Set(ro *AppInstanceSetRequest) (*AppInstance, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AppInstance{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterAppInstanceEndpoints(resp)
	return resp, nil, nil
}

type AppInstanceDeleteRequest struct {
	Ctxt  context.Context `json:"-"`
	Force bool            `json:"force,omitempty" mapstructure:"force"`
}

func (e *AppInstance) Delete(ro *AppInstanceDeleteRequest) (*AppInstance, *ApiErrorResponse, error) {
	rs, apierr, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AppInstance{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterAppInstanceEndpoints(resp)
	return resp, nil, nil
}

type AppInstanceMetadata map[string]string

type AppInstanceMetadataGetRequest struct {
	Ctxt context.Context `json:"-"`
}

func (e *AppInstance) GetMetadata(ro *AppInstanceMetadataGetRequest) (*AppInstanceMetadata, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, "metadata"), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AppInstanceMetadata{}
	for k, v := range rs.Data {
		(*resp)[k] = v.(string)
	}
	return resp, nil, nil
}

type AppInstanceMetadataSetRequest struct {
	Ctxt     context.Context `json:"-"`
	Metadata map[string]string
}

func (e *AppInstance) SetMetadata(ro *AppInstanceMetadataSetRequest) (*AppInstanceMetadata, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro.Metadata}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, _path.Join(e.Path, "metadata"), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AppInstanceMetadata{}
	for k, v := range rs.Data {
		var nv string
		switch t := v.(type) {
		case string:
			nv = v.(string)
		case bool:
			nv = strconv.FormatBool(v.(bool))
		case int:
			nv = strconv.FormatInt(int64(v.(int)), 10)
		default:
			panic(fmt.Sprintf("Don't know this, what do?: %s", t))
		}

		(*resp)[k] = nv
	}
	return resp, nil, nil
}
