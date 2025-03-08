package types

import "reflect"

func IsUnsignedNumber(val interface{}) bool {
	kind := reflect.TypeOf(val).Kind()
	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}
