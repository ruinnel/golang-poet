package poet

type Struct struct {
	Name      string
	Variables Variables
}

type Structs []Struct

func (s Structs) String() string {
	return GetTemplate().Struct(s)
}
