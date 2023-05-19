package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/OmniFlix/streampay/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	fee := m.Keeper.GetStreamPaymentFee(ctx)
	if !msg.Fee.Equal(fee) {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidFee,
			fmt.Sprintf("fee must be %s", fee.String()),
		)
	}
	if err := m.distributionKeeper.FundCommunityPool(ctx, sdk.NewCoins(msg.Fee), sender); err != nil {
		return nil, err
	}

	if err := m.Keeper.CreateStreamPayment(
		ctx,
		sender, recipient,
		msg.Amount,
		msg.StreamType,
		msg.Duration,
		msg.Periods,
		msg.Cancellable,
	); err != nil {
		return nil, err
	}

	return &types.MsgStreamSendResponse{}, nil
}

func (m msgServer) StopStream(goCtx context.Context, msg *types.MsgStopStream) (*types.MsgStopStreamResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	if err := m.Keeper.StopStreamPayment(
		ctx,
		msg.StreamId,
		sender,
	); err != nil {
		return nil, err
	}

	return &types.MsgStopStreamResponse{}, nil
}

func (m msgServer) ClaimStreamedAmount(
	goCtx context.Context, msg *types.MsgClaimStreamedAmount,
) (*types.MsgClaimStreamedAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	claimer, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		return nil, err
	}

	if err := m.Keeper.ClaimStreamedAmount(
		ctx,
		msg.StreamId,
		claimer,
	); err != nil {
		return nil, err
	}

	return &types.MsgClaimStreamedAmountResponse{}, nil
}
