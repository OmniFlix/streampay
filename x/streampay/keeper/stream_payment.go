package keeper

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/OmniFlix/streampay/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
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
func (k Keeper) SetStreamPayment(ctx sdk.Context, PaymentStream types.StreamPayment) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&PaymentStream)
	store.Set(types.KeyPrefixSteamPayment(PaymentStream.Id), b)
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

func (k Keeper) ProcessStreamPayments(ctx sdk.Context) {
	logger := k.Logger(ctx)
	logger.Info("Processing stream payments ..")

	k.IterateStreamPayments(ctx, func(index int64, streamPayment types.StreamPayment) (stop bool) {
		switch streamPayment.GetStreamType() {
		case types.TypeDelayed:
			if ctx.BlockTime().Unix() < streamPayment.EndTime.Unix() {
				return false
			}
			if err := k.processDelayedStreamPayment(ctx, streamPayment); err != nil {
				panic(err)
			}
			logger.Debug(
				fmt.Sprintf(
					"Transferred amount %s to %s", streamPayment.TotalAmount.String(), streamPayment.Recipient,
				),
			)
			// Remove stream payment
			k.RemoveStreamPayment(ctx, streamPayment.Id)

			// Emit events
			k.emitStreamPaymentTransferEvent(ctx, streamPayment.Id, streamPayment.Recipient, streamPayment.TotalAmount)
			k.emitStreamPaymentEndEvent(ctx, streamPayment.Id, streamPayment.Sender)

		case types.TypeContinuous:
			if ctx.BlockTime().Unix() < streamPayment.StartTime.Unix() {
				return false
			}
			if streamPayment.TotalTransferred.IsGTE(streamPayment.TotalAmount) {
				k.RemoveStreamPayment(ctx, streamPayment.Id)
				k.emitStreamPaymentEndEvent(ctx, streamPayment.Id, streamPayment.Sender)
				return false
			}
			unlockedAmount := k.getUnlockedAmount(ctx, streamPayment)
			amountToSend := int64(unlockedAmount) - streamPayment.TotalTransferred.Amount.Int64()
			amount := sdk.NewCoin(streamPayment.TotalAmount.Denom, sdk.NewInt(amountToSend))

			if amount.IsZero() || amount.IsNil() {
				return false
			}

			logger.Debug(
				fmt.Sprintf("Total unlocked amount %s for payment %s", amount.String(), streamPayment.Id),
			)
			if err := k.processContinuousStreamPayment(ctx, amount, streamPayment); err != nil {
				panic(err)
			}
			logger.Debug(fmt.Sprintf("Transferred amount %s to %s", amount.String(), streamPayment.Recipient))
			// update stream payment
			streamPayment.TotalTransferred = streamPayment.TotalTransferred.Add(amount)
			streamPayment.LastTransferredAt = ctx.BlockTime()
			k.SetStreamPayment(ctx, streamPayment)
			// emit events
			k.emitStreamPaymentTransferEvent(ctx, streamPayment.Id, streamPayment.Recipient, amount)

		default:
			panic(fmt.Errorf("unknown error type"))
		}

		return false
	})
	logger.Info("Processed stream payments ..")
}

func (k Keeper) processDelayedStreamPayment(ctx sdk.Context, streamPayment types.StreamPayment) error {
	recipient, err := sdk.AccAddressFromBech32(streamPayment.Recipient)
	if err != nil {
		return err
	}
	err = k.TransferAmountFromModuleAccount(ctx, recipient, sdk.NewCoins(streamPayment.TotalAmount))
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) processContinuousStreamPayment(ctx sdk.Context, amount sdk.Coin, streamPayment types.StreamPayment) error {
	recipient, err := sdk.AccAddressFromBech32(streamPayment.Recipient)
	if err != nil {
		return err
	}
	err = k.TransferAmountFromModuleAccount(ctx, recipient, sdk.NewCoins(amount))
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) getUnlockedAmount(ctx sdk.Context, streamPayment types.StreamPayment) float64 {
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
