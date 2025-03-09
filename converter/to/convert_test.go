package to

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/viebiz/clgt/types"
)

func TestConvert(t *testing.T) {
	// When
	result, err := Convert[uint64, int8](types.MaxUint64)

	// Then
	//if tc.expErr != nil {
	//	require.EqualError(t, err, tc.expErr.Error())
	//} else {
	require.NoError(t, err)
	require.Equal(t, 0, result)
	//}

	type arg[TFrom types.Number, TTo types.Number] struct {
		givenFrom TFrom
		expResult TTo
		expErr    error
	}

	dataTypes := map[string]interface{}{
		"int32->uint32": map[string]arg[int32, uint32]{
			"error": {
				givenFrom: -1,
				expErr:    errors.New("cannot convert negative number int32 to unsigned number uint32"),
			},
			"max": {
				givenFrom: types.MaxInt32,
				expResult: types.MaxInt32,
			},
			"ok": {
				givenFrom: 1,
				expResult: 1,
			},
		},
		"int32->uint8": map[string]arg[int32, uint8]{
			"error": {
				givenFrom: -1,
				expErr:    errors.New("cannot convert negative number int32 to unsigned number uint8"),
			},
			"max": {
				givenFrom: types.MaxInt32,
				expResult: types.MaxUint8,
			},
			"ok": {
				givenFrom: 1,
				expResult: 1,
			},
		},
		"uint64->int8": map[string]arg[uint64, int8]{
			"max": {
				givenFrom: types.MaxUint64,
				expResult: types.MaxInt8,
			},
			"ok": {
				givenFrom: 1,
				expResult: 1,
			},
		},
	}
	for dataType, testCases := range dataTypes {
		switch dataType {
		case "int32->uint32":
			tcs := testCases.(map[string]arg[int32, uint32])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					result, err := Convert[int32, uint32](tc.givenFrom)

					// Then
					if tc.expErr != nil {
						require.EqualError(t, err, tc.expErr.Error())
					} else {
						require.NoError(t, err)
						require.Equal(t, tc.expResult, result)
					}
				})
			}
		case "int32->uint8":
			tcs := testCases.(map[string]arg[int32, uint8])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					result, err := Convert[int32, uint8](tc.givenFrom)

					// Then
					if tc.expErr != nil {
						require.EqualError(t, err, tc.expErr.Error())
					} else {
						require.NoError(t, err)
						require.Equal(t, tc.expResult, result)
					}
				})
			}
		case "uint64->int8":
			tcs := testCases.(map[string]arg[uint64, int8])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					result, err := Convert[uint64, int8](tc.givenFrom)

					// Then
					if tc.expErr != nil {
						require.EqualError(t, err, tc.expErr.Error())
					} else {
						require.NoError(t, err)
						require.Equal(t, tc.expResult, result)
					}
				})
			}
			//case "float32":
			//	tcs := testCases.(map[string]arg[float32])
			//	for scenario, tc := range tcs {
			//		t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
			//			// Given
			//
			//			// When
			//			result, err := Cast[float32](tc.givenValue)
			//
			//			// Then
			//			if tc.expErr != nil {
			//				require.EqualError(t, err, tc.expErr.Error())
			//			} else {
			//				require.NoError(t, err)
			//				require.Equal(t, tc.expResult, result)
			//			}
			//		})
			//	}
			//case "struct":
			//	tcs := testCases.(map[string]arg[simpleStruct])
			//	for scenario, tc := range tcs {
			//		t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
			//			// Given
			//
			//			// When
			//			result, err := Cast[simpleStruct](tc.givenValue)
			//
			//			// Then
			//			if tc.expErr != nil {
			//				require.EqualError(t, err, tc.expErr.Error())
			//			} else {
			//				require.NoError(t, err)
			//				require.Equal(t, tc.expResult, result)
			//			}
			//		})
			//	}
		}
	}
}
