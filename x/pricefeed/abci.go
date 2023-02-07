package pricefeed

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	bandtypes "github.com/bandprotocol/consumer/types/band"
	"github.com/bandprotocol/consumer/x/pricefeed/keeper"
	"github.com/bandprotocol/consumer/x/pricefeed/types"
)

func handleBeginBlock(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	symbolsOsMap := make(map[uint64][]string)

	blockHeight := ctx.BlockHeight()

	params := k.GetParams(ctx)
	symbols := k.GetAllSymbolRequest(ctx)

	for _, symbol := range symbols {
		if symbol.BlockInterval != 0 && blockHeight%int64(symbol.BlockInterval) == 0 {
			_, ok := symbolsOsMap[symbol.OracleScriptId]
			if ok {
				symbolsOsMap[symbol.OracleScriptId] = append(symbolsOsMap[symbol.OracleScriptId], symbol.Symbol)
			} else {
				symbolsOsMap[symbol.OracleScriptId] = []string{symbol.Symbol}
			}
		}
	}

	for osID, symbols := range symbolsOsMap {
		calldataByte, err := bandtypes.EncodeCalldata(symbols, uint8(params.MinDsCount))
		if err != nil {
			ctx.EventManager().EmitEvent(sdk.NewEvent(
				types.EventTypeEncodeCalldataFailed,
				sdk.NewAttribute(types.AttributeKeyReason, fmt.Sprintf("Unable to encode calldata: %s", err)),
			))
			continue
		}

		prepareGas := types.CalculateGas(params.PrepareGasBase, params.PrepareGasEach, uint64(len(symbols)))
		executeGas := types.CalculateGas(params.ExecuteGasBase, params.ExecuteGasEach, uint64(len(symbols)))

		oracleRequestPacket := bandtypes.NewOracleRequestPacketData(types.ModuleName, osID, calldataByte, params.AskCount, params.MinCount, params.FeeLimit, prepareGas, executeGas)

		err = k.RequestBandChainData(ctx, params.SourceChannel, oracleRequestPacket)
		if err != nil {
			ctx.EventManager().EmitEvent(sdk.NewEvent(
				types.EventTypeRequestBandChainFailed,
				sdk.NewAttribute(types.AttributeKeyReason, fmt.Sprintf("Unable to request data on BandChain: %s", err)),
			))
		}
	}

	fmt.Print("\n\n*********************************************\n")
	fmt.Printf("%+v %d height %d\n", symbolsOsMap, params.PrepareGasBase, ctx.BlockHeight())
	fmt.Print("*********************************************\n")
}
