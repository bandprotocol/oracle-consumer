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

	requestIntervals, err := k.GetAllRequestInterval(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	params := k.GetParams(ctx)

	blockHeight := ctx.BlockHeight()

	for _, requestInterval := range requestIntervals {
		if blockHeight%int64(requestInterval.BlockInterval) == 0 {
			_, ok := symbolsOsMap[requestInterval.OracleScriptId]
			if ok {
				symbolsOsMap[requestInterval.OracleScriptId] = append(symbolsOsMap[requestInterval.OracleScriptId], requestInterval.Symbol)
			} else {
				symbolsOsMap[requestInterval.OracleScriptId] = []string{requestInterval.Symbol}
			}
		}
	}

	for osID, symbols := range symbolsOsMap {
		calldataByte, err := bandtypes.EncodeCalldata(symbols, params.MinDsCount)
		if err != nil {
			fmt.Printf("request data error: %s", err.Error())
		}

		prepareGas := types.CalculateGas(params.PrepareGasA, params.PrepareGasB, uint64(len(symbols)))
		executeGas := types.CalculateGas(params.ExecuteGasA, params.ExecuteGasB, uint64(len(symbols)))

		oracleRequestPacket := bandtypes.NewOracleRequestPacketData("consumer", osID, calldataByte, params.AskCount, params.MinCount, params.FeeLimit, prepareGas, executeGas)
		err = k.RequestBandChainData(ctx, "channel-5", oracleRequestPacket)
		if err != nil {
			panic(err)
			fmt.Printf("request data error: %s", err.Error())
		}
	}

	fmt.Print("\n\n*********************************************\n")
	fmt.Printf("%+v %d height %d\n", symbolsOsMap, params.PrepareGasA, ctx.BlockHeight())
	fmt.Print("*********************************************\n")
}
