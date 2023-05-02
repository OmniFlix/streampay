package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/OmniFlix/streampay/x/streampay/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreateStreamPayment(ctx sdk.Context,
	sender, recipient sdk.AccAddress, amount sdk.Coin, paymentType types.PaymentType, endTime time.Time) error {
	if ctx.BlockTime().Unix() >= endTime.Unix() {
		return sdkerrors.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("endtime %s is not valid, should be a future timestamp", endTime.String()),
		)
	}
	if amount.IsNil() || amount.IsNegative() || amount.IsZero() {
		return sdkerrors.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("amount %s is not valid format", amount.String()),
		)
	}
	if err := k.TransferAmountToModuleAccount(ctx, sender, sdk.NewCoins(amount)); err != nil {
		return err
	}
	paymentStream := types.NewStreamPayment(sender.String(), recipient.String(), amount, paymentType, endTime)
	pNum := k.GetNextStreamPaymentNumber(ctx)

	paymentStream.Id = types.StreamPaymentPrefix + fmt.Sprint(pNum)
	paymentStream.LockHeight = ctx.BlockHeight()
	paymentStream.StartTime = ctx.BlockTime()
	paymentStream.TotalTransferred = sdk.NewCoin(paymentStream.TotalAmount.Denom, sdk.NewInt(0))

	k.SetStreamPayment(ctx, paymentStream)
	k.SetNextStreamPaymentNumber(ctx, pNum+1)

	return nil
}

func (k Keeper) CreateStreamPaymentFromModule(
	ctx sdk.Context,
	fromModule string,
	recipient sdk.AccAddress,
	amount sdk.Coin,
	paymentType types.PaymentType,
	endTime time.Time,
) error {
	if ctx.BlockTime().Unix() >= endTime.Unix() {
		return sdkerrors.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("endtime %s is not valid, should be a future timestamp", endTime.String()),
		)
	}
	if amount.IsNil() || amount.IsNegative() || amount.IsZero() {
		return sdkerrors.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("amount %s is not valid format", amount.String()),
		)
	}
	if err := k.TransferAmountFromModuleToModule(ctx, fromModule, sdk.NewCoins(amount)); err != nil {
		return err
	}
	fromModuleAddress := k.accountKeeper.GetModuleAccount(ctx, fromModule)
	paymentStream := types.NewStreamPayment(fromModuleAddress.String(), recipient.String(), amount, paymentType, endTime)
	pNum := k.GetNextStreamPaymentNumber(ctx)

	paymentStream.Id = types.StreamPaymentPrefix + fmt.Sprint(pNum)
	paymentStream.LockHeight = ctx.BlockHeight()
	paymentStream.StartTime = ctx.BlockTime()
	paymentStream.TotalTransferred = sdk.NewCoin(paymentStream.TotalAmount.Denom, sdk.NewInt(0))

	k.SetStreamPayment(ctx, paymentStream)
	k.SetNextStreamPaymentNumber(ctx, pNum+1)

	return nil
}
