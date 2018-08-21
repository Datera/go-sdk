package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type Subsystem struct {
	Path        string   `json:"path,omitempty"`
	Causes      []string `json:"causes,omitempty"`
	Fan         string   `json:"fan,omitempty"`
	Health      string   `json:"health,omitempty"`
	Network     string   `json:"network,omitempty"`
	Power       string   `json:"power,omitempty"`
	Temperature string   `json:"temperature,omitempty"`
	Voltage     string   `json:"voltage,omitempty"`
	ctxt        context.Context
	conn        *ApiConnection
}

type Subsystems struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

func newSubsystems(ctxt context.Context, conn *ApiConnection, path string) *Subsystems {
	return &Subsystems{
		Path: _path.Join(path, "subsystem_states"),
		ctxt: ctxt,
		conn: conn,
	}
}

type SubsystemsListRequest struct {
	Params map[string]string
}

type SubsystemsListResponse []Subsystem

func (e *Subsystems) List(ro *SubsystemsListRequest) (*SubsystemsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := SubsystemsListResponse{}
	for _, data := range rs.Data {
		elem := &Subsystem{}
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

type SubsystemsGetRequest struct {
	Id string
}

type SubsystemsGetResponse Subsystem

func (e *Subsystems) Get(ro *SubsystemsGetRequest) (*SubsystemsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &SubsystemsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}
