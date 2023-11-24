package types

import "fmt"

// DefaultNextStreamPaymentNumber is the default number for next stream payment
const DefaultNextStreamPaymentNumber uint64 = 1

func NewGenesisState(streams []StreamPayment, nextStreamPaymentNum uint64) *GenesisState {
	return &GenesisState{
		StreamPayments:          streams,
		NextStreamPaymentNumber: nextStreamPaymentNum,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		StreamPayments:          []StreamPayment{},
		NextStreamPaymentNumber: DefaultNextStreamPaymentNumber,
		Params:                  DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	var streamIds map[string]bool
	for _, sp := range gs.StreamPayments {
		if err := validateStreamPayment(sp); err != nil {
			return err
		}
		if streamIds[sp.Id] {
			return fmt.Errorf("duplicate stream id %s", sp.Id)
		}
	}

	if err := ValidateNextStreamPaymentNumber(gs.NextStreamPaymentNumber); err != nil {
		return err
	}
	return gs.Params.ValidateBasic()
}
