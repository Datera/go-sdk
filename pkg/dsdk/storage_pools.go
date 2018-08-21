package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type StoragePool struct {
	Path    string        `json:"path,omitempty"`
	Members []StorageNode `json:"members,omitempty"`
	Name    string        `json:"name,omitempty"`
	ctxt    context.Context
	conn    *ApiConnection
}

type StoragePools struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type StoragePoolsCreateRequest struct {
	Members []StorageNode `json:"members,omitempty"`
	Name    string        `json:"name,omitempty"`
}

type StoragePoolsCreateResponse StoragePool

func newStoragePools(ctxt context.Context, conn *ApiConnection, path string) *StoragePools {
	return &StoragePools{
		Path: _path.Join(path, "storage_pool"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *StoragePools) Create(ro *StoragePoolsCreateRequest) (*StoragePoolsCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StoragePoolsCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type StoragePoolsListRequest struct {
	Params map[string]string
}

type StoragePoolsListResponse []StoragePool

func (e *StoragePools) List(ro *StoragePoolsListRequest) (*StoragePoolsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
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
	for _, init := range resp {
		init.conn = e.conn
		init.ctxt = e.ctxt
	}
	return &resp, nil
}

type StoragePoolsGetRequest struct {
	Uuid string `json:"id,omitempty"`
}

type StoragePoolsGetResponse StoragePool

func (e *StoragePools) Get(ro *StoragePoolsGetRequest) (*StoragePoolsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Uuid), gro)
	if err != nil {
		return nil, err
	}
	resp := &StoragePoolsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type StoragePoolSetRequest struct {
	Members []StorageNode `json:"members,omitempty"`
}

type StoragePoolSetResponse StoragePool

func (e *StoragePool) Set(ro *StoragePoolSetRequest) (*StoragePoolSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StoragePoolSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil

}

type StoragePoolDeleteRequest struct {
}

type StoragePoolDeleteResponse StoragePool

func (e *StoragePool) Delete(ro *StoragePoolDeleteRequest) (*StoragePoolDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &StoragePoolDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}
