package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type PlacementPolicy struct {
	Path         string        `json:"path,omitempty" mapstructure:"path"`
	Name         string        `json:"name,omitempty" mapstructure:"name"`
        Max          []interface{} `json:"max,omitempty" mapstructure:"max"`
        Min          []interface{} `json:"min,omitempty" mapstructure:"min"`
	Descr        string        `json:"descr,omitempty" mapstructure:"descr"`
}

type PlacementPolicies struct {
	Path string
}

type PlacementPoliciesCreateRequest struct {
	Ctxt  context.Context `json:"-"`
	Id    string          `json:"id,omitempty" mapstructure:"id"`
	Name  string          `json:"name,omitempty" mapstructure:"name"`
	Force bool            `json:"force,omitempty" mapstructure:"force"`
}

func newPlacementPolicies(path string) *PlacementPolicies {
	return &PlacementPolicies{
		Path: _path.Join(path, "placement_policies"),
	}
}

func (e *PlacementPolicies) Create(ro *PlacementPoliciesCreateRequest) (*PlacementPolicy, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &PlacementPolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type PlacementPoliciesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListParams      `json:"params,omitempty"`
}

func (e *PlacementPolicies) List(ro *PlacementPoliciesListRequest) ([]*PlacementPolicy, *ApiErrorResponse, error) {
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
	resp := []*PlacementPolicy{}
	for _, data := range rs.Data {
		elem := &PlacementPolicy{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

type PlacementPoliciesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string          `json:"-"`
}

func (e *PlacementPolicies) Get(ro *PlacementPoliciesGetRequest) (*PlacementPolicy, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Name), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &PlacementPolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type PlacementPolicySetRequest struct {
	Ctxt    context.Context `json:"-"`
	MaxMembers []string     `json:"members,omitempty" mapstructure:"max_members"`
        MinMembers []string     `json:"members,omitempty" mapstructure:"min_members"`
}

func (e *PlacementPolicy) Set(ro *PlacementPolicySetRequest) (*PlacementPolicy, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &PlacementPolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil

}

type PlacementPolicyDeleteRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"id,omitempty" mapstructure:"id"`
}

func (e *PlacementPolicy) Delete(ro *PlacementPolicyDeleteRequest) (*PlacementPolicy, *ApiErrorResponse, error) {
	rs, apierr, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &PlacementPolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}
