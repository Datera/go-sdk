package dsdk

type Access struct {
	Path string   `json:"path,omitempty"`
	Ips  []string `json:"ips,omitempty"`
	Iqn  string   `json:"iqn,omitempty"`
}
