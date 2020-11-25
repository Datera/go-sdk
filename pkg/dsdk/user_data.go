package dsdk

import (
	"context"
	greq "github.com/levigross/grequests"
	_path "path"
)

type UserData struct {
	AppInstanceId  string                  `json:"app_instance_id"`
	Data           map[string]interface{}  `json:"data"`
}

type UserDatas struct {
	Path string
}

func newUserDatas(path string) *UserDatas {
	return &UserDatas{
		Path: _path.Join(path, "user_data"),
	}
}

type UserDataUpdateRequest struct {
	Ctxt                 context.Context        `json:"-"`
	AppInstanceId        string                 `json:"app_instance_id" mapstructure:"app_instance_id"`
	Data                 map[string]interface{} `json:"data" mapstructure:"data"`
}

// Update adds a JSON User Data Record to an App Instance
func (e *UserDatas) Update(ud *UserDataUpdateRequest) (*UserData, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ud}
	rs, apierr, err := GetConn(ud.Ctxt).Put(ud.Ctxt, _path.Join("app_instances", ud.AppInstanceId, e.Path), gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := &UserData{
		AppInstanceId: ud.AppInstanceId,
	}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}

// UserDatasListRequest lists all custom user data on all apps within a tenant
// Params is the normal ListParams, but Sort isn't used/supported.
type UserDatasListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListParams      `json:"params,omitempty"`
}

// List shows all UserData that have been stored
// it can be filtered via a Glob search in ro.Filter field
func (e *UserDatas) List(udlr *UserDatasListRequest) ([]*UserData, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   udlr,
		Params: udlr.Params.ToMap()}
	rs, apierr, err := GetConn(udlr.Ctxt).GetList(udlr.Ctxt, "app_instance_user_data", gro)
	if apierr != nil {
		return nil, apierr, err
	}
	if err != nil {
		return nil, nil, err
	}
	resp := []*UserData{}
	for _, data := range rs.Data {
		elem := &UserData{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
	}
	return resp, nil, nil
}

// UserDataGetRequest gets one AppInstance's uploaded user data
type UserDataGetRequest struct {
	Ctxt context.Context     `json:"-"`
	AppInstanceId string     `json:"app_instance_id"`
}

// Get returns an individual JSON UserData object attached to an AppInstance
func (e *UserDatas) Get(ud *UserDataGetRequest) (*UserData, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{JSON: ud}
	rs, apierr, err := GetConn(ud.Ctxt).Get(ud.Ctxt, _path.Join("app_instances", e.Path, ud.AppInstanceId), gro)
	if apierr != nil || err != nil {
		return nil, apierr, err
	}

	resp := &UserData{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, nil, err
	}
	return resp, nil, nil
}
