package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type Volume struct {
	Path               string        `json:"path,omitempty"`
	ActiveStorageNodes []StorageNode `json:"active_storage_nodes,omitempty"`
	AvailabilityState  string        `json:"availability_state,omitempty"`
	CapacityInUse      int           `json:"capacity_in_use,omitempty"`
	Causes             []string      `json:"causes,omitempty"`
	DeploymentState    string        `json:"deployment_state,omitempty"`
	EffectiveSize      int           `json:"effective_size,omitempty"`
	ExclusiveSize      int           `json:"exclusive_size,omitempty"`
	Health             string        `json:"health,omitempty"`
	LogicalSize        int           `json:"logical_size,omitempty"`
	Name               string        `json:"name,omitempty"`
	OpState            string        `json:"op_state,omitempty"`
	OpStatus           string        `json:"op_status,omitempty"`
	PhysicalSize       int           `json:"physical_size,omitempty"`
	PlacementMode      string        `json:"placement_mode,omitempty"`
	PlacementPolicy    string        `json:"placement_policy,omitempty"`
	RecoveryState      string        `json:"recovery_state,omitempty"`
	ReplicaCount       int           `json:"replica_count,omitempty"`
	RestorePoint       string        `json:"restore_point,omitempty"`
	Size               int           `json:"size,omitempty"`
	Snapshots          []Snapshot    `json:"snapshots,omitempty"`
	StoragePool        []StoragePool `json:"storage_pool,omitempty"`
	StorageState       string        `json:"storage_state,omitempty"`
	Uuid               string        `json:"uuid,omitempty"`
	SnapshotsEp        *Snapshots
	ctxt               context.Context
	conn               *ApiConnection
}

type Volumes struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type VolumesCreateRequest struct {
	Name            string `json:"name,omitempty"`
	ReplicaCount    int    `json:"replica_count,omitempty"`
	Size            int    `json:"size,omitempty"`
	PlacementMode   string `json:"placement_mode,omitempty"`
	PlacementPolicy string `json:"placement_policy,omitempty"`
	Force           bool   `json:"force,omitempty"`
}

type VolumesCreateResponse Volume

func newVolumes(ctxt context.Context, conn *ApiConnection, path string) *Volumes {
	return &Volumes{
		Path: _path.Join(path, "volumes"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *Volumes) Create(ro *VolumesCreateRequest) (*VolumesCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &VolumesCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.SnapshotsEp = newSnapshots(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type VolumesListRequest struct {
	Params map[string]string
}

type VolumesListResponse []Volume

func (e *Volumes) List(ro *VolumesListRequest) (*VolumesListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := VolumesListResponse{}
	for _, data := range rs.Data {
		elem := &Volume{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, *elem)
	}
	for _, r := range resp {
		r.conn = e.conn
		r.ctxt = e.ctxt
		r.SnapshotsEp = newSnapshots(e.ctxt, e.conn, e.Path)
	}
	return &resp, nil
}

type VolumesGetRequest struct {
	Id string
}

type VolumesGetResponse Volume

func (e *Volumes) Get(ro *VolumesGetRequest) (*VolumesGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &VolumesGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.SnapshotsEp = newSnapshots(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type VolumeSetRequest struct {
	Name string `json:"name,omitempty"`
}

type VolumeSetResponse Volume

func (e *Volume) Set(ro *VolumeSetRequest) (*VolumeSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &VolumeSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.SnapshotsEp = newSnapshots(e.ctxt, e.conn, e.Path)
	return resp, nil

}

type VolumeDeleteRequest struct {
}

type VolumeDeleteResponse Volume

func (e *Volume) Delete(ro *VolumeDeleteRequest) (*VolumeDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &VolumeDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.SnapshotsEp = newSnapshots(e.ctxt, e.conn, e.Path)
	return resp, nil
}
