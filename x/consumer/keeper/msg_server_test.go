package keeper_test

import (
	"context"
	"testing"

	keepertest "consumer/testutil/keeper"
	"consumer/x/consumer/keeper"
	"consumer/x/consumer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ConsumerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
