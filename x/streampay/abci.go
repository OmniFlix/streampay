package streampay

import (
	"github.com/OmniFlix/payment-stream/x/streampay/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func EndBlock(ctx sdk.Context, k *keeper.Keeper) []abcitypes.ValidatorUpdate {
	if err := k.ProcessStreamPayments(ctx); err != nil {
		panic(err)
	}
	return []abcitypes.ValidatorUpdate{}
}
