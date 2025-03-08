package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterBy(t *testing.T) {
	type simpleStruct struct {
		fieldA string
	}

	type arg[TItem any] struct {
		givenItems     []TItem
		givenPredicate func(item TItem) bool
		expItems       []TItem
	}

	dataTypes := map[string]interface{}{
		"int": map[string]arg[int]{
			"empty": {
				givenItems:     []int{0, 1, 2, 3, 4, 5},
				givenPredicate: func(item int) bool { return item > 6 },
			},
			"ok": {
				givenItems:     []int{0, 1, 2, 3, 4, 5},
				givenPredicate: func(item int) bool { return item > 2 },
				expItems:       []int{3, 4, 5},
			},
		},
		"float32": map[string]arg[float32]{
			"empty": {
				givenItems:     []float32{0.200, 1.200, 2.200, 3.200, 4.200, 5.200},
				givenPredicate: func(item float32) bool { return item == 6 },
			},
			"ok": {
				givenItems:     []float32{0.200, 1.200, 2.200, 3.200, 4.200, 5.200},
				givenPredicate: func(item float32) bool { return item > 1.2 },
				expItems:       []float32{2.200, 3.200, 4.200, 5.200},
			},
		},
		"struct": map[string]arg[simpleStruct]{
			"empty": {
				givenItems:     []simpleStruct{{"a"}, {"b"}, {"c"}},
				givenPredicate: func(item simpleStruct) bool { return item.fieldA == "d" },
			},
			"ok": {
				givenItems:     []simpleStruct{{"a"}, {"b"}, {"c"}},
				givenPredicate: func(item simpleStruct) bool { return item.fieldA == "b" },
				expItems:       []simpleStruct{{"b"}},
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
					items := FilterBy(tc.givenItems, tc.givenPredicate)

					// Then
					require.Equal(t, tc.expItems, items)
				})
			}
		case "float32":
			tcs := testCases.(map[string]arg[float32])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					items := FilterBy(tc.givenItems, tc.givenPredicate)

					// Then
					require.Equal(t, tc.expItems, items)
				})
			}
		case "struct":
			tcs := testCases.(map[string]arg[simpleStruct])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					items := FilterBy(tc.givenItems, tc.givenPredicate)

					// Then
					require.Equal(t, tc.expItems, items)
				})
			}
		}
	}
}
