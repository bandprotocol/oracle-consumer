package keeper_test

import (
	"context"
	"testing"

	keepertest "consumer/testutil/keeper"
	"consumer/x/pricefeed/keeper"
	"consumer/x/pricefeed/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.OracleconsumerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
