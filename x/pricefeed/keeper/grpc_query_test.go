package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/bandprotocol/oracle-consumer/testutil/keeper"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

func TestParamsQuery(t *testing.T) {
	k, ctx := testkeeper.PriceFeedKeeper(t)
	params := types.DefaultParams()
	k.SetParams(ctx, params)

	storedParams := k.GetParams(ctx)

	require.Equal(t, storedParams, params)
}
