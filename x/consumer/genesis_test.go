package consumer_test

import (
	"testing"

	keepertest "github.com/bandprotocol/consumer/testutil/keeper"
	"github.com/stretchr/testify/require"

	"github.com/bandprotocol/consumer/testutil/nullify"
	"github.com/bandprotocol/consumer/x/consumer"
	"github.com/bandprotocol/consumer/x/consumer/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.ConsumerKeeper(t)
	consumer.InitGenesis(ctx, *k, genesisState)
	got := consumer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
