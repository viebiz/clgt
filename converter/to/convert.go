package to

import (
	"fmt"

	"github.com/viebiz/clgt/types"
)

// Convert converts a value to another type.
func Convert[TFrom types.Number, TTo types.Number](v TFrom) (TTo, error) {
	var result TTo
	if v < 0 && types.IsUnsignedNumber(result) {
		return result, fmt.Errorf("cannot convert negative number %T to unsigned number %T", v, result)
	}
	result = TTo(v)
	return result, nil
}
