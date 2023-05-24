package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
)

func TestParamKeyTable(t *testing.T) {
	require.IsType(t, paramtypes.KeyTable{}, ParamKeyTable())
	require.NotEmpty(t, ParamKeyTable())
}

func TestParamSetPairs(t *testing.T) {
	params := DefaultParams()
	require.NotEmpty(t, params.ParamSetPairs())
}

func TestParams_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name     string
		params   Params
		expError bool
	}{
		{"default params", DefaultParams(), false},
		{"invalid stream fee", Params{StreamPaymentFee: sdk.NewCoin("uspay", sdk.ZeroInt())}, true},
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
