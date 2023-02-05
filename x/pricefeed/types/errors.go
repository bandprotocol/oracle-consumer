package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/pricefeed module sentinel errors
var (
	ErrSymbolNotFound  = sdkerrors.Register(ModuleName, 1, "symbol not found")
	ErrSymbolsNotFound = sdkerrors.Register(ModuleName, 2, "symbols not found")

	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
)
