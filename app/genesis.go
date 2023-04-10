package app

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"

	pricefeedtypes "github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

// The genesis state of the blockchain is represented here as a map of raw json
// messages key'd by a identifier string.
// The identifier is used to determine which module genesis information belongs
// to so it may be appropriately routed during init chain.
// Within this application default genesis information is retrieved from
// the ModuleBasicManager which populates json from each BasicModule
// object provided to it during init.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState(cdc codec.JSONCodec) GenesisState {
	defaultGenesis := ModuleBasics.DefaultGenesis(cdc)

	// Get default genesis states of the modules we are to override.
	pricefeedGenesis := pricefeedtypes.DefaultGenesis()

	// Override the genesis parameters.

	// Initial symbols that want to request to BandChain
	pricefeedGenesis.SymbolRequests = []pricefeedtypes.SymbolRequest{
		{
			Symbol:         "BTC",
			OracleScriptId: 396,
			BlockInterval:  40,
		},
		{
			Symbol:         "ETH",
			OracleScriptId: 396,
			BlockInterval:  40,
		},
		{
			Symbol:         "BAND",
			OracleScriptId: 396,
			BlockInterval:  40,
		},
	}
	defaultGenesis[pricefeedtypes.ModuleName] = cdc.MustMarshalJSON(pricefeedGenesis)

	return defaultGenesis
}
