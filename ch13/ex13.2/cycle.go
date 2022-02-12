package cycle

import (
	"reflect"
	"unsafe"
)

func hasCycle(x reflect.Value, seen map[unsafe.Pointer]bool) bool {
	if !x.IsValid() {
		return false
	}
	switch x.Kind() {
	case reflect.Struct:
		xptr := unsafe.Pointer(x.UnsafeAddr())
		if seen[xptr] {
			return true
		}
		seen[xptr] = true
		for i := 0; i < x.NumField(); i++ {
			if hasCycle(x.Field(i), seen) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func HasCycle(x interface{}) bool {
	seen := make(map[unsafe.Pointer]bool)
	return hasCycle(reflect.ValueOf(x), seen)
}
