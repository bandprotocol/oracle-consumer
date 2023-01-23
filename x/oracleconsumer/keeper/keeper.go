package keeper

import (
	"encoding/binary"
	"fmt"

	"consumer/x/oracleconsumer/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k Keeper) SetRequestIntervalCount(ctx sdk.Context, count uint64) {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	ctx.KVStore(k.storeKey).Set(types.RequestIntervalCountStoreKey, bz)
}

func (k Keeper) GetRequestIntervalCount(ctx sdk.Context) uint64 {
	bz := ctx.KVStore(k.storeKey).Get(types.RequestIntervalCountStoreKey)
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) GetNextRequestIntervalID(ctx sdk.Context) uint64 {
	requestNumber := k.GetRequestIntervalCount(ctx)
	k.SetRequestIntervalCount(ctx, requestNumber+1)
	return requestNumber + 1
}

func (k Keeper) SetRequestInterval(ctx sdk.Context, id uint64, requestInterval types.RequestInterval) {
	ctx.KVStore(k.storeKey).Set(types.RequestIntervalStoreKey(id), k.cdc.MustMarshal(&requestInterval))
}

func (k Keeper) AddRequestInterval(ctx sdk.Context, requestInterval types.RequestInterval) uint64 {
	id := k.GetNextRequestIntervalID(ctx)
	k.SetRequestInterval(ctx, id, requestInterval)
	return id
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
