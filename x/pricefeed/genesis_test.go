package pricefeed_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/bandprotocol/consumer/testutil/keeper"
	"github.com/bandprotocol/consumer/testutil/nullify"
	"github.com/bandprotocol/consumer/x/pricefeed"
	"github.com/bandprotocol/consumer/x/pricefeed/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
	}

	k, ctx := keepertest.PriceFeedKeeper(t)
	pricefeed.InitGenesis(ctx, *k, genesisState)
	got := pricefeed.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)
}
