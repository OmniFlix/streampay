package keeper

import (
	"encoding/binary"
	"fmt"
	"github.com/OmniFlix/payment-stream/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math"
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

// GetAllStreamPayments returns all PaymentStreams
func (k Keeper) GetAllStreamPayments(ctx sdk.Context) (list []types.StreamPayment) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PrefixPaymentStreamId)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StreamPayment
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ProcessStreamPayments(ctx sdk.Context) error {
	k.Logger(ctx).Info("Processing stream payments ..")

	StreamPayments := k.GetAllStreamPayments(ctx)
	for _, streamPayment := range StreamPayments {
		k.Logger(ctx).Debug(fmt.Sprintf(
			"paymentId: %s, type: %s", streamPayment.Id, streamPayment.StreamType.String()))
		if streamPayment.GetStreamType() == types.TypeDelayed {
			if ctx.BlockTime().Unix() >= streamPayment.EndTime.Unix() {
				recipient, err := sdk.AccAddressFromBech32(streamPayment.Recipient)
				if err != nil {
					return err
				}
				err = k.TransferAmountFromModuleAccount(ctx, recipient, sdk.NewCoins(streamPayment.TotalAmount))
				if err != nil {
					return err
				}
				streamPayment.TotalTransferred = streamPayment.TotalAmount
				streamPayment.LastTransferredAt = ctx.BlockTime()
				k.RemoveStreamPayment(ctx, streamPayment.Id)

				ctx.EventManager().EmitEvents(
					sdk.Events{sdk.NewEvent(
						"payment_transfer",
						sdk.NewAttribute("payment-id", streamPayment.Id),
						sdk.NewAttribute("sender", streamPayment.Sender),
					),
					},
				)
				ctx.EventManager().EmitEvents(
					sdk.Events{sdk.NewEvent(
						"payment_complete",
						sdk.NewAttribute("payment-id", streamPayment.Id),
						sdk.NewAttribute("sender", streamPayment.Sender),
					),
					},
				)
			}

		} else if streamPayment.GetStreamType() == types.TypeContinuous {
			recipient, err := sdk.AccAddressFromBech32(streamPayment.Recipient)
			if err != nil {
				return err
			}
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
			unlockedAmount := float64(totalAmount) * percentage
			k.Logger(ctx).Debug(fmt.Sprintf("Total unlocked amount %f for payment %s", unlockedAmount, streamPayment.Id))
			amountToSend := int64(unlockedAmount) - streamPayment.TotalTransferred.Amount.Int64()
			amount := sdk.NewCoin(streamPayment.TotalAmount.Denom, sdk.NewInt(amountToSend))
			if amount.IsZero() || amount.IsNil() {
				continue
			}

			err = k.TransferAmountFromModuleAccount(ctx, recipient, sdk.NewCoins(amount))
			if err != nil {
				return err
			}
			k.Logger(ctx).Debug(fmt.Sprintf("Transferred amount %s to %s", amount.String(), recipient.String()))
			streamPayment.TotalTransferred = streamPayment.TotalTransferred.Add(amount)
			streamPayment.LastTransferredAt = ctx.BlockTime()

			k.SetStreamPayment(ctx, streamPayment)

			ctx.EventManager().EmitEvents(
				sdk.Events{
					sdk.NewEvent(
						"payment_transfer",
						sdk.NewAttribute("payment-id", streamPayment.Id),
						sdk.NewAttribute("sender", streamPayment.Sender),
					),
				},
			)

			if streamPayment.TotalTransferred.IsGTE(streamPayment.TotalAmount) {
				k.RemoveStreamPayment(ctx, streamPayment.Id)

				ctx.EventManager().EmitEvents(
					sdk.Events{sdk.NewEvent(
						"payment_complete",
						sdk.NewAttribute("payment-id", streamPayment.Id),
						sdk.NewAttribute("sender", streamPayment.Sender),
					),
					},
				)
			}

		} else {
			continue
		}
	}
	k.Logger(ctx).Info("Processed payment streams ..")
	return nil
}
