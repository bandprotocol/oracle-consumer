package keeper

import (
	"consumer/x/pricefeed/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryParamsResponse{
		Params: k.GetParams(ctx),
	}, nil
}

func (k Querier) Symbols(c context.Context, req *types.QuerySymbols) (*types.QuerySymbolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	ss, err := k.GetAllSymbol(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QuerySymbolsResponse{
		Symbols: &types.Symbols{
			Symbols: ss,
		},
	}, nil
}

func (k Querier) Price(c context.Context, req *types.QueryPrice) (*types.QueryPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	p, err := k.GetPrice(ctx, req.Symbol)

	if err != nil {
		return nil, err
	}
	return &types.QueryPriceResponse{
		Symbol:      p.Symbol,
		Price:       p.Price,
		ResolveTime: p.ResolveTime,
	}, nil
}
