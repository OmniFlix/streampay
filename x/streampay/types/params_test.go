package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParams_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name     string
		params   Params
		expError bool
	}{
		{"default params", DefaultParams(), false},
		{"valid params", Params{StreamPaymentFeePercentage: sdk.NewDec(0)}, false},
		{"invalid stream fee percentage", Params{StreamPaymentFeePercentage: sdk.NewDec(1)}, true},
		{"invalid stream fee percentage", Params{StreamPaymentFeePercentage: sdk.NewDec(-1)}, true},
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
