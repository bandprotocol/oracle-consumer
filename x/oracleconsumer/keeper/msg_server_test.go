package keeper_test

import (
	"context"
	"testing"

	keepertest "consumer/testutil/keeper"
	"consumer/x/oracleconsumer/keeper"
	"consumer/x/oracleconsumer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.OracleconsumerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
