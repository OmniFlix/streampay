package streampay

import (
	"github.com/OmniFlix/streampay/x/streampay/keeper"
	"github.com/OmniFlix/streampay/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the streampay
	for _, streamPayment := range genState.StreamPaymentsList {
		k.SetStreamPayment(ctx, streamPayment)
	}
	// set next payment number
	k.SetNextStreamPaymentNumber(ctx, genState.NextStreamPaymentNumber)

	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.StreamPaymentsList = k.GetAllStreamPayments(ctx)
	genesis.NextStreamPaymentNumber = k.GetNextStreamPaymentNumber(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
