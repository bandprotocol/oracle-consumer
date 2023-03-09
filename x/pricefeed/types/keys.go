package types

const (
	// ModuleName defines the module name
	ModuleName = "pricefeed"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_pricefeed"

	// Version defines the current version the IBC module supports
	Version = "bandchain-1"

	// PortID is the default port id that module binds to
	PortID = "pricefeed"
)

var (
	GlobalStoreKeyPrefix = []byte{0x00}

	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("pricefeed-port-")

	SymbolRequestStoreKeyPrefix = []byte{0x01}

	PriceStoreKeyPrefix = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func SymbolRequestStoreKey(symbol string) []byte {
	return append(SymbolRequestStoreKeyPrefix, []byte(symbol)...)
}

func PriceStoreKey(symbol string) []byte {
	return append(PriceStoreKeyPrefix, []byte(symbol)...)
}
