package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type StreamPaymentI interface {
	GetId() string
	GetTotalAmount() sdk.Coin
	GetSender() string
	GetRecipient() string
	GetStartTime() time.Time
	GetEndTime() time.Time
	GetTotalTransferred() sdk.Coin
}
