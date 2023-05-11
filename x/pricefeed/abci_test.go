package pricefeed_test

import (
	testingapp "github.com/bandprotocol/oracle-consumer/testing"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
	ibctesting "github.com/cosmos/ibc-go/v5/testing"
)

func (suite *PricefeedTestSuite) TestHandleBeginBlock() {
	// setup path
	path := NewPricefeedPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)
	suite.coordinator.CreateChannels(path)

	// setup keeper
	pricefeedKeeper := testingapp.GetTestingApp(
		suite.chainA,
	).PricefeedKeeper
	params := pricefeedKeeper.GetParams(suite.chainA.GetContext())
	params.SourceChannel = ibctesting.FirstChannelID
	pricefeedKeeper.SetParams(suite.chainA.GetContext(), params)
	pricefeedKeeper.HandleSymbolRequests(suite.chainA.GetContext(), []types.SymbolRequest{
		{
			Symbol:         "BTC",
			OracleScriptID: 1,
			BlockInterval:  1,
		},
		{
			Symbol:         "BAND",
			OracleScriptID: 2,
			BlockInterval:  1,
		},
	})

	// try begin block
	suite.Require().NotPanics(func() {
		pricefeed.HandleBeginBlock(suite.chainA.GetContext(), pricefeedKeeper)
	}, "HandleBeginBlock shouldn't panic")
}
