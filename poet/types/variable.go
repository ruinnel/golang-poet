package types

type Variable struct {
	Names []string
	Type  Type
	Value Value
}

type Variables []Variable

func (v Variable) String() string {
	return GetTemplate().Variable(v)
}

func (v Variables) String() string {
	return GetTemplate().Variables(v)
}
