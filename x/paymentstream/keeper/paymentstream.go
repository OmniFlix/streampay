package keeper

import (
	"github.com/OmniFlix/payment-stream/x/paymentstream/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPaymentStream set a specific PaymentStream in the store from its index
func (k Keeper) SetPaymentStream(ctx sdk.Context, PaymentStream types.PaymentStream) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(""))
	b := k.cdc.MustMarshal(&PaymentStream)
	store.Set(types.KeyPrefix(
		PaymentStream.Id,
	), b)
}

// GetAllPaymentStream returns all PaymentStream
func (k Keeper) GetAllPaymentStream(ctx sdk.Context) (list []types.PaymentStream) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(""))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PaymentStream
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
