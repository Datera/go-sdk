package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type VolumeTemplate struct {
	Path               string        `json:"path,omitempty"`
	Name               string        `json:"name,omitempty"`
	PlacementMode      string        `json:"placement_mode,omitempty"`
	PlacementPolicy    string        `json:"placement_policy,omitempty"`
	ReplicaCount       int           `json:"replica_count,omitempty"`
	Size               int           `json:"size,omitempty"`
	StoragePool        []StoragePool `json:"storage_pool,omitempty"`
	SnapshotPoliciesEp *SnapshotPolicies
	ctxt               context.Context
	conn               *ApiConnection
}

type VolumeTemplates struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type VolumeTemplatesCreateRequest struct {
	Name            string `json:"name,omitempty"`
	ReplicaCount    int    `json:"replica_count,omitempty"`
	Size            int    `json:"size,omitempty"`
	PlacementMode   string `json:"placement_mode,omitempty"`
	PlacementPolicy string `json:"placement_policy,omitempty"`
}

type VolumeTemplatesCreateResponse VolumeTemplate

func newVolumeTemplates(ctxt context.Context, conn *ApiConnection, path string) *VolumeTemplates {
	return &VolumeTemplates{
		Path: _path.Join(path, "volume_templates"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *VolumeTemplates) Create(ro *VolumeTemplatesCreateRequest) (*VolumeTemplatesCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &VolumeTemplatesCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.SnapshotPoliciesEp = newSnapshotPolicies(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type VolumeTemplatesListRequest struct {
	Params map[string]string
}

type VolumeTemplatesListResponse []VolumeTemplate

func (e *VolumeTemplates) List(ro *VolumeTemplatesListRequest) (*VolumeTemplatesListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := VolumeTemplatesListResponse{}
	for _, data := range rs.Data {
		elem := &VolumeTemplate{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, *elem)
	}
	for _, r := range resp {
		r.conn = e.conn
		r.ctxt = e.ctxt
		r.SnapshotPoliciesEp = newSnapshotPolicies(e.ctxt, e.conn, e.Path)
	}
	return &resp, nil
}

type VolumeTemplatesGetRequest struct {
	Name string
}

type VolumeTemplatesGetResponse VolumeTemplate

func (e *VolumeTemplates) Get(ro *VolumeTemplatesGetRequest) (*VolumeTemplatesGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Name), gro)
	if err != nil {
		return nil, err
	}
	resp := &VolumeTemplatesGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.SnapshotPoliciesEp = newSnapshotPolicies(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type VolumeTemplateSetRequest struct {
	PlacementMode   string        `json:"placement_mode,omitempty"`
	PlacementPolicy string        `json:"placement_policy,omitempty"`
	ReplicaCount    int           `json:"replica_count,omitempty"`
	Size            int           `json:"size,omitempty"`
	StoragePool     []StoragePool `json:"storage_pool,omitempty"`
}

type VolumeTemplateSetResponse VolumeTemplate

func (e *VolumeTemplate) Set(ro *VolumeTemplateSetRequest) (*VolumeTemplateSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &VolumeTemplateSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.SnapshotPoliciesEp = newSnapshotPolicies(e.ctxt, e.conn, e.Path)
	return resp, nil

}

type VolumeTemplateDeleteRequest struct {
}

type VolumeTemplateDeleteResponse VolumeTemplate

func (e *VolumeTemplate) Delete(ro *VolumeTemplateDeleteRequest) (*VolumeTemplateDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &VolumeTemplateDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.SnapshotPoliciesEp = newSnapshotPolicies(e.ctxt, e.conn, e.Path)
	return resp, nil
}
