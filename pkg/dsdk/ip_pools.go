package dsdk

type IpPool struct {
	Path           string `json:"path,omitempty"`
	ResolvedPath   string `json:"resolved_path,omitempty"`
	ResolvedTenant string `json:"resolved_tenant,omitempty"`
}
