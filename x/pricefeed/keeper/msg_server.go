package keeper

import (
	"consumer/x/pricefeed/types"
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

func (k msgServer) CreateRequestInterval(goCtx context.Context, msg *types.MsgCreateRequestInterval) (*types.MsgCreateRequestIntervalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.SetRequestInterval(ctx, types.NewRequestInterval(msg.Symbol, msg.OracleScriptId, msg.BlockInterval))
	return &types.MsgCreateRequestIntervalResponse{}, nil
}
