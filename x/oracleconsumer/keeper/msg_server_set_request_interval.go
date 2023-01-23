package keeper

import (
	"context"

	"consumer/x/oracleconsumer/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateRequestInterval(goCtx context.Context, msg *types.MsgCreateRequestInterval) (*types.MsgCreateRequestIntervalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateRequestIntervalResponse{}, nil
}
