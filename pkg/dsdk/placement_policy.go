package dsdk

import (
	"context"
	"encoding/json"
	_path "path"

	greq "github.com/levigross/grequests"
)

type PlacementPolicy struct {
	Path           string `json:"path,omitempty" mapstructure:"path"`
	ResolvedPath   string `json:"resolved_path,omitempty" mapstructure:"resolved_path"`
	ResolvedTenant string `json:"resolved_tenant,omitempty" mapstructure:"resolved_tenant"`
}

func (p PlacementPolicy) MarshalJSON() ([]byte, error) {
	if p.Path == "" && p.ResolvedTenant == "" {
		return []byte(p.ResolvedPath), nil
	}
	m := map[string]string{
		"path":            p.Path,
		"resolved_path":   p.ResolvedPath,
		"resolved_tenant": p.ResolvedTenant,
	}
	return json.Marshal(m)
}

func (p PlacementPolicy) UnmarshalJSON(b []byte) error {
	np := map[string]string{}
	err := json.Unmarshal(b, &np)
	if err != nil {
		p.Path = ""
		p.ResolvedPath = string(b)
		p.ResolvedTenant = ""
	} else {
		p.Path = np["path"]
		p.ResolvedPath = np["resolved_path"]
		p.ResolvedTenant = np["resolved_tenant"]
	}
	return nil
}

type PlacementPolicies struct {
	Path string
}

type PlacementPoliciesCreateRequest struct {
	Ctxt  context.Context `json:"-"`
	Name  string          `json:"name,omitempty" mapstructure:"name"`
	Descr string          `json:"descr,omitempty" mapstructure:"descr"`
	Max   []string        `json:"max,omitempty" mapstructure:"max"`
	Min   []string        `json:"min,omitempty" mapstructure:"min"`
}

func newPlacementPolicies(path string) *PlacementPolicies {
	return &PlacementPolicies{
		Path: _path.Join(path, "PlacementPolicies"),
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
	Name string          `json:"name" mapstructure:"name"`
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
	Ctxt context.Context `json:"-"`
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

type PlacementPolicyReloadRequest struct {
	Ctxt context.Context `json:"-"`
}

func (e *PlacementPolicy) Reload(ro *PlacementPolicyReloadRequest) (*PlacementPolicy, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, e.Path, gro)
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
