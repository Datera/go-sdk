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

type StoragePoolsCreateResponse StoragePool

func newStoragePools(path string) *StoragePools {
	return &StoragePools{
		Path: _path.Join(path, "storage_pools"),
	}
}

func (e *StoragePools) Create(ro *StoragePoolsCreateRequest) (*StoragePoolsCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StoragePoolsCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type StoragePoolsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

type StoragePoolsListResponse []StoragePool

func (e *StoragePools) List(ro *StoragePoolsListRequest) (*StoragePoolsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := StoragePoolsListResponse{}
	for _, data := range rs.Data {
		elem := &StoragePool{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, *elem)
	}
	return &resp, nil
}

type StoragePoolsGetRequest struct {
	Ctxt context.Context `json:"-"`
	Uuid string          `json:"id,omitempty" mapstructure:"id"`
}

type StoragePoolsGetResponse StoragePool

func (e *StoragePools) Get(ro *StoragePoolsGetRequest) (*StoragePoolsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Uuid), gro)
	if err != nil {
		return nil, err
	}
	resp := &StoragePoolsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type StoragePoolSetRequest struct {
	Ctxt    context.Context `json:"-"`
	Members []*StorageNode  `json:"members,omitempty" mapstructure:"members"`
}

type StoragePoolSetResponse StoragePool

func (e *StoragePool) Set(ro *StoragePoolSetRequest) (*StoragePoolSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StoragePoolSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}

type StoragePoolDeleteRequest struct {
	Ctxt context.Context `json:"-"`
}

type StoragePoolDeleteResponse StoragePool

func (e *StoragePool) Delete(ro *StoragePoolDeleteRequest) (*StoragePoolDeleteResponse, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &StoragePoolDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
