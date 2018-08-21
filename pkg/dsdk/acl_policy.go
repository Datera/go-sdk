package dsdk

type AclPolicy struct {
	Path            string           `json:"path,omitempty"`
	Initiators      []Initiator      `json:"initiators,omitempty"`
	InitiatorGroups []InitiatorGroup `json:"initiator_groups,omitempty"`
}
