package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	type simpleAStruct struct {
		fieldA string
	}

	type simpleBStruct struct {
		fieldB string
	}

	type arg[TItem any, TResult any] struct {
		givenItems  []TItem
		givenMapper func(TItem) TResult
		expItems    []TResult
	}

	dataTypes := map[string]interface{}{
		"int": map[string]arg[int, int64]{
			"ok": {
				givenItems: []int{0, 1, 2},
				givenMapper: func(item int) int64 {
					return int64(item)
				},
				expItems: []int64{0, 1, 2},
			},
		},
		"float32": map[string]arg[float32, float64]{
			"ok": {
				givenItems: []float32{0.200, 1.200, 2.200},
				givenMapper: func(item float32) float64 {
					//val := to.String(item)

					return float64(item)
				},
				expItems: []float64{0.200, 1.200, 2.200},
			},
		},
		"struct": map[string]arg[simpleAStruct, simpleBStruct]{
			"ok": {
				givenItems: []simpleAStruct{{"a"}, {"b"}, {"c"}},
				givenMapper: func(item simpleAStruct) simpleBStruct {
					return simpleBStruct{item.fieldA}
				},
				expItems: []simpleBStruct{{"a"}, {"b"}, {"c"}},
			},
		},
	}
	for dataType, testCases := range dataTypes {
		switch dataType {
		case "int":
			tcs := testCases.(map[string]arg[int, int64])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					items := Map(tc.givenItems, tc.givenMapper)

					// Then
					require.Equal(t, tc.expItems, items)
				})
			}
		case "float32":
			tcs := testCases.(map[string]arg[float32, float64])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					items := Map(tc.givenItems, tc.givenMapper)

					// Then
					require.Equal(t, tc.expItems, items)
				})
			}
		case "struct":
			tcs := testCases.(map[string]arg[simpleAStruct, simpleBStruct])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					items := Map(tc.givenItems, tc.givenMapper)

					// Then
					require.Equal(t, tc.expItems, items)
				})
			}
		}
	}
}
