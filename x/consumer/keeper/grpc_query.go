package keeper

import (
	"consumer/x/consumer/types"
)

var _ types.QueryServer = Keeper{}
