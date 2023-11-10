package pricefeed_test

import (
	"testing"

	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v7/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibctmtypes "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
	"github.com/stretchr/testify/suite"
)

type PricefeedTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain
}

func (suite *PricefeedTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(2))

	// set localhost client
	revision := clienttypes.ParseChainID(suite.chainA.GetContext().ChainID())
	localHostClient := ibctmtypes.NewClientState(
		suite.chainA.GetContext().ChainID(),
		ibctmtypes.DefaultTrustLevel,
		ibctesting.TrustingPeriod,
		ibctesting.UnbondingPeriod,
		ibctesting.MaxClockDrift,
		clienttypes.NewHeight(revision, uint64(suite.chainA.GetContext().BlockHeight())),
		commitmenttypes.GetSDKSpecs(),
		ibctesting.UpgradePath,
	)
	suite.chainA.App.GetIBCKeeper().ClientKeeper.SetClientState(
		suite.chainA.GetContext(),
		exported.Localhost,
		localHostClient,
	)
}

func NewPricefeedPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = types.PortID
	path.EndpointB.ChannelConfig.PortID = types.PortID
	path.EndpointA.ChannelConfig.Version = types.Version
	path.EndpointB.ChannelConfig.Version = types.Version

	return path
}

func TestPricefeedTestSuite(t *testing.T) {
	suite.Run(t, new(PricefeedTestSuite))
}
