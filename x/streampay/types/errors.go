package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/streampay module errors
var (
	ErrInvalidAmount            = sdkerrors.Register(ModuleName, 2, "invalid amount")
	ErrInvalidPaymentStreamType = sdkerrors.Register(ModuleName, 3, "invalid payment stream type")
)
