package keeper

import (
	"encoding/binary"
	"github.com/OmniFlix/payment-stream/x/paymentstream/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetPaymentStreamCount get the number of listings
func (k Keeper) GetNextPaymentStreamNumber(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.PrefixPaymentStreamCount
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetNextPaymentStreamNumber(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.PrefixPaymentStreamCount
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetPaymentStream gets a specific PaymentStream in the store from its id
func (k Keeper) GetPaymentStream(ctx sdk.Context, id string) (val types.PaymentStream, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefix(id)
	if !store.Has(key) {
		return val, false
	}
	bz := store.Get(key)
	k.cdc.MustUnmarshal(bz, &val)
	return val, true
}

// SetPaymentStream set a specific PaymentStream in the store with its id
func (k Keeper) SetPaymentStream(ctx sdk.Context, PaymentStream types.PaymentStream) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	b := k.cdc.MustMarshal(&PaymentStream)
	store.Set(types.KeyPrefix(
		PaymentStream.Id,
	), b)
}

// RemovePaymentStream removes a payment-stream from the store
func (k Keeper) RemoveListing(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	store.Delete(types.KeyPrefix(id))
}

// GetAllPaymentStream returns all PaymentStreams
func (k Keeper) GetAllPaymentStreams(ctx sdk.Context) (list []types.PaymentStream) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	iterator := sdk.KVStorePrefixIterator(store, types.PrefixPaymentStreamId)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PaymentStream
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
