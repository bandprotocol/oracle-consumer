package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	"github.com/bandprotocol/oracle-consumer/x/consumer/keeper"
	"github.com/bandprotocol/oracle-consumer/x/consumer/types"
	pricefeedkeeper "github.com/bandprotocol/oracle-consumer/x/pricefeed/keeper"
	pricefeedtypes "github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

func ConsumerKeeper(t testing.TB) (keeper.Keeper, sdk.Context, pricefeedkeeper.Keeper) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	pfStoreKey := sdk.NewKVStoreKey(pricefeedtypes.StoreKey)
	pfMemStoreKey := storetypes.NewMemoryStoreKey(pricefeedtypes.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(pfStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(pfMemStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	capabilityKeeper := capabilitykeeper.NewKeeper(cdc, pfStoreKey, pfMemStoreKey)

	paramsSubspace := typesparams.NewSubspace(
		cdc,
		pricefeedtypes.Amino,
		pfStoreKey,
		pfMemStoreKey,
		"PriceFeedParams",
	)

	pricefeedKeeper := pricefeedkeeper.NewKeeper(
		cdc,
		storeKey,
		paramsSubspace,
		priceFeedChannelKeeper{},
		priceFeedChannelKeeper{},
		priceFeedPortKeeper{},
		capabilityKeeper.ScopeToModule("PriceFeedScopedKeeper"),
	)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		pricefeedKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return k, ctx, pricefeedKeeper
}
