package types

import sdk "github.com/cosmos/cosmos-sdk/types"
import "time"

const (
	MsgRoute          = "paymentstream"
	TypeMsgStreamSend = "stream_send"
	// DoNotModify used to indicate that some field should not be updated
	DoNotModify = "[do-not-modify]"
)

var (
	_ sdk.Msg = &MsgStreamSend{}
)

func NewMsgStreamSend(sender, recipient string, amount sdk.Coin, _type PaymentType, endTime time.Time) *MsgStreamSend {
	return &MsgStreamSend{
		Sender:     sender,
		Recipient:  recipient,
		Amount:     amount,
		StreamType: _type,
		EndTime:    endTime,
	}
}

func (msg MsgStreamSend) Route() string { return MsgRoute }

func (msg MsgStreamSend) Type() string { return TypeMsgStreamSend }

func (msg MsgStreamSend) ValidateBasic() error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgStreamSend) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgStreamSend) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
