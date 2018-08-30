package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type SnapshotPolicy struct {
	Path           string `json:"path,omitempty" mapstructure:"path"`
	Name           string `json:"name,omitempty" mapstructure:"name"`
	Interval       string `json:"interval,omitempty" mapstructure:"interval"`
	RetentionCount int    `json:"retention_count,omitempty" mapstructure:"retention_count"`
	StartTime      string `json:"start_time,omitempty" mapstructure:"start_time"`
}

type SnapshotPolicies struct {
	Path string
}

type SnapshotPoliciesCreateRequest struct {
	Ctxt           context.Context `json:"-"`
	Name           string          `json:"name,omitempty" mapstructure:"name"`
	Interval       string          `json:"interval,omitempty" mapstructure:"interval"`
	RetentionCount string          `json:"retention_count,omitempty" mapstructure:"retention_count"`
	StartTime      string          `json:"start_time,omitempty" mapstructure:"start_time"`
}

func newSnapshotPolicies(path string) *SnapshotPolicies {
	return &SnapshotPolicies{
		Path: _path.Join(path, "snapshot_policies"),
	}
}

func (e *SnapshotPolicies) Create(ro *SnapshotPoliciesCreateRequest) (*SnapshotPolicy, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type SnapshotPoliciesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

func (e *SnapshotPolicies) List(ro *SnapshotPoliciesListRequest) ([]*SnapshotPolicy, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := []*SnapshotPolicy{}
	for _, data := range rs.Data {
		elem := &SnapshotPolicy{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil
}

type SnapshotPoliciesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string
}

func (e *SnapshotPolicies) Get(ro *SnapshotPoliciesGetRequest) (*SnapshotPolicy, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Name), gro)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type SnapshotPolicySetRequest struct {
	Ctxt           context.Context `json:"-"`
	Interval       string          `json:"name,omitempty" mapstructure:"name"`
	RetentionCount int             `json:"retention_count,omitempty" mapstructure:"retention_count"`
	StartTime      string          `json:"start_time,omitempty" mapstructure:"start_time"`
}

func (e *SnapshotPolicy) Set(ro *SnapshotPolicySetRequest) (*SnapshotPolicy, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}

type SnapshotPolicyDeleteRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"id,omitempty" mapstructure:"id"`
}

func (e *SnapshotPolicy) Delete(ro *SnapshotPolicyDeleteRequest) (*SnapshotPolicy, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
