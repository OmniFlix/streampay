package keeper_test

import (
	"context"
	keepertest "github.com/OmniFlix/streampay/testutil/keeper"
	"testing"

	"github.com/OmniFlix/streampay/x/streampay/keeper"
	"github.com/OmniFlix/streampay/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.StreamPayKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
