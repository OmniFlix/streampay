package streampay

import (
	"github.com/OmniFlix/streampay/v2/x/streampay/keeper"
	"github.com/OmniFlix/streampay/v2/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the streampay
	for _, streamPayment := range genState.StreamPayments {
		k.SetStreamPayment(ctx, streamPayment)
	}
	// set next payment number
	k.SetNextStreamPaymentNumber(ctx, genState.NextStreamPaymentNumber)
	// set params
	err := k.SetParams(ctx, genState.Params)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.StreamPayments = k.GetAllStreamPayments(ctx)
	genesis.NextStreamPaymentNumber = k.GetNextStreamPaymentNumber(ctx)
	genesis.Params = k.GetParams(ctx)

	return genesis
}
