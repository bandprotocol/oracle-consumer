package keeper_test

import (
	"testing"

	testkeeper "consumer/testutil/keeper"
	"consumer/x/oracleconsumer/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OracleconsumerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
