package keeper

import (
	"context"

	"consumer/x/oracleconsumer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetRequestInterval(goCtx context.Context, msg *types.MsgSetRequestInterval) (*types.MsgSetRequestIntervalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetRequestIntervalResponse{}, nil
}
