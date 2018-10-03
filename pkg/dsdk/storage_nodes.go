package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type StorageNode struct {
	Path                string             `json:"path,omitempty" mapstructure:"path"`
	AdminState          string             `json:"admin_state,omitempty" mapstructure:"admin_state"`
	AvailableCapacity   int                `json:"available_capacity,omitempty" mapstructure:"available_capacity"`
	BiosVersion         string             `json:"bios_version,omitempty" mapstructure:"bios_version"`
	BootDrives          []*BootDrive       `json:"boot_drives,omitempty" mapstructure:"boot_drives"`
	BuildVersion        string             `json:"build_version,omitempty" mapstructure:"build_version"`
	Causes              []string           `json:"causes,omitempty" mapstructure:"causes"`
	Compression         bool               `json:"compression_enabled,omitempty" mapstructure:"compression_enabled"`
	CompressionRatio    string             `json:"compression_ratio,omitempty" mapstructure:"compression_ratio"`
	Disconnected        bool               `json:"disconnected,omitempty" mapstructure:"disconnected"`
	FailureDomains      []*FailureDomain   `json:"failure_domains,omitempty" mapstructure:"failure_domains"`
	FlashDevices        []*FlashDevice     `json:"flash_devices,omitempty" mapstructure:"flash_devices"`
	Hdds                []*Hdd             `json:"hdds,omitempty" mapstructure:"hdds"`
	Health              string             `json:"health,omitempty" mapstructure:"health"`
	HwHealth            string             `json:"hw_health,omitempty" mapstructure:"hw_health"`
	HwState             string             `json:"hw_state,omitempty" mapstructure:"hw_state"`
	InternalIp1         string             `json:"internal_ip_1,omitempty" mapstructure:"internal_ip_1"`
	InternalIp2         string             `json:"internal_ip_2,omitempty" mapstructure:"internal_ip_2"`
	LastRebootTimestamp string             `json:"last_reboot_timestamp,omitempty" mapstructure:"last_reboot_timestamp"`
	MediaPolicy         string             `json:"media_policy,omitempty" mapstructure:"media_policy"`
	MgmtIp1             string             `json:"mgmt_ip_1,omitempty" mapstructure:"mgmt_ip_1"`
	MgmtIp2             string             `json:"mgmt_ip_2,omitempty" mapstructure:"mgmt_ip_2"`
	Model               string             `json:"model,omitempty" mapstructure:"model"`
	Name                string             `json:"name,omitempty" mapstructure:"name"`
	Nics                []*Nic             `json:"nics,omitempty" mapstructure:"nics"`
	NvmFlashDevices     []*NvmFlashDevice  `json:"nvm_flash_devices,omitempty" mapstructure:"nvm_flash_devices"`
	OpProgress          string             `json:"op_progress,omitempty" mapstructure:"op_progress"`
	OpState             string             `json:"op_state,omitempty" mapstructure:"op_state"`
	OpStatus            string             `json:"op_status,omitempty" mapstructure:"op_status"`
	OsVersion           string             `json:"os_version,omitempty" mapstructure:"os_version"`
	Psus                []*Psu             `json:"psus,omitempty" mapstructure:"psus"`
	SerialNo            string             `json:"serial_no,omitempty" mapstructure:"serial_no"`
	StorageInstances    []*StorageInstance `json:"storage_instances,omitempty" mapstructure:"storage_instances"`
	SubsystemHealth     []*Subsystem       `json:"subsystem_health,omitempty" mapstructure:"subsystem_health"`
	SubsystemStates     *Subsystem         `json:"subsystem_states,omitempty" mapstructure:"subsystem_states"`
	SwHealth            string             `json:"sw_health,omitempty" mapstructure:"sw_health"`
	SwState             string             `json:"sw_state,omitempty" mapstructure:"sw_state"`
	SwVersion           string             `json:"sw_version,omitempty" mapstructure:"sw_version"`
	TotalCapacity       int                `json:"total_capacity,omitempty" mapstructure:"total_capacity"`
	TotalRawCapacity    int                `json:"total_raw_capacity,omitempty" mapstructure:"total_raw_capacity"`
	Type                string             `json:"type,omitempty" mapstructure:"type"`
	Upgrade             *Upgrade           `json:"upgrade,omitempty" mapstructure:"upgrade"`
	Uuid                string             `json:"uuid,omitempty" mapstructure:"uuid"`
	Vendor              string             `json:"vendor,omitempty" mapstructure:"vendor"`
	Volumes             []*Volume          `json:"volumes,omitempty" mapstructure:"volumes"`
	BootDrivesEp        *BootDrives
}

func RegisterStorageNodeEndpoints(a *StorageNode) {
	a.BootDrivesEp = newBootDrives(a.Path)
	for _, si := range a.StorageInstances {
		RegisterStorageInstanceEndpoints(si)
	}
	for _, vol := range a.Volumes {
		RegisterVolumeEndpoints(vol)
	}
}

type StorageNodes struct {
	Path string
}

func newStorageNodes(path string) *StorageNodes {
	return &StorageNodes{
		Path: _path.Join(path, "storage_nodes"),
	}
}

type StorageNodesListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListParams      `json:"params,omitempty"`
}

func (e *StorageNodes) List(ro *StorageNodesListRequest) ([]*StorageNode, *ApiErrorResponse, error) {
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
	resp := []*StorageNode{}
	for _, data := range rs.Data {
		elem := &StorageNode{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
		RegisterStorageNodeEndpoints(elem)
	}
	return resp, nil, nil
}

type StorageNodesGetRequest struct {
	Ctxt context.Context `json:"-"`
	Uuid string          `json:"-"`
}

func (e *StorageNodes) Get(ro *StorageNodesGetRequest) (*StorageNode, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Uuid), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &StorageNode{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterStorageNodeEndpoints(resp)
	return resp, nil, nil
}

type StorageNodeSetRequest struct {
	Ctxt        context.Context `json:"-"`
	AdminState  string          `json:"admin_state,omitempty" mapstructure:"admin_state"`
	MediaPolicy string          `json:"media_policy,omitempty" mapstructure:"media_policy"`
}

func (e *StorageNode) Set(ro *StorageNodeSetRequest) (*StorageNode, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &StorageNode{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterStorageNodeEndpoints(resp)
	return resp, nil, nil

}
