package keeper

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	"github.com/OmniFlix/streampay/v2/x/streampay/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type (
	Keeper struct {
		cdc                codec.BinaryCodec
		storeKey           storetypes.StoreKey
		memKey             storetypes.StoreKey
		accountKeeper      types.AccountKeeper
		bankKeeper         types.BankKeeper
		distributionKeeper types.DistributionKeeper
		authority          string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	distributionKeeper types.DistributionKeeper,
	authority string,
) *Keeper {
	// ensure streampay module account is set
	if addr := accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	return &Keeper{
		cdc:                cdc,
		storeKey:           storeKey,
		memKey:             memKey,
		accountKeeper:      accountKeeper,
		bankKeeper:         bankKeeper,
		distributionKeeper: distributionKeeper,
		authority:          authority,
	}
}

// GetAuthority returns the x/streampay module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreateStreamPayment(ctx sdk.Context,
	sender,
	recipient sdk.AccAddress,
	amount sdk.Coin,
	streamType types.StreamType,
	duration time.Duration,
	periods []*types.Period,
	cancellable bool,
) (string, error) {
	if duration <= 0 {
		return "", errorsmod.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("duration %s is not valid, should be a possitive value", duration.String()),
		)
	}
	if amount.IsNil() || amount.IsNegative() || amount.IsZero() {
		return "", errorsmod.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("amount %s is not valid format", amount.String()),
		)
	}
	endTime := ctx.BlockTime().Add(duration)
	if err := k.TransferAmountToModuleAccount(ctx, sender, sdk.NewCoins(amount)); err != nil {
		return "", err
	}
	streamPayment := types.NewStreamPayment(
		sender.String(),
		recipient.String(),
		amount,
		streamType,
		ctx.BlockTime(),
		endTime,
		periods,
		cancellable,
	)
	pNum := k.GetNextStreamPaymentNumber(ctx)

	// update stream payment
	streamPayment.Id = types.StreamPaymentPrefix + fmt.Sprint(pNum)
	streamPayment.StreamedAmount = sdk.NewCoin(streamPayment.TotalAmount.Denom, sdk.NewInt(0))

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

	return streamPayment.Id, nil
}

func (k Keeper) StopStreamPayment(ctx sdk.Context, streamId string, sender sdk.AccAddress) error {
	streamPayment, ok := k.GetStreamPayment(ctx, streamId)
	if !ok {
		return errorsmod.Wrapf(
			sdkerrors.ErrNotFound,
			fmt.Sprintf("no stream payment found with id %s", streamId),
		)
	}
	if sender.String() != streamPayment.Sender {
		return errorsmod.Wrapf(
			sdkerrors.ErrUnauthorized,
			fmt.Sprintf("address %s is not allowed to stop the stream payment", streamId),
		)
	}
	if !streamPayment.Cancellable {
		return errorsmod.Wrapf(
			sdkerrors.ErrUnauthorized,
			fmt.Sprintf("stream payment %s is not cancellable", streamId),
		)
	}
	if ctx.BlockTime().Unix() > streamPayment.EndTime.Unix() {
		return errorsmod.Wrapf(
			sdkerrors.ErrUnauthorized,
			fmt.Sprintf("ended stream payment cannot be canceled, stream payment %s", streamId),
		)
	}
	streamedAmount := float64(0)
	if streamPayment.StreamType == types.TypeContinuous {
		streamedAmount = k.getStreamedAmount(ctx, streamPayment)
	} else if streamPayment.StreamType == types.TypePeriodic {
		streamedAmount = k.getStreamedAmountForPeriodicStreamPayment(ctx, streamPayment)
	}
	lastStreamedAmount := streamPayment.StreamedAmount
	streamPayment.StreamedAmount.Amount = sdk.NewInt(int64(streamedAmount))
	remainingAmount := streamPayment.TotalAmount.Sub(streamPayment.StreamedAmount)
	// transfer remaining amount to sender
	if err := k.TransferAmountFromModuleAccount(ctx, sender, sdk.NewCoins(remainingAmount)); err != nil {
		return err
	}
	amountToRecipient := streamPayment.StreamedAmount.Sub(lastStreamedAmount)
	recipient, _ := sdk.AccAddressFromBech32(streamPayment.Recipient)
	// transfer streamed amount to recipient
	if err := k.TransferAmountFromModuleAccount(
		ctx, recipient, sdk.NewCoins(amountToRecipient),
	); err != nil {
		return err
	}
	k.RemoveStreamPayment(ctx, streamId)

	// emit events
	k.emitStopStreamPaymentEvent(
		ctx,
		streamPayment.Id,
		streamPayment.Sender,
		streamPayment.Recipient,
		remainingAmount,
		amountToRecipient,
	)
	k.emitStreamPaymentEndEvent(ctx, streamId, streamPayment.Sender)
	return nil
}

func (k Keeper) ClaimStreamedAmount(ctx sdk.Context, streamId string, claimer sdk.AccAddress) error {
	streamPayment, ok := k.GetStreamPayment(ctx, streamId)
	if !ok {
		return errorsmod.Wrapf(
			sdkerrors.ErrNotFound,
			fmt.Sprintf("no stream payment found with id %s", streamId),
		)
	}
	claimerAddr := claimer.String()
	if claimerAddr != streamPayment.Recipient {
		return errorsmod.Wrapf(
			sdkerrors.ErrUnauthorized,
			fmt.Sprintf("address %s is not allowed to claim the stream payment", claimerAddr),
		)
	}
	if ctx.BlockTime().Unix() < streamPayment.StartTime.Unix() {
		return errorsmod.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("stream payment %s is not started yet", streamId),
		)
	}
	if streamPayment.StreamedAmount.IsGTE(streamPayment.TotalAmount) {
		return errorsmod.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("stream payment %s is already fully claimed", streamId),
		)
	}
	switch streamPayment.GetStreamType() {
	case types.TypeDelayed:
		if err := k.claimDelayedStreamPayment(ctx, streamPayment, claimer); err != nil {
			return err
		}
	case types.TypeContinuous:
		if err := k.claimContinuousStreamPayment(ctx, streamPayment, claimer); err != nil {
			return err
		}
	case types.TypePeriodic:
		if err := k.claimPeriodicStreamPayment(ctx, streamPayment, claimer); err != nil {
			return err
		}

	default:
		return errorsmod.Wrapf(
			types.ErrInvalidStreamPaymentType,
			fmt.Sprintf("stream payment %s has invalid type", streamId),
		)
	}
	return nil
}

func (k Keeper) claimDelayedStreamPayment(ctx sdk.Context, streamPayment types.StreamPayment, claimer sdk.AccAddress) error {
	if ctx.BlockTime().Unix() < streamPayment.EndTime.Unix() {
		return errorsmod.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("stream payment %s is delayed type and not ended yet", streamPayment.Id),
		)
	}
	if err := k.TransferAmountFromModuleAccount(ctx, claimer, sdk.NewCoins(streamPayment.TotalAmount)); err != nil {
		return err
	}
	// Remove stream payment
	k.RemoveStreamPayment(ctx, streamPayment.Id)

	// Emit events
	k.emitStreamPaymentClaimEvent(ctx, streamPayment.Id, streamPayment.Recipient, streamPayment.TotalAmount)
	k.emitStreamPaymentEndEvent(ctx, streamPayment.Id, streamPayment.Sender)

	return nil
}

func (k Keeper) claimContinuousStreamPayment(ctx sdk.Context, streamPayment types.StreamPayment, claimer sdk.AccAddress) error {
	streamedAmount := k.getStreamedAmount(ctx, streamPayment)
	amountToSend := int64(streamedAmount) - streamPayment.StreamedAmount.Amount.Int64()
	amount := sdk.NewCoin(streamPayment.TotalAmount.Denom, sdk.NewInt(amountToSend))

	if amount.IsZero() || amount.IsNil() {
		return errorsmod.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("no valid amount to claim for stream payment %s ", streamPayment.Id),
		)
	}
	if err := k.TransferAmountFromModuleAccount(ctx, claimer, sdk.NewCoins(amount)); err != nil {
		return err
	}

	// update stream payment
	streamPayment.StreamedAmount = streamPayment.StreamedAmount.Add(amount)
	streamPayment.LastClaimedAt = ctx.BlockTime()
	if streamPayment.StreamedAmount.IsGTE(streamPayment.TotalAmount) {
		k.RemoveStreamPayment(ctx, streamPayment.Id)
		k.emitStreamPaymentEndEvent(ctx, streamPayment.Id, streamPayment.Sender)
	} else {
		k.SetStreamPayment(ctx, streamPayment)
	}
	// emit events
	k.emitStreamPaymentClaimEvent(ctx, streamPayment.Id, streamPayment.Recipient, amount)
	return nil
}

func (k Keeper) claimPeriodicStreamPayment(ctx sdk.Context, streamPayment types.StreamPayment, claimer sdk.AccAddress) error {
	streamedAmount := k.getStreamedAmountForPeriodicStreamPayment(ctx, streamPayment)
	amountToSend := int64(streamedAmount) - streamPayment.StreamedAmount.Amount.Int64()
	amount := sdk.NewCoin(streamPayment.TotalAmount.Denom, sdk.NewInt(amountToSend))

	if amount.IsZero() || amount.IsNil() {
		return errorsmod.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("no valid amount to claim for stream payment %s ", streamPayment.Id),
		)
	}
	if err := k.TransferAmountFromModuleAccount(ctx, claimer, sdk.NewCoins(amount)); err != nil {
		return err
	}

	// update stream payment
	streamPayment.StreamedAmount = streamPayment.StreamedAmount.Add(amount)
	streamPayment.LastClaimedAt = ctx.BlockTime()
	if streamPayment.StreamedAmount.IsGTE(streamPayment.TotalAmount) {
		k.RemoveStreamPayment(ctx, streamPayment.Id)
		k.emitStreamPaymentEndEvent(ctx, streamPayment.Id, streamPayment.Sender)
	} else {
		k.SetStreamPayment(ctx, streamPayment)
	}
	// emit events
	k.emitStreamPaymentClaimEvent(ctx, streamPayment.Id, streamPayment.Recipient, amount)
	return nil
}

// CreateModuleAccount creates a module account with minting and burning capabilities
// This account isn't intended to store any coins,
// it purely mints and burns them on behalf of the admin of respective denoms,
// and sends to the relevant address.
func (k Keeper) CreateModuleAccount(ctx sdk.Context) {
	moduleAcc := authtypes.NewEmptyModuleAccount(types.ModuleName, authtypes.Minter, authtypes.Burner)
	k.accountKeeper.SetModuleAccount(ctx, moduleAcc)
}
