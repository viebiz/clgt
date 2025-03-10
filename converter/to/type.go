package to

import (
	"fmt"
	"github.com/viebiz/clgt/types"
)

type conversion[TFrom, TTo types.Number] interface {
	Convert(v TFrom) (TTo, error)
	getMin() TTo
	getMax() TTo
}

type conversionStrategy[TFrom, TTo types.Number] struct {
	Min TTo
	Max TTo
}

type intToUintConversion[TUint types.UIntNumber] conversionStrategy[int, TUint]

func (i *intToUintConversion[TUint]) getMin() TUint {
	return i.Min
}

func (i *intToUintConversion[TUint]) getMax() TUint {
	return i.Max
}

func (i *intToUintConversion[TUint]) Convert(v int) (TUint, error) {
	var result TUint
	if v < 0 {
		return 0, fmt.Errorf("cannot convert negative number %T to unsigned number %T", v, result)
	}
	if v > i.getMax() {
		return 0, ErrOverflow
	}
}
