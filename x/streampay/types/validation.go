package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func validateStreamPayment(streamPayment StreamPayment) error {
	if _, err := sdk.AccAddressFromBech32(streamPayment.Sender); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(streamPayment.Recipient); err != nil {
		return err
	}
	if err := validateAmount(streamPayment.TotalAmount); err != nil {
		return err
	}
	if err := ValidateTimestamp(streamPayment.StartTime); err != nil {
		return err
	}
	if err := ValidateTimestamp(streamPayment.EndTime); err != nil {
		return err
	}
	return validateStreamType(streamPayment.StreamType)
}

func validateAmount(amount sdk.Coin) error {
	if !amount.IsValid() || amount.IsNil() || amount.Amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrapf(
			ErrInvalidAmount,
			fmt.Sprintf("amount %s is not valid", amount.String()),
		)
	}
	return nil
}

func validateStreamType(_type StreamType) error {
	if !(_type == TypeDelayed || _type == TypeContinuous || _type == TypePeriodic) {
		return sdkerrors.Wrapf(
			ErrInvalidStreamPaymentType,
			fmt.Sprintf("stream payment type %s is not valid", _type),
		)
	}
	return nil
}

func ValidateNextStreamPaymentNumber(n interface{}) error {
	_, ok := n.(uint64)
	if !ok {
		return sdkerrors.Wrapf(ErrInvalidNextPaymentNumber, "invalid value for next payment number: %v", n)
	}
	return nil
}

func ValidateTimestamp(t interface{}) error {
	_, ok := t.(time.Time)
	if !ok {
		return sdkerrors.Wrapf(ErrInvalidTimestamp, "invalid timestamp: %T", t)
	}
	return nil
}

func validatePeriods(periods []*Period, totalAmount sdk.Coin) error {
	if len(periods) == 0 {
		return sdkerrors.Wrapf(ErrInvalidPeriods, "periods cannot be empty")
	}
	totalAmt := int64(0)
	for _, period := range periods {
		if period.Amount < 1 {
			return sdkerrors.Wrapf(ErrInvalidPeriods, "invalid period amount: %d", period.Amount)
		}
		if period.Duration < 1 {
			return sdkerrors.Wrapf(ErrInvalidPeriods, "invalid period duration: %d", period.Duration)
		}
		totalAmt += period.Amount
	}
	if totalAmt != totalAmount.Amount.Int64() {
		return sdkerrors.Wrapf(ErrInvalidPeriods, "period amounts do not add up to total amount")
	}
	return nil
}
