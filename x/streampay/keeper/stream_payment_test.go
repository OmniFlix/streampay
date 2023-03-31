package keeper_test

import (
	"github.com/OmniFlix/streampay/x/streampay/keeper"
	"github.com/OmniFlix/streampay/x/streampay/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStreamPayments(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StreamPayment {
	items := make([]types.StreamPayment, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)

		keeper.SetStreamPayment(ctx, items[i])
	}
	return items
}
