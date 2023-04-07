package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/bandprotocol/oracle-consumer/testutil/keeper"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/keeper"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

func TestParamsQuery(t *testing.T) {
	k, ctx := testkeeper.PriceFeedKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	q := keeper.Querier{
		Keeper: k,
	}

	// Set params
	params := types.DefaultParams()
	k.SetParams(ctx, params)

	response, err := q.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}

func TestSymbolRequestQuery(t *testing.T) {
	k, ctx := testkeeper.PriceFeedKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	q := keeper.Querier{
		Keeper: k,
	}

	// Set symbol request
	symbolRequest := types.SymbolRequest{
		Symbol:         "BTC",
		OracleScriptId: 1,
		BlockInterval:  60,
	}
	k.SetSymbolRequest(ctx, symbolRequest)

	response, err := q.SymbolRequest(wctx, &types.QuerySymbolRequest{
		Symbol: "BTC",
	})
	require.NoError(t, err)
	require.Equal(t, &types.QuerySymbolRequestResponse{SymbolRequest: &symbolRequest}, response)
}

func TestSymbolRequestsQuery(t *testing.T) {
	k, ctx := testkeeper.PriceFeedKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	q := keeper.Querier{
		Keeper: k,
	}

	// Set symbol requests
	symbolRequests := []types.SymbolRequest{
		{
			Symbol:         "BTC",
			OracleScriptId: 1,
			BlockInterval:  60,
		},
		{
			Symbol:         "ETH",
			OracleScriptId: 2,
			BlockInterval:  100,
		},
	}
	k.SetSymbolRequests(ctx, symbolRequests)

	response, err := q.SymbolRequests(wctx, &types.QuerySymbolRequests{})
	require.NoError(t, err)
	require.Equal(t, &types.QuerySymbolRequestsResponse{SymbolRequests: symbolRequests}, response)
}

func TestPriceQuery(t *testing.T) {
	k, ctx := testkeeper.PriceFeedKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	q := keeper.Querier{
		Keeper: k,
	}

	// Set price
	price := types.Price{
		Symbol:      "BTC",
		Price:       100000,
		ResolveTime: 1690000000,
	}
	k.SetPrice(ctx, price)

	response, err := q.Price(wctx, &types.QueryPrice{
		Symbol: "BTC",
	})
	require.NoError(t, err)
	require.Equal(t, &types.QueryPriceResponse{Price: &price}, response)
}
