package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type StorageNode struct {
	Path                string            `json:"path,omitempty"`
	AdminState          string            `json:"admin_state,omitempty"`
	AvailableCapacity   int               `json:"available_capacity,omitempty"`
	BiosVersion         string            `json:"bios_version,omitempty"`
	BootDrives          []BootDrive       `json:"boot_drives,omitempty"`
	BuildVersion        string            `json:"build_version,omitempty"`
	Causes              []string          `json:"causes,omitempty"`
	Compression         bool              `json:"compression_enabled,omitempty"`
	CompressionRatio    string            `json:"compression_ratio,omitempty"`
	Disconnected        bool              `json:"disconnected,omitempty"`
	FailureDomains      []FailureDomain   `json:"failure_domains,omitempty"`
	FlashDevices        []FlashDevice     `json:"flash_devices,omitempty"`
	Hdds                []Hdd             `json:"hdds,omitempty"`
	Health              string            `json:"health,omitempty"`
	HwHealth            string            `json:"hw_health,omitempty"`
	HwState             string            `json:"hw_state,omitempty"`
	InternalIp1         string            `json:"internal_ip_1,omitempty"`
	InternalIp2         string            `json:"internal_ip_2,omitempty"`
	LastRebootTimestamp string            `json:"last_reboot_timestamp,omitempty"`
	MediaPolicy         string            `json:"media_policy,omitempty"`
	MgmtIp1             string            `json:"mgmt_ip_1,omitempty"`
	MgmtIp2             string            `json:"mgmt_ip_2,omitempty"`
	Model               string            `json:"model,omitempty"`
	Name                string            `json:"name,omitempty"`
	Nics                []Nic             `json:"nics,omitempty"`
	NvmFlashDevices     []NvmFlashDevice  `json:"nvm_flash_devices,omitempty"`
	OpProgress          string            `json:"op_progress,omitempty"`
	OpState             string            `json:"op_state,omitempty"`
	OpStatus            string            `json:"op_status,omitempty"`
	OsVersion           string            `json:"os_version,omitempty"`
	Psus                []Psu             `json:"psus,omitempty"`
	SerialNo            string            `json:"serial_no,omitempty"`
	StorageInstances    []StorageInstance `json:"storage_instances,omitempty"`
	SubsystemHealth     []Subsystem       `json:"subsystem_health,omitempty"`
	SubsystemStates     []Subsystem       `json:"subsystem_states,omitempty"`
	SwHealth            string            `json:"sw_health,omitempty"`
	SwState             string            `json:"sw_state,omitempty"`
	SwVersion           string            `json:"sw_version,omitempty"`
	TotalCapacity       int               `json:"total_capacity,omitempty"`
	TotalRawCapacity    int               `json:"total_raw_capacity,omitempty"`
	Type                string            `json:"type,omitempty"`
	Upgrade             Upgrade           `json:"upgrade,omitempty"`
	Uuid                string            `json:"uuid,omitempty"`
	Vendor              string            `json:"vendor,omitempty"`
	Volumes             []Volume          `json:"volumes,omitempty"`
	BootDrivesEp        *BootDrives
	ctxt                context.Context
	conn                *ApiConnection
}

type StorageNodes struct {
	Path string
	ctxt context.Context
	conn *ApiConnection
}

func newStorageNodes(ctxt context.Context, conn *ApiConnection, path string) *StorageNodes {
	return &StorageNodes{
		Path: _path.Join(path, "storage_nodes"),
		ctxt: ctxt,
		conn: conn,
	}
}

type StorageNodesListRequest struct {
	Params map[string]string
}

type StorageNodesListResponse []StorageNode

func (e *StorageNodes) List(ro *StorageNodesListRequest) (*StorageNodesListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := e.conn.GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := StorageNodesListResponse{}
	for _, data := range rs.Data {
		elem := &StorageNode{}
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

type StorageNodesGetRequest struct {
	Uuid string
}

type StorageNodesGetResponse StorageNode

func (e *StorageNodes) Get(ro *StorageNodesGetRequest) (*StorageNodesGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(_path.Join(e.Path, ro.Uuid), gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageNodesGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type StorageNodeSetRequest struct {
	AdminState  string `json:"admin_state,omitempty"`
	MediaPolicy string `json:"media_policy,omitempty"`
}

type StorageNodeSetResponse StorageNode

func (e *StorageNode) Set(ro *StorageNodeSetRequest) (*StorageNodeSetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &StorageNodeSetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil

}
