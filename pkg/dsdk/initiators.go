package dsdk

import (
	"context"

	greq "github.com/levigross/grequests"
	// log "github.com/sirupsen/logrus"
)

type Initiators struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type InitiatorsCreateRequest struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Force bool   `json:"force"`
}

type InitiatorsCreateResponse Initiator

type InitiatorsListRequest struct {
	Filter string `json:"filter"`
	Limit  string `json:"limit"`
	Sort   string `json:"sort"`
	Offset string `json:"offset"`
}

type InitiatorsListResponse []Initiator

type InitiatorGetResponse Initiator

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
	return &InitiatorsListResponse{}, nil
}

// func (e *Initiators) Get(ro *InitiatorsGetRequest) (*InitiatorsGetResponse, error) {
// 	return &InitiatorsGetResponse{}, nil
// }

type Initiator struct {
	Path string `json:"path"`
	Id   string `json:"id"`
	Name string `json:"name"`
	conn *ApiConnection
}

type InitiatorSetRequest struct {
	Name string `json:"id"`
}

type InititatorSetResponse Initiator

type InitiatorDeleteRequest struct {
	Id string `json:"id"`
}

type InitiatorDeleteResponse Initiator

// func (e *Initiator) Set(ro *InitiatorSetRequest) (*InitiatorSetResponse, error) {
// 	return &InitiatorSetResponse{}, nil

// }

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
