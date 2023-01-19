package keeper

import (
	"consumer/x/oracleconsumer/types"
)

var _ types.QueryServer = Keeper{}
