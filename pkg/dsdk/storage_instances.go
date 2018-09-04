package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type StorageInstance struct {
	Path                 string               `json:"path,omitempty" mapstructure:"path"`
	Access               *Access              `json:"access,omitempty" mapstructure:"access"`
	AccessControlMode    string               `json:"access_control_mode,omitempty" mapstructure:"access_control_mode"`
	AclPolicy            *AclPolicy           `json:"acl_policy,omitempty" mapstructure:"acl_policy"`
	ActiveInitiators     []*Initiator         `json:"active_initiators,omitempty" mapstructure:"active_initiators"`
	ActiveStorageNodes   []*StorageNode       `json:"active_storage_nodes,omitempty" mapstructure:"active_storage_nodes"`
	AdminState           string               `json:"admin_state,omitempty" mapstructure:"admin_state"`
	Auth                 *Auth                `json:"auth,omitempty" mapstructure:"auth"`
	Causes               []string             `json:"causes,omitempty" mapstructure:"causes"`
	DeploymentState      string               `json:"deployment_state,omitempty" mapstructure:"deployment_state"`
	Health               string               `json:"health,omitempty" mapstructure:"health"`
	IpPool               *AccessNetworkIpPool `json:"ip_pool,omitempty" mapstructure:"ip_pool"`
	Name                 string               `json:"name,omitempty" mapstructure:"name"`
	OpState              string               `json:"op_state,omitempty" mapstructure:"op_state"`
	ServiceConfiguration string               `json:"service_configuration,omitempty" mapstructure:"service_configuration"`
	Uuid                 string               `json:"uuid,omitempty" mapstructure:"uuid"`
	Volumes              []*Volume            `json:"volumes,omitempty" mapstructure:"volumes"`
	VolumesEp            *Volumes             `json:"-"`
}

func RegisterStorageInstanceEndpoints(a *StorageInstance) {
	a.VolumesEp = newVolumes(a.Path)
	for _, si := range a.Volumes {
		RegisterVolumeEndpoints(si)
	}
}

type StorageInstances struct {
	Path string
}

type StorageInstancesCreateRequest struct {
	Ctxt                 context.Context      `json:"-"`
	AccessControlMode    string               `json:"access_control_mode,omitempty" mapstructure:"access_control_mode"`
	AclPolicy            *AclPolicy           `json:"acl_policy,omitempty" mapstructure:"acl_policy"`
	AdminState           string               `json:"admin_state,omitempty" mapstructure:"admin_state"`
	Auth                 *Auth                `json:"auth,omitempty" mapstructure:"auth"`
	IpPool               *AccessNetworkIpPool `json:"ip_pool,omitempty" mapstructure:"ip_pool"`
	Name                 string               `json:"name,omitempty" mapstructure:"name"`
	ServiceConfiguration string               `json:"service_configuration,omitempty" mapstructure:"service_configuration"`
	Volumes              []*Volume            `json:"volumes,omitempty" mapstructure:"volumes"`
}

func newStorageInstances(path string) *StorageInstances {
	return &StorageInstances{
		Path: _path.Join(path, "storage_instances"),
	}
}

func (e *StorageInstances) Create(ro *StorageInstancesCreateRequest) (*StorageInstance, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &StorageInstance{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterStorageInstanceEndpoints(resp)
	return resp, nil, nil
}

type StorageInstancesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

func (e *StorageInstances) List(ro *StorageInstancesListRequest) ([]*StorageInstance, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, apierr, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := []*StorageInstance{}
	for _, data := range rs.Data {
		elem := &StorageInstance{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		RegisterStorageInstanceEndpoints(elem)
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

type StorageInstancesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string
}

func (e *StorageInstances) Get(ro *StorageInstancesGetRequest) (*StorageInstance, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Name), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &StorageInstance{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterStorageInstanceEndpoints(resp)
	return resp, nil, nil
}

type StorageInstanceSetRequest struct {
	Ctxt              context.Context      `json:"-"`
	AccessControlMode string               `json:"access_control_mode,omitempty" mapstructure:"access_control_mode"`
	AclPolicy         *AclPolicy           `json:"acl_policy,omitempty" mapstructure:"acl_policy"`
	AdminState        string               `json:"admin_state,omitempty" mapstructure:"admin_state"`
	Auth              *Auth                `json:"auth,omitempty" mapstructure:"auth"`
	Force             bool                 `json:"force,omitempty" mapstructure:"force"`
	IpPool            *AccessNetworkIpPool `json:"ip_pool,omitempty" mapstructure:"ip_pool"`
	Volumes           []*Volume            `json:"volumes,omitempty" mapstructure:"volumes"`
}

func (e *StorageInstance) Set(ro *StorageInstanceSetRequest) (*StorageInstance, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &StorageInstance{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterStorageInstanceEndpoints(resp)
	return resp, nil, nil

}

type StorageInstanceDeleteRequest struct {
	Ctxt  context.Context `json:"-"`
	Force bool            `json:"force,omitempty" mapstructure:"force"`
}

func (e *StorageInstance) Delete(ro *StorageInstanceDeleteRequest) (*StorageInstance, *ApiErrorResponse, error) {
	rs, apierr, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &StorageInstance{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterStorageInstanceEndpoints(resp)
	return resp, nil, nil
}
