package types

type Struct struct {
	Name               string
	ParameterizedTypes ParameterizedTypes
	Fields             StructFields
}

type Structs []Struct

func (s Structs) String() string {
	return GetTemplate().Struct(s)
}
