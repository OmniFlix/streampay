package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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
	if m.bankKeeper.BlockedAddr(recipient) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", msg.Recipient)
	}

	feePercentage := m.Keeper.GetStreamPaymentFeePercentage(ctx)
	requiredFeeAmount := sdk.NewCoin(msg.Amount.Denom, sdk.NewDecFromInt(msg.Amount.Amount).Mul(feePercentage).TruncateInt())
	if !msg.PaymentFee.Equal(requiredFeeAmount) {
		return nil, errorsmod.Wrap(types.ErrInvalidStreamPaymentFee, "fee coin didn't match with stream coin")
	}
	if requiredFeeAmount.Amount.GTE(sdk.NewInt(1)) {
		if err := m.distributionKeeper.FundCommunityPool(ctx, sdk.NewCoins(requiredFeeAmount), sender); err != nil {
			return nil, err
		}
	}
	streamPaymentId, err := m.Keeper.CreateStreamPayment(
		ctx,
		sender, recipient,
		msg.Amount,
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
