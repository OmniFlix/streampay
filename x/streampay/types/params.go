package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var DefaultStreamPaymentFee = sdk.NewInt64Coin("uspay", 10_000_000) // 10SPAY

var ParamStoreKeyStreamPaymentFee = []byte("StreamPaymentFee")

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewStreampayParams(streamPaymentFee sdk.Coin) Params {
	return Params{
		StreamPaymentFee: streamPaymentFee,
	}
}

// DefaultParams returns default streampay parameters
func DefaultParams() Params {
	return NewStreampayParams(
		DefaultStreamPaymentFee,
	)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyStreamPaymentFee, &p.StreamPaymentFee, validateStreamPaymentFee),
	}
}

// ValidateBasic performs basic validation on streampay parameters.
func (p Params) ValidateBasic() error {
	if err := validateStreamPaymentFee(p.StreamPaymentFee); err != nil {
		return err
	}
	return nil
}

// validateStreamPaymentFee performs validation of stream payment fee

func validateStreamPaymentFee(i interface{}) error {
	fee, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !fee.IsValid() || fee.IsZero() {
		return sdkerrors.Wrapf(
			ErrInvalidStreamPaymentFee,
			"invalid fee amount %s, only accepts positive amounts",
			fee.String(),
		)
	}
	return nil
}
