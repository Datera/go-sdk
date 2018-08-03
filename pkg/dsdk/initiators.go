package dsdk

import (
	"context"
	"path"

	greq "github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
)

type Initiator struct {
	Path string `json:"path,omitempty"`
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	conn *ApiConnection
}

type Initiators struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type InitiatorsCreateRequest struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Force bool   `json:"force,omitempty"`
}

type InitiatorsCreateResponse Initiator

type InitiatorsListRequest struct {
	Params map[string]string
}

type InitiatorsListResponse []Initiator

type InitiatorsGetRequest struct {
	Id string
}

type InitiatorsGetResponse Initiator

func newInitiators(ctxt context.Context, conn *ApiConnection) *Initiators {
	return &Initiators{
		Path: "initiators",
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
	return resp, nil
}

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
	}
	log.Debugf("INITIATORS---: %#v", resp)
	return &resp, nil
}

func (e *Initiators) Get(ro *InitiatorsGetRequest) (*InitiatorsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	return resp, nil
}

type InitiatorSetRequest struct {
	Name string `json:"name,omitempty"`
}

type InitiatorSetResponse Initiator

type InitiatorDeleteRequest struct {
	Id string `json:"id,omitempty"`
}

type InitiatorDeleteResponse Initiator

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
	return resp, nil

}

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
	return resp, nil
}
