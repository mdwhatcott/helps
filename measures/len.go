package measures

import (
	"reflect"
)

type Len interface {
	Len() int
}

func Length(a interface{}) (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = -1
		}
	}()
	if l, ok := a.(Len); ok {
		return l.Len()
	}
	return reflect.ValueOf(a).Len()
}
