package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	pricefeedtypes "github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

type PriceFeedKeeper interface {
	GetPrice(ctx sdk.Context, symbol string) (pricefeedtypes.Price, error)
}
