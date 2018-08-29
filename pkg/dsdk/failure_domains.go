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
}

type FailureDomains struct {
	Path string
}

type FailureDomainsCreateRequest struct {
	Ctxt         context.Context `json:"-"`
	Name         string          `json:"name,omitempty" mapstructure:"name"`
	StorageNodes []StorageNode   `json:"storage_nodes,omitempty" mapstructure:"storage_nodes"`
}

type FailureDomainsCreateResponse FailureDomain

func newFailureDomains(path string) *FailureDomains {
	return &FailureDomains{
		Path: _path.Join(path, "failure_domains"),
	}
}

func (e *FailureDomains) Create(ro *FailureDomainsCreateRequest) (*FailureDomainsCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &FailureDomainsCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type FailureDomainsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

type FailureDomainsListResponse []FailureDomain

func (e *FailureDomains) List(ro *FailureDomainsListRequest) (*FailureDomainsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
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
	return &resp, nil
}

type FailureDomainsGetRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string
}

type FailureDomainsGetResponse FailureDomain

func (e *FailureDomains) Get(ro *FailureDomainsGetRequest) (*FailureDomainsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &FailureDomainsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type FailureDomainSetRequest struct {
	Ctxt         context.Context `json:"-"`
	StorageNodes []StorageNode   `json:"storage_nodes,omitempty" mapstructure:"storage_nodes"`
}

type FailureDomainSetResponse FailureDomain

func (e *FailureDomain) Set(ro *FailureDomainSetRequest) (*FailureDomainSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &FailureDomainSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}

type FailureDomainDeleteRequest struct {
	Ctxt context.Context `json:"-"`
	Name string          `json:"id,omitempty" mapstructure:"id"`
}

type FailureDomainDeleteResponse FailureDomain

func (e *FailureDomain) Delete(ro *FailureDomainDeleteRequest) (*FailureDomainDeleteResponse, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &FailureDomainDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
