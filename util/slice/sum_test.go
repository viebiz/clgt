package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/viebiz/clgt/types"
)

func TestSum(t *testing.T) {
	type arg[TNumber types.Number] struct {
		givenItems []TNumber
		expResult  TNumber
	}

	dataTypes := map[string]interface{}{
		"int": map[string]arg[int]{
			"ok": {
				givenItems: []int{1, 2, 3, 4, 5},
				expResult:  15,
			},
		},
		"int8": map[string]arg[int8]{
			"ok": {
				givenItems: []int8{1, 2, 3, 4, 5},
				expResult:  15,
			},
		},
		"int16": map[string]arg[int16]{
			"ok": {
				givenItems: []int16{1, 2, 3, 4, 5},
				expResult:  15,
			},
		},
		"int32": map[string]arg[int32]{
			"ok": {
				givenItems: []int32{1, 2, 3, 4, 5},
				expResult:  15,
			},
		},
		"int64": map[string]arg[int64]{
			"ok": {
				givenItems: []int64{1, 2, 3, 4, 5},
				expResult:  15,
			},
		},
		"uint": map[string]arg[uint]{
			"ok": {
				givenItems: []uint{1, 2, 3, 4, 5},
				expResult:  15,
			},
		},
		"uint8": map[string]arg[uint8]{
			"ok": {
				givenItems: []uint8{1, 2, 3, 4, 5},
				expResult:  15,
			},
		},
		"uint16": map[string]arg[uint16]{
			"ok": {
				givenItems: []uint16{1, 2, 3, 4, 5},
				expResult:  15,
			},
		},
		"uint32": map[string]arg[uint32]{
			"ok": {
				givenItems: []uint32{1, 2, 3, 4, 5},
				expResult:  15,
			},
		},
		"uint64": map[string]arg[uint64]{
			"ok": {
				givenItems: []uint64{1, 2, 3, 4, 5},
				expResult:  15,
			},
		},
		"float32": map[string]arg[float32]{
			"ok": {
				givenItems: []float32{1.2, 2.2, 3.2, 4.2, 5.2},
				expResult:  16,
			},
		},
		"float64": map[string]arg[float64]{
			"ok": {
				givenItems: []float64{1.2, 2.2, 3.2, 4.2, 5.2},
				expResult:  16,
			},
		},
	}
	for dataType, testCases := range dataTypes {
		switch dataType {
		case "int":
			tcs := testCases.(map[string]arg[int])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "int8":
			tcs := testCases.(map[string]arg[int8])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "int16":
			tcs := testCases.(map[string]arg[int16])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "int32":
			tcs := testCases.(map[string]arg[int32])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "int64":
			tcs := testCases.(map[string]arg[int64])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "uint":
			tcs := testCases.(map[string]arg[uint])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "uint8":
			tcs := testCases.(map[string]arg[uint8])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "uint16":
			tcs := testCases.(map[string]arg[uint16])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "uint32":
			tcs := testCases.(map[string]arg[uint32])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "uint64":
			tcs := testCases.(map[string]arg[uint64])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "float32":
			tcs := testCases.(map[string]arg[float32])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		case "float64":
			tcs := testCases.(map[string]arg[float64])
			for scenario, tc := range tcs {
				tc := tc
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					t.Parallel()
					// Given

					// When
					result := Sum(tc.givenItems)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		}
	}
}
