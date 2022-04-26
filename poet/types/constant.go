package types

type Constant struct {
	Name  string
	Value Value
}

type Constants []Constant

func (c Constant) String() string {
	return GetTemplate().Constant(c)
}

func (c Constants) String() string {
	return GetTemplate().Constants(c)
}
