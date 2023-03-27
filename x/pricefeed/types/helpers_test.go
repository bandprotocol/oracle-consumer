package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

func TestCalculateGas(t *testing.T) {
	testCases := []struct {
		name        string
		base        uint64
		each        uint64
		n           uint64
		expectedGas uint64
	}{
		{
			name:        "Base case",
			base:        10,
			each:        5,
			n:           4,
			expectedGas: 30,
		},
		{
			name:        "Zero base",
			base:        0,
			each:        5,
			n:           4,
			expectedGas: 20,
		},
		{
			name:        "Zero each",
			base:        10,
			each:        0,
			n:           4,
			expectedGas: 10,
		},
		{
			name:        "Zero n",
			base:        10,
			each:        5,
			n:           0,
			expectedGas: 10,
		},
		{
			name:        "All zero",
			base:        0,
			each:        0,
			n:           0,
			expectedGas: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calculatedGas := types.CalculateGas(tc.base, tc.each, tc.n)
			require.Equal(t, tc.expectedGas, calculatedGas, "CalculateGas should return the correct value")
		})
	}
}

func TestSortKeysUint64StringMap(t *testing.T) {
	uint64StringMap := map[uint64][]string{
		3: {"value1", "value2"},
		1: {"value3", "value4"},
		5: {"value5", "value6"},
		2: {"value7", "value8"},
	}

	expectedKeys := []uint64{1, 2, 3, 5}
	sortedKeys := types.SortKeysUint64StringMap(uint64StringMap)

	require.Equal(t, expectedKeys, sortedKeys, "SortKeysUint64StringMap should return the correct sorted keys")
}
