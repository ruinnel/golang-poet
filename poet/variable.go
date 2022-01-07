package poet

import "reflect"

type Variable struct {
	StructField bool
	Names       []string
	Type        *reflect.Type
	Value       Value
}

type Variables []Variable

func (v Variables) String() string {
	return GetTemplate().Variable(v)
}
