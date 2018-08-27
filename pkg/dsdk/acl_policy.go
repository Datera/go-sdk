package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type AclPolicy struct {
	Path            string            `json:"path,omitempty" mapstructure:"path"`
	Initiators      []Initiator       `json:"initiators,omitempty" mapstructure:"initiators"`
	InitiatorGroups []InitiatorGroups `json:"initiator_groups,omitempty" mapstructure:"initiator_groups"`
	ctxt            context.Context
	conn            *ApiConnection
}

func newAclPolicy(ctxt context.Context, conn *ApiConnection, path string) *AclPolicy {
	return &AclPolicy{
		Path: _path.Join(path, "acl_policy"),
		ctxt: ctxt,
		conn: conn,
	}
}

type AclPolicyGetRequest struct {
}

type AclPolicyGetResponse AclPolicy

func (e *AclPolicy) Get(ro *AclPolicyGetRequest) (*AclPolicyGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Get(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AclPolicyGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil
}

type AclPolicySetRequest struct {
	Initiators      []Initiator       `json:"initiators,omitempty" mapstructure:"initiators"`
	InitiatorGroups []InitiatorGroups `json:"initiator_groups,omitempty" mapstructure:"initiator_groups"`
}

type AclPolicySetResponse AclPolicy

func (e *AclPolicy) Set(ro *AclPolicySetRequest) (*AclPolicySetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := e.conn.Put(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := &AclPolicySetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	resp.conn = e.conn
	resp.ctxt = e.ctxt
	return resp, nil

}
