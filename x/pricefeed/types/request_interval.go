package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewRequestInterval(
	oracleScriptId uint64, calldata []byte, blockInterval, askCount, minCount uint64, feeLimit sdk.Coins, prepareGas, executeGas uint64, sourceChannel string) RequestInterval {
	return RequestInterval{
		OracleScriptId: oracleScriptId,
		Calldata:       calldata,
		BlockInterval:  blockInterval,
		AskCount:       askCount,
		MinCount:       minCount,
		FeeLimit:       feeLimit,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
		SourceChannel:  sourceChannel,
	}
}
