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

type SnapshotPoliciesCreateResponse SnapshotPolicy

func newSnapshotPolicies(path string) *SnapshotPolicies {
	return &SnapshotPolicies{
		Path: _path.Join(path, "snapshot_policies"),
	}
}

func (e *SnapshotPolicies) Create(ro *SnapshotPoliciesCreateRequest) (*SnapshotPoliciesCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPoliciesCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type SnapshotPoliciesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

type SnapshotPoliciesListResponse []SnapshotPolicy

func (e *SnapshotPolicies) List(ro *SnapshotPoliciesListRequest) (*SnapshotPoliciesListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := SnapshotPoliciesListResponse{}
	for _, data := range rs.Data {
		elem := &SnapshotPolicy{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, *elem)
	}
	return &resp, nil
}

type SnapshotPoliciesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string
}

type SnapshotPoliciesGetResponse SnapshotPolicy

func (e *SnapshotPolicies) Get(ro *SnapshotPoliciesGetRequest) (*SnapshotPoliciesGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(_path.Join(e.Path, ro.Name), gro)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPoliciesGetResponse{}
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

type SnapshotPolicySetResponse SnapshotPolicy

func (e *SnapshotPolicy) Set(ro *SnapshotPolicySetRequest) (*SnapshotPolicySetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPolicySetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}

type SnapshotPolicyDeleteRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"id,omitempty" mapstructure:"id"`
}

type SnapshotPolicyDeleteResponse SnapshotPolicy

func (e *SnapshotPolicy) Delete(ro *SnapshotPolicyDeleteRequest) (*SnapshotPolicyDeleteResponse, error) {
	rs, err := GetConn(ro.Ctxt).Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPolicyDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
