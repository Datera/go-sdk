package dsdk

import (
	"context"
	"fmt"
	"testing"

	udc "github.com/Datera/go-udc/pkg/udc"
	// log "github.com/sirupsen/logrus"
	// greq "github.com/levigross/grequests"
)

func TestInitiatorsCreateDelete(t *testing.T) {
	conf, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	ctxt := context.Background()
	conn := NewApiConnection(ctxt, conf, true)
	initEp := newInitiators(ctxt, conn, "/")
	ro := &InitiatorsCreateRequest{
		Id:   fmt.Sprintf("iqn.1993-08.org.debian:01:%s", RandString(12)),
		Name: RandString(12),
	}
	var resp *InitiatorsCreateResponse
	if resp, err = initEp.Create(ro); err != nil {
		t.Errorf("%s", err)
	}
	init := Initiator(*resp)
	if init.Name == "" || init.Id == "" || init.conn == nil || init.ctxt == nil {
		t.Errorf("InitiatorsCreateResponse object not populated: %#v", init)
	}
	if _, err := init.Delete(nil); err != nil {
		t.Errorf("%s", err)
	}
}

func TestInitiatorsList(t *testing.T) {
	conf, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	ctxt := context.Background()
	conn := NewApiConnection(ctxt, conf, true)

	// Create Initiators
	inits := []Initiator{}
	initEp := newInitiators(ctxt, conn, "/")
	for i := 0; i < 2; i++ {
		ro := &InitiatorsCreateRequest{
			Id:   fmt.Sprintf("iqn.1993-08.org.debian:01:%s", RandString(12)),
			Name: RandString(12),
		}
		var resp *InitiatorsCreateResponse
		if resp, err = initEp.Create(ro); err != nil {
			t.Errorf("%s", err)
		}
		init := Initiator(*resp)
		inits = append(inits, init)

	}
	// List initiators, limit to 2 in case more already exist
	ro := &InitiatorsListRequest{Params: map[string]string{"limit": "2"}}
	var resp *InitiatorsListResponse
	resp, err = initEp.List(ro)

	if len(*resp) != 2 {
		t.Errorf("List returned incorrect number of initiators, %#v", resp)
	}

	// Clean up initiators
	for _, init := range inits {
		if _, err := init.Delete(nil); err != nil {
			t.Errorf("%s", err)
		}
	}
}

func TestInitiatorsGet(t *testing.T) {
	conf, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	ctxt := context.Background()
	conn := NewApiConnection(ctxt, conf, true)
	initEp := newInitiators(ctxt, conn, "/")

	// Create initiator
	ro := &InitiatorsCreateRequest{
		Id:   fmt.Sprintf("iqn.1993-08.org.debian:01:%s", RandString(12)),
		Name: RandString(12),
	}
	if _, err = initEp.Create(ro); err != nil {
		t.Errorf("%s", err)
	}

	// Get Initiator
	rog := &InitiatorsGetRequest{Id: ro.Id}
	resp := &InitiatorsGetResponse{}
	if resp, err = initEp.Get(rog); err != nil {
		t.Errorf("%s", err)
	}
	init := Initiator(*resp)
	if init.Name == "" || init.Id == "" || init.conn == nil {
		t.Errorf("InitiatorsGetResponse object not populated: %#v", init)
	}

	// Delete Initiator
	if _, err := init.Delete(nil); err != nil {
		t.Errorf("%s", err)
	}

}

func TestInitiatorsSet(t *testing.T) {
	conf, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	ctxt := context.Background()
	conn := NewApiConnection(ctxt, conf, true)
	initEp := newInitiators(ctxt, conn, "/")

	// Create initiator
	ro := &InitiatorsCreateRequest{
		Id:   fmt.Sprintf("iqn.1993-08.org.debian:01:%s", RandString(12)),
		Name: RandString(12),
	}
	var resp *InitiatorsCreateResponse
	if resp, err = initEp.Create(ro); err != nil {
		t.Errorf("%s", err)
	}
	init := Initiator(*resp)

	// Change Initiator
	ros := &InitiatorSetRequest{Name: "my-test-init"}
	var resps *InitiatorSetResponse
	if resps, err = init.Set(ros); err != nil {
		t.Errorf("%s", err)
	}
	init = Initiator(*resps)

	if init.Name != ros.Name {
		t.Errorf("Initiator name was not updated. %s != %s", init.Name, ros.Name)
	}

	// Delete Initiator
	if _, err := init.Delete(nil); err != nil {
		t.Errorf("%s", err)
	}
}
