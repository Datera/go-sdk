package dsdk

import (
	"encoding/json"
	"errors"
	"strings"
)

var conn IAPIConnection

type RootEp struct {
	Path string
}

type IEndpoint interface {
	GetEp(string) IEndpoint
	Create(bodyp ...interface{}) (IEntity, error)
	List(queryp ...string) ([]IEntity, error)
	Set(bodyp ...interface{}) (IEntity, error)
}

type Endpoint struct {
	Path string
}

type IEntity interface {
	Get(string) interface{}
	GetEn(string) []IEntity
	GetEp(string) IEndpoint
	Reload() (IEntity, error)
	Set(bodyp ...interface{}) (IEntity, error)
	Delete(bodyp ...interface{}) error
}

type Entity struct {
	Path  string
	Items map[string]interface{}
}

func NewRootEp(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*RootEp, error) {
	var err error
	//Initialize global connection object
	conn, err = NewAPIConnection(hostname, port, username, password, apiVersion, tenant, timeout, headers, secure)
	if err != nil {
		return nil, err
	}
	err = conn.Login()
	if err != nil {
		return nil, err
	}
	return &RootEp{
		Path: "",
	}, nil

}

func NewEp(parent, child string) IEndpoint {
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

func (ep RootEp) GetEp(path string) IEndpoint {
	return NewEp(ep.Path, path)
}

func (ep Endpoint) GetEp(path string) IEndpoint {
	return NewEp(ep.Path, path)
}

func (ep Endpoint) Create(bodyp ...interface{}) (IEntity, error) {
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

func (ep Endpoint) Set(bodyp ...interface{}) (IEntity, error) {
	var n Entity
	r, _ := conn.Put(ep.Path, false, bodyp...)
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

func (en Entity) Get(key string) interface{} {
	return en.Items[key]
}

func (en Entity) GetEp(path string) IEndpoint {
	return NewEp(en.Path, path)
}

func (en Entity) GetEn(enKey string) []IEntity {
	ens := []IEntity{}
	eitems := en.Items[enKey].([]interface{})
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

func (en Entity) Set(bodyp ...interface{}) (IEntity, error) {
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

func (en Entity) Delete(bodyp ...interface{}) error {
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
