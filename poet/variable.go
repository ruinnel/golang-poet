package poet

type Variable struct {
	StructField bool
	Names       []string
	Type        Type
	Value       Value
}

type Variables []Variable

func (v Variables) String() string {
	return GetTemplate().Variable(v)
}
