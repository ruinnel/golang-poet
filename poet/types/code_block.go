package types

type CodeBlock struct {
	Code string
}

func (b CodeBlock) String() string {
	return b.Code
}
