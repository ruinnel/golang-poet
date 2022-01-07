package poet

import "reflect"

type Argument struct {
	VariableNumber bool
	Name           string
	Type           reflect.Type
}

type Arguments []Argument

func (a Arguments) String() string {
	return GetTemplate().Argument(a)
}

func (a Arguments) NeedType(idx int) bool {
	if len(a) <= idx+1 {
		return true
	}
	if len(a) > idx+1 {
		this := a[idx]
		next := a[idx+1]
		return this.Type.Kind() != next.Type.Kind()
	}
	return false
}
