package poet

import (
	"reflect"
	"strings"
)

type Receiver struct {
	Name string
	Type reflect.Type
}

type Statements []CodeBlock

func (s Statements) String() string {
	var buf []string
	for _, code := range s {
		buf = append(buf, code.String())
	}
	return strings.Join(buf, "\n")
}

type Function struct {
	Interface  bool
	Receiver   *Receiver
	Name       string
	Arguments  Arguments
	Returns    Returns
	Statements Statements
}

type Functions []Function

func (f Functions) String() string {
	return GetTemplate().Function(f)
}
