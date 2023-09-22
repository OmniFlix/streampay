package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var ParamStoreKeyStreamPaymentFee = []byte("StreamPaymentFee")

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyStreamPaymentFee, &p.StreamPaymentFee, validateStreamPaymentFee),
	}
}
