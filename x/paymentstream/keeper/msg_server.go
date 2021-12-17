package keeper

import (
	"context"
	"fmt"
	"github.com/OmniFlix/payment-stream/x/paymentstream/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) StreamSend(goCtx context.Context, msg *types.MsgStreamSend) (*types.MsgStreamSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}
	if ctx.BlockTime().Unix() >= msg.EndTime.Unix() {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("endtime %s is not valid, should be a future timestamp", msg.EndTime.String()),
		)
	}
	if msg.Amount.IsNil() || msg.Amount.IsNegative() || msg.Amount.IsZero() {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("amount %s is not valid format", msg.Amount.String()),
		)
	}
	err = m.Keeper.LockAmountToModuleAccount(ctx, sender, sdk.NewCoins(msg.Amount))
	if err != nil {
		return nil, sdkerrors.Wrapf(
			types.ErrUnableToLockAmount,
			fmt.Sprintf("unable to lock amount from address %s", sender.String()),
		)
	}

	paymentStream := types.NewPaymentStream(sender.String(), recipient.String(), msg.Amount, msg.StreamType, msg.EndTime)
	m.Keeper.StartPaymentStream(ctx, paymentStream)

	return &types.MsgStreamSendResponse{}, nil
}
