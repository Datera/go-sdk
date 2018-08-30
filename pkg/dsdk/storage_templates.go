package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type StorageTemplate struct {
	Path                 string               `json:"path,omitempty" mapstructure:"path"`
	Auth                 *Auth                `json:"auth,omitempty" mapstructure:"auth"`
	Name                 string               `json:"name,omitempty" mapstructure:"name"`
	IpPool               *AccessNetworkIpPool `json:"ip_pool,omitempty" mapstructure:"ip_pool"`
	ServiceConfiguration string               `json:"service_configuration,omitempty" mapstructure:"service_configuration"`
	VolumeTemplates      []VolumeTemplates    `json:"volume_templates,omitempty" mapstructure:"volume_templates"`
	VolumeTemplatesEp    *VolumeTemplates     `json:"-"`
}

type StorageTemplates struct {
	Path string
}

type StorageTemplatesCreateRequest struct {
	Ctxt            context.Context `json:"-"`
	Name            string          `json:"name,omitempty" mapstructure:"name"`
	ReplicaCount    int             `json:"replica_count,omitempty" mapstructure:"replica_count"`
	Size            int             `json:"size,omitempty" mapstructure:"size"`
	PlacementMode   string          `json:"placement_mode,omitempty" mapstructure:"placement_mode"`
	PlacementPolicy string          `json:"placement_policy,omitempty" mapstructure:"placement_policy"`
	Force           bool            `json:"force,omitempty" mapstructure:"force"`
}

func newStorageTemplates(path string) *StorageTemplates {
	return &StorageTemplates{
		Path: _path.Join(path, "storage_templates"),
	}
}

func (e *StorageTemplates) Create(ro *StorageTemplatesCreateRequest) (*StorageTemplate, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type StorageTemplatesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

func (e *StorageTemplates) List(ro *StorageTemplatesListRequest) ([]*StorageTemplate, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := []*StorageTemplate{}
	for _, data := range rs.Data {
		elem := &StorageTemplate{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil
}

type StorageTemplatesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string
}

func (e *StorageTemplates) Get(ro *StorageTemplatesGetRequest) (*StorageTemplate, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Name), gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type StorageTemplateSetRequest struct {
	Ctxt            context.Context     `json:"-"`
	Auth            Auth                `json:"auth,omitempty" mapstructure:"auth"`
	IpPool          AccessNetworkIpPool `json:"ip_pool,omitempty" mapstructure:"ip_pool"`
	VolumeTemplates []VolumeTemplates   `json:"volume_templates,omitempty" mapstructure:"volume_templates"`
}

func (e *StorageTemplate) Set(ro *StorageTemplateSetRequest) (*StorageTemplate, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}

type StorageTemplateDeleteRequest struct {
	Ctxt  context.Context `json:"-"`
	Force bool            `json:"force,omitempty" mapstructure:"force"`
}

func (e *StorageTemplate) Delete(ro *StorageTemplateDeleteRequest) (*StorageTemplate, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &StorageTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
