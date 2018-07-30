package dsdk

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	udc "github.com/Datera/go-udc/pkg/udc"
	log "github.com/sirupsen/logrus"
	// greq "github.com/levigross/grequests"
)

func TestInitiatorCreate(t *testing.T) {
	conf, err := udc.GetConfig()
	if err != nil {
		t.Errorf("%s", err)
	}
	ctxt := context.Background()
	conn := NewApiConnection(ctxt, conf, true)
	initEp := newInitiators(ctxt, conn)
	ro := &InitiatorsCreateRequest{
		Id:   fmt.Sprintf("iqn.1993-08.org.debian:01:%s", RandString(12)),
		Name: RandString(12),
	}
	var resp *InitiatorsCreateResponse
	if resp, err = initEp.Create(ro); err != nil {
		t.Errorf("%s", err)
	}
	init := Initiator(*resp)
	if init.Name == "" || init.Id == "" || init.conn == nil {
		t.Errorf("InitiatorsCreateResponse object not populated: %#v", init)
	}
	if _, err := init.Delete(nil); err != nil {
		t.Errorf("%s", err)
	}
}

type MyFormatter struct {
}

func (f *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	msg := entry.Message
	level := entry.Level
	t := entry.Time
	return []byte(fmt.Sprintf("%s %s %s", t.Format(time.RFC3339), strings.ToUpper(level.String()), string(msg))), nil
}

func TestMain(m *testing.M) {
	log.SetFormatter(&MyFormatter{})
	log.SetLevel(log.DebugLevel)
	m.Run()
}
