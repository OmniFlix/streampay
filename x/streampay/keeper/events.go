package keeper

import (
	"time"

	"github.com/OmniFlix/streampay/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) emitStreamPaymentTransferEvent(ctx sdk.Context, streamId, recipient string, amount sdk.Coin) {
	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeTransferStreamPayment,
				sdk.NewAttribute(types.EventAttributePaymentId, streamId),
				sdk.NewAttribute(types.EventAttributeRecipient, recipient),
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
				sdk.NewAttribute(types.EventAttributeEndTime, endTime.String()),
			),
		},
	)
}

func (k Keeper) emitStopStreamPaymentEvent(ctx sdk.Context, streamId, sender string) {
	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeStopStreamPayment,
				sdk.NewAttribute(types.EventAttributePaymentId, streamId),
				sdk.NewAttribute(types.EventAttributeSender, sender),
			),
		},
	)
}
