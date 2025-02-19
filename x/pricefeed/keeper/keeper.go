package keeper

import (
	"fmt"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"

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

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, paramSpace paramtypes.Subspace,
	ics4Wrapper types.ICS4Wrapper, channelKeeper types.ChannelKeeper, portKeeper types.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper, authority string,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		paramSpace:    paramSpace,
		ics4Wrapper:   ics4Wrapper,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
		authority:     authority,
	}
}

// GetAuthority returns the x/pricefeed module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
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

func (k Keeper) DeleteSymbolRequest(ctx sdk.Context, symbol string) {
	ctx.KVStore(k.storeKey).Delete(types.SymbolRequestStoreKey(symbol))
}

func (k Keeper) HandleSymbolRequests(ctx sdk.Context, symbolRequests []types.SymbolRequest) {
	for _, sr := range symbolRequests {
		// delete when block interval is equal to zero
		if sr.BlockInterval == 0 {
			k.DeleteSymbolRequest(ctx, sr.Symbol)
		} else {
			k.SetSymbolRequest(ctx, sr)
		}
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

func (k Keeper) UpdatePrice(ctx sdk.Context, price types.Price) bool {
	old, found := k.GetPrice(ctx, price.Symbol)
	if !found || old.ResolveTime < price.ResolveTime {
		k.setPrice(ctx, price)
		return true
	}
	return false
}

func (k Keeper) setPrice(ctx sdk.Context, price types.Price) {
	ctx.KVStore(k.storeKey).Set(types.PriceStoreKey(price.Symbol), k.cdc.MustMarshal(&price))
}

func (k Keeper) GetPrice(ctx sdk.Context, symbol string) (types.Price, bool) {
	bz := ctx.KVStore(k.storeKey).Get(types.PriceStoreKey(symbol))
	if bz == nil {
		return types.Price{}, false
	}

	var pf types.Price
	k.cdc.MustUnmarshal(bz, &pf)

	return pf, true
}

func (k Keeper) RequestBandChainDataBySymbolRequests(ctx sdk.Context) {
	blockHeight := ctx.BlockHeight()

	params := k.GetParams(ctx)

	// Verify that SourceChannel params is set by open params proposal already
	if params.SourceChannel == types.NotSet {
		return
	}

	symbols := k.GetAllSymbolRequests(ctx)

	// Map symbols that need to request on this block by oracle script ID and symbol block interval
	tasks := types.ComputeOracleTasks(symbols, blockHeight)

	for _, task := range tasks {
		calldataByte, err := bandtypes.EncodeCalldata(task.Symbols, uint8(params.MinDsCount))
		if err != nil {
			// This error don't expect to happen, so just log in case unexpected bug
			ctx.Logger().Error(fmt.Sprintf("Unable to encode calldata: %s", err))
			continue
		}

		// Calculate the prepareGas and executeGas for the oracle request packet based on the module's parameters and
		// the number of symbols to be requested
		prepareGas := types.CalculateGas(params.PrepareGasBase, params.PrepareGasEach, uint64(len(symbols)))
		executeGas := types.CalculateGas(params.ExecuteGasBase, params.ExecuteGasEach, uint64(len(symbols)))

		oracleRequestPacket := bandtypes.NewOracleRequestPacketData(
			types.ModuleName,
			task.OracleScriptID,
			calldataByte,
			params.AskCount,
			params.MinCount,
			params.FeeLimit,
			prepareGas,
			executeGas,
		)

		// Send the oracle request packet to the Band Chain using the RequestBandChainData function from the keeper
		err = k.RequestBandChainData(ctx, params.SourceChannel, oracleRequestPacket)
		// In the normal case, this module should able to create new packet, so just log error to debug should be ok
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("Unable to send oracle request: %s", err))
		}
	}
}

// RequestBandChainData is a function that sends an OracleRequestPacketData to BandChain via IBC.
func (k Keeper) RequestBandChainData(
	ctx sdk.Context,
	sourceChannel string,
	oracleRequestPacket bandtypes.OracleRequestPacketData,
) error {
	portID := k.GetPort(ctx)
	_, found := k.channelKeeper.GetChannel(ctx, portID, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrChannelNotFound,
			"port ID (%s) channel ID (%s)",
			portID,
			sourceChannel,
		)
	}

	// Get the capability associated with the given channel.
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	// Get the next sequence number for the given channel and port.
	_, found = k.channelKeeper.GetNextSequenceSend(
		ctx, portID, sourceChannel,
	)
	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown sequence number for channel %s port %s",
			sourceChannel,
			portID,
		)
	}

	// Send the packet via the channel and capability associated with the given channel.
	if _, err := k.ics4Wrapper.SendPacket(
		ctx,
		channelCap,
		portID,
		sourceChannel,
		clienttypes.ZeroHeight(),
		uint64(ctx.BlockTime().UnixNano()+int64(20*time.Minute)),
		oracleRequestPacket.GetBytes(),
	); err != nil {
		return err
	}

	return nil
}

// StoreOracleResponsePacket is a function that receives an OracleResponsePacketData from BandChain.
func (k Keeper) StoreOracleResponsePacket(ctx sdk.Context, res bandtypes.OracleResponsePacketData) error {
	// Decode the result from the response packet.
	result, err := bandtypes.DecodeResult(res.Result)
	if err != nil {
		return err
	}

	// Loop through the result and set the price in the state for each symbol.
	for _, r := range result {
		if r.ResponseCode == 0 {
			changed := k.UpdatePrice(ctx, types.Price{
				Symbol:      r.Symbol,
				Price:       r.Rate,
				ResolveTime: res.ResolveTime,
			})
			if changed {
				ctx.EventManager().EmitEvent(
					sdk.NewEvent(
						types.EventTypePriceUpdate,
						sdk.NewAttribute(types.AttributeKeySymbol, r.Symbol),
						sdk.NewAttribute(types.AttributeKeyPrice, fmt.Sprintf("%d", r.Rate)),
						sdk.NewAttribute(types.AttributeKeyTimestamp, res.ResolveStatus.String()),
					),
				)
			}
		}
		// TODO: allow to write logic to handle failed symbol now just ignore and skip update
	}

	return nil
}

// ClaimCapability attempts to claim a given Capability. The provided name and
// the scoped module's name tuple are treated as the owner. It will attempt
// to add the owner to the persistent set of capability owners for the capability
// index. If the owner already exists, it will return an error. Otherwise, it will
// also set a forward and reverse index for the capability and capability name.
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}
