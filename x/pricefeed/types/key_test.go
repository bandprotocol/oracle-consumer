package types_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

func TestSymbolRequestStoreKey(t *testing.T) {
	expect, err := hex.DecodeString("01425443")
	require.NoError(t, err)
	require.Equal(t, expect, types.SymbolRequestStoreKey("BTC"))
}

func TestPriceStoreKey(t *testing.T) {
	expect, err := hex.DecodeString("02425443")
	require.NoError(t, err)
	require.Equal(t, expect, types.PriceStoreKey("BTC"))
}
