package to

import (
	"fmt"

	"github.com/viebiz/clgt/types"
)

func Cast[TType types.Struct](val interface{}) (TType, error) {
	result, ok := val.(TType)
	if !ok {
		return TType{}, fmt.Errorf("cannot cast %v to %T", val, TType{})
	}
	return result, nil
}
