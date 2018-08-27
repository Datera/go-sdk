package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type BootDrive struct {
	Path      string   `json:"path,omitempty" mapstructure:"path"`
	Causes    []string `json:"causes,omitempty" mapstructure:"causes"`
	Health    string   `json:"health,omitempty" mapstructure:"health"`
	Id        string   `json:"id,omitempty" mapstructure:"id"`
	OpState   string   `json:"op_state,omitempty" mapstructure:"op_state"`
	Size      int      `json:"size,omitempty" mapstructure:"size"`
	SlotLabel string   `json:"slot_label,omitempty" mapstructure:"slot_label"`
	ctxt      context.Context
	conn      *ApiConnection
}

type BootDrives struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

func newBootDrives(ctxt context.Context, conn *ApiConnection, path string) *BootDrives {
	return &BootDrives{
		Path: _path.Join(path, "boot_drives"),
		ctxt: ctxt,
		conn: conn,
	}
}

type BootDrivesListRequest struct {
	Params map[string]string
}

type BootDrivesListResponse []BootDrive

func (e *BootDrives) List(ro *BootDrivesListRequest) (*BootDrivesListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := BootDrivesListResponse{}
	for _, data := range rs.Data {
		elem := &BootDrive{}
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

type BootDrivesGetRequest struct {
	Id string
}

type BootDrivesGetResponse BootDrive

func (e *BootDrives) Get(ro *BootDrivesGetRequest) (*BootDrivesGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &BootDrivesGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}
