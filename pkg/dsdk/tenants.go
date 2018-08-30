package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type Tenant struct {
	Path             string            `json:"path,omitempty" mapstructure:"path"`
	Descr            string            `json:"descr,omitempty" mapstructure:"descr"`
	InitiatorListSrc string            `json:"initiator_list_src,omitempty" mapstructure:"initiator_list_src"`
	MgmtIps          map[string]string `json:"mgmt_ips,omitempty" mapstructure:"mgmt_ips"`
	Name             string            `json:"name,omitempty" mapstructure:"name"`
	ParentPath       string            `json:"parent_path,omitempty" mapstructure:"parent_path"`
	Quota            Quota             `json:"quota,omitempty" mapstructure:"quota"`
	QuotaStatus      QuotaStatus       `json:"quota_status,omitempty" mapstructure:"quota_status"`
	Subtenants       []Tenant          `json:"subtenants,omitempty" mapstructure:"subtenants"`
}

type Tenants struct {
	Path string
}

type TenantsCreateRequest struct {
	Ctxt  context.Context `json:"-"`
	Id    string          `json:"id,omitempty" mapstructure:"id"`
	Name  string          `json:"name,omitempty" mapstructure:"name"`
	Force bool            `json:"force,omitempty" mapstructure:"force"`
}

func newTenants(path string) *Tenants {
	return &Tenants{
		Path: _path.Join(path, "tenants"),
	}
}

func (e *Tenants) Create(ro *TenantsCreateRequest) (*Tenant, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &Tenant{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type TenantsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

func (e *Tenants) List(ro *TenantsListRequest) ([]*Tenant, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := []*Tenant{}
	for _, data := range rs.Data {
		elem := &Tenant{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil
}

type TenantsGetRequest struct {
	Ctxt context.Context `json:"-"`
	Path string
}

func (e *Tenants) Get(ro *TenantsGetRequest) (*Tenant, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Path), gro)
	if err != nil {
		return nil, err
	}
	resp := &Tenant{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type TenantSetRequest struct {
	Ctxt             context.Context `json:"-"`
	Path             string          `json:"path,omitempty" mapstructure:"path"`
	Descr            string          `json:"descr,omitempty" mapstructure:"descr"`
	InitiatorListSrc string          `json:"initiator_list_src,omitempty" mapstructure:"initiator_list_src"`
	MgmtIps          []string        `json:"mgmt_ips,omitempty" mapstructure:"mgmt_ips"`
	Name             string          `json:"name,omitempty" mapstructure:"name"`
	ParentPath       string          `json:"parent_path,omitempty" mapstructure:"parent_path"`
	Quota            Quota           `json:"quota,omitempty" mapstructure:"quota"`
	QuotaStatus      QuotaStatus     `json:"quota_status,omitempty" mapstructure:"quota_status"`
	Subtenants       []Tenant        `json:"subtenants,omitempty" mapstructure:"subtenants"`
}

func (e *Tenant) Set(ro *TenantSetRequest) (*Tenant, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &Tenant{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}

type TenantDeleteRequest struct {
	Ctxt context.Context `json:"-"`
}

func (e *Tenant) Delete(ro *TenantDeleteRequest) (*Tenant, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &Tenant{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
