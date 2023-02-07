package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	keepertest "github.com/bandprotocol/consumer/testutil/keeper"
	"github.com/bandprotocol/consumer/x/pricefeed/keeper"
	"github.com/bandprotocol/consumer/x/pricefeed/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PriceFeedKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
