package types

type Struct struct {
	Name               string
	ParameterizedTypes ParameterizedTypes
	Fields             StructFields
}

type Structs []Struct

func (s Struct) String() string {
	return GetTemplate().Struct(s)
}

func (s Structs) String() string {
	return GetTemplate().Structs(s)
}
