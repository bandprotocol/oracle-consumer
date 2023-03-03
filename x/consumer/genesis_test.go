package consumer_test

import (
	"testing"

	keepertest "github.com/bandprotocol/oracle-consumer/testutil/keeper"
	"github.com/stretchr/testify/require"

	"github.com/bandprotocol/oracle-consumer/testutil/nullify"
	"github.com/bandprotocol/oracle-consumer/x/consumer"
	"github.com/bandprotocol/oracle-consumer/x/consumer/types"
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
