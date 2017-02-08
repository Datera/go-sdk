package dsdk

import (
	"encoding/json"
	//"fmt"
	"strings"
)

type RootEp struct {
	Path                 string
	conn                 *ApiConnection
	AppInstances         AppInstancesEndpoint
	Api                  ApiEndpoint
	AppTemplates         AppTemplatesEndpoint
	Initiators           InitiatorsEndpoint
	InitiatorGroups      InitiatorGroupsEndpoint
	AccessNetworkIpPools AccessNetworkIpPoolsEndpoint
	StorageNodes         StorageNodesEndpoint
	System               SystemEndpoint
	EventLogs            EventLogsEndpoint
	AuditLogs            AuditLogsEndpoint
	FaultLogs            FaultLogsEndpoint
	Roles                RolesEndpoint
	Users                UsersEndpoint
	Upgrade              UpgradeEndpoint
	Time                 TimeEndpoint
	Tenants              TenantsEndpoint
}

func NewRootEp(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*RootEp, error) {
	conn, err := NewApiConnection(hostname, port, username, password, apiVersion, tenant, timeout, headers, secure)
	if err != nil {
		return nil, err
	}
	err = conn.Login()
	if err != nil {
		return nil, err
	}
	return &RootEp{
		Path:                 "",
		conn:                 conn,
		AppInstances:         NewAppInstancesEndpoint("", conn),
		Api:                  NewApiEndpoint("", conn),
		AppTemplates:         NewAppTemplatesEndpoint("", conn),
		Initiators:           NewInitiatorsEndpoint("", conn),
		InitiatorGroups:      NewInitiatorGroupsEndpoint("", conn),
		AccessNetworkIpPools: NewAccessNetworkIpPoolsEndpoint("", conn),
		StorageNodes:         NewStorageNodesEndpoint("", conn),
		System:               NewSystemEndpoint("", conn),
		EventLogs:            NewEventLogsEndpoint("", conn),
		AuditLogs:            NewAuditLogsEndpoint("", conn),
		FaultLogs:            NewFaultLogsEndpoint("", conn),
		Roles:                NewRolesEndpoint("", conn),
		Users:                NewUsersEndpoint("", conn),
		Upgrade:              NewUpgradeEndpoint("", conn),
		Time:                 NewTimeEndpoint("", conn),
		Tenants:              NewTenantsEndpoint("", conn),
	}, nil
}

type AccessNetworkIpPoolEntity struct {
	conn         *ApiConnection
	Descr        string      `json:"descr,omitempty"`
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

func (en AccessNetworkIpPoolEntity) Reload() (AccessNetworkIpPoolEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n AccessNetworkIpPoolEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type AclPolicyEntity struct {
	conn            *ApiConnection
	InitiatorGroups []InitiatorGroupEntity `json:"initiator_groups,omitempty"`
	Initiators      []InitiatorEntity      `json:"initiators,omitempty"`
	Path            string                 `json:"path,omitempty"`
}

func (en AclPolicyEntity) Reload() (AclPolicyEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n AclPolicyEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type AppInstanceEntity struct {
	conn              *ApiConnection
	AccessControlMode string                  `json:"access_control_mode,omitempty"`
	AdminState        string                  `json:"admin_state,omitempty"`
	AppTemplate       AppTemplateEntity       `json:"app_template,omitempty"`
	Causes            []interface{}           `json:"causes,omitempty"`
	CloneSrc          map[string]string       `json:"clone_src,omitempty"`
	CreateMode        string                  `json:"create_mode,omitempty"`
	Descr             string                  `json:"descr,omitempty"`
	Health            string                  `json:"health,omitempty"`
	Id                string                  `json:"id,omitempty"`
	Name              string                  `json:"name,omitempty"`
	OpStatus          string                  `json:"op_status,omitempty"`
	Path              string                  `json:"path,omitempty"`
	RestorePoint      string                  `json:"restore_point,omitempty"`
	SnapshotPolicies  []SnapshotPolicyEntity  `json:"snapshot_policies,omitempty"`
	Snapshots         []SnapshotEntity        `json:"snapshots,omitempty"`
	StorageInstances  []StorageInstanceEntity `json:"storage_instances,omitempty"`
	Uuid              string                  `json:"uuid,omitempty"`
}

func (en AppInstanceEntity) Reload() (AppInstanceEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n AppInstanceEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type AppTemplateEntity struct {
	conn             *ApiConnection
	AppInstances     []AppInstanceEntity     `json:"app_instances,omitempty"`
	Descr            string                  `json:"descr,omitempty"`
	Name             string                  `json:"name,omitempty"`
	Path             string                  `json:"path,omitempty"`
	SnapshotPolicies []SnapshotPolicyEntity  `json:"snapshot_policies,omitempty"`
	StorageTemplates []StorageTemplateEntity `json:"storage_templates,omitempty"`
}

func (en AppTemplateEntity) Reload() (AppTemplateEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n AppTemplateEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type AuditLogEntity struct {
	conn        *ApiConnection
	Description string     `json:"description,omitempty"`
	Id          string     `json:"id,omitempty"`
	ObjectName  string     `json:"object_name,omitempty"`
	ObjectType  string     `json:"object_type,omitempty"`
	ObjectUrl   string     `json:"object_url,omitempty"`
	Operation   string     `json:"operation,omitempty"`
	ParamInfo   string     `json:"param_info,omitempty"`
	Path        string     `json:"path,omitempty"`
	SessionInfo string     `json:"session_info,omitempty"`
	Timestamp   string     `json:"timestamp,omitempty"`
	User        UserEntity `json:"user,omitempty"`
	Version     string     `json:"version,omitempty"`
}

func (en AuditLogEntity) Reload() (AuditLogEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n AuditLogEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type AuthEntity struct {
	conn              *ApiConnection
	InitiatorPswd     string `json:"initiator_pswd,omitempty"`
	InitiatorUserName string `json:"initiator_user_name,omitempty"`
	Path              string `json:"path,omitempty"`
	TargetPswd        string `json:"target_pswd,omitempty"`
	TargetUserName    string `json:"target_user_name,omitempty"`
	Type              string `json:"type,omitempty"`
}

func (en AuthEntity) Reload() (AuthEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n AuthEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type BootDriveEntity struct {
	conn      *ApiConnection
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func (en BootDriveEntity) Reload() (BootDriveEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n BootDriveEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type DnsEntity struct {
	conn   *ApiConnection
	Domain string `json:"domain,omitempty"`
	Path   string `json:"path,omitempty"`
}

func (en DnsEntity) Reload() (DnsEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n DnsEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type DnsSearchDomainEntity struct {
	conn   *ApiConnection
	Domain string `json:"domain,omitempty"`
	Order  int    `json:"order,omitempty"`
	Path   string `json:"path,omitempty"`
}

func (en DnsSearchDomainEntity) Reload() (DnsSearchDomainEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n DnsSearchDomainEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type DnsServerEntity struct {
	conn  *ApiConnection
	Ip    string `json:"ip,omitempty"`
	Order int    `json:"order,omitempty"`
	Path  string `json:"path,omitempty"`
}

func (en DnsServerEntity) Reload() (DnsServerEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n DnsServerEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type EventLogEntity struct {
	conn         *ApiConnection
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

func (en EventLogEntity) Reload() (EventLogEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n EventLogEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type FaultLogEntity struct {
	conn            *ApiConnection
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

func (en FaultLogEntity) Reload() (FaultLogEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n FaultLogEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type HddEntity struct {
	conn      *ApiConnection
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func (en HddEntity) Reload() (HddEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n HddEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type HttpProxyEntity struct {
	conn     *ApiConnection
	Enabled  bool       `json:"enabled,omitempty"`
	Host     string     `json:"host,omitempty"`
	Password string     `json:"password,omitempty"`
	Path     string     `json:"path,omitempty"`
	Port     int        `json:"port,omitempty"`
	User     UserEntity `json:"user,omitempty"`
}

func (en HttpProxyEntity) Reload() (HttpProxyEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n HttpProxyEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type InitiatorEntity struct {
	conn *ApiConnection
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

func (en InitiatorEntity) Reload() (InitiatorEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n InitiatorEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type InitiatorGroupEntity struct {
	conn    *ApiConnection
	Members []interface{} `json:"members,omitempty"`
	Name    string        `json:"name,omitempty"`
	Path    string        `json:"path,omitempty"`
}

func (en InitiatorGroupEntity) Reload() (InitiatorGroupEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n InitiatorGroupEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type InternalIpBlockEntity struct {
	conn    *ApiConnection
	Gateway string `json:"gateway,omitempty"`
	Mtu     int    `json:"mtu,omitempty"`
	Name    string `json:"name,omitempty"`
	Netmask int    `json:"netmask,omitempty"`
	Path    string `json:"path,omitempty"`
	Range   int    `json:"range,omitempty"`
	StartIp string `json:"start_ip,omitempty"`
	Vlan    int    `json:"vlan,omitempty"`
}

func (en InternalIpBlockEntity) Reload() (InternalIpBlockEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n InternalIpBlockEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type IpAddressEntity struct {
	conn    *ApiConnection
	Gateway string `json:"gateway,omitempty"`
	Ip      string `json:"ip,omitempty"`
	Mtu     int    `json:"mtu,omitempty"`
	Name    string `json:"name,omitempty"`
	Netmask int    `json:"netmask,omitempty"`
	Path    string `json:"path,omitempty"`
	Vlan    int    `json:"vlan,omitempty"`
}

func (en IpAddressEntity) Reload() (IpAddressEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n IpAddressEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type IpBlockEntity struct {
	conn    *ApiConnection
	Gateway string `json:"gateway,omitempty"`
	Mtu     int    `json:"mtu,omitempty"`
	Name    string `json:"name,omitempty"`
	Netmask int    `json:"netmask,omitempty"`
	Path    string `json:"path,omitempty"`
	Range   int    `json:"range,omitempty"`
	StartIp string `json:"start_ip,omitempty"`
	Vlan    int    `json:"vlan,omitempty"`
}

func (en IpBlockEntity) Reload() (IpBlockEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n IpBlockEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type IpPoolEntity struct {
	conn         *ApiConnection
	Descr        string      `json:"descr,omitempty"`
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

func (en IpPoolEntity) Reload() (IpPoolEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n IpPoolEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type MonitoringDestinationEntity struct {
	conn      *ApiConnection
	Facility  string `json:"facility,omitempty"`
	Host      string `json:"host,omitempty"`
	LastMsgTs string `json:"last_msg_ts,omitempty"`
	Name      string `json:"name,omitempty"`
	OpState   string `json:"op_state,omitempty"`
	Path      string `json:"path,omitempty"`
	Port      int    `json:"port,omitempty"`
	Type      string `json:"type,omitempty"`
}

func (en MonitoringDestinationEntity) Reload() (MonitoringDestinationEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n MonitoringDestinationEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type MonitoringPolicyEntity struct {
	conn         *ApiConnection
	Destinations []interface{} `json:"destinations,omitempty"`
	Enabled      bool          `json:"enabled,omitempty"`
	Name         string        `json:"name,omitempty"`
	Path         string        `json:"path,omitempty"`
}

func (en MonitoringPolicyEntity) Reload() (MonitoringPolicyEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n MonitoringPolicyEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type NetworkEntity struct {
	conn            *ApiConnection
	AccessNetworks  []interface{} `json:"access_networks,omitempty"`
	AccessVip       interface{}   `json:"access_vip,omitempty"`
	InternalNetwork interface{}   `json:"internal_network,omitempty"`
	Mapping         interface{}   `json:"mapping,omitempty"`
	MgmtVip         interface{}   `json:"mgmt_vip,omitempty"`
	Path            string        `json:"path,omitempty"`
}

func (en NetworkEntity) Reload() (NetworkEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n NetworkEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type NetworkMappingEntity struct {
	conn      *ApiConnection
	Access1   string `json:"access_1,omitempty"`
	Access2   string `json:"access_2,omitempty"`
	Internal1 string `json:"internal_1,omitempty"`
	Internal2 string `json:"internal_2,omitempty"`
	Mgmt1     string `json:"mgmt_1,omitempty"`
	Mgmt2     string `json:"mgmt_2,omitempty"`
	Path      string `json:"path,omitempty"`
}

func (en NetworkMappingEntity) Reload() (NetworkMappingEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n NetworkMappingEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type NicEntity struct {
	conn      *ApiConnection
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func (en NicEntity) Reload() (NicEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n NicEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type NtpServerEntity struct {
	conn  *ApiConnection
	Ip    string `json:"ip,omitempty"`
	Order int    `json:"order,omitempty"`
	Path  string `json:"path,omitempty"`
}

func (en NtpServerEntity) Reload() (NtpServerEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n NtpServerEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type NvmFlashDeviceEntity struct {
	conn      *ApiConnection
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func (en NvmFlashDeviceEntity) Reload() (NvmFlashDeviceEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n NvmFlashDeviceEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type PerformancePolicyEntity struct {
	conn              *ApiConnection
	Path              string `json:"path,omitempty"`
	ReadBandwidthMax  int    `json:"read_bandwidth_max,omitempty"`
	ReadIopsMax       int    `json:"read_iops_max,omitempty"`
	TotalBandwidthMax int    `json:"total_bandwidth_max,omitempty"`
	TotalIopsMax      int    `json:"total_iops_max,omitempty"`
	WriteBandwidthMax int    `json:"write_bandwidth_max,omitempty"`
	WriteIopsMax      int    `json:"write_iops_max,omitempty"`
}

func (en PerformancePolicyEntity) Reload() (PerformancePolicyEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n PerformancePolicyEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type RoleEntity struct {
	conn       *ApiConnection
	Path       string        `json:"path,omitempty"`
	Privileges []interface{} `json:"privileges,omitempty"`
	RoleId     string        `json:"role_id,omitempty"`
}

func (en RoleEntity) Reload() (RoleEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n RoleEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type SnapshotEntity struct {
	conn      *ApiConnection
	OpState   string `json:"op_state,omitempty"`
	Path      string `json:"path,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	UtcTs     int    `json:"utc_ts,omitempty"`
	Uuid      string `json:"uuid,omitempty"`
}

func (en SnapshotEntity) Reload() (SnapshotEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n SnapshotEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type SnapshotPolicyEntity struct {
	conn           *ApiConnection
	Interval       string `json:"interval,omitempty"`
	Name           string `json:"name,omitempty"`
	Path           string `json:"path,omitempty"`
	RetentionCount int    `json:"retention_count,omitempty"`
	StartTime      string `json:"start_time,omitempty"`
}

func (en SnapshotPolicyEntity) Reload() (SnapshotPolicyEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n SnapshotPolicyEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type SnmpPolicyEntity struct {
	conn     *ApiConnection
	Contact  string       `json:"contact,omitempty"`
	Enabled  bool         `json:"enabled,omitempty"`
	Location string       `json:"location,omitempty"`
	Path     string       `json:"path,omitempty"`
	Users    []UserEntity `json:"users,omitempty"`
}

func (en SnmpPolicyEntity) Reload() (SnmpPolicyEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n SnmpPolicyEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type SnmpUserEntity struct {
	conn          *ApiConnection
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

func (en SnmpUserEntity) Reload() (SnmpUserEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n SnmpUserEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type StorageInstanceEntity struct {
	conn               *ApiConnection
	Access             interface{}     `json:"access,omitempty"`
	AccessControlMode  string          `json:"access_control_mode,omitempty"`
	AclPolicy          AclPolicyEntity `json:"acl_policy,omitempty"`
	ActiveInitiators   []interface{}   `json:"active_initiators,omitempty"`
	ActiveStorageNodes []interface{}   `json:"active_storage_nodes,omitempty"`
	AdminState         string          `json:"admin_state,omitempty"`
	Auth               AuthEntity      `json:"auth,omitempty"`
	Causes             []interface{}   `json:"causes,omitempty"`
	Health             string          `json:"health,omitempty"`
	IpPool             IpPoolEntity    `json:"ip_pool,omitempty"`
	Name               string          `json:"name,omitempty"`
	OpState            string          `json:"op_state,omitempty"`
	Path               string          `json:"path,omitempty"`
	Uuid               string          `json:"uuid,omitempty"`
	Volumes            []VolumeEntity  `json:"volumes,omitempty"`
}

func (en StorageInstanceEntity) Reload() (StorageInstanceEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n StorageInstanceEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type StorageNodeEntity struct {
	conn                *ApiConnection
	AdminState          string                  `json:"admin_state,omitempty"`
	AvailableCapacity   int                     `json:"available_capacity,omitempty"`
	BiosVersion         string                  `json:"bios_version,omitempty"`
	BootDrives          interface{}             `json:"boot_drives,omitempty"`
	BuildVersion        string                  `json:"build_version,omitempty"`
	Causes              []interface{}           `json:"causes,omitempty"`
	Disconnected        bool                    `json:"disconnected,omitempty"`
	FlashDevices        interface{}             `json:"flash_devices,omitempty"`
	Hdds                []HddEntity             `json:"hdds,omitempty"`
	Health              string                  `json:"health,omitempty"`
	HwHealth            string                  `json:"hw_health,omitempty"`
	HwState             string                  `json:"hw_state,omitempty"`
	InternalIp1         string                  `json:"internal_ip_1,omitempty"`
	InternalIp2         string                  `json:"internal_ip_2,omitempty"`
	LastRebootTimestamp int                     `json:"last_reboot_timestamp,omitempty"`
	MgmtIp1             string                  `json:"mgmt_ip_1,omitempty"`
	MgmtIp2             string                  `json:"mgmt_ip_2,omitempty"`
	Model               string                  `json:"model,omitempty"`
	Name                string                  `json:"name,omitempty"`
	Nics                []NicEntity             `json:"nics,omitempty"`
	NvmFlashDevices     []NvmFlashDeviceEntity  `json:"nvm_flash_devices,omitempty"`
	OpProgress          interface{}             `json:"op_progress,omitempty"`
	OpState             string                  `json:"op_state,omitempty"`
	OpStatus            string                  `json:"op_status,omitempty"`
	OsVersion           string                  `json:"os_version,omitempty"`
	Path                string                  `json:"path,omitempty"`
	Psus                interface{}             `json:"psus,omitempty"`
	SerialNo            string                  `json:"serial_no,omitempty"`
	StorageInstances    []StorageInstanceEntity `json:"storage_instances,omitempty"`
	SubsystemHealth     interface{}             `json:"subsystem_health,omitempty"`
	SubsystemStates     interface{}             `json:"subsystem_states,omitempty"`
	SwHealth            string                  `json:"sw_health,omitempty"`
	SwState             string                  `json:"sw_state,omitempty"`
	SwVersion           string                  `json:"sw_version,omitempty"`
	TotalCapacity       int                     `json:"total_capacity,omitempty"`
	TotalRawCapacity    int                     `json:"total_raw_capacity,omitempty"`
	Type                string                  `json:"type,omitempty"`
	Upgrade             interface{}             `json:"upgrade,omitempty"`
	Uuid                string                  `json:"uuid,omitempty"`
	Vendor              string                  `json:"vendor,omitempty"`
	Volumes             []VolumeEntity          `json:"volumes,omitempty"`
}

func (en StorageNodeEntity) Reload() (StorageNodeEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n StorageNodeEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type StorageTemplateEntity struct {
	conn            *ApiConnection
	Auth            AuthEntity             `json:"auth,omitempty"`
	IpPool          IpPoolEntity           `json:"ip_pool,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Path            string                 `json:"path,omitempty"`
	VolumeTemplates []VolumeTemplateEntity `json:"volume_templates,omitempty"`
}

func (en StorageTemplateEntity) Reload() (StorageTemplateEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n StorageTemplateEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type SubsystemEntity struct {
	conn        *ApiConnection
	Causes      []interface{} `json:"causes,omitempty"`
	Fan         string        `json:"fan,omitempty"`
	Health      string        `json:"health,omitempty"`
	Network     NetworkEntity `json:"network,omitempty"`
	Path        string        `json:"path,omitempty"`
	Power       string        `json:"power,omitempty"`
	Temperature string        `json:"temperature,omitempty"`
	Voltage     string        `json:"voltage,omitempty"`
}

func (en SubsystemEntity) Reload() (SubsystemEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n SubsystemEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type SystemEntity struct {
	conn                        *ApiConnection
	AllFlashAvailableCapacity   int               `json:"all_flash_available_capacity,omitempty"`
	AllFlashProvisionedCapacity int               `json:"all_flash_provisioned_capacity,omitempty"`
	AllFlashTotalCapacity       int               `json:"all_flash_total_capacity,omitempty"`
	AvailableCapacity           int               `json:"available_capacity,omitempty"`
	BuildVersion                string            `json:"build_version,omitempty"`
	CallhomeEnabled             bool              `json:"callhome_enabled,omitempty"`
	Causes                      []interface{}     `json:"causes,omitempty"`
	Dns                         DnsEntity         `json:"dns,omitempty"`
	Health                      string            `json:"health,omitempty"`
	HttpProxy                   HttpProxyEntity   `json:"http_proxy,omitempty"`
	HybridAvailableCapacity     int               `json:"hybrid_available_capacity,omitempty"`
	HybridProvisionedCapacity   int               `json:"hybrid_provisioned_capacity,omitempty"`
	HybridTotalCapacity         int               `json:"hybrid_total_capacity,omitempty"`
	LastRebootTimestamp         string            `json:"last_reboot_timestamp,omitempty"`
	Name                        string            `json:"name,omitempty"`
	Network                     NetworkEntity     `json:"network,omitempty"`
	NtpServers                  []NtpServerEntity `json:"ntp_servers,omitempty"`
	OpState                     string            `json:"op_state,omitempty"`
	Path                        string            `json:"path,omitempty"`
	SwVersion                   string            `json:"sw_version,omitempty"`
	TotalCapacity               int               `json:"total_capacity,omitempty"`
	TotalProvisionedCapacity    int               `json:"total_provisioned_capacity,omitempty"`
	Upgrade                     interface{}       `json:"upgrade,omitempty"`
	Uptime                      int               `json:"uptime,omitempty"`
	Uuid                        string            `json:"uuid,omitempty"`
}

func (en SystemEntity) Reload() (SystemEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n SystemEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type TenantEntity struct {
	conn       *ApiConnection
	Descr      string      `json:"descr,omitempty"`
	Name       string      `json:"name,omitempty"`
	ParentPath string      `json:"parent_path,omitempty"`
	Path       string      `json:"path,omitempty"`
	Subtenants interface{} `json:"subtenants,omitempty"`
}

func (en TenantEntity) Reload() (TenantEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n TenantEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type UserEntity struct {
	conn     *ApiConnection
	Email    string         `json:"email,omitempty"`
	Enabled  bool           `json:"enabled,omitempty"`
	FullName string         `json:"full_name,omitempty"`
	Password string         `json:"password,omitempty"`
	Path     string         `json:"path,omitempty"`
	Roles    []RoleEntity   `json:"roles,omitempty"`
	Tenants  []TenantEntity `json:"tenants,omitempty"`
	UserId   string         `json:"user_id,omitempty"`
	Version  string         `json:"version,omitempty"`
}

func (en UserEntity) Reload() (UserEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n UserEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type VipEntity struct {
	conn         *ApiConnection
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

func (en VipEntity) Reload() (VipEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n VipEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type VolumeEntity struct {
	conn               *ApiConnection
	ActiveStorageNodes []interface{}    `json:"active_storage_nodes,omitempty"`
	CapacityInUse      int              `json:"capacity_in_use,omitempty"`
	Causes             []interface{}    `json:"causes,omitempty"`
	Health             string           `json:"health,omitempty"`
	Name               string           `json:"name,omitempty"`
	OpState            string           `json:"op_state,omitempty"`
	OpStatus           string           `json:"op_status,omitempty"`
	Path               string           `json:"path,omitempty"`
	PlacementMode      string           `json:"placement_mode,omitempty"`
	ReplicaCount       int              `json:"replica_count,omitempty"`
	RestorePoint       string           `json:"restore_point,omitempty"`
	Size               int              `json:"size,omitempty"`
	Snapshots          []SnapshotEntity `json:"snapshots,omitempty"`
	Uuid               string           `json:"uuid,omitempty"`
}

func (en VolumeEntity) Reload() (VolumeEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n VolumeEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type VolumeTemplateEntity struct {
	conn          *ApiConnection
	Name          string `json:"name,omitempty"`
	Path          string `json:"path,omitempty"`
	PlacementMode string `json:"placement_mode,omitempty"`
	ReplicaCount  int    `json:"replica_count,omitempty"`
	Size          int    `json:"size,omitempty"`
}

func (en VolumeTemplateEntity) Reload() (VolumeTemplateEntity, error) {
	r, _ := en.conn.Get(en.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var n VolumeTemplateEntity
	err = json.Unmarshal(d, &n)
	n.conn = en.conn
	return n, nil
}

type MetadataEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewMetadataEndpoint(parent string, conn *ApiConnection) MetadataEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "metadata"}, "/"), "/")
	return MetadataEndpoint{
		conn: conn,
		Path: path,
	}
}

type LoginEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewLoginEndpoint(parent string, conn *ApiConnection) LoginEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "login"}, "/"), "/")
	return LoginEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint struct {
	conn       *ApiConnection
	Path       string
	VolumeName AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint) List(queryp ...string) ([]VolumeEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []VolumeEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint{
		conn:       conn,
		Path:       path,
		VolumeName: NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint(path, conn),
	}
}

type MetricsHwMetricEndpoint struct {
	conn   *ApiConnection
	Path   string
	Latest MetricsHwMetricLatestEndpoint
}

func NewMetricsHwMetricEndpoint(parent string, conn *ApiConnection) MetricsHwMetricEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "metrics/hw/:metric"}, "/"), "/")
	return MetricsHwMetricEndpoint{
		conn:   conn,
		Path:   path,
		Latest: NewMetricsHwMetricLatestEndpoint(path, conn),
	}
}

type EventLogsEndpoint struct {
	conn *ApiConnection
	Path string
	Id   EventLogsIdEndpoint
}

func (ep EventLogsEndpoint) List(queryp ...string) ([]EventLogEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []EventLogEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewEventLogsEndpoint(parent string, conn *ApiConnection) EventLogsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "event_logs"}, "/"), "/")
	return EventLogsEndpoint{
		conn: conn,
		Path: path,
		Id:   NewEventLogsIdEndpoint(path, conn),
	}
}

type StorageNodesUuidFlashDevicesEndpoint struct {
	conn *ApiConnection
	Path string
	Id   StorageNodesUuidFlashDevicesIdEndpoint
}

func NewStorageNodesUuidFlashDevicesEndpoint(parent string, conn *ApiConnection) StorageNodesUuidFlashDevicesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/flash_devices"}, "/"), "/")
	return StorageNodesUuidFlashDevicesEndpoint{
		conn: conn,
		Path: path,
		Id:   NewStorageNodesUuidFlashDevicesIdEndpoint(path, conn),
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name/snapshots/:timestamp"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint{
		conn: conn,
		Path: path,
	}
}

type LogoutEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewLogoutEndpoint(parent string, conn *ApiConnection) LogoutEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "logout"}, "/"), "/")
	return LogoutEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemSnmpPolicyUsersEndpoint struct {
	conn   *ApiConnection
	Path   string
	UserId SystemSnmpPolicyUsersUserIdEndpoint
}

func (ep SystemSnmpPolicyUsersEndpoint) List(queryp ...string) ([]UserEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []UserEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewSystemSnmpPolicyUsersEndpoint(parent string, conn *ApiConnection) SystemSnmpPolicyUsersEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/snmp_policy/users"}, "/"), "/")
	return SystemSnmpPolicyUsersEndpoint{
		conn:   conn,
		Path:   path,
		UserId: NewSystemSnmpPolicyUsersUserIdEndpoint(path, conn),
	}
}

type SystemSnmpPolicyUsersUserIdEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewSystemSnmpPolicyUsersUserIdEndpoint(parent string, conn *ApiConnection) SystemSnmpPolicyUsersUserIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/snmp_policy/users/:user_id"}, "/"), "/")
	return SystemSnmpPolicyUsersUserIdEndpoint{
		conn: conn,
		Path: path,
	}
}

type UsersUserIdEndpoint struct {
	conn  *ApiConnection
	Path  string
	Roles UsersUserIdRolesEndpoint
}

func NewUsersUserIdEndpoint(parent string, conn *ApiConnection) UsersUserIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "users/:user_id"}, "/"), "/")
	return UsersUserIdEndpoint{
		conn:  conn,
		Path:  path,
		Roles: NewUsersUserIdRolesEndpoint(path, conn),
	}
}

type AccessNetworkIpPoolsPoolNameEndpoint struct {
	conn         *ApiConnection
	Path         string
	NetworkPaths AccessNetworkIpPoolsPoolNameNetworkPathsEndpoint
}

func NewAccessNetworkIpPoolsPoolNameEndpoint(parent string, conn *ApiConnection) AccessNetworkIpPoolsPoolNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "access_network_ip_pools/:pool_name"}, "/"), "/")
	return AccessNetworkIpPoolsPoolNameEndpoint{
		conn:         conn,
		Path:         path,
		NetworkPaths: NewAccessNetworkIpPoolsPoolNameNetworkPathsEndpoint(path, conn),
	}
}

type ApiEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewApiEndpoint(parent string, conn *ApiConnection) ApiEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "api"}, "/"), "/")
	return ApiEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemDnsSearchDomainsEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewSystemDnsSearchDomainsEndpoint(parent string, conn *ApiConnection) SystemDnsSearchDomainsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/dns/search_domains"}, "/"), "/")
	return SystemDnsSearchDomainsEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdEndpoint struct {
	conn             *ApiConnection
	Path             string
	Metadata         AppInstancesIdMetadataEndpoint
	SnapshotPolicies AppInstancesIdSnapshotPoliciesEndpoint
	Snapshots        AppInstancesIdSnapshotsEndpoint
	StorageInstances AppInstancesIdStorageInstancesEndpoint
}

func NewAppInstancesIdEndpoint(parent string, conn *ApiConnection) AppInstancesIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id"}, "/"), "/")
	return AppInstancesIdEndpoint{
		conn:             conn,
		Path:             path,
		Metadata:         NewAppInstancesIdMetadataEndpoint(path, conn),
		SnapshotPolicies: NewAppInstancesIdSnapshotPoliciesEndpoint(path, conn),
		Snapshots:        NewAppInstancesIdSnapshotsEndpoint(path, conn),
		StorageInstances: NewAppInstancesIdStorageInstancesEndpoint(path, conn),
	}
}

type MetricsIoMetricLatestEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewMetricsIoMetricLatestEndpoint(parent string, conn *ApiConnection) MetricsIoMetricLatestEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "metrics/io/:metric/latest"}, "/"), "/")
	return MetricsIoMetricLatestEndpoint{
		conn: conn,
		Path: path,
	}
}

type AccessNetworkIpPoolsEndpoint struct {
	conn     *ApiConnection
	Path     string
	PoolName AccessNetworkIpPoolsPoolNameEndpoint
}

func (ep AccessNetworkIpPoolsEndpoint) List(queryp ...string) ([]AccessNetworkIpPoolEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []AccessNetworkIpPoolEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAccessNetworkIpPoolsEndpoint(parent string, conn *ApiConnection) AccessNetworkIpPoolsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "access_network_ip_pools"}, "/"), "/")
	return AccessNetworkIpPoolsEndpoint{
		conn:     conn,
		Path:     path,
		PoolName: NewAccessNetworkIpPoolsPoolNameEndpoint(path, conn),
	}
}

type AccessNetworkIpPoolsPoolNameNetworkPathsEndpoint struct {
	conn     *ApiConnection
	Path     string
	PathName AccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint
}

func NewAccessNetworkIpPoolsPoolNameNetworkPathsEndpoint(parent string, conn *ApiConnection) AccessNetworkIpPoolsPoolNameNetworkPathsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "access_network_ip_pools/:pool_name/network_paths"}, "/"), "/")
	return AccessNetworkIpPoolsPoolNameNetworkPathsEndpoint{
		conn:     conn,
		Path:     path,
		PathName: NewAccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint(path, conn),
	}
}

type InitConfigEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewInitConfigEndpoint(parent string, conn *ApiConnection) InitConfigEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "init/config"}, "/"), "/")
	return InitConfigEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppTemplatesNameSnapshotPoliciesEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep AppTemplatesNameSnapshotPoliciesEndpoint) List(queryp ...string) ([]SnapshotPolicyEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []SnapshotPolicyEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppTemplatesNameSnapshotPoliciesEndpoint(parent string, conn *ApiConnection) AppTemplatesNameSnapshotPoliciesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:name/snapshot_policies"}, "/"), "/")
	return AppTemplatesNameSnapshotPoliciesEndpoint{
		conn: conn,
		Path: path,
	}
}

type PrivilegesEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewPrivilegesEndpoint(parent string, conn *ApiConnection) PrivilegesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "privileges"}, "/"), "/")
	return PrivilegesEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemNetworkMgmtVipNetworkPathsEndpoint struct {
	conn        *ApiConnection
	Path        string
	NetworkPath SystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint
}

func NewSystemNetworkMgmtVipNetworkPathsEndpoint(parent string, conn *ApiConnection) SystemNetworkMgmtVipNetworkPathsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/mgmt_vip/network_paths"}, "/"), "/")
	return SystemNetworkMgmtVipNetworkPathsEndpoint{
		conn:        conn,
		Path:        path,
		NetworkPath: NewSystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint(path, conn),
	}
}

type SystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewSystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint(parent string, conn *ApiConnection) SystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/mgmt_vip/network_paths/:network_path"}, "/"), "/")
	return SystemNetworkMgmtVipNetworkPathsNetworkPathEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint struct {
	conn              *ApiConnection
	Path              string
	PerformancePolicy AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint
	SnapshotPolicies  AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint
	Snapshots         AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameEndpoint{
		conn:              conn,
		Path:              path,
		PerformancePolicy: NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint(path, conn),
		SnapshotPolicies:  NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint(path, conn),
		Snapshots:         NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint(path, conn),
	}
}

type AuditLogsIdEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewAuditLogsIdEndpoint(parent string, conn *ApiConnection) AuditLogsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "audit_logs/:id"}, "/"), "/")
	return AuditLogsIdEndpoint{
		conn: conn,
		Path: path,
	}
}

type MonitoringDestinationsDefaultEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewMonitoringDestinationsDefaultEndpoint(parent string, conn *ApiConnection) MonitoringDestinationsDefaultEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "monitoring/destinations/default"}, "/"), "/")
	return MonitoringDestinationsDefaultEndpoint{
		conn: conn,
		Path: path,
	}
}

type EventsSystemEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep EventsSystemEndpoint) List(queryp ...string) ([]SystemEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []SystemEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewEventsSystemEndpoint(parent string, conn *ApiConnection) EventsSystemEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "events/system"}, "/"), "/")
	return EventsSystemEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemDnsServersEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewSystemDnsServersEndpoint(parent string, conn *ApiConnection) SystemDnsServersEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/dns/servers"}, "/"), "/")
	return SystemDnsServersEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemDnsEndpoint struct {
	conn          *ApiConnection
	Path          string
	SearchDomains SystemDnsSearchDomainsEndpoint
	Servers       SystemDnsServersEndpoint
}

func (ep SystemDnsEndpoint) List(queryp ...string) ([]DnsEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []DnsEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewSystemDnsEndpoint(parent string, conn *ApiConnection) SystemDnsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/dns"}, "/"), "/")
	return SystemDnsEndpoint{
		conn:          conn,
		Path:          path,
		SearchDomains: NewSystemDnsSearchDomainsEndpoint(path, conn),
		Servers:       NewSystemDnsServersEndpoint(path, conn),
	}
}

type SystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewSystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint(parent string, conn *ApiConnection) SystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/internal_network/network_paths/:network_path"}, "/"), "/")
	return SystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemNetworkMgmtVipEndpoint struct {
	conn         *ApiConnection
	Path         string
	NetworkPaths SystemNetworkMgmtVipNetworkPathsEndpoint
}

func NewSystemNetworkMgmtVipEndpoint(parent string, conn *ApiConnection) SystemNetworkMgmtVipEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/mgmt_vip"}, "/"), "/")
	return SystemNetworkMgmtVipEndpoint{
		conn:         conn,
		Path:         path,
		NetworkPaths: NewSystemNetworkMgmtVipNetworkPathsEndpoint(path, conn),
	}
}

type AppTemplatesNameStorageTemplatesEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep AppTemplatesNameStorageTemplatesEndpoint) List(queryp ...string) ([]StorageTemplateEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []StorageTemplateEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppTemplatesNameStorageTemplatesEndpoint(parent string, conn *ApiConnection) AppTemplatesNameStorageTemplatesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:name/storage_templates"}, "/"), "/")
	return AppTemplatesNameStorageTemplatesEndpoint{
		conn: conn,
		Path: path,
	}
}

type TenantsTenantPathEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewTenantsTenantPathEndpoint(parent string, conn *ApiConnection) TenantsTenantPathEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "tenants/:tenant_path"}, "/"), "/")
	return TenantsTenantPathEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint struct {
	conn            *ApiConnection
	Path            string
	InitiatorGroups AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint
	Initiators      AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint) List(queryp ...string) ([]AclPolicyEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []AclPolicyEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/acl_policy"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint{
		conn:            conn,
		Path:            path,
		InitiatorGroups: NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint(path, conn),
		Initiators:      NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint(path, conn),
	}
}

type UpgradeAvailableEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewUpgradeAvailableEndpoint(parent string, conn *ApiConnection) UpgradeAvailableEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "upgrade/available"}, "/"), "/")
	return UpgradeAvailableEndpoint{
		conn: conn,
		Path: path,
	}
}

type StorageNodesUuidNicsIdEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewStorageNodesUuidNicsIdEndpoint(parent string, conn *ApiConnection) StorageNodesUuidNicsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/nics/:id"}, "/"), "/")
	return StorageNodesUuidNicsIdEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdSnapshotsTimestampEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewAppInstancesIdSnapshotsTimestampEndpoint(parent string, conn *ApiConnection) AppInstancesIdSnapshotsTimestampEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/snapshots/:timestamp"}, "/"), "/")
	return AppInstancesIdSnapshotsTimestampEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemVersionConfigEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewSystemVersionConfigEndpoint(parent string, conn *ApiConnection) SystemVersionConfigEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/version_config"}, "/"), "/")
	return SystemVersionConfigEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewAppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint(parent string, conn *ApiConnection) AppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/snapshot_policies/:snapshot_policy_name"}, "/"), "/")
	return AppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint) List(queryp ...string) ([]InitiatorEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []InitiatorEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/acl_policy/initiators"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorsEndpoint{
		conn: conn,
		Path: path,
	}
}

type EventsDebugEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewEventsDebugEndpoint(parent string, conn *ApiConnection) EventsDebugEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "events/debug"}, "/"), "/")
	return EventsDebugEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint) List(queryp ...string) ([]PerformancePolicyEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []PerformancePolicyEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name/performance_policy"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNamePerformancePolicyEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint) List(queryp ...string) ([]AuthEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []AuthEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/auth"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint{
		conn: conn,
		Path: path,
	}
}

type UsersUserIdRolesEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep UsersUserIdRolesEndpoint) List(queryp ...string) ([]RoleEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []RoleEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewUsersUserIdRolesEndpoint(parent string, conn *ApiConnection) UsersUserIdRolesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "users/:user_id/roles"}, "/"), "/")
	return UsersUserIdRolesEndpoint{
		conn: conn,
		Path: path,
	}
}

type TenantsEndpoint struct {
	conn       *ApiConnection
	Path       string
	TenantPath TenantsTenantPathEndpoint
}

func (ep TenantsEndpoint) List(queryp ...string) ([]TenantEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []TenantEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewTenantsEndpoint(parent string, conn *ApiConnection) TenantsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "tenants"}, "/"), "/")
	return TenantsEndpoint{
		conn:       conn,
		Path:       path,
		TenantPath: NewTenantsTenantPathEndpoint(path, conn),
	}
}

type InitiatorsEndpoint struct {
	conn *ApiConnection
	Path string
	Id   InitiatorsIdEndpoint
}

func (ep InitiatorsEndpoint) List(queryp ...string) ([]InitiatorEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []InitiatorEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewInitiatorsEndpoint(parent string, conn *ApiConnection) InitiatorsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "initiators"}, "/"), "/")
	return InitiatorsEndpoint{
		conn: conn,
		Path: path,
		Id:   NewInitiatorsIdEndpoint(path, conn),
	}
}

type StorageNodesUuidSubsystemStatesEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewStorageNodesUuidSubsystemStatesEndpoint(parent string, conn *ApiConnection) StorageNodesUuidSubsystemStatesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/subsystem_states"}, "/"), "/")
	return StorageNodesUuidSubsystemStatesEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint struct {
	conn               *ApiConnection
	Path               string
	SnapshotPolicyName AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint) List(queryp ...string) ([]SnapshotPolicyEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []SnapshotPolicyEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name/snapshot_policies"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesEndpoint{
		conn:               conn,
		Path:               path,
		SnapshotPolicyName: NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint(path, conn),
	}
}

type MetricsIoMetricEndpoint struct {
	conn   *ApiConnection
	Path   string
	Latest MetricsIoMetricLatestEndpoint
}

func NewMetricsIoMetricEndpoint(parent string, conn *ApiConnection) MetricsIoMetricEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "metrics/io/:metric"}, "/"), "/")
	return MetricsIoMetricEndpoint{
		conn:   conn,
		Path:   path,
		Latest: NewMetricsIoMetricLatestEndpoint(path, conn),
	}
}

type TimeEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewTimeEndpoint(parent string, conn *ApiConnection) TimeEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "time"}, "/"), "/")
	return TimeEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppTemplatesNameEndpoint struct {
	conn             *ApiConnection
	Path             string
	SnapshotPolicies AppTemplatesNameSnapshotPoliciesEndpoint
	StorageTemplates AppTemplatesNameStorageTemplatesEndpoint
}

func NewAppTemplatesNameEndpoint(parent string, conn *ApiConnection) AppTemplatesNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:name"}, "/"), "/")
	return AppTemplatesNameEndpoint{
		conn:             conn,
		Path:             path,
		SnapshotPolicies: NewAppTemplatesNameSnapshotPoliciesEndpoint(path, conn),
		StorageTemplates: NewAppTemplatesNameStorageTemplatesEndpoint(path, conn),
	}
}

type AuditLogsEndpoint struct {
	conn *ApiConnection
	Path string
	Id   AuditLogsIdEndpoint
}

func (ep AuditLogsEndpoint) List(queryp ...string) ([]AuditLogEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []AuditLogEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAuditLogsEndpoint(parent string, conn *ApiConnection) AuditLogsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "audit_logs"}, "/"), "/")
	return AuditLogsEndpoint{
		conn: conn,
		Path: path,
		Id:   NewAuditLogsIdEndpoint(path, conn),
	}
}

type RolesEndpoint struct {
	conn   *ApiConnection
	Path   string
	RoleId RolesRoleIdEndpoint
}

func (ep RolesEndpoint) List(queryp ...string) ([]RoleEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []RoleEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewRolesEndpoint(parent string, conn *ApiConnection) RolesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "roles"}, "/"), "/")
	return RolesEndpoint{
		conn:   conn,
		Path:   path,
		RoleId: NewRolesRoleIdEndpoint(path, conn),
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint) List(queryp ...string) ([]SnapshotPolicyEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []SnapshotPolicyEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint(parent string, conn *ApiConnection) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name/volume_templates/:volume_template_name/snapshot_policies"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemNetworkInternalNetworkEndpoint struct {
	conn         *ApiConnection
	Path         string
	NetworkPaths SystemNetworkInternalNetworkNetworkPathsEndpoint
}

func NewSystemNetworkInternalNetworkEndpoint(parent string, conn *ApiConnection) SystemNetworkInternalNetworkEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/internal_network"}, "/"), "/")
	return SystemNetworkInternalNetworkEndpoint{
		conn:         conn,
		Path:         path,
		NetworkPaths: NewSystemNetworkInternalNetworkNetworkPathsEndpoint(path, conn),
	}
}

type EventsUserEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep EventsUserEndpoint) List(queryp ...string) ([]UserEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []UserEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewEventsUserEndpoint(parent string, conn *ApiConnection) EventsUserEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "events/user"}, "/"), "/")
	return EventsUserEndpoint{
		conn: conn,
		Path: path,
	}
}

type InitiatorGroupsNameEndpoint struct {
	conn    *ApiConnection
	Path    string
	Members InitiatorGroupsNameMembersEndpoint
}

func NewInitiatorGroupsNameEndpoint(parent string, conn *ApiConnection) InitiatorGroupsNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "initiator_groups/:name"}, "/"), "/")
	return InitiatorGroupsNameEndpoint{
		conn:    conn,
		Path:    path,
		Members: NewInitiatorGroupsNameMembersEndpoint(path, conn),
	}
}

type FaultLogsEndpoint struct {
	conn *ApiConnection
	Path string
	Id   FaultLogsIdEndpoint
}

func (ep FaultLogsEndpoint) List(queryp ...string) ([]FaultLogEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []FaultLogEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewFaultLogsEndpoint(parent string, conn *ApiConnection) FaultLogsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "fault_logs"}, "/"), "/")
	return FaultLogsEndpoint{
		conn: conn,
		Path: path,
		Id:   NewFaultLogsIdEndpoint(path, conn),
	}
}

type StorageNodesUuidHddsIdEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewStorageNodesUuidHddsIdEndpoint(parent string, conn *ApiConnection) StorageNodesUuidHddsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/hdds/:id"}, "/"), "/")
	return StorageNodesUuidHddsIdEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemHttpProxyEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep SystemHttpProxyEndpoint) List(queryp ...string) ([]HttpProxyEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []HttpProxyEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewSystemHttpProxyEndpoint(parent string, conn *ApiConnection) SystemHttpProxyEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/http_proxy"}, "/"), "/")
	return SystemHttpProxyEndpoint{
		conn: conn,
		Path: path,
	}
}

type StorageNodesUuidEndpoint struct {
	conn            *ApiConnection
	Path            string
	BootDrives      StorageNodesUuidBootDrivesEndpoint
	FlashDevices    StorageNodesUuidFlashDevicesEndpoint
	Hdds            StorageNodesUuidHddsEndpoint
	Nics            StorageNodesUuidNicsEndpoint
	SubsystemStates StorageNodesUuidSubsystemStatesEndpoint
}

func NewStorageNodesUuidEndpoint(parent string, conn *ApiConnection) StorageNodesUuidEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid"}, "/"), "/")
	return StorageNodesUuidEndpoint{
		conn:            conn,
		Path:            path,
		BootDrives:      NewStorageNodesUuidBootDrivesEndpoint(path, conn),
		FlashDevices:    NewStorageNodesUuidFlashDevicesEndpoint(path, conn),
		Hdds:            NewStorageNodesUuidHddsEndpoint(path, conn),
		Nics:            NewStorageNodesUuidNicsEndpoint(path, conn),
		SubsystemStates: NewStorageNodesUuidSubsystemStatesEndpoint(path, conn),
	}
}

type EventsUuidEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewEventsUuidEndpoint(parent string, conn *ApiConnection) EventsUuidEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "events/:uuid"}, "/"), "/")
	return EventsUuidEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemEndpoint struct {
	conn          *ApiConnection
	Path          string
	Dns           SystemDnsEndpoint
	HttpProxy     SystemHttpProxyEndpoint
	Network       SystemNetworkEndpoint
	NtpServers    SystemNtpServersEndpoint
	SnmpPolicy    SystemSnmpPolicyEndpoint
	VersionConfig SystemVersionConfigEndpoint
}

func (ep SystemEndpoint) List(queryp ...string) ([]SystemEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []SystemEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewSystemEndpoint(parent string, conn *ApiConnection) SystemEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system"}, "/"), "/")
	return SystemEndpoint{
		conn:          conn,
		Path:          path,
		Dns:           NewSystemDnsEndpoint(path, conn),
		HttpProxy:     NewSystemHttpProxyEndpoint(path, conn),
		Network:       NewSystemNetworkEndpoint(path, conn),
		NtpServers:    NewSystemNtpServersEndpoint(path, conn),
		SnmpPolicy:    NewSystemSnmpPolicyEndpoint(path, conn),
		VersionConfig: NewSystemVersionConfigEndpoint(path, conn),
	}
}

type StorageNodesEndpoint struct {
	conn *ApiConnection
	Path string
	Uuid StorageNodesUuidEndpoint
}

func (ep StorageNodesEndpoint) List(queryp ...string) ([]StorageNodeEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []StorageNodeEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewStorageNodesEndpoint(parent string, conn *ApiConnection) StorageNodesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes"}, "/"), "/")
	return StorageNodesEndpoint{
		conn: conn,
		Path: path,
		Uuid: NewStorageNodesUuidEndpoint(path, conn),
	}
}

type AppTemplatesEndpoint struct {
	conn *ApiConnection
	Path string
	Name AppTemplatesNameEndpoint
}

func (ep AppTemplatesEndpoint) List(queryp ...string) ([]AppTemplateEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []AppTemplateEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppTemplatesEndpoint(parent string, conn *ApiConnection) AppTemplatesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates"}, "/"), "/")
	return AppTemplatesEndpoint{
		conn: conn,
		Path: path,
		Name: NewAppTemplatesNameEndpoint(path, conn),
	}
}

type UpgradeEndpoint struct {
	conn      *ApiConnection
	Path      string
	Available UpgradeAvailableEndpoint
}

func NewUpgradeEndpoint(parent string, conn *ApiConnection) UpgradeEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "upgrade"}, "/"), "/")
	return UpgradeEndpoint{
		conn:      conn,
		Path:      path,
		Available: NewUpgradeAvailableEndpoint(path, conn),
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint struct {
	conn      *ApiConnection
	Path      string
	Timestamp AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint) List(queryp ...string) ([]SnapshotEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []SnapshotEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name/snapshots"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsEndpoint{
		conn:      conn,
		Path:      path,
		Timestamp: NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotsTimestampEndpoint(path, conn),
	}
}

type InitiatorsIdEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewInitiatorsIdEndpoint(parent string, conn *ApiConnection) InitiatorsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "initiators/:id"}, "/"), "/")
	return InitiatorsIdEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemNetworkAccessVipNetworkPathsEndpoint struct {
	conn        *ApiConnection
	Path        string
	NetworkPath SystemNetworkAccessVipNetworkPathsNetworkPathEndpoint
}

func NewSystemNetworkAccessVipNetworkPathsEndpoint(parent string, conn *ApiConnection) SystemNetworkAccessVipNetworkPathsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/access_vip/network_paths"}, "/"), "/")
	return SystemNetworkAccessVipNetworkPathsEndpoint{
		conn:        conn,
		Path:        path,
		NetworkPath: NewSystemNetworkAccessVipNetworkPathsNetworkPathEndpoint(path, conn),
	}
}

type PolicyadmEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewPolicyadmEndpoint(parent string, conn *ApiConnection) PolicyadmEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "policyadm"}, "/"), "/")
	return PolicyadmEndpoint{
		conn: conn,
		Path: path,
	}
}

type EventLogsIdEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewEventLogsIdEndpoint(parent string, conn *ApiConnection) EventLogsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "event_logs/:id"}, "/"), "/")
	return EventLogsIdEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint) List(queryp ...string) ([]InitiatorGroupEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []InitiatorGroupEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/acl_policy/initiator_groups"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyInitiatorGroupsEndpoint{
		conn: conn,
		Path: path,
	}
}

type MonitoringAlertsEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewMonitoringAlertsEndpoint(parent string, conn *ApiConnection) MonitoringAlertsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "monitoring/alerts"}, "/"), "/")
	return MonitoringAlertsEndpoint{
		conn: conn,
		Path: path,
	}
}

type StorageNodesUuidFlashDevicesIdEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewStorageNodesUuidFlashDevicesIdEndpoint(parent string, conn *ApiConnection) StorageNodesUuidFlashDevicesIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/flash_devices/:id"}, "/"), "/")
	return StorageNodesUuidFlashDevicesIdEndpoint{
		conn: conn,
		Path: path,
	}
}

type MetricsHwMetricLatestEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewMetricsHwMetricLatestEndpoint(parent string, conn *ApiConnection) MetricsHwMetricLatestEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "metrics/hw/:metric/latest"}, "/"), "/")
	return MetricsHwMetricLatestEndpoint{
		conn: conn,
		Path: path,
	}
}

type FaultLogsIdEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewFaultLogsIdEndpoint(parent string, conn *ApiConnection) FaultLogsIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "fault_logs/:id"}, "/"), "/")
	return FaultLogsIdEndpoint{
		conn: conn,
		Path: path,
	}
}

type InitiatorGroupsEndpoint struct {
	conn *ApiConnection
	Path string
	Name InitiatorGroupsNameEndpoint
}

func (ep InitiatorGroupsEndpoint) List(queryp ...string) ([]InitiatorGroupEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []InitiatorGroupEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewInitiatorGroupsEndpoint(parent string, conn *ApiConnection) InitiatorGroupsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "initiator_groups"}, "/"), "/")
	return InitiatorGroupsEndpoint{
		conn: conn,
		Path: path,
		Name: NewInitiatorGroupsNameEndpoint(path, conn),
	}
}

type SystemNetworkMappingEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewSystemNetworkMappingEndpoint(parent string, conn *ApiConnection) SystemNetworkMappingEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/mapping"}, "/"), "/")
	return SystemNetworkMappingEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppTemplatesAtNameStorageTemplatesStNameVolumeTemplatesVtNameSnapshotPoliciesSnapshotPolicyNameEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewAppTemplatesAtNameStorageTemplatesStNameVolumeTemplatesVtNameSnapshotPoliciesSnapshotPolicyNameEndpoint(parent string, conn *ApiConnection) AppTemplatesAtNameStorageTemplatesStNameVolumeTemplatesVtNameSnapshotPoliciesSnapshotPolicyNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:at_name/storage_templates/:st_name/volume_templates/:vt_name/snapshot_policies/:snapshot_policy_name"}, "/"), "/")
	return AppTemplatesAtNameStorageTemplatesStNameVolumeTemplatesVtNameSnapshotPoliciesSnapshotPolicyNameEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdSnapshotsEndpoint struct {
	conn      *ApiConnection
	Path      string
	Timestamp AppInstancesIdSnapshotsTimestampEndpoint
}

func (ep AppInstancesIdSnapshotsEndpoint) List(queryp ...string) ([]SnapshotEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []SnapshotEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesIdSnapshotsEndpoint(parent string, conn *ApiConnection) AppInstancesIdSnapshotsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/snapshots"}, "/"), "/")
	return AppInstancesIdSnapshotsEndpoint{
		conn:      conn,
		Path:      path,
		Timestamp: NewAppInstancesIdSnapshotsTimestampEndpoint(path, conn),
	}
}

type SystemSnmpPolicyEndpoint struct {
	conn  *ApiConnection
	Path  string
	Users SystemSnmpPolicyUsersEndpoint
}

func (ep SystemSnmpPolicyEndpoint) List(queryp ...string) ([]SnmpPolicyEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []SnmpPolicyEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewSystemSnmpPolicyEndpoint(parent string, conn *ApiConnection) SystemSnmpPolicyEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/snmp_policy"}, "/"), "/")
	return SystemSnmpPolicyEndpoint{
		conn:  conn,
		Path:  path,
		Users: NewSystemSnmpPolicyUsersEndpoint(path, conn),
	}
}

type UserinfoEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewUserinfoEndpoint(parent string, conn *ApiConnection) UserinfoEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "userinfo"}, "/"), "/")
	return UserinfoEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint struct {
	conn               *ApiConnection
	Path               string
	VolumeTemplateName AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint
}

func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint) List(queryp ...string) ([]VolumeTemplateEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []VolumeTemplateEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint(parent string, conn *ApiConnection) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name/volume_templates"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint{
		conn:               conn,
		Path:               path,
		VolumeTemplateName: NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint(path, conn),
	}
}

type HealthAttrsEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewHealthAttrsEndpoint(parent string, conn *ApiConnection) HealthAttrsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "health_attrs"}, "/"), "/")
	return HealthAttrsEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameEndpoint struct {
	conn      *ApiConnection
	Path      string
	AclPolicy AppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint
	Auth      AppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint
	Volumes   AppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameEndpoint{
		conn:      conn,
		Path:      path,
		AclPolicy: NewAppInstancesIdStorageInstancesStorageInstanceNameAclPolicyEndpoint(path, conn),
		Auth:      NewAppInstancesIdStorageInstancesStorageInstanceNameAuthEndpoint(path, conn),
		Volumes:   NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesEndpoint(path, conn),
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint) List(queryp ...string) ([]AuthEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []AuthEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint(parent string, conn *ApiConnection) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name/auth"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint{
		conn: conn,
		Path: path,
	}
}

type StorageNodesUuidHddsEndpoint struct {
	conn *ApiConnection
	Path string
	Id   StorageNodesUuidHddsIdEndpoint
}

func (ep StorageNodesUuidHddsEndpoint) List(queryp ...string) ([]HddEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []HddEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewStorageNodesUuidHddsEndpoint(parent string, conn *ApiConnection) StorageNodesUuidHddsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/hdds"}, "/"), "/")
	return StorageNodesUuidHddsEndpoint{
		conn: conn,
		Path: path,
		Id:   NewStorageNodesUuidHddsIdEndpoint(path, conn),
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameEndpoint struct {
	conn            *ApiConnection
	Path            string
	Auth            AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint
	VolumeTemplates AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameEndpoint(parent string, conn *ApiConnection) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameEndpoint{
		conn:            conn,
		Path:            path,
		Auth:            NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameAuthEndpoint(path, conn),
		VolumeTemplates: NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesEndpoint(path, conn),
	}
}

type MonitoringPoliciesDefaultEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewMonitoringPoliciesDefaultEndpoint(parent string, conn *ApiConnection) MonitoringPoliciesDefaultEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "monitoring/policies/default"}, "/"), "/")
	return MonitoringPoliciesDefaultEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemNetworkAccessVipEndpoint struct {
	conn         *ApiConnection
	Path         string
	NetworkPaths SystemNetworkAccessVipNetworkPathsEndpoint
}

func NewSystemNetworkAccessVipEndpoint(parent string, conn *ApiConnection) SystemNetworkAccessVipEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/access_vip"}, "/"), "/")
	return SystemNetworkAccessVipEndpoint{
		conn:         conn,
		Path:         path,
		NetworkPaths: NewSystemNetworkAccessVipNetworkPathsEndpoint(path, conn),
	}
}

type AppInstancesIdSnapshotPoliciesEndpoint struct {
	conn               *ApiConnection
	Path               string
	SnapshotPolicyName AppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint
}

func (ep AppInstancesIdSnapshotPoliciesEndpoint) List(queryp ...string) ([]SnapshotPolicyEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []SnapshotPolicyEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesIdSnapshotPoliciesEndpoint(parent string, conn *ApiConnection) AppInstancesIdSnapshotPoliciesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/snapshot_policies"}, "/"), "/")
	return AppInstancesIdSnapshotPoliciesEndpoint{
		conn:               conn,
		Path:               path,
		SnapshotPolicyName: NewAppInstancesIdSnapshotPoliciesSnapshotPolicyNameEndpoint(path, conn),
	}
}

type AppInstancesAiIdStorageInstancesSiIdMetadataEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewAppInstancesAiIdStorageInstancesSiIdMetadataEndpoint(parent string, conn *ApiConnection) AppInstancesAiIdStorageInstancesSiIdMetadataEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:ai_id/storage_instances/:si_id/metadata"}, "/"), "/")
	return AppInstancesAiIdStorageInstancesSiIdMetadataEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdStorageInstancesEndpoint struct {
	conn                *ApiConnection
	Path                string
	StorageInstanceName AppInstancesIdStorageInstancesStorageInstanceNameEndpoint
}

func (ep AppInstancesIdStorageInstancesEndpoint) List(queryp ...string) ([]StorageInstanceEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []StorageInstanceEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesIdStorageInstancesEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances"}, "/"), "/")
	return AppInstancesIdStorageInstancesEndpoint{
		conn:                conn,
		Path:                path,
		StorageInstanceName: NewAppInstancesIdStorageInstancesStorageInstanceNameEndpoint(path, conn),
	}
}

type StorageNodesUuidBootDrivesIdEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewStorageNodesUuidBootDrivesIdEndpoint(parent string, conn *ApiConnection) StorageNodesUuidBootDrivesIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/boot_drives/:id"}, "/"), "/")
	return StorageNodesUuidBootDrivesIdEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewAppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint(parent string, conn *ApiConnection) AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/storage_instances/:storage_instance_name/volumes/:volume_name/snapshot_policies/:snapshot_policy_name"}, "/"), "/")
	return AppInstancesIdStorageInstancesStorageInstanceNameVolumesVolumeNameSnapshotPoliciesSnapshotPolicyNameEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemNetworkEndpoint struct {
	conn            *ApiConnection
	Path            string
	AccessVip       SystemNetworkAccessVipEndpoint
	InternalNetwork SystemNetworkInternalNetworkEndpoint
	Mapping         SystemNetworkMappingEndpoint
	MgmtVip         SystemNetworkMgmtVipEndpoint
}

func (ep SystemNetworkEndpoint) List(queryp ...string) ([]NetworkEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []NetworkEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewSystemNetworkEndpoint(parent string, conn *ApiConnection) SystemNetworkEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network"}, "/"), "/")
	return SystemNetworkEndpoint{
		conn:            conn,
		Path:            path,
		AccessVip:       NewSystemNetworkAccessVipEndpoint(path, conn),
		InternalNetwork: NewSystemNetworkInternalNetworkEndpoint(path, conn),
		Mapping:         NewSystemNetworkMappingEndpoint(path, conn),
		MgmtVip:         NewSystemNetworkMgmtVipEndpoint(path, conn),
	}
}

type UsersEndpoint struct {
	conn   *ApiConnection
	Path   string
	UserId UsersUserIdEndpoint
}

func (ep UsersEndpoint) List(queryp ...string) ([]UserEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []UserEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewUsersEndpoint(parent string, conn *ApiConnection) UsersEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "users"}, "/"), "/")
	return UsersEndpoint{
		conn:   conn,
		Path:   path,
		UserId: NewUsersUserIdEndpoint(path, conn),
	}
}

type InitiatorGroupsNameMembersEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewInitiatorGroupsNameMembersEndpoint(parent string, conn *ApiConnection) InitiatorGroupsNameMembersEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "initiator_groups/:name/members"}, "/"), "/")
	return InitiatorGroupsNameMembersEndpoint{
		conn: conn,
		Path: path,
	}
}

type StorageNodesUuidBootDrivesEndpoint struct {
	conn *ApiConnection
	Path string
	Id   StorageNodesUuidBootDrivesIdEndpoint
}

func NewStorageNodesUuidBootDrivesEndpoint(parent string, conn *ApiConnection) StorageNodesUuidBootDrivesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/boot_drives"}, "/"), "/")
	return StorageNodesUuidBootDrivesEndpoint{
		conn: conn,
		Path: path,
		Id:   NewStorageNodesUuidBootDrivesIdEndpoint(path, conn),
	}
}

type SystemNtpServersEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep SystemNtpServersEndpoint) List(queryp ...string) ([]NtpServerEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []NtpServerEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewSystemNtpServersEndpoint(parent string, conn *ApiConnection) SystemNtpServersEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/ntp_servers"}, "/"), "/")
	return SystemNtpServersEndpoint{
		conn: conn,
		Path: path,
	}
}

type SystemNetworkInternalNetworkNetworkPathsEndpoint struct {
	conn        *ApiConnection
	Path        string
	NetworkPath SystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint
}

func NewSystemNetworkInternalNetworkNetworkPathsEndpoint(parent string, conn *ApiConnection) SystemNetworkInternalNetworkNetworkPathsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/internal_network/network_paths"}, "/"), "/")
	return SystemNetworkInternalNetworkNetworkPathsEndpoint{
		conn:        conn,
		Path:        path,
		NetworkPath: NewSystemNetworkInternalNetworkNetworkPathsNetworkPathEndpoint(path, conn),
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint struct {
	conn              *ApiConnection
	Path              string
	PerformancePolicy AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint
	SnapshotPolicies  AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint(parent string, conn *ApiConnection) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name/volume_templates/:volume_template_name"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameEndpoint{
		conn:              conn,
		Path:              path,
		PerformancePolicy: NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint(path, conn),
		SnapshotPolicies:  NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNameSnapshotPoliciesEndpoint(path, conn),
	}
}

type AppInstancesEndpoint struct {
	conn *ApiConnection
	Path string
	Id   AppInstancesIdEndpoint
}

func (ep AppInstancesEndpoint) List(queryp ...string) ([]AppInstanceEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []AppInstanceEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppInstancesEndpoint(parent string, conn *ApiConnection) AppInstancesEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances"}, "/"), "/")
	return AppInstancesEndpoint{
		conn: conn,
		Path: path,
		Id:   NewAppInstancesIdEndpoint(path, conn),
	}
}

type AppInstancesIdMetadataEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewAppInstancesIdMetadataEndpoint(parent string, conn *ApiConnection) AppInstancesIdMetadataEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_instances/:id/metadata"}, "/"), "/")
	return AppInstancesIdMetadataEndpoint{
		conn: conn,
		Path: path,
	}
}

type StorageNodesUuidNicsEndpoint struct {
	conn *ApiConnection
	Path string
	Id   StorageNodesUuidNicsIdEndpoint
}

func (ep StorageNodesUuidNicsEndpoint) List(queryp ...string) ([]NicEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []NicEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewStorageNodesUuidNicsEndpoint(parent string, conn *ApiConnection) StorageNodesUuidNicsEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "storage_nodes/:uuid/nics"}, "/"), "/")
	return StorageNodesUuidNicsEndpoint{
		conn: conn,
		Path: path,
		Id:   NewStorageNodesUuidNicsIdEndpoint(path, conn),
	}
}

type SystemNetworkAccessVipNetworkPathsNetworkPathEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewSystemNetworkAccessVipNetworkPathsNetworkPathEndpoint(parent string, conn *ApiConnection) SystemNetworkAccessVipNetworkPathsNetworkPathEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "system/network/access_vip/network_paths/:network_path"}, "/"), "/")
	return SystemNetworkAccessVipNetworkPathsNetworkPathEndpoint{
		conn: conn,
		Path: path,
	}
}

type RolesRoleIdEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewRolesRoleIdEndpoint(parent string, conn *ApiConnection) RolesRoleIdEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "roles/:role_id"}, "/"), "/")
	return RolesRoleIdEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint struct {
	conn *ApiConnection
	Path string
}

func (ep AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint) List(queryp ...string) ([]PerformancePolicyEntity, error) {
	r, _ := ep.conn.Get(ep.Path)
	d, err := getData(r)
	if err != nil {
		panic(err)
	}
	var ens []PerformancePolicyEntity
	err = json.Unmarshal(d, &ens)
	if err != nil {
		panic(err)
	}
	for _, en := range ens {
		en.conn = ep.conn
	}
	return ens, nil
}

func NewAppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint(parent string, conn *ApiConnection) AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/storage_templates/:storage_template_name/volume_templates/:volume_template_name/performance_policy"}, "/"), "/")
	return AppTemplatesAppTemplateNameStorageTemplatesStorageTemplateNameVolumeTemplatesVolumeTemplateNamePerformancePolicyEndpoint{
		conn: conn,
		Path: path,
	}
}

type AccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewAccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint(parent string, conn *ApiConnection) AccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "access_network_ip_pools/:pool_name/network_paths/:path_name"}, "/"), "/")
	return AccessNetworkIpPoolsPoolNameNetworkPathsPathNameEndpoint{
		conn: conn,
		Path: path,
	}
}

type AppTemplatesAppTemplateNameSnapshotPoliciesSnapshotPolicyNameEndpoint struct {
	conn *ApiConnection
	Path string
}

func NewAppTemplatesAppTemplateNameSnapshotPoliciesSnapshotPolicyNameEndpoint(parent string, conn *ApiConnection) AppTemplatesAppTemplateNameSnapshotPoliciesSnapshotPolicyNameEndpoint {
	path := strings.Trim(strings.Join([]string{parent, "app_templates/:app_template_name/snapshot_policies/:snapshot_policy_name"}, "/"), "/")
	return AppTemplatesAppTemplateNameSnapshotPoliciesSnapshotPolicyNameEndpoint{
		conn: conn,
		Path: path,
	}
}
