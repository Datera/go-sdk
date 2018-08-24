package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type PerformancePolicy struct {
	Path              string `json:"path,omitempty"`
	WriteIopsMax      int    `json:"write_iops_max,omitempty"`
	ReadIopsMax       int    `json:"read_iops_max,omitempty"`
	TotalIopsMax      int    `json:"total_iops_max,omitempty"`
	WriteBandwidthMax int    `json:"write_bandwidth_max,omitempty"`
	ReadBandwidthMax  int    `json:"read_bandwidth_max,omitempty"`
	TotalBandwidthMax int    `json:"total_bandwidth_max,omitempty"`
	ctxt              context.Context
	conn              *ApiConnection
}

type PerformancePolicyCreateRequest struct {
	WriteIopsMax      int `json:"write_iops_max,omitempty"`
	ReadIopsMax       int `json:"read_iops_max,omitempty"`
	TotalIopsMax      int `json:"total_iops_max,omitempty"`
	WriteBandwidthMax int `json:"write_bandwidth_max,omitempty"`
	ReadBandwidthMax  int `json:"read_bandwidth_max,omitempty"`
	TotalBandwidthMax int `json:"total_bandwidth_max,omitempty"`
}

type PerformancePolicyCreateResponse PerformancePolicy

func newPerformancePolicy(ctxt context.Context, conn *ApiConnection, path string) *PerformancePolicy {
	return &PerformancePolicy{
		Path: _path.Join(path, "performance_policy"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *PerformancePolicy) Create(ro *PerformancePolicyCreateRequest) (*PerformancePolicyCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicyCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type PerformancePolicyListRequest struct {
	Params map[string]string
}

type PerformancePolicyListResponse []PerformancePolicy

func (e *PerformancePolicy) List(ro *PerformancePolicyListRequest) (*PerformancePolicyListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := PerformancePolicyListResponse{}
	for _, data := range rs.Data {
		elem := &PerformancePolicy{}
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

type PerformancePolicyGetRequest struct {
}

type PerformancePolicyGetResponse PerformancePolicy

func (e *PerformancePolicy) Get(ro *PerformancePolicyGetRequest) (*PerformancePolicyGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicyGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type PerformancePolicySetRequest struct {
	WriteIopsMax      int `json:"write_iops_max,omitempty"`
	ReadIopsMax       int `json:"read_iops_max,omitempty"`
	TotalIopsMax      int `json:"total_iops_max,omitempty"`
	WriteBandwidthMax int `json:"write_bandwidth_max,omitempty"`
	ReadBandwidthMax  int `json:"read_bandwidth_max,omitempty"`
	TotalBandwidthMax int `json:"total_bandwidth_max,omitempty"`
}

type PerformancePolicySetResponse PerformancePolicy

func (e *PerformancePolicy) Set(ro *PerformancePolicySetRequest) (*PerformancePolicySetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicySetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil

}

type PerformancePolicyDeleteRequest struct {
}

type PerformancePolicyDeleteResponse PerformancePolicy

func (e *PerformancePolicy) Delete(ro *PerformancePolicyDeleteRequest) (*PerformancePolicyDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicyDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}
