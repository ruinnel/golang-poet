package types

type Interface struct {
	Name               string
	ParameterizedTypes ParameterizedTypes
	GenericTypes       GenericTypes
	Functions          Functions
}

type Interfaces []Interface

func (i Interface) String() string {
	return GetTemplate().Interface(i)
}

func (i Interfaces) String() string {
	return GetTemplate().Interfaces(i)
}
