package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	"github.com/tendermint/tendermint/libs/log"

	bandtypes "github.com/bandprotocol/oracle-consumer/types/band"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

type Keeper struct {
	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	paramSpace paramtypes.Subspace

	ics4Wrapper   types.ICS4Wrapper
	channelKeeper types.ChannelKeeper
	portKeeper    types.PortKeeper
	scopedKeeper  capabilitykeeper.ScopedKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, paramSpace paramtypes.Subspace,
	ics4Wrapper types.ICS4Wrapper, channelKeeper types.ChannelKeeper, portKeeper types.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		paramSpace:    paramSpace,
		ics4Wrapper:   ics4Wrapper,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// IsBound checks if the pricefeed module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the ort Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the pricefeed module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the pricefeed module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
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

func (k Keeper) GetAllSymbolRequests(ctx sdk.Context) []types.SymbolRequest {
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

func (k Keeper) GetPrice(ctx sdk.Context, symbol string) (types.Price, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.PriceStoreKey(symbol))

	if bz == nil {
		return types.Price{}, sdkerrors.Wrapf(types.ErrPriceNotFound, "symbol: %s", symbol)
	}
	pf := types.Price{}
	k.cdc.MustUnmarshal(bz, &pf)

	return pf, nil
}

func (k Keeper) RequestBandChainDataBySymbolRequests(ctx sdk.Context) {
	blockHeight := ctx.BlockHeight()

	params := k.GetParams(ctx)
	symbols := k.GetAllSymbolRequests(ctx)

	// Map symbols that need to request on this block by oracle script ID and symbol block interval
	symbolsOsMap := types.MapSymbolsByOsIDAndCheckBlockIntervalRequest(symbols, blockHeight)
	// Sort keys map for deterministic value
	symbolsOsMapKeys := types.SortKeysUint64StringMap(symbolsOsMap)

	for _, osID := range symbolsOsMapKeys {
		calldataByte, err := bandtypes.EncodeCalldata(symbolsOsMap[osID], uint8(params.MinDsCount))
		if err != nil {
			ctx.EventManager().EmitEvent(sdk.NewEvent(
				types.EventTypeEncodeCalldataFailed,
				sdk.NewAttribute(types.AttributeKeyReason, fmt.Sprintf("Unable to encode calldata: %s", err)),
			))
			continue
		}

		// Calculate the prepareGas and executeGas for the oracle request packet based on the module's parameters and
		// the number of symbols to be requested
		prepareGas := types.CalculateGas(params.PrepareGasBase, params.PrepareGasEach, uint64(len(symbols)))
		executeGas := types.CalculateGas(params.ExecuteGasBase, params.ExecuteGasEach, uint64(len(symbols)))

		oracleRequestPacket := bandtypes.NewOracleRequestPacketData(types.ModuleName, osID, calldataByte, params.AskCount, params.MinCount, params.FeeLimit, prepareGas, executeGas)

		// Send the oracle request packet to the Band Chain using the RequestBandChainData function from the keeper
		err = k.RequestBandChainData(ctx, params.SourceChannel, oracleRequestPacket)
		if err != nil {
			ctx.EventManager().EmitEvent(sdk.NewEvent(
				types.EventTypeRequestBandChainFailed,
				sdk.NewAttribute(types.AttributeKeyReason, fmt.Sprintf("Unable to request data on BandChain: %s", err)),
			))
		}

	}
}

// RequestBandChainData is a function that sends an OracleRequestPacketData to BandChain via IBC.
func (k Keeper) RequestBandChainData(ctx sdk.Context, sourceChannel string, oracleRequestPacket bandtypes.OracleRequestPacketData) error {
	channel, found := k.channelKeeper.GetChannel(ctx, types.PortID, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", types.PortID, sourceChannel)
	}

	destinationPort := channel.GetCounterparty().GetPortID()
	destinationChannel := channel.GetCounterparty().GetChannelID()

	// Get the capability associated with the given channel.
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(types.PortID, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	// Get the next sequence number for the given channel and port.
	sequence, found := k.channelKeeper.GetNextSequenceSend(
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

	// Create a new packet with the oracle request packet data and the sequence number.
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

	// Send the packet via the channel and capability associated with the given channel.
	if err := k.ics4Wrapper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// StoreOracleResponsePacket is a function that receives an OracleResponsePacketData from BandChain.
func (k Keeper) StoreOracleResponsePacket(ctx sdk.Context, res bandtypes.OracleResponsePacketData) {
	// Decode the result from the response packet.
	result, err := bandtypes.DecodeResult(res.Result)
	if err != nil {
		// Emit an event if decoding the result fails and return.
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeDecodeBandChainResultFailed,
			sdk.NewAttribute(types.AttributeKeyReason, fmt.Sprintf("Unable to decode result from BandChain: %s", err)),
		))
		return
	}

	// Loop through the result and set the price in the state for each symbol.
	for _, r := range result {
		k.SetPrice(ctx, types.Price{
			Symbol:      r.Symbol,
			Price:       r.Rate,
			ResolveTime: res.ResolveTime,
		})
	}
}

// ClaimCapability allows the pricefeed module that can claim a capability that IBC module
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}
