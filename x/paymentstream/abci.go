package paymentstream

import (
	"fmt"
	"github.com/OmniFlix/payment-stream/x/paymentstream/keeper"
	"github.com/OmniFlix/payment-stream/x/paymentstream/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func EndBlock(ctx sdk.Context, k *keeper.Keeper) []abcitypes.ValidatorUpdate {
	var (
		log = k.Logger(ctx)
	)
	log.Info("Starting active streaming  payments iteration")

	paymentStreams := k.GetAllPaymentStreams(ctx)
	for _, ps := range paymentStreams {
		if ps.GetStatus() != types.StatusOpen {
			continue
		}
		log.Info(fmt.Sprintf("paymentId: %s, type: %s, status: %s", ps.Id, ps.StreamType.String(), ps.Status.String()))
		if ps.GetStreamType() == types.TypeDelayed {
			if ctx.BlockTime().Unix() >= ps.EndTime.Unix() {
				recipient, err := sdk.AccAddressFromBech32(ps.Recipient)
				if err != nil {
					panic(err)
				}
				err = k.TransferAmountFromModuleAccount(ctx, recipient, sdk.NewCoins(ps.TotalAmount))
				if err != nil {
					panic(err)
				}
				ps.TotalTransferred = ps.TotalAmount
				ps.LastTransferredAt = ctx.BlockTime()
				ps.Status = types.StatusCompleted
				k.SetPaymentStream(ctx, ps)
				k.SetInActivePayment(ctx, ps.Id)

				ctx.EventManager().EmitEvents(
					sdk.Events{sdk.NewEvent(
						"payment_transfer",
						sdk.NewAttribute("payment-id", ps.Id),
						sdk.NewAttribute("sender", ps.Sender),
					),
					},
				)
				ctx.EventManager().EmitEvents(
					sdk.Events{sdk.NewEvent(
						"payment_complete",
						sdk.NewAttribute("payment-id", ps.Id),
						sdk.NewAttribute("sender", ps.Sender),
					),
					},
				)
			}

		} else if ps.GetStreamType() == types.TypeContinuous {
			recipient, err := sdk.AccAddressFromBech32(ps.Recipient)
			if err != nil {
				panic(err)
			}
			nowTime := ctx.BlockTime().Unix()
			startTime := ps.StartTime.Unix()
			endTime := ps.EndTime.Unix()
			if nowTime >= endTime {
				nowTime = endTime
			}
			totalAmount := ps.TotalAmount.Amount.Int64()
			percentage := float64(nowTime-startTime) / float64(endTime-startTime)
			unlockedAmount := float64(totalAmount) * percentage
			log.Info(fmt.Sprintf("Total unlocked amount %f for payment %s", unlockedAmount, ps.Id))
			amountTosend := int64(unlockedAmount) - ps.TotalTransferred.Amount.Int64()
			amount := sdk.NewCoin(ps.TotalAmount.Denom, sdk.NewInt(amountTosend))
			if amount.IsZero() || amount.IsNil() {
				continue
			}
			log.Info(fmt.Sprintf("Transferring amount %s to %s", amount.String(), recipient.String()))
			err = k.TransferAmountFromModuleAccount(ctx, recipient, sdk.NewCoins(amount))
			if err != nil {
				panic(err)
			}
			ps.TotalTransferred = ps.TotalTransferred.Add(amount)
			ps.LastTransferredAt = ctx.BlockTime()
			if ps.TotalTransferred.IsGTE(ps.TotalAmount) {
				log.Info(fmt.Sprintf("Updating payment status.. paymentId: %s", ps.Id))
				ps.Status = types.StatusCompleted
				k.SetInActivePayment(ctx, ps.Id)

				ctx.EventManager().EmitEvents(
					sdk.Events{sdk.NewEvent(
						"payment_complete",
						sdk.NewAttribute("payment-id", ps.Id),
						sdk.NewAttribute("sender", ps.Sender),
					),
					},
				)
			}
			k.SetPaymentStream(ctx, ps)

			ctx.EventManager().EmitEvents(
				sdk.Events{
					sdk.NewEvent(
						"payment_transfer",
						sdk.NewAttribute("payment-id", ps.Id),
						sdk.NewAttribute("sender", ps.Sender),
					),
				},
			)
		} else {
			continue
		}
	}
	log.Info("Completed streaming payments iteration ..")
	return nil
}
