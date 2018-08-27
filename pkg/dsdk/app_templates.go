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
	ctxt               context.Context
	conn               *ApiConnection
}

type AppTemplates struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type AppTemplatesCreateRequest struct {
	CopyFrom         *AppTemplate       `json:"copy_from,omitempty" mapstructure:"copy_from"`
	Name             string             `json:"name,omitempty" mapstructure:"name"`
	Descr            string             `json:"descr,omitempty" mapstructure:"descr"`
	SnapshotPolicies []*SnapshotPolicy  `json:"snapshot_policies,omitempty" mapstructure:"snapshot_policies"`
	StorageTemplates []*StorageTemplate `json:"storage_templates,omitempty" mapstructure:"storage_templates"`
}

type AppTemplatesCreateResponse AppTemplate

func newAppTemplates(ctxt context.Context, conn *ApiConnection, path string) *AppTemplates {
	return &AppTemplates{
		Path: _path.Join(path, "app_templates"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *AppTemplates) Create(ro *AppTemplatesCreateRequest) (*AppTemplatesCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AppTemplatesCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.StorageTemplatesEp = newStorageTemplates(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type AppTemplatesListRequest struct {
	Params map[string]string
}

type AppTemplatesListResponse []AppTemplate

func (e *AppTemplates) List(ro *AppTemplatesListRequest) (*AppTemplatesListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := AppTemplatesListResponse{}
	for _, data := range rs.Data {
		elem := &AppTemplate{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, *elem)
	}
	for _, r := range resp {
		r.conn = e.conn
		r.ctxt = e.ctxt
		r.StorageTemplatesEp = newStorageTemplates(e.ctxt, e.conn, e.Path)
	}
	return &resp, nil
}

type AppTemplatesGetRequest struct {
	Name string
}

type AppTemplatesGetResponse AppTemplate

func (e *AppTemplates) Get(ro *AppTemplatesGetRequest) (*AppTemplatesGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Name), gro)
	if err != nil {
		return nil, err
	}
	resp := &AppTemplatesGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.StorageTemplatesEp = newStorageTemplates(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type AppTemplateSetRequest struct {
	Descr            string             `json:"descr,omitempty" mapstructure:"descr"`
	SnapshotPolicies []*SnapshotPolicy  `json:"snapshot_policies,omitempty" mapstructure:"snapshot_policies"`
	StorageTemplates []*StorageTemplate `json:"storage_templates,omitempty" mapstructure:"storage_templates"`
}

type AppTemplateSetResponse AppTemplate

func (e *AppTemplate) Set(ro *AppTemplateSetRequest) (*AppTemplateSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AppTemplateSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.StorageTemplatesEp = newStorageTemplates(e.ctxt, e.conn, e.Path)
	return resp, nil

}

type AppTemplateDeleteRequest struct {
	Force bool `json:"force,omitempty" mapstructure:"force"`
}

type AppTemplateDeleteResponse AppTemplate

func (e *AppTemplate) Delete(ro *AppTemplateDeleteRequest) (*AppTemplateDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &AppTemplateDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.StorageTemplatesEp = newStorageTemplates(e.ctxt, e.conn, e.Path)
	return resp, nil
}
