package dsdk

type Access struct {
	Path string   `json:"path,omitempty" mapstructure:"path"`
	Ips  []string `json:"ips,omitempty" mapstructure:"ips"`
	Iqn  string   `json:"iqn,omitempty" mapstructure:"iqn"`
}
