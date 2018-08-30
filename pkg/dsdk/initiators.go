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

func newInitiators(path string) *Initiators {
	return &Initiators{
		Path: _path.Join(path, "initiators"),
	}
}

func (e *Initiators) Create(ro *InitiatorsCreateRequest) (*Initiator, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &Initiator{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type InitiatorsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

func (e *Initiators) List(ro *InitiatorsListRequest) ([]*Initiator, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := []*Initiator{}
	for _, data := range rs.Data {
		elem := &Initiator{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil
}

type InitiatorsGetRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string
}

func (e *Initiators) Get(ro *InitiatorsGetRequest) (*Initiator, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &Initiator{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type InitiatorSetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string          `json:"name,omitempty" mapstructure:"name"`
}

func (e *Initiator) Set(ro *InitiatorSetRequest) (*Initiator, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &Initiator{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}

type InitiatorDeleteRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"id,omitempty" mapstructure:"id"`
}

func (e *Initiator) Delete(ro *InitiatorDeleteRequest) (*Initiator, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &Initiator{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
