package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/oracle-consumer/x/consumer/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Price(c context.Context, req *types.QueryPriceRequest) (*types.QueryPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	pf, err := k.PriceFeedKeeper.GetPrice(ctx, req.Symbol)
	if err != nil {
		return nil, err
	}

	return &types.QueryPriceResponse{
		Symbol:      pf.Symbol,
		Price:       pf.Price,
		ResolveTime: pf.ResolveTime,
	}, nil
}
