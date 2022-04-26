package builder

import (
	"github.com/ruinnel/golang-poet/poet/types"
)

type InterfaceBuilder struct {
	interfaces types.Interfaces
}

func NewInterfaceBuilder() *InterfaceBuilder {
	var interfaces types.Interfaces
	return &InterfaceBuilder{interfaces}
}

func (b *InterfaceBuilder) AddInterface(name string, parameterizedTypes types.ParameterizedTypes, genericTypes types.GenericTypes, functions types.Functions) *InterfaceBuilder {
	b.interfaces = append(b.interfaces, types.Interface{
		Name:               name,
		ParameterizedTypes: parameterizedTypes,
		GenericTypes:       genericTypes,
		Functions:          functions,
	})
	return b
}

func (b *InterfaceBuilder) WithInterfaces(interfaces ...types.Interface) *InterfaceBuilder {
	b.interfaces = interfaces
	return b
}

func (b InterfaceBuilder) Build() string {
	return b.interfaces.String()
}
