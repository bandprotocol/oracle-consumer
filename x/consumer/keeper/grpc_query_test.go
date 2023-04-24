package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/bandprotocol/oracle-consumer/testutil/keeper"
	"github.com/bandprotocol/oracle-consumer/x/consumer/keeper"
	"github.com/bandprotocol/oracle-consumer/x/consumer/types"
	pricefeedtypes "github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

func TestPriceQuery(t *testing.T) {
	k, ctx, pfk := testkeeper.ConsumerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	q := keeper.Querier{
		Keeper: k,
	}

	// Set price
	price := pricefeedtypes.Price{
		Symbol:      "BTC",
		Price:       100000,
		ResolveTime: 1690000000,
	}
	pfk.SetPrice(ctx, price)

	response, err := q.Price(wctx, &types.QueryPriceRequest{
		Symbol: "BTC",
	})
	require.NoError(t, err)
	require.Equal(t, &types.QueryPriceResponse{
		Symbol:      price.Symbol,
		Price:       price.Price,
		ResolveTime: price.ResolveTime,
	}, response)
}
