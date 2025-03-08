package to

import "github.com/viebiz/clgt/types"

// Ptr converts a value to a pointer.
func Ptr[TType types.PrimitiveType](val TType) *TType {
	return &val
}
