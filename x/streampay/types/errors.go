package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/streampay module errors
var (
	ErrInvalidAmount            = sdkerrors.Register(ModuleName, 2, "invalid amount")
	ErrInvalidStreamPaymentType = sdkerrors.Register(ModuleName, 3, "invalid stream payment type")
	ErrInvalidNextPaymentNumber = sdkerrors.Register(ModuleName, 4, "invalid next payment number")
	ErrInvalidTimestamp         = sdkerrors.Register(ModuleName, 5, "invalid timestamp")
	ErrInvalidStreamPaymentFee  = sdkerrors.Register(ModuleName, 6, "invalid stream payment fee")
	ErrInvalidFee               = sdkerrors.Register(ModuleName, 7, "invalid fee")
	ErrInvalidPeriods           = sdkerrors.Register(ModuleName, 8, "invalid periods")
	ErrInvalidDuration          = sdkerrors.Register(ModuleName, 9, "invalid duration")
)
