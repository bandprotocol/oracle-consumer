package pricefeedx_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/bandprotocol/oracle-consumer/testutil/keeper"
	"github.com/bandprotocol/oracle-consumer/testutil/nullify"
	"github.com/bandprotocol/oracle-consumer/x/pricefeedx"
	"github.com/bandprotocol/oracle-consumer/x/pricefeedx/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
	}

	k, ctx := keepertest.PriceFeedKeeper(t)
	pricefeedx.InitGenesis(ctx, *k, genesisState)
	got := pricefeedx.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)
}
