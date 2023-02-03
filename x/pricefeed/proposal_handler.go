package pricefeed

import (
	"consumer/x/pricefeed/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"consumer/x/pricefeed/types"
)

// NewParamChangeProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewUpdateSymbolRequestProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.UpdateSymbolRequestProposal:
			return handleUpdateSymbolRequestProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}

func handleUpdateSymbolRequestProposal(ctx sdk.Context, k keeper.Keeper, p *types.UpdateSymbolRequestProposal) error {
	k.SetSymbols(ctx, types.Symbols{
		Symbols: p.Symbols,
	})
	return nil
}
