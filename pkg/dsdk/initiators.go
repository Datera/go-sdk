package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type Initiator struct {
	Path string `json:"path,omitempty" mapstructure:"path"`
	Id   string `json:"id,omitempty" mapstructure:"id"`
	Name string `json:"name,omitempty" mapstructure:"name"`
	ctxt context.Context
	conn *ApiConnection
}

type Initiators struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type InitiatorsCreateRequest struct {
	Id    string `json:"id,omitempty" mapstructure:"id"`
	Name  string `json:"name,omitempty" mapstructure:"name"`
	Force bool   `json:"force,omitempty" mapstructure:"force"`
}

type InitiatorsCreateResponse Initiator

func newInitiators(ctxt context.Context, conn *ApiConnection, path string) *Initiators {
	return &Initiators{
		Path: _path.Join(path, "initiators"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *Initiators) Create(ro *InitiatorsCreateRequest) (*InitiatorsCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorsCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type InitiatorsListRequest struct {
	Params map[string]string
}

type InitiatorsListResponse []Initiator

func (e *Initiators) List(ro *InitiatorsListRequest) (*InitiatorsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := InitiatorsListResponse{}
	for _, data := range rs.Data {
		elem := &Initiator{}
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

type InitiatorsGetRequest struct {
	Id string
}

type InitiatorsGetResponse Initiator

func (e *Initiators) Get(ro *InitiatorsGetRequest) (*InitiatorsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type InitiatorSetRequest struct {
	Name string `json:"name,omitempty" mapstructure:"name"`
}

type InitiatorSetResponse Initiator

func (e *Initiator) Set(ro *InitiatorSetRequest) (*InitiatorSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil

}

type InitiatorDeleteRequest struct {
	Id string `json:"id,omitempty" mapstructure:"id"`
}

type InitiatorDeleteResponse Initiator

func (e *Initiator) Delete(ro *InitiatorDeleteRequest) (*InitiatorDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}
