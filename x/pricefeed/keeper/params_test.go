package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/bandprotocol/consumer/testutil/keeper"
	"github.com/bandprotocol/consumer/x/pricefeed/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PriceFeedKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
