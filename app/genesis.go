package app

import (
	"encoding/json"
)

// GenesisState The genesis state of the blockchain is represented here as a map of raw json
// messages key'd by a identifier string.
// The identifier is used to determine which module genesis information belongs
// to so it may be appropriately routed during init chain.
// Within this application default genesis information is retrieved from
// the ModuleBasicManager which populates json from each BasicModule
// object provided to it during init.
type GenesisState map[string]json.RawMessage

var defaultGenesisState GenesisState = nil

// CloneGenesisState creates a deep clone of the provided GenesisState.
func cloneGenesisState(original GenesisState) GenesisState {
	clone := make(GenesisState, len(original))
	for key, value := range original {
		// Make a copy of the json.RawMessage (which is a []byte slice).
		copiedValue := make(json.RawMessage, len(value))
		copy(copiedValue, value)
		if len(copiedValue) == 0 {
			// If the value is empty, set it to nil.
			copiedValue = nil
		}
		clone[key] = copiedValue
	}
	return clone
}

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState() GenesisState {
	if defaultGenesisState != nil {
		return cloneGenesisState(defaultGenesisState)
	}
	encCfg := GetEncodingConfig()
	gen := ModuleBasics.DefaultGenesis(encCfg.Marshaler)
	defaultGenesisState = cloneGenesisState(gen)
	return gen
}
