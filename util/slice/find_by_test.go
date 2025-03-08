package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindBy(t *testing.T) {
	type simpleStruct struct {
		fieldA string
	}

	type arg[TItem any] struct {
		givenItems     []TItem
		givenPredicate func(item TItem) bool
		expItem        TItem
		expFound       bool
	}

	dataTypes := map[string]interface{}{
		"int": map[string]arg[int]{
			"empty": {
				givenItems:     []int{0, 1, 2, 3, 4, 5},
				givenPredicate: func(item int) bool { return item == 6 },
				expItem:        0,
				expFound:       false,
			},
			"ok": {
				givenItems:     []int{0, 1, 2, 3, 4, 5},
				givenPredicate: func(item int) bool { return item == 2 },
				expItem:        2,
				expFound:       true,
			},
		},
		"float32": map[string]arg[float32]{
			"empty": {
				givenItems:     []float32{0.200, 1.200, 2.200, 3.200, 4.200, 5.200},
				givenPredicate: func(item float32) bool { return item == 6 },
				expItem:        0,
				expFound:       false,
			},
			"ok": {
				givenItems:     []float32{0.200, 1.200, 2.200, 3.200, 4.200, 5.200},
				givenPredicate: func(item float32) bool { return item == 1.2 },
				expItem:        1.2,
				expFound:       true,
			},
		},
		"struct": map[string]arg[simpleStruct]{
			"empty": {
				givenItems:     []simpleStruct{{"a"}, {"b"}, {"c"}},
				givenPredicate: func(item simpleStruct) bool { return item.fieldA == "d" },
				expItem:        simpleStruct{},
				expFound:       false,
			},
			"ok": {
				givenItems:     []simpleStruct{{"a"}, {"b"}, {"c"}},
				givenPredicate: func(item simpleStruct) bool { return item.fieldA == "b" },
				expItem:        simpleStruct{"b"},
				expFound:       true,
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
					item, found := FindBy(tc.givenItems, tc.givenPredicate)

					// Then
					require.Equal(t, tc.expItem, item)
					require.Equal(t, tc.expFound, found)
				})
			}
		case "float32":
			tcs := testCases.(map[string]arg[float32])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					item, found := FindBy(tc.givenItems, tc.givenPredicate)

					// Then
					require.Equal(t, tc.expItem, item)
					require.Equal(t, tc.expFound, found)
				})
			}
		case "struct":
			tcs := testCases.(map[string]arg[simpleStruct])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					item, found := FindBy(tc.givenItems, tc.givenPredicate)

					// Then
					require.Equal(t, tc.expItem, item)
					require.Equal(t, tc.expFound, found)
				})
			}
		}
	}
}
