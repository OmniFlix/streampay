package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

var _ proto.Message = &StreamPayment{}

func NewStreamPayment(
	sender, recipient string,
	amount sdk.Coin,
	_type StreamType,
	endTime time.Time,
	periods []*Period,
) StreamPayment {
	return StreamPayment{
		Sender:      sender,
		Recipient:   recipient,
		TotalAmount: amount,
		EndTime:     endTime,
		StreamType:  _type,
		Periods:     periods,
	}
}

func (p StreamPayment) GetAmount() string {
	return p.TotalAmount.String()
}
