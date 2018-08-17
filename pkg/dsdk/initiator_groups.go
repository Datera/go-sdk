package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type InitiatorGroup struct {
	Path    string      `json:"path,omitempty"`
	Name    string      `json:"name,omitempty"`
	Members []Initiator `json:"members,omitempty"`
	ctxt    context.Context
	conn    *ApiConnection
}

type InitiatorGroups struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type InitiatorGroupsCreateRequest struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Force bool   `json:"force,omitempty"`
}

type InitiatorGroupsCreateResponse InitiatorGroup

func newInitiatorGroups(ctxt context.Context, conn *ApiConnection, path string) *InitiatorGroups {
	return &InitiatorGroups{
		Path: _path.Join(path, "initiator_groups"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *InitiatorGroups) Create(ro *InitiatorGroupsCreateRequest) (*InitiatorGroupsCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorGroupsCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type InitiatorGroupsListRequest struct {
	Params map[string]string
}

type InitiatorGroupsListResponse []InitiatorGroup

func (e *InitiatorGroups) List(ro *InitiatorGroupsListRequest) (*InitiatorGroupsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := InitiatorGroupsListResponse{}
	for _, data := range rs.Data {
		elem := &InitiatorGroup{}
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

type InitiatorGroupsGetRequest struct {
	Id string
}

type InitiatorGroupsGetResponse InitiatorGroup

func (e *InitiatorGroups) Get(ro *InitiatorGroupsGetRequest) (*InitiatorGroupsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorGroupsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type InitiatorGroupSetRequest struct {
	Members []Initiator `json:"members,omitempty"`
}

type InitiatorGroupSetResponse InitiatorGroup

func (e *InitiatorGroup) Set(ro *InitiatorGroupSetRequest) (*InitiatorGroupSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorGroupSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil

}

type InitiatorGroupDeleteRequest struct {
	Id string `json:"id,omitempty"`
}

type InitiatorGroupDeleteResponse InitiatorGroup

func (e *InitiatorGroup) Delete(ro *InitiatorGroupDeleteRequest) (*InitiatorGroupDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorGroupDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}
