package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindIndexBy(t *testing.T) {
	type simpleStruct struct {
		fieldA string
	}

	type arg[TItem any] struct {
		givenItems     []TItem
		givenPredicate func(item TItem) bool
		expIndex       int
	}

	dataTypes := map[string]interface{}{
		"int": map[string]arg[int]{
			"-1": {
				givenItems:     []int{0, 1, 2, 3, 4, 5},
				givenPredicate: func(item int) bool { return item == 6 },
				expIndex:       -1,
			},
			"ok": {
				givenItems:     []int{0, 1, 2, 3, 4, 5},
				givenPredicate: func(item int) bool { return item == 2 },
				expIndex:       2,
			},
		},
		"float32": map[string]arg[float32]{
			"empty": {
				givenItems:     []float32{0.200, 1.200, 2.200, 3.200, 4.200, 5.200},
				givenPredicate: func(item float32) bool { return item == 6 },
				expIndex:       -1,
			},
			"ok": {
				givenItems:     []float32{0.200, 1.200, 2.200, 3.200, 4.200, 5.200},
				givenPredicate: func(item float32) bool { return item == 1.2 },
				expIndex:       1,
			},
		},
		"struct": map[string]arg[simpleStruct]{
			"empty": {
				givenItems:     []simpleStruct{{"a"}, {"b"}, {"c"}},
				givenPredicate: func(item simpleStruct) bool { return item.fieldA == "d" },
				expIndex:       -1,
			},
			"ok": {
				givenItems:     []simpleStruct{{"a"}, {"b"}, {"c"}},
				givenPredicate: func(item simpleStruct) bool { return item.fieldA == "b" },
				expIndex:       1,
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
					result := FindIndexBy(tc.givenItems, tc.givenPredicate)

					// Then
					require.Equal(t, tc.expIndex, result)
				})
			}
		case "float32":
			tcs := testCases.(map[string]arg[float32])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					result := FindIndexBy(tc.givenItems, tc.givenPredicate)

					// Then
					require.Equal(t, tc.expIndex, result)
				})
			}
		case "struct":
			tcs := testCases.(map[string]arg[simpleStruct])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					result := FindIndexBy(tc.givenItems, tc.givenPredicate)

					// Then
					require.Equal(t, tc.expIndex, result)
				})
			}
		}
	}
}
