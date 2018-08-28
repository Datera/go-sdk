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

type AccessNetworkIpPoolsCreateResponse AccessNetworkIpPool

func newAccessNetworkIpPools(path string) *AccessNetworkIpPools {
	return &AccessNetworkIpPools{
		Path: _path.Join(path, "access_network_ip_pools"),
	}
}

func (e *AccessNetworkIpPools) Create(ro *AccessNetworkIpPoolsCreateRequest) (*AccessNetworkIpPoolsCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AccessNetworkIpPoolsCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type AccessNetworkIpPoolsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

type AccessNetworkIpPoolsListResponse []AccessNetworkIpPool

func (e *AccessNetworkIpPools) List(ro *AccessNetworkIpPoolsListRequest) (*AccessNetworkIpPoolsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := AccessNetworkIpPoolsListResponse{}
	for _, data := range rs.Data {
		elem := &AccessNetworkIpPool{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, *elem)
	}
	return &resp, nil
}

type AccessNetworkIpPoolsGetRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string
}

type AccessNetworkIpPoolsGetResponse AccessNetworkIpPool

func (e *AccessNetworkIpPools) Get(ro *AccessNetworkIpPoolsGetRequest) (*AccessNetworkIpPoolsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &AccessNetworkIpPoolsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type AccessNetworkIpPoolSetRequest struct {
	Ctxt    context.Context `json:"-"`
	Members []Initiator     `json:"members,omitempty" mapstructure:"members"`
}

type AccessNetworkIpPoolSetResponse AccessNetworkIpPool

func (e *AccessNetworkIpPool) Set(ro *AccessNetworkIpPoolSetRequest) (*AccessNetworkIpPoolSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AccessNetworkIpPoolSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}

type AccessNetworkIpPoolDeleteRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"id,omitempty" mapstructure:"id"`
}

type AccessNetworkIpPoolDeleteResponse AccessNetworkIpPool

func (e *AccessNetworkIpPool) Delete(ro *AccessNetworkIpPoolDeleteRequest) (*AccessNetworkIpPoolDeleteResponse, error) {
	rs, err := GetConn(ro.Ctxt).Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &AccessNetworkIpPoolDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
