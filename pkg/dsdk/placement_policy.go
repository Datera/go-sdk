package dsdk

import (
	"encoding/json"
)

type PlacementPolicy struct {
	Path           string `json:"path,omitempty" mapstructure:"path"`
	ResolvedPath   string `json:"resolved_path,omitempty" mapstructure:"resolved_path"`
	ResolvedTenant string `json:"resolved_tenant,omitempty" mapstructure:"resolved_tenant"`
}

func (p PlacementPolicy) MarshalJSON() ([]byte, error) {
	if p.Path == "" && p.ResolvedTenant == "" {
		return []byte(p.ResolvedPath), nil
	}
	m := map[string]string{
		"path":            p.Path,
		"resolved_path":   p.ResolvedPath,
		"resolved_tenant": p.ResolvedTenant,
	}
	return json.Marshal(m)
}

func (p PlacementPolicy) UnmarshalJSON(b []byte) error {
	np := map[string]string{}
	err := json.Unmarshal(b, &np)
	if err != nil {
		p.Path = ""
		p.ResolvedPath = string(b)
		p.ResolvedTenant = ""
	} else {
		p.Path = np["path"]
		p.ResolvedPath = np["resolved_path"]
		p.ResolvedTenant = np["resolved_tenant"]
	}
	return nil
}
