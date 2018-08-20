package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type SnapshotPolicy struct {
	Path           string `json:"path,omitempty"`
	Name           string `json:"name,omitempty"`
	Interval       string `json:"name,omitempty"`
	RetentionCount int    `json:"retention_count,omitempty"`
	StartTime      string `json:"start_time,omitempty"`
	ctxt           context.Context
	conn           *ApiConnection
}

type SnapshotPolicies struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type SnapshotPoliciesCreateRequest struct {
	Name           string `json:"name,omitempty"`
	Interval       string `json:"name,omitempty"`
	RetentionCount string `json:"retention_count,omitempty"`
	StartTime      string `json:"start_time,omitempty"`
}

type SnapshotPoliciesCreateResponse SnapshotPolicy

func newSnapshotPolicies(ctxt context.Context, conn *ApiConnection, path string) *SnapshotPolicies {
	return &SnapshotPolicies{
		Path: _path.Join(path, "snapshot_policies"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *SnapshotPolicies) Create(ro *SnapshotPoliciesCreateRequest) (*SnapshotPoliciesCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPoliciesCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type SnapshotPoliciesListRequest struct {
	Params map[string]string
}

type SnapshotPoliciesListResponse []SnapshotPolicy

func (e *SnapshotPolicies) List(ro *SnapshotPoliciesListRequest) (*SnapshotPoliciesListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
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
	for _, init := range resp {
		init.conn = e.conn
		init.ctxt = e.ctxt
	}
	return &resp, nil
}

type SnapshotPoliciesGetRequest struct {
	Id string
}

type SnapshotPoliciesGetResponse SnapshotPolicy

func (e *SnapshotPolicies) Get(ro *SnapshotPoliciesGetRequest) (*SnapshotPoliciesGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPoliciesGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type SnapshotPolicySetRequest struct {
	Interval       string `json:"name,omitempty"`
	RetentionCount int    `json:"retention_count,omitempty"`
	StartTime      string `json:"start_time,omitempty"`
}

type SnapshotPolicySetResponse SnapshotPolicy

func (e *SnapshotPolicy) Set(ro *SnapshotPolicySetRequest) (*SnapshotPolicySetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPolicySetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil

}

type SnapshotPolicyDeleteRequest struct {
	Id string `json:"id,omitempty"`
}

type SnapshotPolicyDeleteResponse SnapshotPolicy

func (e *SnapshotPolicy) Delete(ro *SnapshotPolicyDeleteRequest) (*SnapshotPolicyDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotPolicyDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}
