package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

const (
	ProviderAWS    = "AWS S3"
	ProviderGoogle = "Google Cloud"
	ProviderS3     = "S3 Object Store"
)

type RemoteProvider struct {
	Path              string                   `json:"path,omitempty" mapstructure:"path"`
	Uuid              string                   `json:"uuid,omitempty" mapstructure:"uuid"`
	AccountId         string                   `json:"account_id,omitempty" mapstructure:"account_id"`
	RemoteType        string                   `json:"remote_type,omitempty" mapstructure:"remote_type"`
	LastSeenTimestamp string                   `json:"last_seen_timestamp,omitempty" mapstructure:"last_seen_timestamp"`
	Operations        []map[string]interface{} `json:"operations,omitempty" mapstructure:"operations"`
	Snapshots         []*Snapshot              `json:"snapshots,omitempty" mapstructure:"snapshots"`
	Label             string                   `json:"label,omitempty" mapstructure:"label"`
	Status            string                   `json:"status,omitempty" mapstructure:"status"`
	Host              string                   `json:"host,omitempty" mapstructure:"host"`
	Port              string                   `json:"port,omitempty" mapstructure:"port"`
	OperationsEp      string
	SnapshotsEp       *Snapshots

	// Present only when the RemoteProvider is a subresource of a snapshot. Indicates the replication state of the
	// snapshot on this RemoteProvider.
	OpState string
}

func RegisterRemoteProviderEndpoints(rp *RemoteProvider) {
	//a.OperationsEp = newOperations(a.Path)
	rp.SnapshotsEp = newSnapshots(rp.Path)
}

type RemoteProviders struct {
	Path string
}

type RemoteProvidersCreateRequest struct {
	Ctxt        context.Context `json:"-"`
	ProjectName string          `json:"project_name,omitempty" mapstructure:"project_name"`
	AccountId   string          `json:"account_id,omitempty" mapstructure:"account_id"`
	RemoteType  string          `json:"remote_type,omitempty" mapstructure:"remote_type"`
	PrivateKey  string          `json:"private_key,omitempty" mapstructure:"private_key"`
	Label       string          `json:"label,omitempty" mapstructure:"label"`
	Host        string          `json:"host,omitempty" mapstructure:"host"`
	Port        int             `json:"port,omitempty" mapstructure:"port"`
	AccessKey   string          `json:"access_key,omitempty" mapstructure:"access_key"`
	SecretKey   string          `json:"secret_key,omitempty" mapstructure:"secret_key"`
}

func newRemoteProviders(path string) *RemoteProviders {
	return &RemoteProviders{
		Path: _path.Join(path, "remote_providers"),
	}
}

func (e *RemoteProviders) Create(ro *RemoteProvidersCreateRequest) (*RemoteProvider, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Post(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &RemoteProvider{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterRemoteProviderEndpoints(resp)
	return resp, nil, nil
}

type RemoteProvidersListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListParams      `json:"params,omitempty"`
}

func (e *RemoteProviders) List(ro *RemoteProvidersListRequest) ([]*RemoteProvider, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params.ToMap()}
	rs, apierr, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := []*RemoteProvider{}
	for _, data := range rs.Data {
		elem := &RemoteProvider{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		RegisterRemoteProviderEndpoints(elem)
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

type RemoteProvidersGetRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string          `json:"-"`
}

func (e *RemoteProviders) Get(ro *RemoteProvidersGetRequest) (*RemoteProvider, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, _path.Join(e.Path, ro.Id), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &RemoteProvider{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterRemoteProviderEndpoints(resp)
	return resp, nil, nil
}

type RemoteProvidersRefreshRequest struct {
	Ctxt context.Context `json:"-"`
	Uuid string          `json:"-"`
}

type RemoteProvidersRefreshResponse struct {
	Uuid string `json:"uuid,omitempty" mapstructure:"uuid"`
}

func (e *RemoteProviders) Refresh(ro *RemoteProvidersRefreshRequest) (*RemoteProvidersRefreshResponse, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, _path.Join(e.Path, ro.Uuid, "refresh"), gro)

	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}

	resp := &RemoteProvidersRefreshResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}

	return resp, nil, nil
}

type RemoteProviderSetRequest struct {
	Ctxt        context.Context `json:"-"`
	ProjectName string          `json:"project_name,omitempty" mapstructure:"project_name"`
	AccountId   string          `json:"account_id,omitempty" mapstructure:"account_id"`
	PrivateKey  string          `json:"private_key,omitempty" mapstructure:"private_key"`
	Label       string          `json:"label,omitempty" mapstructure:"label"`
	Host        string          `json:"host,omitempty" mapstructure:"host"`
	Port        int             `json:"port,omitempty" mapstructure:"port"`
	AccessKey   string          `json:"access_key,omitempty" mapstructure:"access_key"`
	SecretKey   string          `json:"secret_key,omitempty" mapstructure:"secret_key"`
}

func (e *RemoteProvider) Set(ro *RemoteProviderSetRequest) (*RemoteProvider, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Put(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &RemoteProvider{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterRemoteProviderEndpoints(resp)
	return resp, nil, nil
}

type RemoteProviderDeleteRequest struct {
	Ctxt  context.Context `json:"-"`
	Force bool            `json:"force,omitempty" mapstructure:"force"`
}

func (e *RemoteProvider) Delete(ro *RemoteProviderDeleteRequest) (*RemoteProvider, *ApiErrorResponse, error) {
	rs, apierr, err := GetConn(ro.Ctxt).Delete(ro.Ctxt, e.Path, nil)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &RemoteProvider{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterRemoteProviderEndpoints(resp)
	return resp, nil, nil
}

type RemoteProviderAppTemplate struct {
	Path           string `json:"path,omitempty" mapstructure:"path"`
	ResolvedPath   string `json:"resolved_path,omitempty" mapstructure:"resolved_path"`
	ResolvedTenant string `json:"resolved_tenant,omitempty" mapstructure:"resolved_tenant"`
}

type RemoteProviderReloadRequest struct {
	Ctxt context.Context `json:"-"`
}

func (e *RemoteProvider) Reload(ro *RemoteProviderReloadRequest) (*RemoteProvider, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, apierr, err := GetConn(ro.Ctxt).Get(ro.Ctxt, e.Path, gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &RemoteProvider{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	RegisterRemoteProviderEndpoints(resp)
	return resp, nil, nil
}

type RemoteProviderOperationsSetRequest struct {
	Ctxt        context.Context `json:"-"`
	OperationId string          `json:"-"`
	Action      string          `json:"action"` //available options are 'clear' and 'abort'
}

type RemoteOperation struct {
	Path               string `json:"path" mapstructure:"path"`
	Uuid               string `json:"uuid" mapstructure:"uuid"`
	RemoteProviderUuid string `json:"remote_provider_uuid" mapstructure:"remote_provider_uuid"`
	AppInstanceUuid    string `json:"app_instance_uuid" mapstructure:"app_instance_uuid"`
	OpState            string `json:"op_state" mapstructure:"op_state"`
	OpType             string `json:"op_type" mapstructure:"op_type"`
	PercentDone        int    `json:"percent_done" mapstructure:"percent_done"`
	TotalTasksDone     int    `json:"total_tasks_done" mapstructure:"total_tasks_done"`
	TotalTasksIssued   int    `json:"total_tasks_issued" mapstructure:"total_tasks_issued"`
	References         struct {
		SnapshotAppInstancePath string `json:"snapshot_app_instance_path" mapstructure:"snapshot_app_instance_path"`
	} `json:"references" mapstructure:"references"`
}

func (e *RemoteProvider) SetOperation(ao *RemoteProviderOperationsSetRequest) (*RemoteOperation, *ApiErrorResponse, error) {

	gro := &greq.RequestOptions{JSON: ao}
	rs, apierr, err := GetConn(ao.Ctxt).Put(ao.Ctxt, _path.Join(e.Path, "operations", ao.OperationId), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &RemoteOperation{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}

	return resp, nil, nil
}
