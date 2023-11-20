package types

const (
	// ModuleName defines the module name
	ModuleName = "streampay"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_streampay"

	// StreamPaymentPrefix payment streams id prefix
	StreamPaymentPrefix = "sp"
)

var (
	PrefixPaymentStreamId    = []byte{0x01}
	PrefixPaymentStreamCount = []byte{0x02}

	ParamsKey = []byte{0x03}
)

func KeyPrefixSteamPayment(p string) []byte {
	return append(PrefixPaymentStreamId, []byte(p)...)
}
