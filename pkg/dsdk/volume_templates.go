package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type VolumeTemplate struct {
	Path               string            `json:"path,omitempty" mapstructure:"path"`
	Name               string            `json:"name,omitempty" mapstructure:"name"`
	PlacementMode      string            `json:"placement_mode,omitempty" mapstructure:"placement_mode"`
	PlacementPolicy    string            `json:"placement_policy,omitempty" mapstructure:"placement_policy"`
	ReplicaCount       int               `json:"replica_count,omitempty" mapstructure:"replica_count"`
	Size               int               `json:"size,omitempty" mapstructure:"size"`
	StoragePool        []StoragePool     `json:"storage_pool,omitempty" mapstructure:"storage_pool"`
	SnapshotPoliciesEp *SnapshotPolicies `json:"-"`
}

type VolumeTemplates struct {
	Path string
}

type VolumeTemplatesCreateRequest struct {
	Ctxt            context.Context `json:"-"`
	Name            string          `json:"name,omitempty" mapstructure:"name"`
	ReplicaCount    int             `json:"replica_count,omitempty" mapstructure:"replica_count"`
	Size            int             `json:"size,omitempty" mapstructure:"size"`
	PlacementMode   string          `json:"placement_mode,omitempty" mapstructure:"placement_mode"`
	PlacementPolicy string          `json:"placement_policy,omitempty" mapstructure:"placement_policy"`
}

func newVolumeTemplates(path string) *VolumeTemplates {
	return &VolumeTemplates{
		Path: _path.Join(path, "volume_templates"),
	}
}

func (e *VolumeTemplates) Create(ro *VolumeTemplatesCreateRequest) (*VolumeTemplate, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &VolumeTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type VolumeTemplatesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

func (e *VolumeTemplates) List(ro *VolumeTemplatesListRequest) ([]*VolumeTemplate, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := []*VolumeTemplate{}
	for _, data := range rs.Data {
		elem := &VolumeTemplate{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil
}

type VolumeTemplatesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Name string
}

func (e *VolumeTemplates) Get(ro *VolumeTemplatesGetRequest) (*VolumeTemplate, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Name), gro)
	if err != nil {
		return nil, err
	}
	resp := &VolumeTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.SnapshotPoliciesEp = newSnapshotPolicies(e.Path)
	return resp, nil
}

type VolumeTemplateSetRequest struct {
	Ctxt            context.Context `json:"-"`
	PlacementMode   string          `json:"placement_mode,omitempty" mapstructure:"placement_mode"`
	PlacementPolicy string          `json:"placement_policy,omitempty" mapstructure:"placement_policy"`
	ReplicaCount    int             `json:"replica_count,omitempty" mapstructure:"replica_count"`
	Size            int             `json:"size,omitempty" mapstructure:"size"`
	StoragePool     []StoragePool   `json:"storage_pool,omitempty" mapstructure:"storage_pool"`
}

func (e *VolumeTemplate) Set(ro *VolumeTemplateSetRequest) (*VolumeTemplate, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &VolumeTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.SnapshotPoliciesEp = newSnapshotPolicies(e.Path)
	return resp, nil

}

type VolumeTemplateDeleteRequest struct {
	Ctxt context.Context `json:"-"`
}

func (e *VolumeTemplate) Delete(ro *VolumeTemplateDeleteRequest) (*VolumeTemplate, error) {
	rs, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &VolumeTemplate{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.SnapshotPoliciesEp = newSnapshotPolicies(e.Path)
	return resp, nil
}
