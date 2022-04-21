package types

type Type interface {
	String() string
}

type TypeName string

func (t TypeName) String() string {
	return string(t)
}
