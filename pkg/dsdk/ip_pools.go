package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type AccessNetworkIpPool struct {
	Path         string        `json:"path,omitempty"`
	Name         string        `json:"name,omitempty"`
	NetworkPaths []interface{} `json:"network_paths,omitempty"`
	Descr        string        `json:"descr,omitempty"`
	ctxt         context.Context
	conn         *ApiConnection
}

type AccessNetworkIpPools struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type AccessNetworkIpPoolsCreateRequest struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Force bool   `json:"force,omitempty"`
}

type AccessNetworkIpPoolsCreateResponse AccessNetworkIpPool

func newAccessNetworkIpPools(ctxt context.Context, conn *ApiConnection, path string) *AccessNetworkIpPools {
	return &AccessNetworkIpPools{
		Path: _path.Join(path, "access_network_ip_pools"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *AccessNetworkIpPools) Create(ro *AccessNetworkIpPoolsCreateRequest) (*AccessNetworkIpPoolsCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AccessNetworkIpPoolsCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type AccessNetworkIpPoolsListRequest struct {
	Params map[string]string
}

type AccessNetworkIpPoolsListResponse []AccessNetworkIpPool

func (e *AccessNetworkIpPools) List(ro *AccessNetworkIpPoolsListRequest) (*AccessNetworkIpPoolsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
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
	for _, init := range resp {
		init.conn = e.conn
		init.ctxt = e.ctxt
	}
	return &resp, nil
}

type AccessNetworkIpPoolsGetRequest struct {
	Id string
}

type AccessNetworkIpPoolsGetResponse AccessNetworkIpPool

func (e *AccessNetworkIpPools) Get(ro *AccessNetworkIpPoolsGetRequest) (*AccessNetworkIpPoolsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &AccessNetworkIpPoolsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type AccessNetworkIpPoolSetRequest struct {
	Members []Initiator `json:"members,omitempty"`
}

type AccessNetworkIpPoolSetResponse AccessNetworkIpPool

func (e *AccessNetworkIpPool) Set(ro *AccessNetworkIpPoolSetRequest) (*AccessNetworkIpPoolSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AccessNetworkIpPoolSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil

}

type AccessNetworkIpPoolDeleteRequest struct {
	Id string `json:"id,omitempty"`
}

type AccessNetworkIpPoolDeleteResponse AccessNetworkIpPool

func (e *AccessNetworkIpPool) Delete(ro *AccessNetworkIpPoolDeleteRequest) (*AccessNetworkIpPoolDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &AccessNetworkIpPoolDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}
