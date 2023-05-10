package keeper_test

import (
	"testing"
	"time"

	"github.com/bandprotocol/oracle-consumer/obi"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/bandprotocol/oracle-consumer/testutil/keeper"
	bandtypes "github.com/bandprotocol/oracle-consumer/types/band"
)

func TestSetSymbolRequest(t *testing.T) {
	// Initialize the testing environment.
	k, ctx := testkeeper.PriceFeedKeeper(t)

	// Define symbol request
	symbolRequest := types.SymbolRequest{
		Symbol:         "BTC",
		OracleScriptID: 1,
		BlockInterval:  60,
	}

	// Set symbol request
	k.SetSymbolRequest(ctx, symbolRequest)

	storedSymbolRequest, err := k.GetSymbolRequest(ctx, "BTC")
	require.NoError(t, err)

	require.EqualValues(t, symbolRequest, storedSymbolRequest)
}

func TestHandleSymbolRequests(t *testing.T) {
	// Initialize the testing environment.
	k, ctx := testkeeper.PriceFeedKeeper(t)

	// Define symbol requests
	symbolRequests := []types.SymbolRequest{
		{
			Symbol:         "BTC",
			OracleScriptID: 1,
			BlockInterval:  60,
		},
		{
			Symbol:         "BTCC",
			OracleScriptID: 1,
			BlockInterval:  60,
		},
	}

	// Handle symbol requests
	k.HandleSymbolRequests(ctx, symbolRequests)

	storedSymbolRequest1, err := k.GetSymbolRequest(ctx, "BTC")
	require.NoError(t, err)
	storedSymbolRequest2, err := k.GetSymbolRequest(ctx, "BTCC")
	require.NoError(t, err)

	require.EqualValues(t, symbolRequests[0], storedSymbolRequest1)
	require.EqualValues(t, symbolRequests[1], storedSymbolRequest2)
}

func TestDeleteSymbolRequest(t *testing.T) {
	// Initialize the testing environment.
	k, ctx := testkeeper.PriceFeedKeeper(t)

	// Define symbol requests
	symbolRequest := types.SymbolRequest{
		Symbol:         "BTC",
		OracleScriptID: 1,
		BlockInterval:  60,
	}

	// Set symbol request
	k.SetSymbolRequest(ctx, symbolRequest)

	// Delete symbol request
	k.DeleteSymbolRequest(ctx, symbolRequest.Symbol)

	_, err := k.GetSymbolRequest(ctx, "BTC")
	require.ErrorIs(t, types.ErrSymbolRequestNotFound, err)
}

func TestDeleteSymbolRequestsBySetBlockIntervalToZero(t *testing.T) {
	// Initialize the testing environment.
	k, ctx := testkeeper.PriceFeedKeeper(t)

	// Define symbol request
	symbolRequests := []types.SymbolRequest{
		{
			Symbol:         "BTC",
			OracleScriptID: 1,
			BlockInterval:  60,
		},
		{
			Symbol:         "BTCC",
			OracleScriptID: 1,
			BlockInterval:  60,
		},
	}

	// Handle symbol request
	k.HandleSymbolRequests(ctx, symbolRequests)

	// Define symbol request with zero block interval to delete symbol
	symbolRequestsZeroBlockInterval := []types.SymbolRequest{
		{
			Symbol:         "BTC",
			OracleScriptID: 1,
			BlockInterval:  0,
		},
		{
			Symbol:         "BTCC",
			OracleScriptID: 1,
			BlockInterval:  0,
		},
	}

	// Handle symbol request
	k.HandleSymbolRequests(ctx, symbolRequestsZeroBlockInterval)

	_, err := k.GetSymbolRequest(ctx, "BTC")
	require.ErrorIs(t, types.ErrSymbolRequestNotFound, err)
	_, err = k.GetSymbolRequest(ctx, "BTCC")
	require.ErrorIs(t, types.ErrSymbolRequestNotFound, err)
}

func TestUpdatePrice(t *testing.T) {
	// Initialize the testing environment.
	k, ctx := testkeeper.PriceFeedKeeper(t)

	// Define price
	price := types.Price{
		Symbol:      "BTC",
		Price:       uint64(50000),
		ResolveTime: 100,
	}

	// Update the first price
	changed := k.UpdatePrice(ctx, price)
	require.True(t, changed)

	storedPrice, found := k.GetPrice(ctx, "BTC")
	require.True(t, found)
	require.EqualValues(t, price, storedPrice)

	// Define new price
	newPrice := types.Price{
		Symbol:      "BTC",
		Price:       uint64(52000),
		ResolveTime: 110,
	}

	// Update new price
	changed = k.UpdatePrice(ctx, newPrice)
	require.True(t, changed)

	storedPrice, found = k.GetPrice(ctx, "BTC")
	require.True(t, found)
	require.EqualValues(t, newPrice, storedPrice)

	// Update with old price
	changed = k.UpdatePrice(ctx, price)
	require.False(t, changed)

	storedPrice, found = k.GetPrice(ctx, "BTC")
	require.True(t, found)
	require.EqualValues(t, newPrice, storedPrice)
}

func TestStoreOracleResponsePacket(t *testing.T) {
	// Initialize the testing environment.
	k, ctx := testkeeper.PriceFeedKeeper(t)

	resolveTime := time.Now().Unix()

	// Define the test data.
	price := types.Price{
		Symbol:      "test",
		Price:       1,
		ResolveTime: resolveTime,
	}

	// Initialize BandChain Responses
	bResult, err := obi.Encode([]bandtypes.Response{
		{
			Symbol:       "test",
			ResponseCode: 0,
			Rate:         1,
		},
	})
	require.NoError(t, err)

	oracleResponsePacket := bandtypes.OracleResponsePacketData{
		RequestID:   1,
		Result:      bResult,
		ResolveTime: resolveTime,
	}

	// Call the function with the test data.
	k.StoreOracleResponsePacket(ctx, oracleResponsePacket)

	// Retrieve the price from the store.
	storedPrice, found := k.GetPrice(ctx, "test")
	require.True(t, found)
	require.Equal(t, price, storedPrice)
}
