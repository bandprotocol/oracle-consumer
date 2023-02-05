package keeper

import (
	"fmt"
	"time"

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
	"consumer/x/pricefeed/types"
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

func (k Keeper) SetSymbol(ctx sdk.Context, symbol types.Symbol) {
	ctx.KVStore(k.storeKey).Set(types.SymbolStoreKey(symbol.Symbol), k.cdc.MustMarshal(&symbol))
}

func (k Keeper) GetSymbol(ctx sdk.Context, symbol string) (types.Symbol, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.SymbolStoreKey(symbol))
	if bz == nil {
		return types.Symbol{}, sdkerrors.Wrapf(types.ErrSymbolNotFound, "symbol: %s", symbol)
	}
	var s types.Symbol
	k.cdc.MustUnmarshal(bz, &s)
	return s, nil
}

func (k Keeper) SetSymbols(ctx sdk.Context, symbols types.Symbols) {
	for _, s := range symbols.Symbols {
		k.SetSymbol(ctx, s)
	}
}

func (k Keeper) GetAllSymbol(ctx sdk.Context) ([]types.Symbol, error) {
	store := ctx.KVStore(k.storeKey)

	iterator := storetypes.KVStorePrefixIterator(store, types.SymbolStoreKeyPrefix)
	defer iterator.Close()
	var ss []types.Symbol
	for ; iterator.Valid(); iterator.Next() {
		var s types.Symbol
		k.cdc.MustUnmarshal(iterator.Value(), &s)
		ss = append(ss, s)
	}
	return ss, nil
}

func (k Keeper) SetPrice(ctx sdk.Context, price types.Price) {
	ctx.KVStore(k.storeKey).Set(types.PriceStoreKey(price.Symbol), k.cdc.MustMarshal(&price))
}

func (k Keeper) GetPrice(ctx sdk.Context, symbol string) (*types.Price, error) {
	pf := &types.Price{}
	bz := ctx.KVStore(k.storeKey).Get(types.PriceStoreKey(symbol))

	if err := k.cdc.Unmarshal(bz, pf); err != nil {
		return nil, err
	}

	return pf, nil
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

func (k Keeper) RecvIbcOracleResponsePacket(ctx sdk.Context, res bandtypes.OracleResponsePacketData) {
	result, err := bandtypes.DecodeResult(res.Result)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, r := range result {
		k.SetPrice(ctx, types.Price{
			Symbol:      r.Symbol,
			Price:       r.Rate,
			ResolveTime: res.ResolveTime,
		})
	}

	fmt.Print("\n\n*********************************************\n")
	fmt.Printf("%+v got result from BandChain\n", result)
	fmt.Print("*********************************************\n")
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
