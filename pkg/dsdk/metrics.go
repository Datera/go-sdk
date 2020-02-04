package dsdk

import (
	"context"
	"fmt"
	_path "path"

	greq "github.com/levigross/grequests"
)

type IOMetric string
type HWMetric string

const (
	Reads        IOMetric = "reads"
	Writes       IOMetric = "writes"
	BytesRead    IOMetric = "bytes_read"
	BytesWritten IOMetric = "bytes_written"
	IOPSRead     IOMetric = "iops_read"
	IOPSWrite    IOMetric = "iops_write"
	ThptRead     IOMetric = "thpt_read"
	ThptWrite    IOMetric = "thpt_write"
	LatAvgRead   IOMetric = "lat_avg_read"
	LatAvgWrite  IOMetric = "lat_avg_write"
	Lat50Read    IOMetric = "lat_50_read"
	Lat90Read    IOMetric = "lat_90_read"
	Lat100Read   IOMetric = "lat_100_read"
	Lat50Write   IOMetric = "lat_50_write"
	Lat90Write   IOMetric = "lat_90_write"
	Lat100Write  IOMetric = "lat_100_write"
)

func (io IOMetric) Validate() error {
	switch io {
	case Reads, Writes, BytesRead, BytesWritten, IOPSRead, IOPSWrite, ThptRead, ThptWrite, LatAvgRead,
		LatAvgWrite, Lat50Read, Lat50Write, Lat90Read, Lat90Write, Lat100Read, Lat100Write:
		return nil
	default:
		return fmt.Errorf("%s is not a valid IO metric", io)
	}
}

const (
	CPUUsage HWMetric = "cpu_usage"
)

func (hw HWMetric) Validate() error {
	switch hw {
	case CPUUsage:
		return nil
	default:
		return fmt.Errorf("%s is not a valid HW metric", hw)
	}
}

type MetricsParams struct {
	ListRangeParams
	Ival string
	UUID string
	Path string
}

func (mp MetricsParams) ToMap() map[string]string {
	r := mp.ListRangeParams.ToMap()
	if mp.Ival != "" {
		r["ival"] = mp.Ival
	}
	if mp.UUID != "" {
		r["uuid"] = mp.UUID
	}
	if mp.Path != "" {
		r["path"] = mp.Path
	}

	return r
}

type Metrics struct {
	EntityPath string
	Tenant     string
	Points     []Point
}

type Point struct {
	Time  int64
	Value float64
}

type IOMetricsRequest struct {
	Ctxt   context.Context `json:"-"`
	Type   IOMetric        `json:"-"`
	Params MetricsParams   `json:"params,omitempty"`
}

type HWMetricsRequest struct {
	Ctxt   context.Context `json:"-"`
	Type   IOMetric        `json:"-"`
	Params MetricsParams   `json:"params,omitempty"`
}

type IOMetrics struct {
	Path string
}

type HWMetrics struct {
	Path string
}

func newIOMetrics(path string) *IOMetrics {
	return &IOMetrics{
		Path: _path.Join(path, "metrics", "io"),
	}
}

func newHWMetrics(path string) *HWMetrics {
	return &HWMetrics{
		Path: _path.Join(path, "metrics", "hw"),
	}
}

func (m *IOMetrics) Get(ro *IOMetricsRequest) ([]*Metrics, *ApiErrorResponse, error) {
	if err := ro.Type.Validate(); err != nil {
		return nil, nil, err
	}

	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params.ToMap(),
	}

	rs, apierr, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, _path.Join(m.Path, string(ro.Type)), gro)
	if apierr != nil {
		return nil, apierr, err
	}

	if err != nil {
		return nil, nil, err
	}

	resp := []*Metrics{}
	for _, data := range rs.Data {
		elem := &Metrics{}
		edata := data.(map[string]interface{})
		if err = FillStruct(edata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
	}

	return resp, nil, nil
}

func (m *HWMetrics) Get(ro *HWMetricsRequest) ([]*Metrics, *ApiErrorResponse, error) {
	if err := ro.Type.Validate(); err != nil {
		return nil, nil, err
	}

	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params.ToMap(),
	}

	rs, apierr, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, _path.Join(m.Path, string(ro.Type)), gro)
	if apierr != nil {
		return nil, apierr, err
	}

	if err != nil {
		return nil, nil, err
	}

	resp := []*Metrics{}
	for _, data := range rs.Data {
		elem := &Metrics{}
		edata := data.(map[string]interface{})
		if err = FillStruct(edata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
	}

	return resp, nil, nil
}
