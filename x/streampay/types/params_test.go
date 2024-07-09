package types

import (
	sdkmath "cosmossdk.io/math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParams_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name     string
		params   Params
		expError bool
	}{
		{"default params", DefaultParams(), false},
		{"valid params", Params{StreamPaymentFeePercentage: sdkmath.LegacyNewDec(0)}, false},
		{"invalid stream fee percentage", Params{StreamPaymentFeePercentage: sdkmath.LegacyNewDec(1)}, true},
		{"invalid stream fee percentage", Params{StreamPaymentFeePercentage: sdkmath.LegacyNewDec(-1)}, true},
	}
	for _, tc := range testCases {
		err := tc.params.ValidateBasic()

		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
		}
	}
}
