package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type StorageTemplate struct {
	Path                 string               `json:"path,omitempty" mapstructure:"path"`
	Auth                 *Auth                `json:"auth,omitempty" mapstructure:"auth"`
	Name                 string               `json:"name,omitempty" mapstructure:"name"`
	IpPool               *AccessNetworkIpPool `json:"ip_pool,omitempty" mapstructure:"ip_pool"`
	ServiceConfiguration string               `json:"service_configuration,omitempty" mapstructure:"service_configuration"`
	VolumeTemplates      []VolumeTemplates    `json:"volume_templates,omitempty" mapstructure:"volume_templates"`
	VolumeTemplatesEp    *VolumeTemplates     `json:"-"`
	ctxt                 context.Context
	conn                 *ApiConnection
}

type StorageTemplates struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

type StorageTemplatesCreateRequest struct {
	Name            string `json:"name,omitempty" mapstructure:"name"`
	ReplicaCount    int    `json:"replica_count,omitempty" mapstructure:"replica_count"`
	Size            int    `json:"size,omitempty" mapstructure:"size"`
	PlacementMode   string `json:"placement_mode,omitempty" mapstructure:"placement_mode"`
	PlacementPolicy string `json:"placement_policy,omitempty" mapstructure:"placement_policy"`
	Force           bool   `json:"force,omitempty" mapstructure:"force"`
}

type StorageTemplatesCreateResponse StorageTemplate

func newStorageTemplates(ctxt context.Context, conn *ApiConnection, path string) *StorageTemplates {
	return &StorageTemplates{
		Path: _path.Join(path, "storage_templates"),
		ctxt: ctxt,
		conn: conn,
	}
}

func (e *StorageTemplates) Create(ro *StorageTemplatesCreateRequest) (*StorageTemplatesCreateResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Post(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageTemplatesCreateResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumeTemplatesEp = newVolumeTemplates(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type StorageTemplatesListRequest struct {
	Params map[string]string
}

type StorageTemplatesListResponse []StorageTemplate

func (e *StorageTemplates) List(ro *StorageTemplatesListRequest) (*StorageTemplatesListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := StorageTemplatesListResponse{}
	for _, data := range rs.Data {
		elem := &StorageTemplate{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, *elem)
	}
	for _, r := range resp {
		r.conn = e.conn
		r.ctxt = e.ctxt
		r.VolumeTemplatesEp = newVolumeTemplates(e.ctxt, e.conn, e.Path)
	}
	return &resp, nil
}

type StorageTemplatesGetRequest struct {
	Name string
}

type StorageTemplatesGetResponse StorageTemplate

func (e *StorageTemplates) Get(ro *StorageTemplatesGetRequest) (*StorageTemplatesGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Name), gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageTemplatesGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumeTemplatesEp = newVolumeTemplates(e.ctxt, e.conn, e.Path)
	return resp, nil
}

type StorageTemplateSetRequest struct {
	Auth            Auth                `json:"auth,omitempty" mapstructure:"auth"`
	IpPool          AccessNetworkIpPool `json:"ip_pool,omitempty" mapstructure:"ip_pool"`
	VolumeTemplates []VolumeTemplates   `json:"volume_templates,omitempty" mapstructure:"volume_templates"`
}

type StorageTemplateSetResponse StorageTemplate

func (e *StorageTemplate) Set(ro *StorageTemplateSetRequest) (*StorageTemplateSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageTemplateSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumeTemplatesEp = newVolumeTemplates(e.ctxt, e.conn, e.Path)
	return resp, nil

}

type StorageTemplateDeleteRequest struct {
	Force bool `json:"force,omitempty" mapstructure:"force"`
}

type StorageTemplateDeleteResponse StorageTemplate

func (e *StorageTemplate) Delete(ro *StorageTemplateDeleteRequest) (*StorageTemplateDeleteResponse, error) {
	rs, err := e.conn.Delete(e.Path, nil)
	if err != nil {
		return nil, err
	}
	resp := &StorageTemplateDeleteResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	resp.VolumeTemplatesEp = newVolumeTemplates(e.ctxt, e.conn, e.Path)
	return resp, nil
}
