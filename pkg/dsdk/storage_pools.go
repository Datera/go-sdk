package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type StoragePool struct {
	Path    string         `json:"path,omitempty" mapstructure:"path"`
	Members []*StorageNode `json:"members,omitempty" mapstructure:"members"`
	Name    string         `json:"name,omitempty" mapstructure:"name"`
}

type StoragePools struct {
	Path string
}

type StoragePoolsCreateRequest struct {
	Ctxt    context.Context `json:"-"`
	Members []*StorageNode  `json:"members,omitempty" mapstructure:"members"`
	Name    string          `json:"name,omitempty" mapstructure:"name"`
}

func newStoragePools(path string) *StoragePools {
	return &StoragePools{
		Path: _path.Join(path, "storage_pools"),
	}
}

func (e *StoragePools) Create(ro *StoragePoolsCreateRequest) (*StoragePool, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &StoragePool{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type StoragePoolsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListParams
}

func (e *StoragePools) List(ro *StoragePoolsListRequest) ([]*StoragePool, *ApiErrorResponse, error) {
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
	resp := []*StoragePool{}
	for _, data := range rs.Data {
		elem := &StoragePool{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

type StoragePoolsGetRequest struct {
	Ctxt context.Context `json:"-"`
	Uuid string          `json:"-"`
}

func (e *StoragePools) Get(ro *StoragePoolsGetRequest) (*StoragePool, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Uuid), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &StoragePool{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type StoragePoolSetRequest struct {
	Ctxt    context.Context `json:"-"`
	Members []*StorageNode  `json:"members,omitempty" mapstructure:"members"`
}

func (e *StoragePool) Set(ro *StoragePoolSetRequest) (*StoragePool, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &StoragePool{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil

}

type StoragePoolDeleteRequest struct {
	Ctxt context.Context `json:"-"`
}

func (e *StoragePool) Delete(ro *StoragePoolDeleteRequest) (*StoragePool, *ApiErrorResponse, error) {
	rs, apierr, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &StoragePool{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}
