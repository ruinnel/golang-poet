package types

import (
	"fmt"
)

type Value struct {
	Value interface{}
}

func (v Value) String() string {
	switch val := v.Value.(type) {
	case string:
		return fmt.Sprintf("\"%s\"", val)
	case float64:
		return fmt.Sprintf("%f", val)
	case float32:
		return fmt.Sprintf("%f", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

type Expression string
