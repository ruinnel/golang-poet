package types

import (
	"reflect"
)

type Return struct {
	Name string
	Type reflect.Type
}

type Returns []Return

func (r Return) String() string {
	return GetTemplate().Return(r)
}

func (r Returns) String() string {
	return GetTemplate().Returns(r)
}

func (r Returns) NeedType(idx int) bool {
	if len(r) <= idx+1 {
		return true
	}
	if len(r) > idx+1 {
		this := r[idx]
		next := r[idx+1]
		return this.Type.Kind() != next.Type.Kind()
	}
	return false
}
