package streampay

import (
	"github.com/OmniFlix/streampay/x/streampay/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func EndBlock(ctx sdk.Context, k *keeper.Keeper) []abcitypes.ValidatorUpdate {
	k.ProcessStreamPayments(ctx)
	return []abcitypes.ValidatorUpdate{}
}
