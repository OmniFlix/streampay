package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/OmniFlix/streampay/v2/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
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

func (m msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if m.authority != req.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", m.authority, req.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

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
	feePercentage := m.Keeper.GetStreamPaymentFeePercentage(ctx)
	feeAmount := sdk.NewCoin(msg.Amount.Denom, msg.Amount.Amount.ToLegacyDec().Mul(feePercentage).TruncateInt())
	amountToSend := msg.Amount.SubAmount(feeAmount.Amount)

	if err := m.distributionKeeper.FundCommunityPool(ctx, sdk.NewCoins(feeAmount), sender); err != nil {
		return nil, err
	}

	streamPaymentId, err := m.Keeper.CreateStreamPayment(
		ctx,
		sender, recipient,
		amountToSend,
		msg.StreamType,
		msg.Duration,
		msg.Periods,
		msg.Cancellable,
	)
	if err != nil {
		return nil, err
	}
	return &types.MsgStreamSendResponse{
		StreamId: streamPaymentId,
	}, nil
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
