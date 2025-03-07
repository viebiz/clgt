package to

// Ptr converts a value to a pointer.
func Ptr[TType Type](val TType) *TType {
	return &val
}
