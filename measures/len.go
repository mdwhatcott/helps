package measures

import (
	"fmt"
	"reflect"
)

type Len interface {
	Len() int
}

func Length(a interface{}) int {
	if l, ok := a.(Len); ok {
		return l.Len()
	}
	switch v := reflect.ValueOf(a); v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return v.Len()
	}
	panic(fmt.Sprintf("The value (%v) was not compatible with len() builtin or does not implement Len interface.", a))
	return -1
}
