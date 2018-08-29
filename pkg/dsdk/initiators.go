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
}

type Initiators struct {
	Path string
}

type InitiatorsCreateRequest struct {
	Ctxt  context.Context `json:"-"`
	Id    string          `json:"id,omitempty" mapstructure:"id"`
	Name  string          `json:"name,omitempty" mapstructure:"name"`
	Force bool            `json:"force,omitempty" mapstructure:"force"`
}

type InitiatorsCreateResponse Initiator

func newInitiators(path string) *Initiators {
	return &Initiators{
		Path: _path.Join(path, "initiators"),
	}
}

func (e *Initiators) Create(ro *InitiatorsCreateRequest) (*InitiatorsCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorsCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type InitiatorsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

type InitiatorsListResponse []Initiator

func (e *Initiators) List(ro *InitiatorsListRequest) (*InitiatorsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
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
	return &resp, nil
}

type InitiatorsGetRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string
}

type InitiatorsGetResponse Initiator

func (e *Initiators) Get(ro *InitiatorsGetRequest) (*InitiatorsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type InitiatorSetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string          `json:"name,omitempty" mapstructure:"name"`
}

type InitiatorSetResponse Initiator

func (e *Initiator) Set(ro *InitiatorSetRequest) (*InitiatorSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}

type InitiatorDeleteRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"id,omitempty" mapstructure:"id"`
}

type InitiatorDeleteResponse Initiator

func (e *Initiator) Delete(ro *InitiatorDeleteRequest) (*InitiatorDeleteResponse, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &InitiatorDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
