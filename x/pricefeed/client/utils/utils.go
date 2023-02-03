package utils

import (
	"consumer/x/pricefeed/types"
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
)

type (
	// ParamChangesJSON defines a slice of ParamChangeJSON objects which can be
	// converted to a slice of ParamChange objects.
	SymbolsJSON []SymbolJSON

	// ParamChangeJSON defines a parameter change used in JSON input. This
	// allows values to be specified in raw JSON instead of being string encoded.
	SymbolJSON struct {
		Symbol         string `json:"symbol" yaml:"symbol"`
		OracleScriptId uint64 `json:"oracle_script_id" yaml:"oracle_script_id"`
		BlockInterval  uint64 `json:"block_interval" yaml:"block_interval"`
	}

	// ParamChangeProposalJSON defines a ParameterChangeProposal with a deposit used
	// to parse parameter change proposals from a JSON file.
	UpdateSymbolRequestProposalJSON struct {
		Title       string      `json:"title" yaml:"title"`
		Description string      `json:"description" yaml:"description"`
		Symbols     SymbolsJSON `json:"symbols" yaml:"symbols"`
		Deposit     string      `json:"deposit" yaml:"deposit"`
	}
)

func NewSymbolJSON(symbol string, oracleScriptId, blockInterval uint64) SymbolJSON {
	return SymbolJSON{symbol, oracleScriptId, blockInterval}
}

// ToParamChange converts a ParamChangeJSON object to ParamChange.
func (sj SymbolJSON) ToSymbol() types.Symbol {
	return types.Symbol{
		Symbol:         sj.Symbol,
		OracleScriptId: sj.OracleScriptId,
		BlockInterval:  sj.BlockInterval,
	}
}

// ToParamChanges converts a slice of ParamChangeJSON objects to a slice of
// ParamChange.
func (sj SymbolsJSON) ToSymbols() types.Symbols {
	res := make([]types.Symbol, len(sj))
	for i, s := range sj {
		res[i] = s.ToSymbol()
	}
	return types.Symbols{
		Symbols: res,
	}
}

// ParseParamChangeProposalJSON reads and parses a ParamChangeProposalJSON from
// file.
func ParseUpdateSymbolRequestProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (UpdateSymbolRequestProposalJSON, error) {
	proposal := UpdateSymbolRequestProposalJSON{}
	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
