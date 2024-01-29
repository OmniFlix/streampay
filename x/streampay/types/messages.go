package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)
import "time"

const (
	MsgRoute                   = "streampay"
	TypeMsgStreamSend          = "stream_send"
	TypeMsgStopStream          = "stop_stream"
	TypeMsgClaimStreamedAmount = "claim_streamed_amount"
)

var (
	_ sdk.Msg = &MsgUpdateParams{}
	_ sdk.Msg = &MsgStreamSend{}
	_ sdk.Msg = &MsgStopStream{}
	_ sdk.Msg = &MsgClaimStreamedAmount{}
)

func NewMsgStreamSend(
	sender,
	recipient string,
	amount sdk.Coin,
	_type StreamType,
	duration time.Duration,
	periods []*Period,
	cancellable bool,
	paymentFee sdk.Coin,
) *MsgStreamSend {
	return &MsgStreamSend{
		Sender:      sender,
		Recipient:   recipient,
		Amount:      amount,
		StreamType:  _type,
		Duration:    duration,
		Periods:     periods,
		Cancellable: cancellable,
		PaymentFee:  paymentFee,
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
	if err := validateStreamAmount(msg.Amount); err != nil {
		return err
	}
	if err := validateFeeAmount(msg.PaymentFee); err != nil {
		return err
	}
	if err := ValidateDuration(msg.Duration); err != nil {
		return err
	}
	if msg.StreamType == TypePeriodic {
		if err := validatePeriods(msg.Periods, msg.Amount, msg.Duration); err != nil {
			return err
		}
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

func NewMsgClaimStreamedAmount(streamId string, claimer string) *MsgClaimStreamedAmount {
	return &MsgClaimStreamedAmount{
		StreamId: streamId,
		Claimer:  claimer,
	}
}

func (msg MsgClaimStreamedAmount) Route() string { return MsgRoute }

func (msg MsgClaimStreamedAmount) Type() string { return TypeMsgClaimStreamedAmount }

func (msg MsgClaimStreamedAmount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		return err
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg *MsgClaimStreamedAmount) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgClaimStreamedAmount) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

// MsgUpdateParams

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}

	return m.Params.ValidateBasic()
}
