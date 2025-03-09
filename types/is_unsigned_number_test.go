package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsUnsignedNumber(t *testing.T) {
	type arg struct {
		givenValue interface{}
		expResult  bool
	}

	tcs := map[string]arg{
		"int": {
			givenValue: 1,
			expResult:  false,
		},
		"uint": {
			givenValue: uint(1),
			expResult:  true,
		},
		"uint8": {
			givenValue: uint8(1),
			expResult:  true,
		},
		"uint16": {
			givenValue: uint16(1),
			expResult:  true,
		},
		"uint32": {
			givenValue: uint32(1),
			expResult:  true,
		},
		"uint64": {
			givenValue: uint64(1),
			expResult:  true,
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given

			// When
			result := IsUnsignedNumber(tc.givenValue)

			// Then
			require.Equal(t, tc.expResult, result)
		})
	}
}
