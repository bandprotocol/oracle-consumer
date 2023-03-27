package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/bandprotocol/oracle-consumer/testutil/keeper"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.PriceFeedKeeper(t)
	delfaultParams := types.DefaultParams()
	keeper.SetParams(ctx, delfaultParams)

	params := keeper.GetParams(ctx)

	require.Equal(t, delfaultParams, params)
}
