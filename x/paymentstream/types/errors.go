package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/paymentstream module sentinel errors
var (
	ErrSample             = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidAmount      = sdkerrors.Register(ModuleName, 1101, "invalid amount")
	ErrUnableToLockAmount = sdkerrors.Register(ModuleName, 1102, "unable to lock amount")
)
