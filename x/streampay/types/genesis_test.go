package types_test

import (
	"testing"

	"github.com/OmniFlix/streampay/v2/x/streampay/types"
	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	// generate a private/public key pair and get the respective address
	pk1 := ed25519.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pk1.Address())

	pk2 := ed25519.GenPrivKey().PubKey()
	addr2 := sdk.AccAddress(pk2.Address())

	defaultAmount := sdk.NewInt64Coin("uspay", 100_000_000)

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
						Sender:         addr1.String(),
						Recipient:      addr2.String(),
						StreamType:     types.TypeDelayed,
						TotalAmount:    defaultAmount,
						StreamedAmount: sdk.NewCoin(defaultAmount.Denom, sdk.ZeroInt()),
					},
					{
						Id:             "sp2",
						Sender:         addr1.String(),
						Recipient:      addr2.String(),
						StreamType:     types.TypeContinuous,
						TotalAmount:    defaultAmount,
						StreamedAmount: sdk.NewCoin(defaultAmount.Denom, sdk.ZeroInt()),
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
