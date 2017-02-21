package dsdk

import "encoding/json"

// AccessNetworkIpPool Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type AccessNetworkIpPool struct {
	Descr        string      `json:"descr,omitempty"`
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

func NewAccessNetworkIpPool(arg []byte) (AccessNetworkIpPool, error) {
	var tmp AccessNetworkIpPool
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *AccessNetworkIpPool) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *AccessNetworkIpPool) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// AclPolicy Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type AclPolicy struct {
	InitiatorGroups *[]InitiatorGroup `json:"initiator_groups,omitempty"`
	Initiators      *[]Initiator      `json:"initiators,omitempty"`
	Path            string            `json:"path,omitempty"`
}

func NewAclPolicy(arg []byte) (AclPolicy, error) {
	var tmp AclPolicy
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *AclPolicy) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *AclPolicy) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// AppInstance Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewAppInstance(arg []byte) (AppInstance, error) {
	var tmp AppInstance
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *AppInstance) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *AppInstance) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// AppTemplate Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type AppTemplate struct {
	AppInstances     *[]AppInstance     `json:"app_instances,omitempty"`
	Descr            string             `json:"descr,omitempty"`
	Name             string             `json:"name,omitempty"`
	Path             string             `json:"path,omitempty"`
	SnapshotPolicies *[]SnapshotPolicy  `json:"snapshot_policies,omitempty"`
	StorageTemplates *[]StorageTemplate `json:"storage_templates,omitempty"`
}

func NewAppTemplate(arg []byte) (AppTemplate, error) {
	var tmp AppTemplate
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *AppTemplate) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *AppTemplate) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// AuditLog Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewAuditLog(arg []byte) (AuditLog, error) {
	var tmp AuditLog
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *AuditLog) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *AuditLog) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Auth Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type Auth struct {
	InitiatorPswd     string `json:"initiator_pswd,omitempty"`
	InitiatorUserName string `json:"initiator_user_name,omitempty"`
	Path              string `json:"path,omitempty"`
	TargetPswd        string `json:"target_pswd,omitempty"`
	TargetUserName    string `json:"target_user_name,omitempty"`
	Type              string `json:"type,omitempty"`
}

func NewAuth(arg []byte) (Auth, error) {
	var tmp Auth
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Auth) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Auth) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// BootDrive Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type BootDrive struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func NewBootDrive(arg []byte) (BootDrive, error) {
	var tmp BootDrive
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *BootDrive) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *BootDrive) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Dns Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type Dns struct {
	Domain string `json:"domain,omitempty"`
	Path   string `json:"path,omitempty"`
}

func NewDns(arg []byte) (Dns, error) {
	var tmp Dns
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Dns) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Dns) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// DnsSearchDomain Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type DnsSearchDomain struct {
	Domain string `json:"domain,omitempty"`
	Order  int    `json:"order,omitempty"`
	Path   string `json:"path,omitempty"`
}

func NewDnsSearchDomain(arg []byte) (DnsSearchDomain, error) {
	var tmp DnsSearchDomain
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *DnsSearchDomain) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *DnsSearchDomain) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// DnsServer Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type DnsServer struct {
	Ip    string `json:"ip,omitempty"`
	Order int    `json:"order,omitempty"`
	Path  string `json:"path,omitempty"`
}

func NewDnsServer(arg []byte) (DnsServer, error) {
	var tmp DnsServer
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *DnsServer) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *DnsServer) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// EventLog Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewEventLog(arg []byte) (EventLog, error) {
	var tmp EventLog
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *EventLog) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *EventLog) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// FaultLog Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewFaultLog(arg []byte) (FaultLog, error) {
	var tmp FaultLog
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *FaultLog) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *FaultLog) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Hdd Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type Hdd struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func NewHdd(arg []byte) (Hdd, error) {
	var tmp Hdd
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Hdd) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Hdd) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// HttpProxy Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type HttpProxy struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Host     string `json:"host,omitempty"`
	Password string `json:"password,omitempty"`
	Path     string `json:"path,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     *User  `json:"user,omitempty"`
}

func NewHttpProxy(arg []byte) (HttpProxy, error) {
	var tmp HttpProxy
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *HttpProxy) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *HttpProxy) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Initiator Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type Initiator struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

func NewInitiator(arg []byte) (Initiator, error) {
	var tmp Initiator
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Initiator) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Initiator) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// InitiatorGroup Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type InitiatorGroup struct {
	Members []interface{} `json:"members,omitempty"`
	Name    string        `json:"name,omitempty"`
	Path    string        `json:"path,omitempty"`
}

func NewInitiatorGroup(arg []byte) (InitiatorGroup, error) {
	var tmp InitiatorGroup
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *InitiatorGroup) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *InitiatorGroup) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// InternalIpBlock Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewInternalIpBlock(arg []byte) (InternalIpBlock, error) {
	var tmp InternalIpBlock
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *InternalIpBlock) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *InternalIpBlock) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// IpAddress Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type IpAddress struct {
	Gateway string `json:"gateway,omitempty"`
	Ip      string `json:"ip,omitempty"`
	Mtu     int    `json:"mtu,omitempty"`
	Name    string `json:"name,omitempty"`
	Netmask int    `json:"netmask,omitempty"`
	Path    string `json:"path,omitempty"`
	Vlan    int    `json:"vlan,omitempty"`
}

func NewIpAddress(arg []byte) (IpAddress, error) {
	var tmp IpAddress
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *IpAddress) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *IpAddress) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// IpBlock Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewIpBlock(arg []byte) (IpBlock, error) {
	var tmp IpBlock
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *IpBlock) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *IpBlock) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// IpPool Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type IpPool struct {
	Descr        string      `json:"descr,omitempty"`
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

func NewIpPool(arg []byte) (IpPool, error) {
	var tmp IpPool
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *IpPool) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *IpPool) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// MonitoringDestination Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewMonitoringDestination(arg []byte) (MonitoringDestination, error) {
	var tmp MonitoringDestination
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *MonitoringDestination) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *MonitoringDestination) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// MonitoringPolicy Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type MonitoringPolicy struct {
	Destinations []interface{} `json:"destinations,omitempty"`
	Enabled      bool          `json:"enabled,omitempty"`
	Name         string        `json:"name,omitempty"`
	Path         string        `json:"path,omitempty"`
}

func NewMonitoringPolicy(arg []byte) (MonitoringPolicy, error) {
	var tmp MonitoringPolicy
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *MonitoringPolicy) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *MonitoringPolicy) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Network Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type Network struct {
	AccessNetworks  []interface{} `json:"access_networks,omitempty"`
	AccessVip       interface{}   `json:"access_vip,omitempty"`
	InternalNetwork interface{}   `json:"internal_network,omitempty"`
	Mapping         interface{}   `json:"mapping,omitempty"`
	MgmtVip         interface{}   `json:"mgmt_vip,omitempty"`
	Path            string        `json:"path,omitempty"`
}

func NewNetwork(arg []byte) (Network, error) {
	var tmp Network
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Network) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Network) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// NetworkMapping Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type NetworkMapping struct {
	Access1   string `json:"access_1,omitempty"`
	Access2   string `json:"access_2,omitempty"`
	Internal1 string `json:"internal_1,omitempty"`
	Internal2 string `json:"internal_2,omitempty"`
	Mgmt1     string `json:"mgmt_1,omitempty"`
	Mgmt2     string `json:"mgmt_2,omitempty"`
	Path      string `json:"path,omitempty"`
}

func NewNetworkMapping(arg []byte) (NetworkMapping, error) {
	var tmp NetworkMapping
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *NetworkMapping) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *NetworkMapping) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Nic Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type Nic struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func NewNic(arg []byte) (Nic, error) {
	var tmp Nic
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Nic) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Nic) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// NtpServer Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type NtpServer struct {
	Ip    string `json:"ip,omitempty"`
	Order int    `json:"order,omitempty"`
	Path  string `json:"path,omitempty"`
}

func NewNtpServer(arg []byte) (NtpServer, error) {
	var tmp NtpServer
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *NtpServer) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *NtpServer) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// NvmFlashDevice Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type NvmFlashDevice struct {
	Causes    []interface{} `json:"causes,omitempty"`
	Health    string        `json:"health,omitempty"`
	Id        string        `json:"id,omitempty"`
	OpState   string        `json:"op_state,omitempty"`
	Path      string        `json:"path,omitempty"`
	Size      int           `json:"size,omitempty"`
	SlotLabel string        `json:"slot_label,omitempty"`
}

func NewNvmFlashDevice(arg []byte) (NvmFlashDevice, error) {
	var tmp NvmFlashDevice
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *NvmFlashDevice) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *NvmFlashDevice) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// PerformancePolicy Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type PerformancePolicy struct {
	Path              string `json:"path,omitempty"`
	ReadBandwidthMax  int    `json:"read_bandwidth_max,omitempty"`
	ReadIopsMax       int    `json:"read_iops_max,omitempty"`
	TotalBandwidthMax int    `json:"total_bandwidth_max,omitempty"`
	TotalIopsMax      int    `json:"total_iops_max,omitempty"`
	WriteBandwidthMax int    `json:"write_bandwidth_max,omitempty"`
	WriteIopsMax      int    `json:"write_iops_max,omitempty"`
}

func NewPerformancePolicy(arg []byte) (PerformancePolicy, error) {
	var tmp PerformancePolicy
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *PerformancePolicy) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *PerformancePolicy) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Role Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type Role struct {
	Path       string        `json:"path,omitempty"`
	Privileges []interface{} `json:"privileges,omitempty"`
	RoleId     string        `json:"role_id,omitempty"`
}

func NewRole(arg []byte) (Role, error) {
	var tmp Role
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Role) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Role) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Snapshot Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type Snapshot struct {
	OpState   string `json:"op_state,omitempty"`
	Path      string `json:"path,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	UtcTs     int    `json:"utc_ts,omitempty"`
	Uuid      string `json:"uuid,omitempty"`
}

func NewSnapshot(arg []byte) (Snapshot, error) {
	var tmp Snapshot
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Snapshot) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Snapshot) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// SnapshotPolicy Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type SnapshotPolicy struct {
	Interval       string `json:"interval,omitempty"`
	Name           string `json:"name,omitempty"`
	Path           string `json:"path,omitempty"`
	RetentionCount int    `json:"retention_count,omitempty"`
	StartTime      string `json:"start_time,omitempty"`
}

func NewSnapshotPolicy(arg []byte) (SnapshotPolicy, error) {
	var tmp SnapshotPolicy
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *SnapshotPolicy) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *SnapshotPolicy) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// SnmpPolicy Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type SnmpPolicy struct {
	Contact  string  `json:"contact,omitempty"`
	Enabled  bool    `json:"enabled,omitempty"`
	Location string  `json:"location,omitempty"`
	Path     string  `json:"path,omitempty"`
	Users    *[]User `json:"users,omitempty"`
}

func NewSnmpPolicy(arg []byte) (SnmpPolicy, error) {
	var tmp SnmpPolicy
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *SnmpPolicy) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *SnmpPolicy) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// SnmpUser Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewSnmpUser(arg []byte) (SnmpUser, error) {
	var tmp SnmpUser
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *SnmpUser) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *SnmpUser) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// StorageInstance Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewStorageInstance(arg []byte) (StorageInstance, error) {
	var tmp StorageInstance
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *StorageInstance) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *StorageInstance) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// StorageNode Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewStorageNode(arg []byte) (StorageNode, error) {
	var tmp StorageNode
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *StorageNode) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *StorageNode) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// StorageTemplate Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type StorageTemplate struct {
	Auth            *Auth             `json:"auth,omitempty"`
	IpPool          *IpPool           `json:"ip_pool,omitempty"`
	Name            string            `json:"name,omitempty"`
	Path            string            `json:"path,omitempty"`
	VolumeTemplates *[]VolumeTemplate `json:"volume_templates,omitempty"`
}

func NewStorageTemplate(arg []byte) (StorageTemplate, error) {
	var tmp StorageTemplate
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *StorageTemplate) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *StorageTemplate) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Subsystem Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewSubsystem(arg []byte) (Subsystem, error) {
	var tmp Subsystem
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Subsystem) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Subsystem) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// System Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewSystem(arg []byte) (System, error) {
	var tmp System
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *System) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *System) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Tenant Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type Tenant struct {
	Descr      string      `json:"descr,omitempty"`
	Name       string      `json:"name,omitempty"`
	ParentPath string      `json:"parent_path,omitempty"`
	Path       string      `json:"path,omitempty"`
	Subtenants interface{} `json:"subtenants,omitempty"`
}

func NewTenant(arg []byte) (Tenant, error) {
	var tmp Tenant
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Tenant) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Tenant) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// User Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewUser(arg []byte) (User, error) {
	var tmp User
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *User) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *User) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Vip Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type Vip struct {
	Name         string      `json:"name,omitempty"`
	NetworkPaths interface{} `json:"network_paths,omitempty"`
	Path         string      `json:"path,omitempty"`
}

func NewVip(arg []byte) (Vip, error) {
	var tmp Vip
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Vip) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Vip) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// Volume Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
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

func NewVolume(arg []byte) (Volume, error) {
	var tmp Volume
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *Volume) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *Volume) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}

// VolumeTemplate Entity Type for use when unpacking IEntity objects
// returned from an Endpoint call or in a Create or Set request
// to an Endpoint
type VolumeTemplate struct {
	Name          string `json:"name,omitempty"`
	Path          string `json:"path,omitempty"`
	PlacementMode string `json:"placement_mode,omitempty"`
	ReplicaCount  int    `json:"replica_count,omitempty"`
	Size          int    `json:"size,omitempty"`
}

func NewVolumeTemplate(arg []byte) (VolumeTemplate, error) {
	var tmp VolumeTemplate
	err := tmp.UnpackB(arg)
	return tmp, err
}

func (en *VolumeTemplate) Unpack(arg map[string]interface{}) error {
	tmp, err := json.Marshal(arg)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmp, &en)
}

func (en *VolumeTemplate) UnpackB(arg []byte) error {
	return json.Unmarshal(arg, &en)
}
