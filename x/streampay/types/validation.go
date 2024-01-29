package types

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func validateStreamPayment(streamPayment StreamPayment) error {
	if _, err := sdk.AccAddressFromBech32(streamPayment.Sender); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(streamPayment.Recipient); err != nil {
		return err
	}
	if err := validateStreamAmount(streamPayment.TotalAmount); err != nil {
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

func validateStreamAmount(amount sdk.Coin) error {
	if !amount.IsValid() || amount.IsNil() || amount.Amount.LTE(sdk.ZeroInt()) {
		return errorsmod.Wrapf(
			ErrInvalidAmount,
			fmt.Sprintf("amount %s is not valid", amount.String()),
		)
	}
	return nil
}

func validateFeeAmount(amount sdk.Coin) error {
	if !amount.IsValid() || amount.IsNil() || amount.Amount.LT(sdk.ZeroInt()) {
		return errorsmod.Wrapf(
			ErrInvalidAmount,
			fmt.Sprintf("fee amount %s is not valid", amount.String()),
		)
	}
	return nil
}

func validateStreamType(_type StreamType) error {
	if !(_type == TypeDelayed || _type == TypeContinuous || _type == TypePeriodic) {
		return errorsmod.Wrapf(
			ErrInvalidStreamPaymentType,
			fmt.Sprintf("stream payment type %s is not valid", _type),
		)
	}
	return nil
}

func ValidateNextStreamPaymentNumber(n interface{}) error {
	_, ok := n.(uint64)
	if !ok {
		return errorsmod.Wrapf(ErrInvalidNextPaymentNumber, "invalid value for next payment number: %v", n)
	}
	return nil
}

func ValidateTimestamp(t interface{}) error {
	_, ok := t.(time.Time)
	if !ok {
		return errorsmod.Wrapf(ErrInvalidTimestamp, "invalid timestamp: %T", t)
	}
	return nil
}

func ValidateDuration(d interface{}) error {
	duration, ok := d.(time.Duration)
	if !ok {
		return errorsmod.Wrapf(ErrInvalidDuration, "invalid duration: %v", d)
	}
	if duration < 1 {
		return errorsmod.Wrapf(ErrInvalidDuration, "invalid duration: %v", duration)
	}
	return nil
}

func validatePeriods(periods []*Period, totalAmount sdk.Coin, totalDuration time.Duration) error {
	if len(periods) == 0 {
		return errorsmod.Wrapf(ErrInvalidPeriods, "periods cannot be empty")
	}
	totalAmt := int64(0)
	totalDur := int64(0)
	for _, period := range periods {
		if period.Amount < 1 {
			return errorsmod.Wrapf(ErrInvalidPeriods, "invalid period amount: %d", period.Amount)
		}
		if period.Duration < 1 {
			return errorsmod.Wrapf(ErrInvalidPeriods, "invalid period duration: %d", period.Duration)
		}
		totalAmt += period.Amount
		totalDur += period.Duration
	}
	if totalAmt != totalAmount.Amount.Int64() {
		return errorsmod.Wrapf(ErrInvalidPeriods, "period amounts do not add up to total amount")
	}
	if totalDur != int64(totalDuration.Seconds()) {
		return errorsmod.Wrapf(ErrInvalidPeriods, "period durations do not add up to total duration")
	}
	return nil
}
