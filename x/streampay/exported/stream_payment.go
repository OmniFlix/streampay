package exported

import (
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
