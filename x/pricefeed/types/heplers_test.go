package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bandprotocol/consumer/x/pricefeed/types"
)

func TestCalculateGas(t *testing.T) {
	gas := types.CalculateGas(5, 1, 1)

	require.Equal(t, uint64(6), gas)
}
