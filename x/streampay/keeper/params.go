package keeper

import (
	"github.com/OmniFlix/streampay/v2/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams returns the total set of streampay parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return params
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams sets the streampay parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.ValidateBasic(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, bz)

	return nil
}

// GetStreamPaymentFeePercentage returns the current stream payment fee.
func (k Keeper) GetStreamPaymentFeePercentage(ctx sdk.Context) sdk.Dec {
	params := k.GetParams(ctx)
	return params.StreamPaymentFeePercentage
}
