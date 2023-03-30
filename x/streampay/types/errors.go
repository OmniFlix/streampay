package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/streampay module sentinel errors
var (
	ErrInvalidAmount            = sdkerrors.Register(ModuleName, 1, "invalid amount")
	ErrInvalidPaymentStreamType = sdkerrors.Register(ModuleName, 2, "invalid payment stream type")
)
