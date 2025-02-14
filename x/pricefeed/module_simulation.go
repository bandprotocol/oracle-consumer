package pricefeed

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/bandprotocol/oracle-consumer/testutil/sample"
	pricefeedsimulation "github.com/bandprotocol/oracle-consumer/x/pricefeed/simulation"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = pricefeedsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgSetRequestInterval = "op_weight_msg_set_request_interval"

	defaultWeightMsgSetRequestInterval int = 100
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	priceFeedGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&priceFeedGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSetRequestInterval int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetRequestInterval, &weightMsgSetRequestInterval, nil,
		func(_ *rand.Rand) {
			weightMsgSetRequestInterval = defaultWeightMsgSetRequestInterval
		},
	)

	return operations
}
