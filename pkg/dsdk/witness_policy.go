package dsdk

type WitnessPolicy struct {
	Path               string `json:"path,omitempty" mapstructure:"path"`
	PreferredSite      string `json:"preferred_site,omitempty" mapstructure:"preferred_site"`
	HeartbeatFrequency int    `json:"heartbeat_frequency,omitempty" mapstructure:"heartbeat_frequency"`
	Enabled            bool   `json:"enabled,omitempty" mapstructure:"enabled"`
	Host               string `json:"host,omitempty" mapstructure:"host"`
	Port               int    `json:"port,omitempty" mapstructure:"port"`
	Site1Fd            string `json:"site_1_fd,omitempty" mapstructure:"site_1_fd"`
	Site2Fd            string `json:"site_2_fd,omitempty" mapstructure:"site_2_fd"`
	VerifyCert         bool   `json:"verify_cert,omitempty" mapstructure:"verify_cert"`
	UseProxy           bool   `json:"use_proxy,omitempty" mapstructure:"use_proxy"`
}
