package testutils

import "reflect"

func GetFnPtr(fn interface{}) uintptr {
	return reflect.ValueOf(fn).Pointer()
}
