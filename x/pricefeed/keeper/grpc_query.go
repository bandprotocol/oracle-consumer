package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/consumer/x/pricefeed/types"
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

func (k Querier) SymbolRequest(c context.Context, req *types.QuerySymbolRequest) (*types.QuerySymbolRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	sr, err := k.GetSymbolRequest(ctx, req.Symbol)
	if err != nil {
		return nil, err
	}

	return &types.QuerySymbolRequestResponse{
		SymbolRequest: &sr,
	}, nil
}

func (k Querier) SymbolRequests(c context.Context, req *types.QuerySymbolRequests) (*types.QuerySymbolRequestsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	srs := k.GetAllSymbolRequest(ctx)

	result := types.QuerySymbolRequestsResponse{}
	for _, sr := range srs {
		result.SymbolRequests = append(result.SymbolRequests, &sr)
	}

	return &result, nil
}

func (k Querier) Price(c context.Context, req *types.QueryPrice) (*types.QueryPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	p, err := k.GetPrice(ctx, req.Symbol)

	if err != nil {
		return nil, err
	}
	return &types.QueryPriceResponse{
		Price: p,
	}, nil
}
