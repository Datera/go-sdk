package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type InitiatorGroup struct {
	Path    string      `json:"path,omitempty" mapstructure:"path"`
	Name    string      `json:"name,omitempty" mapstructure:"name"`
	Members []Initiator `json:"members,omitempty" mapstructure:"members"`
}

type InitiatorGroups struct {
	Path string
}

type InitiatorGroupsCreateRequest struct {
	Ctxt  context.Context `json:"-"`
	Id    string          `json:"id,omitempty" mapstructure:"id"`
	Name  string          `json:"name,omitempty" mapstructure:"name"`
	Force bool            `json:"force,omitempty" mapstructure:"force"`
}

func newInitiatorGroups(path string) *InitiatorGroups {
	return &InitiatorGroups{
		Path: _path.Join(path, "initiator_groups"),
	}
}

func (e *InitiatorGroups) Create(ro *InitiatorGroupsCreateRequest) (*InitiatorGroup, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &InitiatorGroup{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type InitiatorGroupsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListParams      `json:"params,omitempty"`
}

func (e *InitiatorGroups) List(ro *InitiatorGroupsListRequest) ([]*InitiatorGroup, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params.ToMap()}
	rs, apierr, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := []*InitiatorGroup{}
	for _, data := range rs.Data {
		elem := &InitiatorGroup{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

type InitiatorGroupsGetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string          `json:"-"`
}

func (e *InitiatorGroups) Get(ro *InitiatorGroupsGetRequest) (*InitiatorGroup, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Name), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &InitiatorGroup{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type InitiatorGroupSetRequest struct {
	Ctxt    context.Context `json:"-"`
	Members []Initiator     `json:"members,omitempty" mapstructure:"members"`
}

func (e *InitiatorGroup) Set(ro *InitiatorGroupSetRequest) (*InitiatorGroup, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &InitiatorGroup{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil

}

type InitiatorGroupDeleteRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"id,omitempty" mapstructure:"id"`
}

func (e *InitiatorGroup) Delete(ro *InitiatorGroupDeleteRequest) (*InitiatorGroup, *ApiErrorResponse, error) {
	rs, apierr, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &InitiatorGroup{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}
