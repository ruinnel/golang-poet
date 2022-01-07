package poet

type Constant struct {
	Name  string
	Value Value
}

type Constants []Constant

func (c Constants) String() string {
	return GetTemplate().Constant(c)
}
