package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type Subsystem struct {
	Path        string   `json:"path,omitempty" mapstructure:"path"`
	Causes      []string `json:"causes,omitempty" mapstructure:"causes"`
	Fan         string   `json:"fan,omitempty" mapstructure:"fan"`
	Health      string   `json:"health,omitempty" mapstructure:"health"`
	Network     string   `json:"network,omitempty" mapstructure:"network"`
	Power       string   `json:"power,omitempty" mapstructure:"power"`
	Temperature string   `json:"temperature,omitempty" mapstructure:"temperature"`
	Voltage     string   `json:"voltage,omitempty" mapstructure:"voltage"`
}

type Subsystems struct {
	Path string
}

func newSubsystems(path string) *Subsystems {
	return &Subsystems{
		Path: _path.Join(path, "subsystem_states"),
	}
}

type SubsystemsListRequest struct {
	Ctxt   context.Context `json:"-"`
	Params map[string]string
}

type SubsystemsListResponse []Subsystem

func (e *Subsystems) List(ro *SubsystemsListRequest) (*SubsystemsListResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params}
	rs, err := GetConn(ro.Ctxt).GetList(e.Path, gro)
	if err != nil {
		return nil, err
	}
	resp := SubsystemsListResponse{}
	for _, data := range rs.Data {
		elem := &Subsystem{}
		adata := data.(map[string]interface{})
		if err = FillStruct(adata, elem); err != nil {
			return nil, err
		}
		resp = append(resp, *elem)
	}
	return &resp, nil
}

type SubsystemsGetRequest struct {
	Ctxt context.Context `json:"-"`
	Id   string
}

type SubsystemsGetResponse Subsystem

func (e *Subsystems) Get(ro *SubsystemsGetRequest) (*SubsystemsGetResponse, error) {
	gro := &greq.RequestOptions{JSON: ro}
	rs, err := GetConn(ro.Ctxt).Get(_path.Join(e.Path, ro.Id), gro)
	if err != nil {
		return nil, err
	}
	resp := &SubsystemsGetResponse{}
	if err = FillStruct(rs.Data, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
