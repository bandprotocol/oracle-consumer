package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bandprotocol/oracle-consumer/x/consumer/types"
	pricefeedtypes "github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (k Querier) Price(c context.Context, req *types.QueryPriceRequest) (*types.QueryPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	pf, found := k.priceFeedKeeper.GetPrice(ctx, req.Symbol)
	if !found {
		return nil, sdkerrors.Wrapf(pricefeedtypes.ErrPriceNotFound, "symbol: %s", req.Symbol)
	}

	return &types.QueryPriceResponse{
		Symbol:      pf.Symbol,
		Price:       pf.Price,
		ResolveTime: pf.ResolveTime,
	}, nil
}
