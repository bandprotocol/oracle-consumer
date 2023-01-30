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

func (k Querier) RequestInterval(c context.Context, req *types.QueryRequestInterval) (*types.QueryRequestIntervalResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	requestInterval, err := k.GetRequestInterval(ctx, req.Symbol)

	if err != nil {
		return nil, err
	}
	return &types.QueryRequestIntervalResponse{
		OracleScriptId: requestInterval.OracleScriptId,
		BlockInterval:  requestInterval.BlockInterval,
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
