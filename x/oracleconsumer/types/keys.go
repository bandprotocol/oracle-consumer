package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "oracleconsumer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oracleconsumer"

	// Version defines the current version the IBC module supports
	Version = "oracleconsumer-1"

	// PortID is the default port id that module binds to
	PortID = "oracleconsumer"
)

var (
	GlobalStoreKeyPrefix = []byte{0x00}

	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("oracleconsumer-port-")

	RequestIntervalCountStoreKey = append(GlobalStoreKeyPrefix, []byte("RequestIntervalCount")...)

	RequestIntervalStoreKeyPrefix = []byte{0x01}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func RequestIntervalStoreKey(requestIntervalID uint64) []byte {
	return append(RequestIntervalStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(requestIntervalID))...)
}
