package to

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/viebiz/clgt/types"
)

func TestString(t *testing.T) {
	type arg[TType types.Number] struct {
		givenValue TType
		expResult  string
	}
	dataTypes := map[string]interface{}{
		"int": map[string]arg[int]{
			"ok": {
				givenValue: 1,
				expResult:  "1",
			},
		},
		"int8": map[string]arg[int8]{
			"ok": {
				givenValue: 1,
				expResult:  "1",
			},
		},
		"int16": map[string]arg[int16]{
			"ok": {
				givenValue: 1,
				expResult:  "1",
			},
		},
		"int32": map[string]arg[int32]{
			"ok": {
				givenValue: 1,
				expResult:  "1",
			},
		},
		"int64": map[string]arg[int64]{
			"ok": {
				givenValue: 1,
				expResult:  "1",
			},
		},
		"uint": map[string]arg[uint]{
			"ok": {
				givenValue: 1,
				expResult:  "1",
			},
		},
		"uint8": map[string]arg[uint8]{
			"ok": {
				givenValue: 1,
				expResult:  "1",
			},
		},
		"uint16": map[string]arg[uint16]{
			"ok": {
				givenValue: 1,
				expResult:  "1",
			},
		},
		"uint32": map[string]arg[uint32]{
			"ok": {
				givenValue: 1,
				expResult:  "1",
			},
		},
		"uint64": map[string]arg[uint64]{
			"ok": {
				givenValue: 1,
				expResult:  "1",
			},
		},
		"float32": map[string]arg[float32]{
			"ok": {
				givenValue: 1.100,
				expResult:  "1.1",
			},
		},
		"float64": map[string]arg[float64]{
			"ok": {
				givenValue: 1.100,
				expResult:  "1.1",
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
					result := String(tc.givenValue)

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
					result := String(tc.givenValue)

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
					result := String(tc.givenValue)

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
					result := String(tc.givenValue)

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
					result := String(tc.givenValue)

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
					result := String(tc.givenValue)

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
					result := String(tc.givenValue)

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
					result := String(tc.givenValue)

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
					result := String(tc.givenValue)

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
					result := String(tc.givenValue)

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
					result := String(tc.givenValue)

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
					result := String(tc.givenValue)

					// Then
					require.Equal(t, tc.expResult, result)
				})
			}
		}
	}
}
