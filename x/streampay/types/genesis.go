package types

// DefaultNextStreamPaymentNumber is the default number for next stream payment
const DefaultNextStreamPaymentNumber uint64 = 1

func NewGenesisState(streams []StreamPayment, nextStreamPaymentNum uint64) *GenesisState {
	return &GenesisState{
		StreamPaymentsList:      streams,
		NextStreamPaymentNumber: nextStreamPaymentNum,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		StreamPaymentsList:      []StreamPayment{},
		NextStreamPaymentNumber: DefaultNextStreamPaymentNumber,
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in streampay
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
