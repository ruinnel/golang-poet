package types

import (
	"reflect"
)

type GenericType struct {
	Approximation bool
	Type          reflect.Type
}

type GenericTypes []GenericType

func (g GenericTypes) String() string {
	return GetTemplate().GenericType(g)
}
