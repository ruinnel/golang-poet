package builder

import (
	"github.com/ruinnel/golang-poet/poet/types"
)

type FunctionBuilder struct {
	function types.Function
}

func NewFunctionBuilder(name string) *FunctionBuilder {
	return &FunctionBuilder{types.Function{Name: name}}
}

func (b *FunctionBuilder) WithName(name string) *FunctionBuilder {
	b.function.Name = name
	return b
}

func (b FunctionBuilder) Build() string {
	return b.function.String()
}
