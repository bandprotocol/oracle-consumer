package oracleconsumer_test

import (
	"testing"

	keepertest "consumer/testutil/keeper"
	"consumer/testutil/nullify"
	"consumer/x/oracleconsumer"
	"consumer/x/oracleconsumer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OracleconsumerKeeper(t)
	oracleconsumer.InitGenesis(ctx, *k, genesisState)
	got := oracleconsumer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
