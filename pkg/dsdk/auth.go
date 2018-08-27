package dsdk

type Auth struct {
	Path              string `json:"path,omitempty" mapstructure:"path"`
	Type              string `json:"type,omitempty" mapstructure:"type"`
	InitiatorUserName string `json:"initiator_user_name,omitempty" mapstructure:"initiator_user_name"`
	InitiatorPassword string `json:"initiator_pswd,omitempty" mapstructure:"initiator_pswd"`
	TargetUserName    string `json:"target_user_name,omitempty" mapstructure:"target_user_name"`
	TargetPassword    string `json:"target_pswd,omitempty" mapstructure:"target_pswd"`
	AccessKey         string `json:"access_key,omitempty" mapstructure:"access_key"`
	SecretKey         string `json:"secret_key,omitempty" mapstructure:"secret_key"`
}
