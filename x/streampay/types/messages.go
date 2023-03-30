package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/ibc-go/modules/apps/transfer/types"
)
import "time"

const (
	MsgRoute          = "streampay"
	TypeMsgStreamSend = "stream_send"
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
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return err
	}
	if !msg.Amount.IsValid() || msg.Amount.IsNil() || msg.Amount.Amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrapf(
			types.ErrInvalidAmount,
			fmt.Sprintf("amount %s is not valid", msg.Amount.String()),
		)
	}
	if !(msg.StreamType == TypeDelayed || msg.StreamType == TypeContinuous) {
		return sdkerrors.Wrapf(
			ErrInvalidPaymentStreamType,
			fmt.Sprintf("payment stream type %s is not valid", msg.Type()),
		)
	}
	return nil
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
