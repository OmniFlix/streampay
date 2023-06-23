package keeper

import (
	"time"

	"github.com/OmniFlix/streampay/v2/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) emitStreamPaymentClaimEvent(ctx sdk.Context, streamId, claimer string, amount sdk.Coin) {
	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeClaimStreamedAmount,
				sdk.NewAttribute(types.EventAttributePaymentId, streamId),
				sdk.NewAttribute(types.EventAttributeClaimer, claimer),
				sdk.NewAttribute(types.EventAttributeAmount, amount.String()),
			),
		},
	)
}

func (k Keeper) emitStreamPaymentEndEvent(ctx sdk.Context, streamId, sender string) {
	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeEndStreamPayment,
				sdk.NewAttribute(types.EventAttributePaymentId, streamId),
				sdk.NewAttribute(types.EventAttributeSender, sender),
			),
		},
	)
}

func (k Keeper) emitCreateStreamPaymentEvent(ctx sdk.Context,
	streamId, sender, recipient string,
	amount sdk.Coin,
	paymentType string,
	endTime time.Time,
) {
	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeCreateStreamPayment,
				sdk.NewAttribute(types.EventAttributePaymentId, streamId),
				sdk.NewAttribute(types.EventAttributeSender, sender),
				sdk.NewAttribute(types.EventAttributeRecipient, recipient),
				sdk.NewAttribute(types.EventAttributeAmount, amount.String()),
				sdk.NewAttribute(types.EventAttributePaymentType, paymentType),
				sdk.NewAttribute(types.EventAttributeEndTime, endTime.Format(time.RFC3339)),
			),
		},
	)
}

func (k Keeper) emitStopStreamPaymentEvent(
	ctx sdk.Context,
	streamId, sender, recipient string,
	returnedAmount, streamedAmount sdk.Coin,
) {
	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeStopStreamPayment,
				sdk.NewAttribute(types.EventAttributePaymentId, streamId),
				sdk.NewAttribute(types.EventAttributeSender, sender),
				sdk.NewAttribute(types.EventAttributeRecipient, recipient),
				sdk.NewAttribute(types.EventAttributeReturnedAmount, returnedAmount.String()),
				sdk.NewAttribute(types.EventAttributeStreamedAmount, streamedAmount.String()),
			),
		},
	)
}
