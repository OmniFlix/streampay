package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/streampay module errors
var (
	ErrInvalidAmount            = errorsmod.Register(ModuleName, 2, "invalid amount")
	ErrInvalidStreamPaymentType = errorsmod.Register(ModuleName, 3, "invalid stream payment type")
	ErrInvalidNextPaymentNumber = errorsmod.Register(ModuleName, 4, "invalid next payment number")
	ErrInvalidTimestamp         = errorsmod.Register(ModuleName, 5, "invalid timestamp")
	ErrInvalidStreamPaymentFee  = errorsmod.Register(ModuleName, 6, "invalid stream payment fee")
	ErrInvalidFee               = errorsmod.Register(ModuleName, 7, "invalid fee")
	ErrInvalidPeriods           = errorsmod.Register(ModuleName, 8, "invalid periods")
	ErrInvalidDuration          = errorsmod.Register(ModuleName, 9, "invalid duration")
)
