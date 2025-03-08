package to

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCast(t *testing.T) {
	type simpleStruct struct {
		fieldA string
	}

	type arg[TType any] struct {
		givenValue interface{}
		expResult  TType
		expErr     error
	}

	dataTypes := map[string]interface{}{
		"int": map[string]arg[int]{
			"error": {
				givenValue: "1",
				expErr:     errors.New("cannot cast string to int"),
			},
			"ok": {
				givenValue: 1,
				expResult:  1,
			},
		},
		"float32": map[string]arg[float32]{
			"error": {
				givenValue: "1",
				expErr:     errors.New("cannot cast string to float32"),
			},
			"ok": {
				givenValue: float32(1.1),
				expResult:  1.1,
			},
		},
		"struct": map[string]arg[simpleStruct]{
			"error": {
				givenValue: "1",
				expErr:     errors.New("cannot cast string to to.simpleStruct"),
			},
			"ok": {
				givenValue: simpleStruct{"a"},
				expResult:  simpleStruct{"a"},
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
					result, err := Cast[int](tc.givenValue)

					// Then
					if tc.expErr != nil {
						require.EqualError(t, err, tc.expErr.Error())
					} else {
						require.NoError(t, err)
						require.Equal(t, tc.expResult, result)
					}
				})
			}
		case "float32":
			tcs := testCases.(map[string]arg[float32])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					result, err := Cast[float32](tc.givenValue)

					// Then
					if tc.expErr != nil {
						require.EqualError(t, err, tc.expErr.Error())
					} else {
						require.NoError(t, err)
						require.Equal(t, tc.expResult, result)
					}
				})
			}
		case "struct":
			tcs := testCases.(map[string]arg[simpleStruct])
			for scenario, tc := range tcs {
				t.Run(fmt.Sprintf("[%s] %s", dataType, scenario), func(t *testing.T) {
					// Given

					// When
					result, err := Cast[simpleStruct](tc.givenValue)

					// Then
					if tc.expErr != nil {
						require.EqualError(t, err, tc.expErr.Error())
					} else {
						require.NoError(t, err)
						require.Equal(t, tc.expResult, result)
					}
				})
			}
		}
	}
}
