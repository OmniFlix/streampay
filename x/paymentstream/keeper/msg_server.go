package keeper

import (
	"context"
	"github.com/OmniFlix/payment-stream/x/paymentstream/types"
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

	// TODO: Handling the message
	_ = ctx
	return &types.MsgStreamSendResponse{}, nil
}
