package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/oracle-consumer/x/consumer/types"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (k Querier) Price(c context.Context, req *types.QueryPriceRequest) (*types.QueryPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	pf, err := k.priceFeedKeeper.GetPrice(ctx, req.Symbol)
	if err != nil {
		return nil, err
	}

	return &types.QueryPriceResponse{
		Symbol:      pf.Symbol,
		Price:       pf.Price,
		ResolveTime: pf.ResolveTime,
	}, nil
}
