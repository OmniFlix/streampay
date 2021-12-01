package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/OmniFlix/payment-stream/testutil/keeper"
	"github.com/OmniFlix/payment-stream/x/paymentstream/keeper"
	"github.com/OmniFlix/payment-stream/x/paymentstream/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PaymentstreamKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
