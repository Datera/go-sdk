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
	GetM() map[string]interface{}
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

func newEp(parent, child string) IEndpoint {
	path := strings.Trim(strings.Join([]string{parent, child}, "/"), "/")
	return Endpoint{
		Path: path,
	}
}

func newEntity(path string, items map[string]interface{}) IEntity {
	return Entity{
		Path:  path,
		Items: items,
	}
}

func (ep Endpoint) GetEp(path string) IEndpoint {
	return newEp(ep.Path, path)
}

func (ep Endpoint) GetPath() string {
	return ep.Path
}

// Create an Entity via this Endpoint.  The IEntity object can be unmarshalled
// into the matching Entity from the entity.go file
// bodyp arguments can be in one of two forms
//
// 1. Vararg strings follwing this pattern: "key=value"
//    These strings have a limitation in that they cannot be arbitrarily nested
//    JSON values, instead they must be simple strings
//    Eg.  "key=value" is fine, but `key=["some", "list"]` will fail
//    the arbitrary JSON usecase is handled by #2
//
// 2. A single map[string]interface{} argument.  This handles the case where
//    we need to send arbitrarily nested JSON as an argument
//
// Function arguments are setup this way to provide an easy way to handle 90%
// of the use cases (where we're just passing key, value string pairs) but that
// remaining 10% we need to pass something more complex
func (ep Endpoint) Create(bodyp ...interface{}) (IEntity, error) {
	// We actually create a concrete entity to return in failure conditions
	// so it can be deleted without nul pointer panics
	en := Entity{}
	conn := Cpool.getConn()
	defer Cpool.releaseConn(conn)
	r, _ := conn.post(ep.Path, bodyp...)
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
	en = newEntity(p, i).(Entity)
	if err != nil {
		panic(err)
	}
	return en, nil
}

// List the Entities hosted on this Endpoint.  The IEntity objects can be unmarshalled
// into the matching Entity from the entity.go file
func (ep Endpoint) List(queryp ...string) ([]IEntity, error) {
	ens := []IEntity{}
	conn := Cpool.getConn()
	defer Cpool.releaseConn(conn)
	r, _ := conn.get(ep.Path, queryp...)
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
		en := newEntity(p, v).(Entity)
		ens = append(ens, en)
	}
	return ens, nil
}

// Get the Entity hosted on this Endpoint.  The IEntity object can be unmarshalled
// into the matching Entity from the entity.go file
func (ep Endpoint) Get(queryp ...string) (IEntity, error) {
	// We actually create a concrete entity to return in failure conditions
	// so it can be deleted without nul pointer panics
	en := Entity{}
	conn := Cpool.getConn()
	defer Cpool.releaseConn(conn)
	r, _ := conn.get(ep.Path, queryp...)
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
	en = newEntity(p, i).(Entity)
	return en, nil
}

// Update attributes of the Entity hosted on this Endpoint.  The IEntity object can be
// unmarshalled into the matching Entity from the entity.go file
// bodyp arguments can be in one of two forms
//
// 1. Vararg strings follwing this pattern: "key=value"
//    These strings have a limitation in that they cannot be arbitrarily nested
//    JSON values, instead they must be simple strings
//    Eg.  "key=value" is fine, but `key=["some", "list"]` will fail
//    the arbitrary JSON usecase is handled by #2
//
// 2. A single map[string]interface{} argument.  This handles the case where
//    we need to send arbitrarily nested JSON as an argument
//
// Function arguments are setup this way to provide an easy way to handle 90%
// of the use cases (where we're just passing key, value string pairs) but that
// remaining 10% we need to pass something more complex
func (ep Endpoint) Set(bodyp ...interface{}) (IEntity, error) {
	// We actually create a concrete entity to return in failure conditions
	// so it can be deleted without nul pointer panics
	en := Entity{}
	conn := Cpool.getConn()
	defer Cpool.releaseConn(conn)
	r, _ := conn.put(ep.Path, false, bodyp...)
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
	en = newEntity(p, i).(Entity)
	return en, nil
}

// Access a key in this entity by name.  Must be type asserted after
// recieving it
func (en Entity) Get(key string) interface{} {
	return en.Items[key]
}

// Short for "Get Map", returns the entire map of this Entity
func (en Entity) GetM() map[string]interface{} {
	return en.Items
}

// Short for "Get Bytes".  Used for Unmarshalling into an Entity from the
// entity.go file
func (en Entity) GetB() []byte {
	b, err := json.Marshal(en.Items)
	if err != nil {
		panic(err)
	}
	return b
}

// Short for "Get Endpoint".  Does not make any request to the backend, just
// used for constructing the path to that Endpoint.  Can be chained together.
func (en Entity) GetEp(path string) IEndpoint {
	return newEp(en.Path, path)
}

// Returns the full path to this Entity
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
//      unpackedSi.UnpackB(ai.GetB())
//      fmt.Println("SI Name", unpackedSI.Name)
func (en Entity) GetEn(enKey string) []IEntity {
	ens := []IEntity{}
	eitems := en.Items[enKey].([]interface{})
	for _, i := range eitems {
		v := i.(map[string]interface{})
		p := v["path"].(string)
		n := newEntity(p, v).(Entity)
		ens = append(ens, n)
	}
	return ens
}

// Pull all the attributes of this Entity.  Useful if it has been changed at
// some point and a newly updated version of the Entity is needed
func (en Entity) Reload() (IEntity, error) {
	// We actually create a concrete entity to return in failure conditions
	// so it can be deleted without nul pointer panics
	n := Entity{}
	conn := Cpool.getConn()
	defer Cpool.releaseConn(conn)
	r, _ := conn.get(en.Path)
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
	n = newEntity(p, i).(Entity)
	return n, nil
}

// Update attributes of this Entity.  The IEntity object can be
// unmarshalled into the matching Entity from the entity.go file
// bodyp arguments can be in one of two forms
//
// 1. Vararg strings follwing this pattern: "key=value"
//    These strings have a limitation in that they cannot be arbitrarily nested
//    JSON values, instead they must be simple strings
//    Eg.  "key=value" is fine, but `key=["some", "list"]` will fail
//    the arbitrary JSON usecase is handled by #2
//
// 2. A single map[string]interface{} argument.  This handles the case where
//    we need to send arbitrarily nested JSON as an argument
//
// Function arguments are setup this way to provide an easy way to handle 90%
// of the use cases (where we're just passing key, value string pairs) but that
// remaining 10% we need to pass something more complex
func (en Entity) Set(bodyp ...interface{}) (IEntity, error) {
	// We actually create a concrete entity to return in failure conditions
	// so it can be deleted without nul pointer panics
	n := Entity{}
	conn := Cpool.getConn()
	defer Cpool.releaseConn(conn)
	r, _ := conn.put(en.Path, false, bodyp...)
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
	n = newEntity(p, i).(Entity)
	return n, nil
}

// Delete this Entity.
// bodyp arguments can be in one of two forms
//
// 1. Vararg strings follwing this pattern: "key=value"
//    These strings have a limitation in that they cannot be arbitrarily nested
//    JSON values, instead they must be simple strings
//    Eg.  "key=value" is fine, but `key=["some", "list"]` will fail
//    the arbitrary JSON usecase is handled by #2
//
// 2. A single map[string]interface{} argument.  This handles the case where
//    we need to send arbitrarily nested JSON as an argument
//
// Function arguments are setup this way to provide an easy way to handle 90%
// of the use cases (where we're just passing key, value string pairs) but that
// remaining 10% we need to pass something more complex
func (en Entity) Delete(bodyp ...interface{}) error {
	conn := Cpool.getConn()
	defer Cpool.releaseConn(conn)
	r, _ := conn.delete(en.Path, bodyp...)
	_, _, e, err := getData(r)
	if e.Message != "" {
		return errors.New(strings.Join(append([]string{e.Message}, e.Errors...), ":"))
	}
	if err != nil {
		panic(err)
	}
	return nil
}
