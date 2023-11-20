package keeper

import (
	"encoding/binary"
	"math"

	"github.com/OmniFlix/streampay/v2/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

// GetNextStreamPaymentNumber get next stream payment number
func (k Keeper) GetNextStreamPaymentNumber(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := types.PrefixPaymentStreamCount
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetNextStreamPaymentNumber sets next stream payment number
func (k Keeper) SetNextStreamPaymentNumber(ctx sdk.Context, count uint64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := types.PrefixPaymentStreamCount
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetStreamPayment gets a specific StreamPayment in the store from its id
func (k Keeper) GetStreamPayment(ctx sdk.Context, id string) (val types.StreamPayment, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPrefixSteamPayment(id)
	if !store.Has(key) {
		return val, false
	}
	bz := store.Get(key)
	k.cdc.MustUnmarshal(bz, &val)
	return val, true
}

// SetStreamPayment set a specific StreamPayment in the store with its id
func (k Keeper) SetStreamPayment(ctx sdk.Context, streamPayment types.StreamPayment) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&streamPayment)
	store.Set(types.KeyPrefixSteamPayment(streamPayment.Id), b)
}

// RemoveStreamPayment removes a stream-payment from the store
func (k Keeper) RemoveStreamPayment(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyPrefixSteamPayment(id))
}

func (k Keeper) IterateStreamPayments(ctx sdk.Context, fn func(index int64, streamPayment types.StreamPayment) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PrefixPaymentStreamId)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		sp := types.StreamPayment{}
		err := proto.Unmarshal(iterator.Value(), &sp)
		if err != nil {
			panic(err)
		}
		stop := fn(i, sp)

		if stop {
			break
		}
		i++
	}
}

// GetAllStreamPayments returns all PaymentStreams
func (k Keeper) GetAllStreamPayments(ctx sdk.Context) (streamPayments []types.StreamPayment) {
	k.IterateStreamPayments(ctx, func(index int64, streamPayment types.StreamPayment) (stop bool) {
		streamPayments = append(streamPayments, streamPayment)
		return false
	})

	return streamPayments
}

func (k Keeper) getStreamedAmount(ctx sdk.Context, streamPayment types.StreamPayment) float64 {
	nowTime := ctx.BlockTime().Unix()
	startTime := streamPayment.StartTime.Unix()
	endTime := streamPayment.EndTime.Unix()
	totalAmount := streamPayment.TotalAmount.Amount.Int64()
	var percentage float64
	if nowTime >= endTime {
		percentage = 1.0
	} else {
		percentage = math.Abs(float64(nowTime-startTime) / float64(endTime-startTime))
	}
	return float64(totalAmount) * percentage
}

func (k Keeper) getStreamedAmountForPeriodicStreamPayment(ctx sdk.Context, streamPayment types.StreamPayment) float64 {
	streamedAmount := int64(0)
	nowTime := ctx.BlockTime().Unix()
	startTime := streamPayment.StartTime.Unix()
	totalDuration := int64(0)
	for _, period := range streamPayment.Periods {
		totalDuration += period.Duration
		if (startTime + totalDuration) <= nowTime {
			streamedAmount += period.Amount
		} else {
			break
		}
	}
	return float64(streamedAmount)
}
