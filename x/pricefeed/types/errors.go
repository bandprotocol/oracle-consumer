package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/pricefeed module sentinel errors
var (
	ErrSymbolRequestNotFound  = sdkerrors.Register(ModuleName, 2, "symbol request not found")
	ErrSymbolRequestsNotFound = sdkerrors.Register(ModuleName, 3, "symbol requests not found")
	ErrPriceNotFound          = sdkerrors.Register(ModuleName, 4, "price not found")

	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
)
