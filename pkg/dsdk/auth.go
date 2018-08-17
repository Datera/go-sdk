package dsdk

type Auth struct {
	Path              string `json:"path,omitempty"`
	Type              string `json:"type,omitempty"`
	InitiatorUserName string `json:"initiator_user_name,omitempty"`
	InitiatorPassword string `json:"initiator_pswd,omitempty"`
	TargetUserName    string `json:"target_user_name,omitempty"`
	TargetPassword    string `json:"target_pswd,omitempty"`
	AccessKey         string `json:"access_key,omitempty"`
	SecretKey         string `json:"secret_key,omitempty"`
}
