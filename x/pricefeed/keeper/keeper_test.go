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
		OracleScriptId: 1,
		BlockInterval:  60,
	}

	// Set symbol request
	k.SetSymbolRequest(ctx, symbolRequest)

	storedSymbolRequest, err := k.GetSymbolRequest(ctx, "BTC")
	require.NoError(t, err)

	require.EqualValues(t, symbolRequest, storedSymbolRequest)
}

func TestSetPrice(t *testing.T) {
	// Initialize the testing environment.
	k, ctx := testkeeper.PriceFeedKeeper(t)

	// Define price
	price := types.Price{
		Symbol:      "BTC",
		Price:       uint64(50000),
		ResolveTime: time.Now().Unix(),
	}

	// Set price
	k.SetPrice(ctx, price)

	storedPrice, err := k.GetPrice(ctx, "BTC")
	require.NoError(t, err)

	require.EqualValues(t, price, storedPrice)
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
	storedPrice, err := k.GetPrice(ctx, "test")
	require.NoError(t, err)

	require.Equal(t, price, storedPrice)
}
