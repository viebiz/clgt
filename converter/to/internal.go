package to

import (
	"reflect"

	"github.com/viebiz/clgt/types"
)

type conversion[TFrom types.Number, TTo types.Number] interface {
	Convert(v TFrom) (TTo, error)
}

type impl[TFrom types.Number, TTo types.Number] struct {
}

func (i *impl[TFrom, TTo]) Convert(v TFrom) (TTo, error) {
	return TTo(v), nil
}

//type conversion[TFrom types.Number, TTo types.Number] struct{}
//
//func (c *conversion[TFrom, TTo]) Convert(v TFrom) (TTo, error) {
//	return TTo(v), nil
//}
//
//type intUint8Conversion conversion[int, int8]
//
//func (i *intUint8Conversion) Convert(v int) (int8, error) {
//
//}

var (
	conversionMap = map[reflect.Kind]map[reflect.Kind]interface{}{
		reflect.Int: {
			reflect.Int8: &impl[int, int8]{},
		},
	}
)

func taaa() {
	i := intUint8Conversion{}
	result, err := i.Convert(1)
}
