package types

const (
	// ModuleName defines the module name
	ModuleName = "paymentstream"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_paymentstream"

	// payment streams id prefix
	PaymentStreamPrefix = "ps"
)

var (
	PrefixPaymentStreamId        = []byte{0x01}
	PrefixPaymentStreamCount     = []byte{0x02}
	PrefixPaymentActive          = []byte{0x03}
	PrefixPaymentInActive        = []byte{0x04}
	PrefixSenderPaymentStream    = []byte{0x05}
	PrefixRecipientPaymentStream = []byte{0x06}
)

func KeyPrefix(p string) []byte {
	return append(PrefixPaymentStreamId, []byte(p)...)
}
func ActivePaymentPrefix(id string) []byte {
	return append(PrefixPaymentActive, []byte(id)...)
}

func InActivePaymentPrefix(id string) []byte {
	return append(PrefixPaymentInActive, []byte(id)...)
}
func IdFromPaymentStatusKey(key []byte) string {
	return string(key[1:])
}

func KeySenderPaymentStream(address, id string) []byte {
	return append(append(PrefixSenderPaymentStream, []byte(address)...), []byte(id)...)
}

func KeyRecipientPaymentStream(address, id string) []byte {
	return append(append(PrefixRecipientPaymentStream, []byte(address)...), []byte(id)...)
}
