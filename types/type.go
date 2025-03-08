package types

// PrimitiveType is the type of the value to convert.
type PrimitiveType interface {
	Number | string | bool
}

// Number is the type of the number to convert.
type Number interface {
	IntNumber | float32 | float64
}

// IntNumber is the type of the integer number to convert.
type IntNumber interface {
	int | int8 | int16 | int32 | int64 | UIntNumber
}

// UIntNumber is the type of the unsigned integer number to convert.
type UIntNumber interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Struct interface {
	struct{}
}
