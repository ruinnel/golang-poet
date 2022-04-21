package types

type ImportSpec struct {
	Package string
	Alias   string
}

type Imports []ImportSpec

func (i Imports) String() string {
	return GetTemplate().Import(i)
}
