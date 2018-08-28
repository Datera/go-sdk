package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type PerformancePolicy struct {
	Path              string `json:"path,omitempty" mapstructure:"path"`
	WriteIopsMax      int    `json:"write_iops_max,omitempty" mapstructure:"write_iops_max"`
	ReadIopsMax       int    `json:"read_iops_max,omitempty" mapstructure:"read_iops_max"`
	TotalIopsMax      int    `json:"total_iops_max,omitempty" mapstructure:"total_iops_max"`
	WriteBandwidthMax int    `json:"write_bandwidth_max,omitempty" mapstructure:"write_bandwidth_max"`
	ReadBandwidthMax  int    `json:"read_bandwidth_max,omitempty" mapstructure:"read_bandwidth_max"`
	TotalBandwidthMax int    `json:"total_bandwidth_max,omitempty" mapstructure:"total_bandwidth_max"`
}

type PerformancePolicyCreateRequest struct {
	Ctxt              context.Context `json:"-"`
	WriteIopsMax      int             `json:"write_iops_max,omitempty" mapstructure:"write_iops_max"`
	ReadIopsMax       int             `json:"read_iops_max,omitempty" mapstructure:"read_iops_max"`
	TotalIopsMax      int             `json:"total_iops_max,omitempty" mapstructure:"total_iops_max"`
	WriteBandwidthMax int             `json:"write_bandwidth_max,omitempty" mapstructure:"write_bandwidth_max"`
	ReadBandwidthMax  int             `json:"read_bandwidth_max,omitempty" mapstructure:"read_bandwidth_max"`
	TotalBandwidthMax int             `json:"total_bandwidth_max,omitempty" mapstructure:"total_bandwidth_max"`
}

type PerformancePolicyCreateResponse PerformancePolicy

func newPerformancePolicy(path string) *PerformancePolicy {
	return &PerformancePolicy{
		Path: _path.Join(path, "performance_policy"),
	}
}

func (e *PerformancePolicy) Create(ro *PerformancePolicyCreateRequest) (*PerformancePolicyCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicyCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type PerformancePolicyListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

type PerformancePolicyListResponse []PerformancePolicy

func (e *PerformancePolicy) List(ro *PerformancePolicyListRequest) (*PerformancePolicyListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(e.Path, gro)
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
	return &resp, nil
}

type PerformancePolicyGetRequest struct {
	Ctxt context.Context `json:"-"`
}

type PerformancePolicyGetResponse PerformancePolicy

func (e *PerformancePolicy) Get(ro *PerformancePolicyGetRequest) (*PerformancePolicyGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicyGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type PerformancePolicySetRequest struct {
	Ctxt              context.Context `json:"-"`
	WriteIopsMax      int             `json:"write_iops_max,omitempty" mapstructure:"write_iops_max"`
	ReadIopsMax       int             `json:"read_iops_max,omitempty" mapstructure:"read_iops_max"`
	TotalIopsMax      int             `json:"total_iops_max,omitempty" mapstructure:"total_iops_max"`
	WriteBandwidthMax int             `json:"write_bandwidth_max,omitempty" mapstructure:"write_bandwidth_max"`
	ReadBandwidthMax  int             `json:"read_bandwidth_max,omitempty" mapstructure:"read_bandwidth_max"`
	TotalBandwidthMax int             `json:"total_bandwidth_max,omitempty" mapstructure:"total_bandwidth_max"`
}

type PerformancePolicySetResponse PerformancePolicy

func (e *PerformancePolicy) Set(ro *PerformancePolicySetRequest) (*PerformancePolicySetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicySetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}

type PerformancePolicyDeleteRequest struct {
	Ctxt context.Context `json:"-"`
}

type PerformancePolicyDeleteResponse PerformancePolicy

func (e *PerformancePolicy) Delete(ro *PerformancePolicyDeleteRequest) (*PerformancePolicyDeleteResponse, error) {
	rs, err := GetConn(ro.Ctxt).Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicyDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
