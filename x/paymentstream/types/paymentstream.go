package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"time"
)

var (
	_ proto.Message = &PaymentStream{}
)

func NewPaymentStream(sender, recipient string, amount sdk.Coin, _type PaymentType, endTime time.Time) PaymentStream {
	return PaymentStream{
		Sender:      sender,
		Recipient:   recipient,
		TotalAmount: amount,
		EndTime:     endTime,
		StreamType:  _type,
	}
}

func (p PaymentStream) GetAmount() string {
	return p.TotalAmount.String()
}
