package types

func NewRequestInterval(
	symbol string, oracleScriptId uint64, blockInterval uint64) RequestInterval {
	return RequestInterval{
		Symbol:         symbol,
		OracleScriptId: oracleScriptId,
		BlockInterval:  blockInterval,
	}
}
