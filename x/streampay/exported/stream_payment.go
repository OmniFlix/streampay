package exported

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StreamPaymentI interface {
	GetId() string
	GetTotalAmount() sdk.Coin
	GetSender() string
	GetRecipient() string
	GetStartTime() time.Time
	GetEndTime() time.Time
	GetStreamedAmount() sdk.Coin
}

type (
	ParamSet = paramtypes.ParamSet

	// Subspace defines an interface that implements the legacy x/params Subspace
	// type.
	//
	// NOTE: This is used solely for migration of x/params managed parameters.
	Subspace interface {
		GetParamSet(ctx sdk.Context, ps ParamSet)
	}
)
