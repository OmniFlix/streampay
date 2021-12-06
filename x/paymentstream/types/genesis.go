package types

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

func NewGenesisState(streams []PaymentStream, nextPaymentStreamNum uint64) *GenesisState {
	return &GenesisState{
		PaymentStreamsList:     streams,
		NextPaymentStreamCount: nextPaymentStreamNum,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PaymentStreamsList:     []PaymentStream{},
		NextPaymentStreamCount: DefaultIndex,
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in paymentstream
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
