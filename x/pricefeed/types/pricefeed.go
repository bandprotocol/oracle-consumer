package types

func MapSymbolsByOsIDAndCheckBlockIntervalRequest(symbols []SymbolRequest, blockHeight int64) map[uint64][]string {
	symbolsOsMap := make(map[uint64][]string)
	for _, symbol := range symbols {
		if symbol.BlockInterval != 0 && blockHeight%int64(symbol.BlockInterval) == 0 {
			symbolsOsMap[symbol.OracleScriptId] = append(symbolsOsMap[symbol.OracleScriptId], symbol.Symbol)
		}
	}
	return symbolsOsMap
}
