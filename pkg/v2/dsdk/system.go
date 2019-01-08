package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type System struct {
	Path                        string           `json:"path,omitempty" mapstructure:"path"`
	AccessInterfaceAggrType     string           `json:"access_interface_aggr_type,omitempty" mapstructure:"access_interface_aggr_type"`
	AllFlashCapacity            int              `json:"all_flash_available_capacity,omitempty" mapstructure:"all_flash_available_capacity"`
	AllFlashProvisionedCapacity int              `json:"all_flash_provisioned_capacity,omitempty" mapstructure:"all_flash_provisioned_capacity"`
	AllFlashTotalCapacity       int              `json:"all_flash_total_capacity,omitempty" mapstructure:"all_flash_total_capacity"`
	AvailableCapacity           int              `json:"available_capacity,omitempty" mapstructure:"available_capacity"`
	BuildVersion                string           `json:"build_version,omitempty" mapstructure:"build_version"`
	CallhomeEnabled             bool             `json:"callhome_enabled,omitempty" mapstructure:"callhome_enabled"`
	Causes                      []string         `json:"causes,omitempty" mapstructure:"causes"`
	CompressionEnabled          bool             `json:"compression_enabled,omitempty" mapstructure:"compression_enabled"`
	CompressionRatio            string           `json:"compression_ratio,omitempty" mapstructure:"compression_ratio"`
	Dns                         *Dns             `json:"dns,omitempty" mapstructure:"dns"`
	Health                      string           `json:"health,omitempty" mapstructure:"health"`
	HttpProxy                   *HttpProxy       `json:"http_proxy,omitempty" mapstructure:"http_proxy"`
	HybridAvailableCapacity     int              `json:"hybrid_available_capacity,omitempty" mapstructure:"hybrid_available_capacity"`
	HybridProvisionedCapacity   int              `json:"hybrid_provisioned_capacity,omitempty" mapstructure:"hybrid_provisioned_capacity"`
	HybridTotalCapacity         int              `json:"hybrid_total_capacity,omitempty" mapstructure:"hybrid_total_capacity"`
	InterfaceAggregationMode    string           `json:"interface_aggregation_mode,omitempty" mapstructure:"interface_aggregation_mode"`
	InternalInterfaceAggrType   string           `json:"internal_interface_aggr_type,omitempty" mapstructure:"internal_interface_aggr_type"`
	L3Enabled                   bool             `json:"l3_enabled,omitempty" mapstructure:"l3_enabled"`
	LastRebootTimestamp         string           `json:"last_reboot_timestamp,omitempty" mapstructure:"last_reboot_timestamp"`
	Name                        string           `json:"name,omitempty" mapstructure:"name"`
	Network                     *Network         `json:"network,omitempty" mapstructure:"network"`
	NetworkDevices              []*NetworkDevice `json:"network_devices,omitempty" mapstructure:"network_devices"`
	NtpServers                  []string         `json:"ntp_servers,omitempty" mapstructure:"ntp_servers"`
	OpState                     string           `json:"op_state,omitempty" mapstructure:"op_state"`
	SwVersion                   string           `json:"sw_version,omitempty" mapstructure:"sw_version"`
	Timezone                    string           `json:"timezone,omitempty" mapstructure:"timezone"`
	TotalCapacity               int              `json:"total_capacity,omitempty" mapstructure:"total_capacity"`
	TotalProvisionedCapacity    int              `json:"total_provisioned_capacity,omitempty" mapstructure:"total_provisioned_capacity"`
	Upgrade                     *Upgrade         `json:"upgrade,omitempty" mapstructure:"upgrade"`
	Uptime                      int              `json:"uptime,omitempty" mapstructure:"uptime"`
	Uuid                        string           `json:"uuid,omitempty" mapstructure:"uuid"`
	WitnessPolicy               *WitnessPolicy   `json:"witness_policy,omitempty" mapstructure:"witness_policy"`
}

func RegisterSystemEndpoints(a *System) {
}

func newSystem(path string) *System {
	return &System{
		Path: _path.Join(path, "system"),
	}
}

type SystemGetRequest struct {
	Ctxt context.Context `json:"-"`
}

func (e *System) Get(ro *SystemGetRequest) (*System, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &System{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterSystemEndpoints(resp)
	return resp, nil, nil
}

type SystemSetRequest struct {
	Ctxt                             context.Context  `json:"-"`
	AccessInterfaceAggrType          string           `json:"access_interface_aggr_type,omitempty" mapstructure:"access_interface_aggr_type"`
	CallhomeEnabled                  bool             `json:"callhome_enabled,omitempty" mapstructure:"callhome_enabled"`
	CompressionEnabled               bool             `json:"compression_enabled,omitempty" mapstructure:"compression_enabled"`
	InterfaceAggregationMode         string           `json:"interface_aggregation_mode,omitempty" mapstructure:"interface_aggregation_mode"`
	InternalInterfaceAggregationType string           `json:"internal_interface_aggr_type,omitempty" mapstructure:"internal_interface_aggr_type"`
	NetworkDevices                   []*NetworkDevice `json:"network_devices,omitempty" mapstructure:"network_devices"`
}

func (e *System) Set(ro *SystemSetRequest) (*System, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &System{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterSystemEndpoints(resp)
	return resp, nil, nil

}
