package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type Tenant struct {
	Path             string      `json:"path,omitempty"`
	Descr            string      `json:"descr,omitempty"`
	InitiatorListSrc string      `json:"initiator_list_src,omitempty"`
	MgmtIps          []string    `json:"mgmt_ips,omitempty"`
	Name             string      `json:"name,omitempty"`
	ParentPath       string      `json:"parent_path,omitempty"`
	Quota            Quota       `json:"quota,omitempty"`
	QuotaStatus      QuotaStatus `json:"quota_status,omitempty"`
	Subtenants       []Tenant    `json:"subtenants,omitempty"`
	ctxt             context.Context
	conn             *ApiConnection
}

type Tenants struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type TenantsCreateRequest struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Force bool   `json:"force,omitempty"`
}

type TenantsCreateResponse Tenant

func newTenants(ctxt context.Context, conn *ApiConnection, path string) *Tenants {
	return &Tenants{
		Path: _path.Join(path, "tenants"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *Tenants) Create(ro *TenantsCreateRequest) (*TenantsCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &TenantsCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type TenantsListRequest struct {
	Params map[string]string
}

type TenantsListResponse []Tenant

func (e *Tenants) List(ro *TenantsListRequest) (*TenantsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := TenantsListResponse{}
	for _, data := range rs.Data {
		elem := &Tenant{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, *elem)
	}
	for _, init := range resp {
		init.conn = e.conn
		init.ctxt = e.ctxt
	}
	return &resp, nil
}

type TenantsGetRequest struct {
	Path string
}

type TenantsGetResponse Tenant

func (e *Tenants) Get(ro *TenantsGetRequest) (*TenantsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Path), gro)
	if err != nil {
		return nil, err
	}
	resp := &TenantsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type TenantSetRequest struct {
	Path             string      `json:"path,omitempty"`
	Descr            string      `json:"descr,omitempty"`
	InitiatorListSrc string      `json:"initiator_list_src,omitempty"`
	MgmtIps          []string    `json:"mgmt_ips,omitempty"`
	Name             string      `json:"name,omitempty"`
	ParentPath       string      `json:"parent_path,omitempty"`
	Quota            Quota       `json:"quota,omitempty"`
	QuotaStatus      QuotaStatus `json:"quota_status,omitempty"`
	Subtenants       []Tenant    `json:"subtenants,omitempty"`
}

type TenantSetResponse Tenant

func (e *Tenant) Set(ro *TenantSetRequest) (*TenantSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &TenantSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil

}

type TenantDeleteRequest struct {
}

type TenantDeleteResponse Tenant

func (e *Tenant) Delete(ro *TenantDeleteRequest) (*TenantDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &TenantDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}
