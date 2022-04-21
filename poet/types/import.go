package types

type Import struct {
	Package string
	Alias   string
}

type Imports []Import

func (i Imports) String() string {
	return GetTemplate().Import(i)
}
