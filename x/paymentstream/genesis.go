package paymentstream

import (
	"github.com/OmniFlix/payment-stream/x/paymentstream/keeper"
	"github.com/OmniFlix/payment-stream/x/paymentstream/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the paymentstream
	for _, elem := range genState.PaymentStreamsList {
		k.InitPaymentStreamFromGenesis(ctx, elem)
	}
	// set next payment number
	k.SetNextPaymentStreamNumber(ctx, genState.NextPaymentStreamCount)

	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.PaymentStreamsList = k.GetAllPaymentStreams(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
