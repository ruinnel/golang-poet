package types

import (
	"fmt"
	"strings"
)

type Receiver struct {
	Name string
	Type Type
}

type Statements []CodeBlock

func (s Statements) String() string {
	var buf []string
	for _, code := range s {
		buf = append(buf, fmt.Sprintf("  %s", code.String()))
	}
	return strings.Join(buf, "\n")
}

type Function struct {
	Interface          bool
	Receiver           *Receiver
	Name               string
	ParameterizedTypes ParameterizedTypes
	Arguments          Arguments
	Returns            Returns
	Statements         Statements
}

type Functions []Function

func (f Functions) String() string {
	return GetTemplate().Function(f)
}
