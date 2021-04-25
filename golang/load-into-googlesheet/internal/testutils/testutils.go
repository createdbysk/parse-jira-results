package testutils

import "reflect"

type MockFunction interface {
	Unpatch()
}

func GetFnPtr(fn interface{}) uintptr {
	return reflect.ValueOf(fn).Pointer()
}
