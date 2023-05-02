package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)
import "time"

const (
	MsgRoute          = "streampay"
	TypeMsgStreamSend = "stream_send"
	TypeMsgStopStream = "stop_stream"
)

var (
	_ sdk.Msg = &MsgStreamSend{}
	_ sdk.Msg = &MsgStopStream{}
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
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return err
	}
	if err := validateAmount(msg.Amount); err != nil {
		return err
	}
	if err := ValidateTimestamp(msg.EndTime); err != nil {
		return err
	}
	return validateStreamType(msg.StreamType)
}

// GetSignBytes Implements Msg.
func (msg *MsgStreamSend) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgStreamSend) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgStopStream(streamId string, sender string) *MsgStopStream {
	return &MsgStopStream{
		StreamId: streamId,
		Sender:   sender,
	}
}

func (msg MsgStopStream) Route() string { return MsgRoute }

func (msg MsgStopStream) Type() string { return TypeMsgStopStream }

func (msg MsgStopStream) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg *MsgStopStream) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgStopStream) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
