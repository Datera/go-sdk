package dsdk

import (
	"encoding/json"
	"errors"
	"strings"
	// "fmt"
)

var conn *ApiConnection

type RootEndpoint struct {
	Path string
}

type IEndpoint interface {
	GetEndpoint(string) IEndpoint
	Create(bodyp ...string) (IEntity, error)
	List(queryp ...string) ([]IEntity, error)
}

type Endpoint struct {
	Path string
}

type IEntity interface {
	Get(string) interface{}
	GetEntities(string) []IEntity
	GetEndpoint(string) IEndpoint
	Reload() (IEntity, error)
	Set(bodyp ...string) (IEntity, error)
	Delete(bodyp ...string) error
}

type Entity struct {
	Path  string
	Items map[string]interface{}
}

func NewRootEndpoint(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*RootEndpoint, error) {
	var err error
	//Initialize global connection object
	conn, err = NewApiConnection(hostname, port, username, password, apiVersion, tenant, timeout, headers, secure)
	if err != nil {
		return nil, err
	}
	err = conn.Login()
	if err != nil {
		return nil, err
	}
	return &RootEndpoint{
		Path: "",
	}, nil

}

func NewEndpoint(parent, child string) IEndpoint {
	path := strings.Trim(strings.Join([]string{parent, child}, "/"), "/")
	return Endpoint{
		Path: path,
	}
}

func NewEntity(path string, items map[string]interface{}) IEntity {
	return Entity{
		Path:  path,
		Items: items,
	}
}

func (ep RootEndpoint) GetEndpoint(path string) IEndpoint {
	return NewEndpoint(ep.Path, path)
}

func (ep Endpoint) GetEndpoint(path string) IEndpoint {
	return NewEndpoint(ep.Path, path)
}

func (ep Endpoint) Create(bodyp ...string) (IEntity, error) {
	var en Entity
	r, _ := conn.Post(ep.Path, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	var i map[string]interface{}
	err = json.Unmarshal(d, &i)
	p := i["path"].(string)
	en = NewEntity(p, i).(Entity)
	if err != nil {
		panic(err)
	}
	return en, nil
}

func (ep Endpoint) List(queryp ...string) ([]IEntity, error) {
	ens := []IEntity{}
	r, _ := conn.Get(ep.Path, queryp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return ens, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	var j []interface{}
	err = json.Unmarshal(d, &j)
	if err != nil {
		panic(err)
	}
	for _, val := range j {
		var en Entity
		v := val.(map[string]interface{})
		p := v["path"].(string)
		en = NewEntity(p, v).(Entity)
		ens = append(ens, en)
	}
	return ens, nil
}

func (en Entity) Get(key string) interface{} {
	return en.Items[key]
}

func (en Entity) GetEndpoint(path string) IEndpoint {
	return NewEndpoint(en.Path, path)
}

func (en Entity) GetEntities(en_key string) []IEntity {
	ens := []IEntity{}
	eitems := en.Items[en_key].([]interface{})
	for _, i := range eitems {
		var en Entity
		v := i.(map[string]interface{})
		p := v["path"].(string)
		en = NewEntity(p, v).(Entity)
		ens = append(ens, en)
	}
	return ens
}

func (en Entity) Reload() (IEntity, error) {
	var n Entity
	r, _ := conn.Get(en.Path)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	var i map[string]interface{}
	err = json.Unmarshal(d, &i)
	if err != nil {
		return n, err
	}
	p := i["path"].(string)
	n = NewEntity(p, i).(Entity)
	return n, nil
}

func (en Entity) Set(bodyp ...string) (IEntity, error) {
	var n Entity
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, e, err := getData(r)
	if e.Message != "" {
		return n, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	var i map[string]interface{}
	err = json.Unmarshal(d, &i)
	if err != nil {
		return n, err
	}
	p := i["path"].(string)
	n = NewEntity(p, i).(Entity)
	return n, nil
}

func (en Entity) Delete(bodyp ...string) error {
	r, _ := conn.Delete(en.Path, bodyp...)
	_, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}
