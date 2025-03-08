package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsEmpty(t *testing.T) {
	type simpleStruct struct {
		fieldA string
	}

	type arg[TItem any] struct {
		givenItems []TItem
		expResult  bool
	}

	dataTypes := map[string]interface{}{
		"int": map[string]arg[int]{
			"false": {
				givenItems: []int{0, 1, 2, 3, 4, 5},
				expResult:  false,
			},
			"true": {
				expResult: true,
			},
		},
		"float32": map[string]arg[float32]{
			"false": {
				givenItems: []float32{0.200, 1.200, 2.200, 3.200, 4.200, 5.200},
				expResult:  false,
			},
			"true": {
				expResult: true,
			},
		},
		"struct": map[string]arg[simpleStruct]{
			"false": {
				givenItems: []simpleStruct{{"a"}, {"b"}, {"c"}},
				expResult:  false,
			},
			"true": {
				expResult: true,
			},
		},
	}
	for dataType, testCases := range dataTypes {
		switch dataType {
		case "int":
			tcs := testCases.(map[string]arg[int])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					result := IsEmpty(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "float32":
			tcs := testCases.(map[string]arg[float32])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					result := IsEmpty(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "struct":
			tcs := testCases.(map[string]arg[simpleStruct])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					result := IsEmpty(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		}
	}
}
