package keeper_test

import (
	"testing"

	testkeeper "github.com/bandprotocol/consumer/testutil/keeper"
	"github.com/bandprotocol/consumer/x/consumer/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ConsumerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
