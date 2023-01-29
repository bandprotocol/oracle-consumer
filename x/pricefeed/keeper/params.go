package keeper

import (
	"consumer/x/pricefeed/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.AskCount(ctx),
		k.MinCount(ctx),
		k.MinDsCount(ctx),
		k.PrepareGasA(ctx),
		k.PrepareGasB(ctx),
		k.ExecuteGasA(ctx),
		k.ExecuteGasB(ctx),
		k.SourceChannel(ctx),
		k.FeeLimit(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) AskCount(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyAskCount, &res)
	return
}

func (k Keeper) MinCount(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMinCount, &res)
	return
}

func (k Keeper) MinDsCount(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMinDsCount, &res)
	return
}

func (k Keeper) PrepareGasA(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyPrepareGasA, &res)
	return
}

func (k Keeper) PrepareGasB(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyPrepareGasB, &res)
	return
}

func (k Keeper) ExecuteGasA(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyExecuteGasA, &res)
	return
}

func (k Keeper) ExecuteGasB(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyExecuteGasB, &res)
	return
}

func (k Keeper) SourceChannel(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeySourceChannel, &res)
	return
}

func (k Keeper) FeeLimit(ctx sdk.Context) (res sdk.Coins) {
	k.paramstore.Get(ctx, types.KeyFeeLimit, &res)
	return
}
