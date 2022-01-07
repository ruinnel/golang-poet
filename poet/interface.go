package poet

type Interface struct {
	Name      string
	Functions Functions
}

type Interfaces []Interface

func (i Interfaces) String() string {
	return GetTemplate().Interface(i)
}
