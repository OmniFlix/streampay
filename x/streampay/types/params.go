package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var DefaultStreamPaymentFeePercentage = sdk.NewDecWithPrec(1, 2) // 1%

func NewStreampayParams(streamPaymentFeePercentage sdk.Dec) Params {
	return Params{
		StreamPaymentFeePercentage: streamPaymentFeePercentage,
	}
}

// DefaultParams returns default streampay parameters
func DefaultParams() Params {
	return NewStreampayParams(
		DefaultStreamPaymentFeePercentage,
	)
}

// ValidateBasic performs basic validation on streampay parameters.
func (p Params) ValidateBasic() error {
	if err := validateStreamPaymentFeePercentage(p.StreamPaymentFeePercentage); err != nil {
		return err
	}
	return nil
}

// validateStreamPaymentFee performs validation of stream payment fee

func validateStreamPaymentFeePercentage(i interface{}) error {
	fee, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if fee.IsNegative() || fee.GTE(sdk.OneDec()) {
		return errorsmod.Wrapf(
			ErrInvalidStreamPaymentFee,
			"invalid fee percentage %s, only accepts value which is positive and less than 1.00",
			fee.String(),
		)
	}
	return nil
}
