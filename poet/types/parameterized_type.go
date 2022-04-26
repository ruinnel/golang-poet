package types

import (
	"reflect"
)

type ParameterizedType struct {
	Symbol        string
	Approximation bool
	Type          reflect.Type
}

type ParameterizedTypes []ParameterizedType

func (p ParameterizedType) ToType() Type {
	return TypeName(p.Symbol)
}

func (p ParameterizedType) String() string {
	return GetTemplate().ParameterizedType(p)
}

func (p ParameterizedTypes) String() string {
	return GetTemplate().ParameterizedTypes(p)
}

func (p ParameterizedTypes) NeedType(idx int) bool {
	if len(p) <= idx+1 {
		return true
	}
	if len(p) > idx+1 {
		this := p[idx]
		next := p[idx+1]
		return this.Type.Kind() != next.Type.Kind()
	}
	return false
}
