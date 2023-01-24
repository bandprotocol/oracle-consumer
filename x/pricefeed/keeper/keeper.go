package keeper

import (
	"fmt"

	"consumer/x/pricefeed/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ignite/cli/ignite/pkg/cosmosibckeeper"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		*cosmosibckeeper.Keeper
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper cosmosibckeeper.ChannelKeeper,
	portKeeper cosmosibckeeper.PortKeeper,
	scopedKeeper cosmosibckeeper.ScopedKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		Keeper: cosmosibckeeper.NewKeeper(
			types.PortKey,
			storeKey,
			channelKeeper,
			portKeeper,
			scopedKeeper,
		),
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) SetRequestInterval(ctx sdk.Context, requestInterval types.RequestInterval) {
	ctx.KVStore(k.storeKey).Set(types.RequestIntervalStoreKey(requestInterval.Symbol), k.cdc.MustMarshal(&requestInterval))
}

func (k Keeper) GetRequestInterval(ctx sdk.Context, symbol string) (types.RequestInterval, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.RequestIntervalStoreKey(symbol))
	if bz == nil {
		return types.RequestInterval{}, sdkerrors.Wrapf(types.ErrRequestIntervalNotFound, "symbol: %s", symbol)
	}
	var requestInterval types.RequestInterval
	k.cdc.MustUnmarshal(bz, &requestInterval)
	return requestInterval, nil
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
