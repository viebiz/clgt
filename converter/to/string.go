package to

import (
	"fmt"
	"github.com/viebiz/clgt/types"
)

// String converts a value to a string.
func String[TType types.Number](val TType) string {
	return fmt.Sprintf("%v", val)
}
