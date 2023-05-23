package types_test

import (
	"testing"

	"github.com/OmniFlix/streampay/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				StreamPayments: []types.StreamPayment{
					{
						Id:             "sp1",
						Sender:         "cosmos1amnszjdguwlxaawqg9ey6axcyqk38vcegm8zm5",
						Recipient:      "cosmos12ecdcddd4rk0zhkazfvj6d37zwyhylhh9y6feg",
						StreamType:     types.TypeDelayed,
						TotalAmount:    sdk.NewCoin("uspay", sdk.NewInt(1000000)),
						StreamedAmount: sdk.NewCoin("uspay", sdk.ZeroInt()),
					},
					{
						Id:             "sp2",
						Sender:         "cosmos1amnszjdguwlxaawqg9ey6axcyqk38vcegm8zm5",
						Recipient:      "cosmos12ecdcddd4rk0zhkazfvj6d37zwyhylhh9y6feg",
						StreamType:     types.TypeContinuous,
						TotalAmount:    sdk.NewCoin("uspay", sdk.NewInt(1000000)),
						StreamedAmount: sdk.NewCoin("uspay", sdk.ZeroInt()),
					},
				},
				NextStreamPaymentNumber: 3,
				Params:                  types.DefaultParams(),
			},
			valid: true,
		},
		{
			desc: "partial values",
			genState: &types.GenesisState{
				StreamPayments: []types.StreamPayment{
					{
						Id: "sp1",
					},
				},
				NextStreamPaymentNumber: 2,
				Params:                  types.DefaultParams(),
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
