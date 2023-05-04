package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/OmniFlix/streampay/x/streampay/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	streamPayment := types.NewStreamPayment(sender.String(), recipient.String(), amount, paymentType, endTime)
	pNum := k.GetNextStreamPaymentNumber(ctx)

	// update stream payment
	streamPayment.Id = types.StreamPaymentPrefix + fmt.Sprint(pNum)
	streamPayment.LockHeight = ctx.BlockHeight()
	streamPayment.StartTime = ctx.BlockTime()
	streamPayment.TotalTransferred = sdk.NewCoin(streamPayment.TotalAmount.Denom, sdk.NewInt(0))

	k.SetStreamPayment(ctx, streamPayment)
	k.SetNextStreamPaymentNumber(ctx, pNum+1)

	// emit events
	k.emitCreateStreamPaymentEvent(
		ctx,
		streamPayment.Id,
		streamPayment.Sender,
		streamPayment.Recipient,
		streamPayment.TotalAmount,
		streamPayment.StreamType.String(),
		streamPayment.EndTime,
	)

	return nil
}

func (k Keeper) StopStreamPayment(ctx sdk.Context, streamId string, sender sdk.AccAddress) error {
	streamPayment, ok := k.GetStreamPayment(ctx, streamId)
	if !ok {
		return sdkerrors.Wrapf(
			sdkerrors.ErrNotFound,
			fmt.Sprintf("no stream payment found with id %s", streamId),
		)
	}
	if sender.String() != streamPayment.Sender {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnauthorized,
			fmt.Sprintf("address %s is not allowed to stop the stream payment", streamId),
		)
	}
	remainingAmount := streamPayment.TotalAmount.Sub(streamPayment.TotalTransferred)
	if err := k.TransferAmountFromModuleAccount(ctx, sender, sdk.NewCoins(remainingAmount)); err != nil {
		return err
	}
	k.RemoveStreamPayment(ctx, streamId)

	// emit events
	k.emitStopStreamPaymentEvent(
		ctx,
		streamPayment.Id,
		streamPayment.Sender,
	)

	return nil
}
