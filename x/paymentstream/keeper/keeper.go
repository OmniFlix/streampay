package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/OmniFlix/payment-stream/x/paymentstream/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		accountKeeper: accountKeeper, bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) StartPaymentStream(ctx sdk.Context, payment types.PaymentStream) {

	pNum := k.GetNextPaymentStreamNumber(ctx)

	payment.Id = "payment" + fmt.Sprint(pNum)
	payment.LockHeight = ctx.BlockHeight()
	payment.StartTime = ctx.BlockTime()
	payment.TotalTransferred = sdk.NewCoin(payment.TotalAmount.Denom, sdk.NewInt(0))
	payment.Status = types.StatusOpen

	k.SetPaymentStream(ctx, payment)
	k.SetNextPaymentStreamNumber(ctx, pNum+1)
}
