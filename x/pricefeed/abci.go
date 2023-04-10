package pricefeed

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/oracle-consumer/x/pricefeed/keeper"
)

// handleBeginBlock is a handler function for the BeginBlock ABCI request.
func HandleBeginBlock(ctx sdk.Context, k keeper.Keeper) {
	// fetches price data from the BandChain
	// at the start of a new block.
	k.RequestBandChainDataBySymbolRequests(ctx)
}
