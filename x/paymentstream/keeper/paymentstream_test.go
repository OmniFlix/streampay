package keeper_test

import (
	"github.com/OmniFlix/payment-stream/x/paymentstream/keeper"
	"github.com/OmniFlix/payment-stream/x/paymentstream/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPaymentstreams(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PaymentStream {
	items := make([]types.PaymentStream, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)

		keeper.SetPaymentStream(ctx, items[i])
	}
	return items
}
