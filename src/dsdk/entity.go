package dsdk

type AccessNetworkIpPool struct {
	Descr        string      `json:"descr,omitempty"`
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

type AclPolicy struct {
	InitiatorGroups *[]InitiatorGroup `json:"initiator_groups,omitempty"`
	Initiators      *[]Initiator      `json:"initiators,omitempty"`
	Path            string            `json:"path,omitempty"`
}

type AppInstance struct {
	AccessControlMode string             `json:"access_control_mode,omitempty"`
	AdminState        string             `json:"admin_state,omitempty"`
	AppTemplate       *AppTemplate       `json:"app_template,omitempty"`
	Causes            []interface{}      `json:"causes,omitempty"`
	CloneSrc          map[string]string  `json:"clone_src,omitempty"`
	CreateMode        string             `json:"create_mode,omitempty"`
	Descr             string             `json:"descr,omitempty"`
	Health            string             `json:"health,omitempty"`
	Id                string             `json:"id,omitempty"`
	Name              string             `json:"name,omitempty"`
	OpStatus          string             `json:"op_status,omitempty"`
	Path              string             `json:"path,omitempty"`
	RestorePoint      string             `json:"restore_point,omitempty"`
	SnapshotPolicies  *[]SnapshotPolicy  `json:"snapshot_policies,omitempty"`
	Snapshots         *[]Snapshot        `json:"snapshots,omitempty"`
	StorageInstances  *[]StorageInstance `json:"storage_instances,omitempty"`
	Uuid              string             `json:"uuid,omitempty"`
}

type AppTemplate struct {
	AppInstances     *[]AppInstance     `json:"app_instances,omitempty"`
	Descr            string             `json:"descr,omitempty"`
	Name             string             `json:"name,omitempty"`
	Path             string             `json:"path,omitempty"`
	SnapshotPolicies *[]SnapshotPolicy  `json:"snapshot_policies,omitempty"`
	StorageTemplates *[]StorageTemplate `json:"storage_templates,omitempty"`
}

type AuditLog struct {
	Description string `json:"description,omitempty"`
	Id          string `json:"id,omitempty"`
	ObjectName  string `json:"object_name,omitempty"`
	ObjectType  string `json:"object_type,omitempty"`
	ObjectUrl   string `json:"object_url,omitempty"`
	Operation   string `json:"operation,omitempty"`
	ParamInfo   string `json:"param_info,omitempty"`
	Path        string `json:"path,omitempty"`
	SessionInfo string `json:"session_info,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	User        *User  `json:"user,omitempty"`
	Version     string `json:"version,omitempty"`
}

type Auth struct {
	InitiatorPswd     string `json:"initiator_pswd,omitempty"`
	InitiatorUserName string `json:"initiator_user_name,omitempty"`
	Path              string `json:"path,omitempty"`
	TargetPswd        string `json:"target_pswd,omitempty"`
	TargetUserName    string `json:"target_user_name,omitempty"`
	Type              string `json:"type,omitempty"`
}

type BootDrive struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

type Dns struct {
	Domain string `json:"domain,omitempty"`
	Path   string `json:"path,omitempty"`
}

type DnsSearchDomain struct {
	Domain string `json:"domain,omitempty"`
	Order  int    `json:"order,omitempty"`
	Path   string `json:"path,omitempty"`
}

type DnsServer struct {
	Ip    string `json:"ip,omitempty"`
	Order int    `json:"order,omitempty"`
	Path  string `json:"path,omitempty"`
}

type EventLog struct {
	Cause        string `json:"cause,omitempty"`
	Code         string `json:"code,omitempty"`
	Id           string `json:"id,omitempty"`
	ObjectId     string `json:"object_id,omitempty"`
	ObjectLbl    string `json:"object_lbl,omitempty"`
	ObjectPath   string `json:"object_path,omitempty"`
	ObjectTenant string `json:"object_tenant,omitempty"`
	ObjectType   string `json:"object_type,omitempty"`
	ObjectUrl    string `json:"object_url,omitempty"`
	Path         string `json:"path,omitempty"`
	Severity     string `json:"severity,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
	Type         string `json:"type,omitempty"`
}

type FaultLog struct {
	Acknowledged    bool   `json:"acknowledged,omitempty"`
	CallhomeEnabled bool   `json:"callhome_enabled,omitempty"`
	Cause           string `json:"cause,omitempty"`
	Cleared         bool   `json:"cleared,omitempty"`
	Code            string `json:"code,omitempty"`
	Count           int    `json:"count,omitempty"`
	Id              string `json:"id,omitempty"`
	ObjectId        string `json:"object_id,omitempty"`
	ObjectLbl       string `json:"object_lbl,omitempty"`
	ObjectPath      string `json:"object_path,omitempty"`
	ObjectTenant    string `json:"object_tenant,omitempty"`
	ObjectType      string `json:"object_type,omitempty"`
	ObjectUrl       string `json:"object_url,omitempty"`
	Path            string `json:"path,omitempty"`
	Repeat          string `json:"repeat,omitempty"`
	Severity        string `json:"severity,omitempty"`
	Timestamp       int    `json:"timestamp,omitempty"`
	Type            string `json:"type,omitempty"`
}

type Hdd struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

type HttpProxy struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Host     string `json:"host,omitempty"`
	Password string `json:"password,omitempty"`
	Path     string `json:"path,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     *User  `json:"user,omitempty"`
}

type Initiator struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

type InitiatorGroup struct {
	Members []interface{} `json:"members,omitempty"`
	Name    string        `json:"name,omitempty"`
	Path    string        `json:"path,omitempty"`
}

type InternalIpBlock struct {
	Gateway string `json:"gateway,omitempty"`
	Mtu     int    `json:"mtu,omitempty"`
	Name    string `json:"name,omitempty"`
	Netmask int    `json:"netmask,omitempty"`
	Path    string `json:"path,omitempty"`
	Range   int    `json:"range,omitempty"`
	StartIp string `json:"start_ip,omitempty"`
	Vlan    int    `json:"vlan,omitempty"`
}

type IpAddress struct {
	Gateway string `json:"gateway,omitempty"`
	Ip      string `json:"ip,omitempty"`
	Mtu     int    `json:"mtu,omitempty"`
	Name    string `json:"name,omitempty"`
	Netmask int    `json:"netmask,omitempty"`
	Path    string `json:"path,omitempty"`
	Vlan    int    `json:"vlan,omitempty"`
}

type IpBlock struct {
	Gateway string `json:"gateway,omitempty"`
	Mtu     int    `json:"mtu,omitempty"`
	Name    string `json:"name,omitempty"`
	Netmask int    `json:"netmask,omitempty"`
	Path    string `json:"path,omitempty"`
	Range   int    `json:"range,omitempty"`
	StartIp string `json:"start_ip,omitempty"`
	Vlan    int    `json:"vlan,omitempty"`
}

type IpPool struct {
	Descr        string      `json:"descr,omitempty"`
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

type MonitoringDestination struct {
	Facility  string `json:"facility,omitempty"`
	Host      string `json:"host,omitempty"`
	LastMsgTs string `json:"last_msg_ts,omitempty"`
	Name      string `json:"name,omitempty"`
	OpState   string `json:"op_state,omitempty"`
	Path      string `json:"path,omitempty"`
	Port      int    `json:"port,omitempty"`
	Type      string `json:"type,omitempty"`
}

type MonitoringPolicy struct {
	Destinations []interface{} `json:"destinations,omitempty"`
	Enabled      bool          `json:"enabled,omitempty"`
	Name         string        `json:"name,omitempty"`
	Path         string        `json:"path,omitempty"`
}

type Network struct {
	AccessNetworks  []interface{} `json:"access_networks,omitempty"`
	AccessVip       interface{}   `json:"access_vip,omitempty"`
	InternalNetwork interface{}   `json:"internal_network,omitempty"`
	Mapping         interface{}   `json:"mapping,omitempty"`
	MgmtVip         interface{}   `json:"mgmt_vip,omitempty"`
	Path            string        `json:"path,omitempty"`
}

type NetworkMapping struct {
	Access1   string `json:"access_1,omitempty"`
	Access2   string `json:"access_2,omitempty"`
	Internal1 string `json:"internal_1,omitempty"`
	Internal2 string `json:"internal_2,omitempty"`
	Mgmt1     string `json:"mgmt_1,omitempty"`
	Mgmt2     string `json:"mgmt_2,omitempty"`
	Path      string `json:"path,omitempty"`
}

type Nic struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

type NtpServer struct {
	Ip    string `json:"ip,omitempty"`
	Order int    `json:"order,omitempty"`
	Path  string `json:"path,omitempty"`
}

type NvmFlashDevice struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

type PerformancePolicy struct {
	Path              string `json:"path,omitempty"`
	ReadBandwidthMax  int    `json:"read_bandwidth_max,omitempty"`
	ReadIopsMax       int    `json:"read_iops_max,omitempty"`
	TotalBandwidthMax int    `json:"total_bandwidth_max,omitempty"`
	TotalIopsMax      int    `json:"total_iops_max,omitempty"`
	WriteBandwidthMax int    `json:"write_bandwidth_max,omitempty"`
	WriteIopsMax      int    `json:"write_iops_max,omitempty"`
}

type Role struct {
	Path       string        `json:"path,omitempty"`
	Privileges []interface{} `json:"privileges,omitempty"`
	RoleId     string        `json:"role_id,omitempty"`
}

type Snapshot struct {
	OpState   string `json:"op_state,omitempty"`
	Path      string `json:"path,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	UtcTs     int    `json:"utc_ts,omitempty"`
	Uuid      string `json:"uuid,omitempty"`
}

type SnapshotPolicy struct {
	Interval       string `json:"interval,omitempty"`
	Name           string `json:"name,omitempty"`
	Path           string `json:"path,omitempty"`
	RetentionCount int    `json:"retention_count,omitempty"`
	StartTime      string `json:"start_time,omitempty"`
}

type SnmpPolicy struct {
	Contact  string  `json:"contact,omitempty"`
	Enabled  bool    `json:"enabled,omitempty"`
	Location string  `json:"location,omitempty"`
	Path     string  `json:"path,omitempty"`
	Users    *[]User `json:"users,omitempty"`
}

type SnmpUser struct {
	AuthPass      string `json:"auth_pass,omitempty"`
	AuthProtocol  string `json:"auth_protocol,omitempty"`
	EncrPass      string `json:"encr_pass,omitempty"`
	EncrProtocol  string `json:"encr_protocol,omitempty"`
	Path          string `json:"path,omitempty"`
	SecurityLevel string `json:"security_level,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	Version       string `json:"version,omitempty"`
	View          string `json:"view,omitempty"`
}

type StorageInstance struct {
	Access             interface{}   `json:"access,omitempty"`
	AccessControlMode  string        `json:"access_control_mode,omitempty"`
	AclPolicy          *AclPolicy    `json:"acl_policy,omitempty"`
	ActiveInitiators   []interface{} `json:"active_initiators,omitempty"`
	ActiveStorageNodes []interface{} `json:"active_storage_nodes,omitempty"`
	AdminState         string        `json:"admin_state,omitempty"`
	Auth               *Auth         `json:"auth,omitempty"`
	Causes             []interface{} `json:"causes,omitempty"`
	Health             string        `json:"health,omitempty"`
	IpPool             *IpPool       `json:"ip_pool,omitempty"`
	Name               string        `json:"name,omitempty"`
	OpState            string        `json:"op_state,omitempty"`
	Path               string        `json:"path,omitempty"`
	Uuid               string        `json:"uuid,omitempty"`
	Volumes            *[]Volume     `json:"volumes,omitempty"`
}

type StorageNode struct {
	AdminState          string             `json:"admin_state,omitempty"`
	AvailableCapacity   int                `json:"available_capacity,omitempty"`
	BiosVersion         string             `json:"bios_version,omitempty"`
	BootDrives          interface{}        `json:"boot_drives,omitempty"`
	BuildVersion        string             `json:"build_version,omitempty"`
	Causes              []interface{}      `json:"causes,omitempty"`
	Disconnected        bool               `json:"disconnected,omitempty"`
	FlashDevices        interface{}        `json:"flash_devices,omitempty"`
	Hdds                *[]Hdd             `json:"hdds,omitempty"`
	Health              string             `json:"health,omitempty"`
	HwHealth            string             `json:"hw_health,omitempty"`
	HwState             string             `json:"hw_state,omitempty"`
	InternalIp1         string             `json:"internal_ip_1,omitempty"`
	InternalIp2         string             `json:"internal_ip_2,omitempty"`
	LastRebootTimestamp int                `json:"last_reboot_timestamp,omitempty"`
	MgmtIp1             string             `json:"mgmt_ip_1,omitempty"`
	MgmtIp2             string             `json:"mgmt_ip_2,omitempty"`
	Model               string             `json:"model,omitempty"`
	Name                string             `json:"name,omitempty"`
	Nics                *[]Nic             `json:"nics,omitempty"`
	NvmFlashDevices     *[]NvmFlashDevice  `json:"nvm_flash_devices,omitempty"`
	OpProgress          interface{}        `json:"op_progress,omitempty"`
	OpState             string             `json:"op_state,omitempty"`
	OpStatus            string             `json:"op_status,omitempty"`
	OsVersion           string             `json:"os_version,omitempty"`
	Path                string             `json:"path,omitempty"`
	Psus                interface{}        `json:"psus,omitempty"`
	SerialNo            string             `json:"serial_no,omitempty"`
	StorageInstances    *[]StorageInstance `json:"storage_instances,omitempty"`
	SubsystemHealth     interface{}        `json:"subsystem_health,omitempty"`
	SubsystemStates     interface{}        `json:"subsystem_states,omitempty"`
	SwHealth            string             `json:"sw_health,omitempty"`
	SwState             string             `json:"sw_state,omitempty"`
	SwVersion           string             `json:"sw_version,omitempty"`
	TotalCapacity       int                `json:"total_capacity,omitempty"`
	TotalRawCapacity    int                `json:"total_raw_capacity,omitempty"`
	Type                string             `json:"type,omitempty"`
	Upgrade             interface{}        `json:"upgrade,omitempty"`
	Uuid                string             `json:"uuid,omitempty"`
	Vendor              string             `json:"vendor,omitempty"`
	Volumes             *[]Volume          `json:"volumes,omitempty"`
}

type StorageTemplate struct {
	Auth            *Auth             `json:"auth,omitempty"`
	IpPool          *IpPool           `json:"ip_pool,omitempty"`
	Name            string            `json:"name,omitempty"`
	Path            string            `json:"path,omitempty"`
	VolumeTemplates *[]VolumeTemplate `json:"volume_templates,omitempty"`
}

type Subsystem struct {
	Causes      []interface{} `json:"causes,omitempty"`
	Fan         string        `json:"fan,omitempty"`
	Health      string        `json:"health,omitempty"`
	Network     *Network      `json:"network,omitempty"`
	Path        string        `json:"path,omitempty"`
	Power       string        `json:"power,omitempty"`
	Temperature string        `json:"temperature,omitempty"`
	Voltage     string        `json:"voltage,omitempty"`
}

type System struct {
	AllFlashAvailableCapacity   int           `json:"all_flash_available_capacity,omitempty"`
	AllFlashProvisionedCapacity int           `json:"all_flash_provisioned_capacity,omitempty"`
	AllFlashTotalCapacity       int           `json:"all_flash_total_capacity,omitempty"`
	AvailableCapacity           int           `json:"available_capacity,omitempty"`
	BuildVersion                string        `json:"build_version,omitempty"`
	CallhomeEnabled             bool          `json:"callhome_enabled,omitempty"`
	Causes                      []interface{} `json:"causes,omitempty"`
	Dns                         *Dns          `json:"dns,omitempty"`
	Health                      string        `json:"health,omitempty"`
	HttpProxy                   *HttpProxy    `json:"http_proxy,omitempty"`
	HybridAvailableCapacity     int           `json:"hybrid_available_capacity,omitempty"`
	HybridProvisionedCapacity   int           `json:"hybrid_provisioned_capacity,omitempty"`
	HybridTotalCapacity         int           `json:"hybrid_total_capacity,omitempty"`
	LastRebootTimestamp         string        `json:"last_reboot_timestamp,omitempty"`
	Name                        string        `json:"name,omitempty"`
	Network                     *Network      `json:"network,omitempty"`
	NtpServers                  *[]NtpServer  `json:"ntp_servers,omitempty"`
	OpState                     string        `json:"op_state,omitempty"`
	Path                        string        `json:"path,omitempty"`
	SwVersion                   string        `json:"sw_version,omitempty"`
	TotalCapacity               int           `json:"total_capacity,omitempty"`
	TotalProvisionedCapacity    int           `json:"total_provisioned_capacity,omitempty"`
	Upgrade                     interface{}   `json:"upgrade,omitempty"`
	Uptime                      int           `json:"uptime,omitempty"`
	Uuid                        string        `json:"uuid,omitempty"`
}

type Tenant struct {
	Descr      string      `json:"descr,omitempty"`
	Name       string      `json:"name,omitempty"`
	ParentPath string      `json:"parent_path,omitempty"`
	Path       string      `json:"path,omitempty"`
	Subtenants interface{} `json:"subtenants,omitempty"`
}

type User struct {
	Email    string    `json:"email,omitempty"`
	Enabled  bool      `json:"enabled,omitempty"`
	FullName string    `json:"full_name,omitempty"`
	Password string    `json:"password,omitempty"`
	Path     string    `json:"path,omitempty"`
	Roles    *[]Role   `json:"roles,omitempty"`
	Tenants  *[]Tenant `json:"tenants,omitempty"`
	UserId   string    `json:"user_id,omitempty"`
	Version  string    `json:"version,omitempty"`
}

type Vip struct {
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

type Volume struct {
	ActiveStorageNodes []interface{} `json:"active_storage_nodes,omitempty"`
	CapacityInUse      int           `json:"capacity_in_use,omitempty"`
	Causes             []interface{} `json:"causes,omitempty"`
	Health             string        `json:"health,omitempty"`
	Name               string        `json:"name,omitempty"`
	OpState            string        `json:"op_state,omitempty"`
	OpStatus           string        `json:"op_status,omitempty"`
	Path               string        `json:"path,omitempty"`
	PlacementMode      string        `json:"placement_mode,omitempty"`
	ReplicaCount       int           `json:"replica_count,omitempty"`
	RestorePoint       string        `json:"restore_point,omitempty"`
	Size               int           `json:"size,omitempty"`
	Snapshots          *[]Snapshot   `json:"snapshots,omitempty"`
	Uuid               string        `json:"uuid,omitempty"`
}

type VolumeTemplate struct {
	Name          string `json:"name,omitempty"`
	Path          string `json:"path,omitempty"`
	PlacementMode string `json:"placement_mode,omitempty"`
	ReplicaCount  int    `json:"replica_count,omitempty"`
	Size          int    `json:"size,omitempty"`
}
