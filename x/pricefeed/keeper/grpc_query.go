package keeper

import (
	"consumer/x/pricefeed/types"
)

var _ types.QueryServer = Keeper{}
