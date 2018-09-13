package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type AppTemplate struct {
	Path               string             `json:"path,omitempty" mapstructure:"path"`
	AppInstances       []*AppInstance     `json:"app_instances,omitempty" mapstructure:"app_instances"`
	Name               string             `json:"name,omitempty" mapstructure:"name"`
	Descr              string             `json:"descr,omitempty" mapstructure:"descr"`
	SnapshotPolicies   []*SnapshotPolicy  `json:"snapshot_policies,omitempty" mapstructure:"snapshot_policies"`
	StorageTemplates   []*StorageTemplate `json:"storage_templates,omitempty" mapstructure:"storage_templates"`
	StorageTemplatesEp *StorageTemplates  `json:"-"`
}

func RegisterAppTemplateEndpoints(a *AppTemplate) {
	a.StorageTemplatesEp = newStorageTemplates(a.Path)
	for _, si := range a.AppInstances {
		RegisterAppInstanceEndpoints(si)
	}
	for _, si := range a.StorageTemplates {
		RegisterStorageTemplateEndpoints(si)
	}
}

type AppTemplates struct {
	Path string
}

type AppTemplatesCreateRequest struct {
	Ctxt             context.Context    `json:"-"`
	CopyFrom         *AppTemplate       `json:"copy_from,omitempty" mapstructure:"copy_from"`
	Name             string             `json:"name,omitempty" mapstructure:"name"`
	Descr            string             `json:"descr,omitempty" mapstructure:"descr"`
	SnapshotPolicies []*SnapshotPolicy  `json:"snapshot_policies,omitempty" mapstructure:"snapshot_policies"`
	StorageTemplates []*StorageTemplate `json:"storage_templates,omitempty" mapstructure:"storage_templates"`
}

func newAppTemplates(path string) *AppTemplates {
	return &AppTemplates{
		Path: _path.Join(path, "app_templates"),
	}
}

func (e *AppTemplates) Create(ro *AppTemplatesCreateRequest) (*AppTemplate, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AppTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	resp.StorageTemplatesEp = newStorageTemplates(e.Path)
	RegisterAppTemplateEndpoints(resp)
	return resp, nil, nil
}

type AppTemplatesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListParams
}

func (e *AppTemplates) List(ro *AppTemplatesListRequest) ([]*AppTemplate, *ApiErrorResponse, error) {
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
	resp := []*AppTemplate{}
	for _, data := range rs.Data {
		elem := &AppTemplate{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		RegisterAppTemplateEndpoints(elem)
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

type AppTemplatesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string          `json:"-"`
}

func (e *AppTemplates) Get(ro *AppTemplatesGetRequest) (*AppTemplate, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Name), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AppTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterAppTemplateEndpoints(resp)
	return resp, nil, nil
}

type AppTemplateSetRequest struct {
	Ctxt             context.Context    `json:"-"`
	Descr            string             `json:"descr,omitempty" mapstructure:"descr"`
	SnapshotPolicies []*SnapshotPolicy  `json:"snapshot_policies,omitempty" mapstructure:"snapshot_policies"`
	StorageTemplates []*StorageTemplate `json:"storage_templates,omitempty" mapstructure:"storage_templates"`
}

func (e *AppTemplate) Set(ro *AppTemplateSetRequest) (*AppTemplate, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AppTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterAppTemplateEndpoints(resp)
	return resp, nil, nil

}

type AppTemplateDeleteRequest struct {
	Ctxt  context.Context `json:"-"`
	Force bool            `json:"force,omitempty" mapstructure:"force"`
}

func (e *AppTemplate) Delete(ro *AppTemplateDeleteRequest) (*AppTemplate, *ApiErrorResponse, error) {
	rs, apierr, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AppTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterAppTemplateEndpoints(resp)
	return resp, nil, nil
}
