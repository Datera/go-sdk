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

func newPerformancePolicy(path string) *PerformancePolicy {
	return &PerformancePolicy{
		Path: _path.Join(path, "performance_policy"),
	}
}

func (e *PerformancePolicy) Create(ro *PerformancePolicyCreateRequest) (*PerformancePolicy, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type PerformancePolicyListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

func (e *PerformancePolicy) List(ro *PerformancePolicyListRequest) ([]*PerformancePolicy, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := []*PerformancePolicy{}
	for _, data := range rs.Data {
		elem := &PerformancePolicy{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil
}

type PerformancePolicyGetRequest struct {
	Ctxt context.Context `json:"-"`
}

func (e *PerformancePolicy) Get(ro *PerformancePolicyGetRequest) (*PerformancePolicy, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicy{}
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

func (e *PerformancePolicy) Set(ro *PerformancePolicySetRequest) (*PerformancePolicy, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}

type PerformancePolicyDeleteRequest struct {
	Ctxt context.Context `json:"-"`
}

func (e *PerformancePolicy) Delete(ro *PerformancePolicyDeleteRequest) (*PerformancePolicy, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &PerformancePolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
