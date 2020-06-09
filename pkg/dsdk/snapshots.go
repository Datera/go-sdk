package dsdk

import (
	"context"
	_path "path"
	"reflect"

	greq "github.com/levigross/grequests"
)

type Snapshot struct {
	Path            string            `json:"path,omitempty" mapstructure:"path"`
	Timestamp       string            `json:"timestamp,omitempty" mapstructure:"timestamp"`
	Uuid            string            `json:"uuid,omitempty" mapstructure:"uuid"`
	RemoteProviders []*RemoteProvider `json:"remote_providers,omitempty" mapstructure:"remote_providers"`
	OpState         string            `json:"op_state,omitempty" mapstructure:"op_state"`
	UtcTs           string            `json:"utc_ts,omitempty" mapstructure:"utc_ts"`
	PhysicalSize    int               `json:"physical_size,omitempty" mapstructure:"physical_size"`
	LogicalSize     int               `json:"logical_size,omitempty" mapstructure:"logical_size"`
	ExclusiveSize   int               `json:"exclusive_size,omitempty" mapstructure:"exclusive_size"`
	EffectiveSize   int               `json:"effective_size,omitempty" mapstructure:"effective_size"`
	Local           bool              `json:"local,omitempty" mapstructure:"local"`
	AppStructure    interface{}       `json:"app_structure,omitempty" mapstructure:"app_structure"`
	TsVersion       string            `json:"ts_version,omitempty" mapstructure:"ts_version"`
	Version         string            `json:"version,omitempty" mapstructure:"version"`
}

type Snapshots struct {
	Path string
}

type SnapshotsCreateRequest struct {
	Ctxt               context.Context `json:"-"`
	Uuid               string          `json:"uuid,omitempty" mapstructure:"uuid"`
	RemoteProviderUuid string          `json:"remote_provider_uuid,omitempty" mapstructure:"remote_provider_uuid"`
	Type               string          `json:"type,omitempty" mapstructure:"type"`
}

func newSnapshots(path string) *Snapshots {
	return &Snapshots{
		Path: _path.Join(path, "snapshots"),
	}
}

func (e *Snapshots) Create(ro *SnapshotsCreateRequest) (*Snapshot, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &Snapshot{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type SnapshotsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListParams      `json:"params,omitempty"`
}

func (e *Snapshots) List(ro *SnapshotsListRequest) ([]*Snapshot, *ApiErrorResponse, error) {
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
	resp := []*Snapshot{}
	for _, data := range rs.Data {
		elem := &Snapshot{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

type SnapshotsGetRequest struct {
	Ctxt      context.Context `json:"-"`
	Timestamp string          `json:"-"`
}

func (e *Snapshots) Get(ro *SnapshotsGetRequest) (*Snapshot, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Timestamp), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &Snapshot{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type SnapshotDeleteRequest struct {
	Ctxt               context.Context `json:"-"`
	RemoteProviderUuid string          `json:"remote_provider_uuid,omitempty" mapstructure:"remote_provider_uuid"`
	Force              bool            `json:"force,omitempty" mapstructure:"force"`
}

func (e *Snapshot) Delete(ro *SnapshotDeleteRequest) (*Snapshot, *ApiErrorResponse, error) {
	if ro == nil {
		return nil, nil, badStatus[InvalidRequest]
	}
	v := reflect.ValueOf(*ro)
	t := reflect.TypeOf(*ro)
	gro := &greq.RequestOptions{
		JSON: ro,
	}
	formatQueryParams(gro, v, t)

	rs, apierr, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &Snapshot{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

type SnapshotReloadRequest struct {
	Ctxt context.Context `json:"-"`
}

func (e *Snapshot) Reload(ro *SnapshotReloadRequest) (*Snapshot, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &Snapshot{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}
