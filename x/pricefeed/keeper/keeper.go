package keeper

import (
	"fmt"
	"time"

	"consumer/x/pricefeed/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	"github.com/ignite/cli/ignite/pkg/cosmosibckeeper"
	"github.com/tendermint/tendermint/libs/log"

	bandtypes "consumer/types/band"
)

const SRC_PORT = "pricefeed"

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

func (k Keeper) GetAllRequestInterval(ctx sdk.Context) ([]types.RequestInterval, error) {
	store := ctx.KVStore(k.storeKey)

	iterator := storetypes.KVStorePrefixIterator(store, types.RequestIntervalStoreKeyPrefix)
	defer iterator.Close()
	var requestIntervals []types.RequestInterval
	for ; iterator.Valid(); iterator.Next() {
		var requestInterval types.RequestInterval
		k.cdc.MustUnmarshal(iterator.Value(), &requestInterval)
		requestIntervals = append(requestIntervals, requestInterval)
	}
	return requestIntervals, nil
}

func (k Keeper) RequestBandChainData(ctx sdk.Context, sourceChannel string, oracleRequestPacket bandtypes.OracleRequestPacketData) error {
	channel, found := k.ChannelKeeper.GetChannel(ctx, SRC_PORT, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", SRC_PORT, sourceChannel)
	}

	destinationPort := channel.GetCounterparty().GetPortID()
	destinationChannel := channel.GetCounterparty().GetChannelID()

	channelCap, ok := k.ScopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(SRC_PORT, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	sequence, found := k.ChannelKeeper.GetNextSequenceSend(
		ctx, SRC_PORT, sourceChannel,
	)
	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown sequence number for channel %s port oracle",
			sourceChannel,
		)
	}

	packet := channeltypes.NewPacket(
		oracleRequestPacket.GetBytes(),
		sequence,
		SRC_PORT,
		sourceChannel,
		destinationPort,
		destinationChannel,
		clienttypes.ZeroHeight(),
		uint64(ctx.BlockTime().UnixNano()+int64(20*time.Minute)),
	)

	if err := k.ChannelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

func (k Keeper) RecvIbcOracleResponsePacket(ctx sdk.Context, data bandtypes.OracleResponsePacketData) {
	fmt.Print("\n\n*********************************************\n")
	fmt.Printf("%+v got data from BandChain\n", data)
	fmt.Print("*********************************************\n")
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
