package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type RemoteProvider string

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
	AppStructure    string            `json:"app_structure,omitempty" mapstructure:"app_structure"`
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

func (e *Snapshots) Create(ro *SnapshotsCreateRequest) (*Snapshot, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &Snapshot{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type SnapshotsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

func (e *Snapshots) List(ro *SnapshotsListRequest) ([]*Snapshot, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := []*Snapshot{}
	for _, data := range rs.Data {
		elem := &Snapshot{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil
}

type SnapshotsGetRequest struct {
	Ctxt      context.Context `json:"-"`
	Timestamp string
}

func (e *Snapshots) Get(ro *SnapshotsGetRequest) (*Snapshot, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Timestamp), gro)
	if err != nil {
		return nil, err
	}
	resp := &Snapshot{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type SnapshotDeleteRequest struct {
	Ctxt               context.Context `json:"-"`
	RemoteProviderUuid string          `json:"remote_provider_uuid,omitempty" mapstructure:"remote_provider_uuid"`
}

func (e *Snapshot) Delete(ro *SnapshotDeleteRequest) (*Snapshot, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &Snapshot{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
