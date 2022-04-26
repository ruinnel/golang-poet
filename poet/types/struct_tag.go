package types

import (
	"bytes"
	"fmt"
)

type StructTag struct {
	Name       string `json:"omitempty"`
	Parameters Parameters
}

type Parameters map[string]string
type StructTags []StructTag

func (p Parameters) String() string {
	code := bytes.NewBufferString("")
	for key, value := range p {
		if code.Len() > 0 {
			code.WriteString(",")
		}
		code.WriteString(key)
		if len(value) > 0 {
			code.WriteString(fmt.Sprintf("=%v", value))
		}
	}

	return code.String()
}

func (v StructTags) String() string {
	return GetTemplate().StructTags(v)
}
