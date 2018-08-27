package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type FailureDomain struct {
	Path         string        `json:"path,omitempty" mapstructure:"path"`
	Name         string        `json:"name,omitempty" mapstructure:"name"`
	StorageNodes []StorageNode `json:"storage_nodes,omitempty" mapstructure:"storage_nodes"`
	ctxt         context.Context
	conn         *ApiConnection
}

type FailureDomains struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type FailureDomainsCreateRequest struct {
	Name         string        `json:"name,omitempty" mapstructure:"name"`
	StorageNodes []StorageNode `json:"storage_nodes,omitempty" mapstructure:"storage_nodes"`
}

type FailureDomainsCreateResponse FailureDomain

func newFailureDomains(ctxt context.Context, conn *ApiConnection, path string) *FailureDomains {
	return &FailureDomains{
		Path: _path.Join(path, "failure_domains"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *FailureDomains) Create(ro *FailureDomainsCreateRequest) (*FailureDomainsCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &FailureDomainsCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type FailureDomainsListRequest struct {
	Params map[string]string
}

type FailureDomainsListResponse []FailureDomain

func (e *FailureDomains) List(ro *FailureDomainsListRequest) (*FailureDomainsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := FailureDomainsListResponse{}
	for _, data := range rs.Data {
		elem := &FailureDomain{}
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

type FailureDomainsGetRequest struct {
	Id string
}

type FailureDomainsGetResponse FailureDomain

func (e *FailureDomains) Get(ro *FailureDomainsGetRequest) (*FailureDomainsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &FailureDomainsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type FailureDomainSetRequest struct {
	StorageNodes []StorageNode `json:"storage_nodes,omitempty" mapstructure:"storage_nodes"`
}

type FailureDomainSetResponse FailureDomain

func (e *FailureDomain) Set(ro *FailureDomainSetRequest) (*FailureDomainSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &FailureDomainSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil

}

type FailureDomainDeleteRequest struct {
	Name string `json:"id,omitempty" mapstructure:"id"`
}

type FailureDomainDeleteResponse FailureDomain

func (e *FailureDomain) Delete(ro *FailureDomainDeleteRequest) (*FailureDomainDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &FailureDomainDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}
