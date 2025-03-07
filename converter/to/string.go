package to

import "fmt"

// String converts a value to a string.
func String[TType Type](val TType) string {
	return fmt.Sprintf("%v", val)
}
