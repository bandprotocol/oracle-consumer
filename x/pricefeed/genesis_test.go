package pricefeed_test

import (
	"testing"

	keepertest "consumer/testutil/keeper"
	"consumer/testutil/nullify"
	"consumer/x/pricefeed"
	"consumer/x/pricefeed/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OracleconsumerKeeper(t)
	pricefeed.InitGenesis(ctx, *k, genesisState)
	got := pricefeed.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
