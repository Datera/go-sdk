package dsdk

import (
	"encoding/json"
	"errors"
	"strings"
)

type IEndpoint interface {
	GetEp(string) IEndpoint
	Create(bodyp ...interface{}) (IEntity, error)
	List(queryp ...string) ([]IEntity, error)
	Set(bodyp ...interface{}) (IEntity, error)
	Get(queryp ...string) (IEntity, error)
	GetPath() string
}

type Endpoint struct {
	Path string
}

type IEntity interface {
	Get(string) interface{}
	GetA() map[string]interface{}
	GetB() []byte
	GetEn(string) []IEntity
	GetEp(string) IEndpoint
	GetPath() string
	Reload() (IEntity, error)
	Set(bodyp ...interface{}) (IEntity, error)
	Delete(bodyp ...interface{}) error
}

type Entity struct {
	Path  string
	Items map[string]interface{}
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

func (ep Endpoint) GetEp(path string) IEndpoint {
	return NewEp(ep.Path, path)
}

func (ep Endpoint) GetPath() string {
	return ep.Path
}

func (ep Endpoint) Create(bodyp ...interface{}) (IEntity, error) {
	// We actually create a concrete entity to return in failure conditions
	// so it can be deleted without nul pointer panics
	en := Entity{}
	conn := Cpool.GetConn()
	defer Cpool.ReleaseConn(conn)
	r, _ := conn.Post(ep.Path, bodyp...)
	d, _, e, err := getData(r)
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
	conn := Cpool.GetConn()
	defer Cpool.ReleaseConn(conn)
	r, _ := conn.Get(ep.Path, queryp...)
	d, _, e, err := getData(r)
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
		v := val.(map[string]interface{})
		p := v["path"].(string)
		en := NewEntity(p, v).(Entity)
		ens = append(ens, en)
	}
	return ens, nil
}

func (ep Endpoint) Get(queryp ...string) (IEntity, error) {
	// We actually create a concrete entity to return in failure conditions
	// so it can be deleted without nul pointer panics
	en := Entity{}
	conn := Cpool.GetConn()
	defer Cpool.ReleaseConn(conn)
	r, _ := conn.Get(ep.Path, queryp...)
	d, _, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	var i map[string]interface{}
	err = json.Unmarshal(d, &i)
	if err != nil {
		return en, err
	}
	p := i["path"].(string)
	en = NewEntity(p, i).(Entity)
	return en, nil
}

func (ep Endpoint) Set(bodyp ...interface{}) (IEntity, error) {
	// We actually create a concrete entity to return in failure conditions
	// so it can be deleted without nul pointer panics
	en := Entity{}
	conn := Cpool.GetConn()
	defer Cpool.ReleaseConn(conn)
	r, _ := conn.Put(ep.Path, false, bodyp...)
	d, _, e, err := getData(r)
	if e.Message != "" {
		return en, errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	var i map[string]interface{}
	err = json.Unmarshal(d, &i)
	if err != nil {
		return en, err
	}
	p := i["path"].(string)
	en = NewEntity(p, i).(Entity)
	return en, nil
}

func (en Entity) Get(key string) interface{} {
	return en.Items[key]
}

// Short for "Get All"
func (en Entity) GetA() map[string]interface{} {
	return en.Items
}

// Short for "Get Bytes"
func (en Entity) GetB() []byte {
	b, err := json.Marshal(en.Items)
	if err != nil {
		panic(err)
	}
	return b
}

// Short for "Get Endpoint"
func (en Entity) GetEp(path string) IEndpoint {
	return NewEp(en.Path, path)
}

func (en Entity) GetPath() string {
	return en.Path
}

// Short for "Get Entities"
// Usage:
//		ai, _ = client.GetEp("app_instances").Create("name=my-app")
// 		ai.GetEp("storage_instances").Create("name=my-stor")
//		ai, _ = ai.Reload()
//		si = ai.GetEn("storage_instances")[0]
//      // Optionally you could unpack the IEntity object into a struct
//      var unpackedSI StorageInstance
//      json.Unmarshal(ai.GetB(), &unpackedSI)
//      fmt.Println("SI Name", unpackedSI.Name)
func (en Entity) GetEn(enKey string) []IEntity {
	ens := []IEntity{}
	eitems := en.Items[enKey].([]interface{})
	for _, i := range eitems {
		v := i.(map[string]interface{})
		p := v["path"].(string)
		n := NewEntity(p, v).(Entity)
		ens = append(ens, n)
	}
	return ens
}

func (en Entity) Reload() (IEntity, error) {
	// We actually create a concrete entity to return in failure conditions
	// so it can be deleted without nul pointer panics
	n := Entity{}
	conn := Cpool.GetConn()
	defer Cpool.ReleaseConn(conn)
	r, _ := conn.Get(en.Path)
	d, _, e, err := getData(r)
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
	// We actually create a concrete entity to return in failure conditions
	// so it can be deleted without nul pointer panics
	n := Entity{}
	conn := Cpool.GetConn()
	defer Cpool.ReleaseConn(conn)
	r, _ := conn.Put(en.Path, false, bodyp...)
	d, _, e, err := getData(r)
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
	conn := Cpool.GetConn()
	defer Cpool.ReleaseConn(conn)
	r, _ := conn.Delete(en.Path, bodyp...)
	_, _, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}
