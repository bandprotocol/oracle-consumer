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

	bandtypes "github.com/bandprotocol/oracle-consumer/types/band"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
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

func (k Keeper) SetSymbolRequest(ctx sdk.Context, symbolRequest types.SymbolRequest) {
	ctx.KVStore(k.storeKey).Set(types.SymbolRequestStoreKey(symbolRequest.Symbol), k.cdc.MustMarshal(&symbolRequest))
}

func (k Keeper) GetSymbolRequest(ctx sdk.Context, symbol string) (types.SymbolRequest, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.SymbolRequestStoreKey(symbol))
	if bz == nil {
		return types.SymbolRequest{}, sdkerrors.Wrapf(types.ErrSymbolRequestNotFound, "symbol: %s", symbol)
	}
	var sr types.SymbolRequest
	k.cdc.MustUnmarshal(bz, &sr)
	return sr, nil
}

func (k Keeper) SetSymbolRequests(ctx sdk.Context, symbolRequests []types.SymbolRequest) {
	for _, sr := range symbolRequests {
		k.SetSymbolRequest(ctx, sr)
	}
}

func (k Keeper) GetAllSymbolRequest(ctx sdk.Context) []types.SymbolRequest {
	store := ctx.KVStore(k.storeKey)

	iterator := storetypes.KVStorePrefixIterator(store, types.SymbolRequestStoreKeyPrefix)
	defer iterator.Close()
	var srs []types.SymbolRequest
	for ; iterator.Valid(); iterator.Next() {
		var sr types.SymbolRequest
		k.cdc.MustUnmarshal(iterator.Value(), &sr)
		srs = append(srs, sr)
	}
	return srs
}

func (k Keeper) SetPrice(ctx sdk.Context, price types.Price) {
	ctx.KVStore(k.storeKey).Set(types.PriceStoreKey(price.Symbol), k.cdc.MustMarshal(&price))
}

func (k Keeper) GetPrice(ctx sdk.Context, symbol string) (*types.Price, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.PriceStoreKey(symbol))

	if bz == nil {
		return &types.Price{}, sdkerrors.Wrapf(types.ErrPriceNotFound, "symbol: %s", symbol)
	}
	pf := &types.Price{}
	k.cdc.MustUnmarshal(bz, pf)

	return pf, nil
}

func (k Keeper) RequestBandChainData(ctx sdk.Context, sourceChannel string, oracleRequestPacket bandtypes.OracleRequestPacketData) error {
	channel, found := k.ChannelKeeper.GetChannel(ctx, types.PortID, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", types.PortID, sourceChannel)
	}

	destinationPort := channel.GetCounterparty().GetPortID()
	destinationChannel := channel.GetCounterparty().GetChannelID()

	channelCap, ok := k.ScopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(types.PortID, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	sequence, found := k.ChannelKeeper.GetNextSequenceSend(
		ctx, types.PortID, sourceChannel,
	)
	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown sequence number for channel %s port %s",
			sourceChannel,
			types.PortID,
		)
	}

	packet := channeltypes.NewPacket(
		oracleRequestPacket.GetBytes(),
		sequence,
		types.PortID,
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
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeDecodeBandChainResultFailed,
			sdk.NewAttribute(types.AttributeKeyReason, fmt.Sprintf("Unable to decode result from BandChain: %s", err)),
		))
		return
	}

	for _, r := range result {
		k.SetPrice(ctx, types.Price{
			Symbol:      r.Symbol,
			Price:       r.Rate,
			ResolveTime: res.ResolveTime,
		})
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
