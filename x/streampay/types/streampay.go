package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"time"
)

var (
	_ proto.Message = &StreamPayment{}
)

func NewStreamPayment(sender, recipient string, amount sdk.Coin, _type PaymentType, endTime time.Time) StreamPayment {
	return StreamPayment{
		Sender:      sender,
		Recipient:   recipient,
		TotalAmount: amount,
		EndTime:     endTime,
		StreamType:  _type,
	}
}

func (p StreamPayment) GetAmount() string {
	return p.TotalAmount.String()
}
