package pricefeed

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	bandtypes "consumer/types/band"
	"consumer/x/pricefeed/keeper"
	"consumer/x/pricefeed/types"
)

func handleBeginBlock(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	symbolsOsMap := make(map[uint64][]string)

	symbols, err := k.GetAllSymbolRequest(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	params := k.GetParams(ctx)

	blockHeight := ctx.BlockHeight()

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
			fmt.Printf("request data error: %s", err.Error())
		}

		prepareGas := types.CalculateGas(params.PrepareGasA, params.PrepareGasB, uint64(len(symbols)))
		executeGas := types.CalculateGas(params.ExecuteGasA, params.ExecuteGasB, uint64(len(symbols)))

		oracleRequestPacket := bandtypes.NewOracleRequestPacketData(types.ModuleName, osID, calldataByte, params.AskCount, params.MinCount, params.FeeLimit, prepareGas, executeGas)

		err = k.RequestBandChainData(ctx, params.SourceChannel, oracleRequestPacket)
		if err != nil {
			fmt.Printf("request data error: %s", err.Error())
		}
	}

	fmt.Print("\n\n*********************************************\n")
	fmt.Printf("%+v %d height %d\n", symbolsOsMap, params.PrepareGasA, ctx.BlockHeight())
	fmt.Print("*********************************************\n")
}
