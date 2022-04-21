package builder

import "github.com/ruinnel/golang-poet/poet/types"

type ConstantBuilder struct {
	constants types.Constants
}

func NewConstantBuilder() *ConstantBuilder {
	var constants types.Constants
	return &ConstantBuilder{constants}
}

func (b *ConstantBuilder) AddConstant(name string, value types.Value) *ConstantBuilder {
	b.constants = append(b.constants, types.Constant{Name: name, Value: value})
	return b
}

func (b *ConstantBuilder) WithConstants(Constants ...types.Constant) *ConstantBuilder {
	b.constants = Constants
	return b
}

func (b ConstantBuilder) Build() string {
	return b.constants.String()
}
