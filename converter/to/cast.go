package to

import (
	"fmt"
)

// Cast converts a value to a struct.
func Cast[TType any](v interface{}) (TType, error) {
	result, ok := v.(TType)
	if !ok {
		return result, fmt.Errorf("cannot cast %T to %T", v, result)
	}
	return result, nil
}
