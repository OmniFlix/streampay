package keeper

import (
	"github.com/OmniFlix/streampay/v2/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams returns the total set of streampay parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the streampay parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetStreamPaymentFee returns the current stream payment fee.
func (k Keeper) GetStreamPaymentFee(ctx sdk.Context) sdk.Coin {
	params := k.GetParams(ctx)
	return params.StreamPaymentFee
}
