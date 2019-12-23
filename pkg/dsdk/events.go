package dsdk

import (
	"context"
	_path "path"

	greq "github.com/levigross/grequests"
)

type SystemEvent struct {
	Time        string `json:"time,omitempty" mapstructure:"time"`
	Code        string `json:"code,omitempty" mapstructure:"code"`
	Context     string `json:"context,omitempty" mapstructure:"context"`
	Debug       string `json:"debug,omitempty" mapstructure:"debug"`
	Description string `json:"description,omitempty" mapstructure:"description"`
	LastSeenTs  string `json:"last_seen_ts,omitempty" mapstructure:"last_seen_ts"`
	Message     string `json:"message,omitempty" mapstructure:"message"`
	RepeatCount int    `json:"repeat_count,omitempty" mapstructure:"repeat_count"`
	Severity    string `json:"severity,omitempty" mapstructure:"severity"`
	Tenant      string `json:"tenant,omitempty" mapstructure:"tenant"`
	Uuid        string `json:"uuid,omitempty" mapstructure:"uuid"`
	ObjectPath  string `json:"object_path,omitempty" mapstructure:"object_path"`
}

type SystemEventsRequest struct {
	Ctxt   context.Context `json:"-"`
	Params ListRangeParams `json:"params,omitempty"`
}

type SystemEvents struct {
	Path string
}

func newSystemEvents(path string) *SystemEvents {
	return &SystemEvents{
		Path: _path.Join(path, "events", "system"),
	}
}

func (e *SystemEvents) List(ro *SystemEventsRequest) ([]*SystemEvent, *ApiErrorResponse, error) {
	gro := &greq.RequestOptions{
		JSON:   ro,
		Params: ro.Params.ToMap(),
	}

	rs, apierr, err := GetConn(ro.Ctxt).GetList(ro.Ctxt, "/events/system", gro)
	if apierr != nil {
		return nil, apierr, err
	}

	if err != nil {
		return nil, nil, err
	}

	resp := []*SystemEvent{}
	for _, data := range rs.Data {
		elem := &SystemEvent{}
		edata := data.(map[string]interface{})
		if err = FillStruct(edata, elem); err != nil {
			return nil, nil, err
		}
		resp = append(resp, elem)
	}

	return resp, nil, nil
}
