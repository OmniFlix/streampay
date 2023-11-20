package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

var _ proto.Message = &StreamPayment{}

func NewStreamPayment(
	sender, recipient string,
	amount sdk.Coin,
	_type StreamType,
	startTime time.Time,
	endTime time.Time,
	periods []*Period,
	cancellable bool,
) StreamPayment {
	return StreamPayment{
		Sender:      sender,
		Recipient:   recipient,
		TotalAmount: amount,
		StartTime:   startTime,
		EndTime:     endTime,
		StreamType:  _type,
		Periods:     periods,
		Cancellable: cancellable,
	}
}

func (p StreamPayment) GetAmount() string {
	return p.TotalAmount.String()
}
