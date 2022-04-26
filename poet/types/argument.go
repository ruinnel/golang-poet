package types

type Argument struct {
	VariableNumber bool
	Name           string
	Type           Type
}

type Arguments []Argument

func (a Argument) String() string {
	return GetTemplate().Argument(a)
}

func (a Arguments) String() string {
	return GetTemplate().Arguments(a)
}

func (a Arguments) NeedType(idx int) bool {
	if len(a) <= idx+1 {
		return true
	}
	if len(a) > idx+1 {
		this := a[idx]
		next := a[idx+1]
		return this.Type.String() != next.Type.String()
	}
	return false
}
