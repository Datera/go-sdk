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

func newFailureDomains(path string) *FailureDomains {
	return &FailureDomains{
		Path: _path.Join(path, "failure_domains"),
	}
}

func (e *FailureDomains) Create(ro *FailureDomainsCreateRequest) (*FailureDomain, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &FailureDomain{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type FailureDomainsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

func (e *FailureDomains) List(ro *FailureDomainsListRequest) ([]*FailureDomain, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, apierr, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := []*FailureDomain{}
	for _, data := range rs.Data {
		elem := &FailureDomain{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

type FailureDomainsGetRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string
}

func (e *FailureDomains) Get(ro *FailureDomainsGetRequest) (*FailureDomain, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Id), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &FailureDomain{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type FailureDomainSetRequest struct {
	Ctxt         context.Context `json:"-"`
	StorageNodes []StorageNode   `json:"storage_nodes,omitempty" mapstructure:"storage_nodes"`
}

func (e *FailureDomain) Set(ro *FailureDomainSetRequest) (*FailureDomain, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &FailureDomain{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil

}

type FailureDomainDeleteRequest struct {
	Ctxt context.Context `json:"-"`
	Name string          `json:"id,omitempty" mapstructure:"id"`
}

func (e *FailureDomain) Delete(ro *FailureDomainDeleteRequest) (*FailureDomain, *ApiErrorResponse, error) {
	rs, apierr, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &FailureDomain{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}
