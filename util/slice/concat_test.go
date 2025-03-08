package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConcat(t *testing.T) {
	type simpleStruct struct {
		fieldA string
	}

	type arg[TItem any] struct {
		givenItems1 []TItem
		givenItems2 []TItem
		expItems    []TItem
	}

	dataTypes := map[string]interface{}{
		"int": map[string]arg[int]{
			"ok": {
				givenItems1: []int{0, 1, 2},
				givenItems2: []int{3, 4, 5},
				expItems:    []int{0, 1, 2, 3, 4, 5},
			},
		},
		"float32": map[string]arg[float32]{
			"ok": {
				givenItems1: []float32{0.200, 1.200, 2.200},
				givenItems2: []float32{3.200, 4.200, 5.200},
				expItems:    []float32{0.200, 1.200, 2.200, 3.200, 4.200, 5.200},
			},
		},
		"struct": map[string]arg[simpleStruct]{
			"ok": {
				givenItems1: []simpleStruct{{"a"}, {"b"}, {"c"}},
				givenItems2: []simpleStruct{{"d"}, {"e"}, {"f"}},
				expItems:    []simpleStruct{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}, {"f"}},
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
					items := Concat(tc.givenItems1, tc.givenItems2)

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
					items := Concat(tc.givenItems1, tc.givenItems2)

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
					items := Concat(tc.givenItems1, tc.givenItems2)

					// Then
					require.Equal(t, tc.expItems, items)
				})
			}
		}
	}
}
