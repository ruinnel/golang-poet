package poet

type Interface struct {
	Name               string
	ParameterizedTypes ParameterizedTypes
	GenericTypes       GenericTypes
	Functions          Functions
}

type Interfaces []Interface

func (i Interfaces) String() string {
	return GetTemplate().Interface(i)
}
