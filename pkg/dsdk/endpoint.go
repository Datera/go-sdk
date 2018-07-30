package dsdk

import (
// "encoding/json"
// "errors"
// "strings"
)

type IEntity interface {
}

type IEndpoint interface {
	Create(interface{}) (IEntity, error)
	Update(interface{}) (IEntity, error)
	List(interface{}) ([]IEntity, error)
	Delete(interface{}) error
}
