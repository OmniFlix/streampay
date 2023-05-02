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
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	for _, sp := range gs.StreamPaymentsList {
		if err := validateStreamPayment(sp); err != nil {
			return err
		}
	}

	return ValidateNextStreamPaymentNumber(gs.NextStreamPaymentNumber)
}
