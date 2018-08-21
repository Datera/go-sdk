package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type RemoteProvider string

type Snapshot struct {
	Path            string           `json:"path,omitempty"`
	Timestamp       string           `json:"timestamp,omitempty"`
	Uuid            string           `json:"uuid,omitempty"`
	RemoteProviders []RemoteProvider `json:"remote_providers,omitempty"`
	OpState         string           `json:"op_state,omitempty"`
	UtcTs           string           `json:"utc_ts,omitempty"`
	PhysicalSize    int              `json:"physical_size,omitempty"`
	LogicalSize     int              `json:"logical_size,omitempty"`
	ExclusiveSize   int              `json:"exclusive_size,omitempty"`
	EffectiveSize   int              `json:"effective_size,omitempty"`
	Local           bool             `json:"local,omitempty"`
	AppStructure    string           `json:"app_structure,omitempty"`
	ctxt            context.Context
	conn            *ApiConnection
}

type Snapshots struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type SnapshotsCreateRequest struct {
	Uuid               string `json:"uuid,omitempty"`
	RemoteProviderUuid string `json:"remote_provider_uuid,omitempty"`
	Type               string `json:"type,omitempty"`
}

type SnapshotsCreateResponse Snapshot

func newSnapshots(ctxt context.Context, conn *ApiConnection, path string) *Snapshots {
	return &Snapshots{
		Path: _path.Join(path, "snapshots"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *Snapshots) Create(ro *SnapshotsCreateRequest) (*SnapshotsCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotsCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type SnapshotsListRequest struct {
	Params map[string]string
}

type SnapshotsListResponse []Snapshot

func (e *Snapshots) List(ro *SnapshotsListRequest) (*SnapshotsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := SnapshotsListResponse{}
	for _, data := range rs.Data {
		elem := &Snapshot{}
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

type SnapshotsGetRequest struct {
	Timestamp string
}

type SnapshotsGetResponse Snapshot

func (e *Snapshots) Get(ro *SnapshotsGetRequest) (*SnapshotsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Timestamp), gro)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type SnapshotDeleteRequest struct {
	RemoteProviderUuid string `json:"remote_provider_uuid,omitempty"`
}

type SnapshotDeleteResponse Snapshot

func (e *Snapshot) Delete(ro *SnapshotDeleteRequest) (*SnapshotDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &SnapshotDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}
