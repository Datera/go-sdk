package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type AccessNetworkIpPool struct {
	Path         string        `json:"path,omitempty" mapstructure:"path"`
	Name         string        `json:"name,omitempty" mapstructure:"name"`
	NetworkPaths []interface{} `json:"network_paths,omitempty" mapstructure:"network_paths"`
	Descr        string        `json:"descr,omitempty" mapstructure:"descr"`
}

type AccessNetworkIpPools struct {
	Path string
}

type AccessNetworkIpPoolsCreateRequest struct {
	Ctxt  context.Context `json:"-"`
	Id    string          `json:"id,omitempty" mapstructure:"id"`
	Name  string          `json:"name,omitempty" mapstructure:"name"`
	Force bool            `json:"force,omitempty" mapstructure:"force"`
}

func newAccessNetworkIpPools(path string) *AccessNetworkIpPools {
	return &AccessNetworkIpPools{
		Path: _path.Join(path, "access_network_ip_pools"),
	}
}

func (e *AccessNetworkIpPools) Create(ro *AccessNetworkIpPoolsCreateRequest) (*AccessNetworkIpPool, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AccessNetworkIpPool{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type AccessNetworkIpPoolsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListParams `json:"params,omitempty"`
}

func (e *AccessNetworkIpPools) List(ro *AccessNetworkIpPoolsListRequest) ([]*AccessNetworkIpPool, *ApiErrorResponse, error) {
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
	resp := []*AccessNetworkIpPool{}
	for _, data := range rs.Data {
		elem := &AccessNetworkIpPool{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

type AccessNetworkIpPoolsGetRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"-"`
}

func (e *AccessNetworkIpPools) Get(ro *AccessNetworkIpPoolsGetRequest) (*AccessNetworkIpPool, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Id), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AccessNetworkIpPool{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type AccessNetworkIpPoolSetRequest struct {
	Ctxt    context.Context `json:"-"`
	Members []Initiator     `json:"members,omitempty" mapstructure:"members"`
}

func (e *AccessNetworkIpPool) Set(ro *AccessNetworkIpPoolSetRequest) (*AccessNetworkIpPool, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AccessNetworkIpPool{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil

}

type AccessNetworkIpPoolDeleteRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"id,omitempty" mapstructure:"id"`
}

func (e *AccessNetworkIpPool) Delete(ro *AccessNetworkIpPoolDeleteRequest) (*AccessNetworkIpPool, *ApiErrorResponse, error) {
	rs, apierr, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &AccessNetworkIpPool{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}
