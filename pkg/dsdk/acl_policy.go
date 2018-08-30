package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type AclPolicy struct {
	Path            string             `json:"path,omitempty" mapstructure:"path"`
	Initiators      []*Initiator       `json:"initiators,omitempty" mapstructure:"initiators"`
	InitiatorGroups []*InitiatorGroups `json:"initiator_groups,omitempty" mapstructure:"initiator_groups"`
}

func newAclPolicy(path string) *AclPolicy {
	return &AclPolicy{
		Path: _path.Join(path, "acl_policy"),
	}
}

type AclPolicyGetRequest struct {
	Ctxt context.Context `json:"-"`
}


func (e *AclPolicy) Get(ro *AclPolicyGetRequest) (*AclPolicy, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AclPolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type AclPolicySetRequest struct {
	Ctxt            context.Context    `json:"-"`
	Initiators      []*Initiator       `json:"initiators,omitempty" mapstructure:"initiators"`
	InitiatorGroups []*InitiatorGroups `json:"initiator_groups,omitempty" mapstructure:"initiator_groups"`
}


func (e *AclPolicy) Set(ro *AclPolicySetRequest) (*AclPolicy, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AclPolicy{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil

}
