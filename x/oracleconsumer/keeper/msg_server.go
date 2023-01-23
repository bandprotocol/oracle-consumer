package keeper

import (
	"consumer/x/oracleconsumer/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const SRC_PORT = "oracle_consumer"

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k Keeper) CreateRequestInterval(goCtx context.Context, msg *types.MsgSetRequestInterval) (*types.MsgSetRequestIntervalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AddRequestInterval(ctx, types.NewRequestInterval(msg.OracleScriptId, msg.Calldata, msg.BlockInterval, msg.AskCount, msg.MinCount, msg.FeeLimit, msg.PrepareGas, msg.ExecuteGas, msg.SourceChannel))
	return &types.MsgSetRequestIntervalResponse{RequestIntervalId: id}, nil
}
