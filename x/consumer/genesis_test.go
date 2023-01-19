package consumer_test

import (
	"testing"

	keepertest "consumer/testutil/keeper"
	"consumer/testutil/nullify"
	"consumer/x/consumer"
	"consumer/x/consumer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ConsumerKeeper(t)
	consumer.InitGenesis(ctx, *k, genesisState)
	got := consumer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
