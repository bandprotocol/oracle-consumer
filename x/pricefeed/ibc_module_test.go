package pricefeed_test

import (
	"math"

	bandtypes "github.com/bandprotocol/oracle-consumer/types/band"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	testingapp "github.com/bandprotocol/oracle-consumer/testing"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	ibctesting "github.com/cosmos/ibc-go/v5/testing"
)

func (suite *PricefeedTestSuite) TestOnChanOpenInit() {
	var (
		channel      *channeltypes.Channel
		path         *ibctesting.Path
		chanCap      *capabilitytypes.Capability
		counterparty channeltypes.Counterparty
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success", func() {}, true,
		},
		{
			"empty version string", func() {
				channel.Version = ""
			}, true,
		},
		{
			"max channels reached", func() {
				path.EndpointA.ChannelID = channeltypes.FormatChannelIdentifier(math.MaxUint32 + 1)
			}, false,
		},
		{
			"invalid order - ORDERED", func() {
				channel.Ordering = channeltypes.ORDERED
			}, false,
		},
		{
			"invalid port ID", func() {
				path.EndpointA.ChannelConfig.PortID = "mock"
			}, false,
		},
		{
			"invalid version", func() {
				channel.Version = "version"
			}, false,
		},
		{
			"capability already claimed", func() {
				err := testingapp.GetTestingApp(suite.chainA).ScopedPricefeedKeeper.ClaimCapability(
					suite.chainA.GetContext(),
					chanCap,
					host.ChannelCapabilityPath(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID),
				)
				suite.Require().NoError(err)
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			// setup
			suite.SetupTest() // reset
			path = NewPricefeedPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)
			path.EndpointA.ChannelID = ibctesting.FirstChannelID

			counterparty = channeltypes.NewCounterparty(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID)
			channel = &channeltypes.Channel{
				State:          channeltypes.INIT,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{path.EndpointA.ConnectionID},
				Version:        types.Version,
			}

			var err error
			chanCap, err = suite.chainA.App.GetScopedIBCKeeper().
				NewCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(types.PortID, path.EndpointA.ChannelID))
			suite.Require().NoError(err)

			// modify test data
			tc.malleate()

			// call function
			pricefeedModule := pricefeed.NewIBCModule(testingapp.GetTestingApp(suite.chainA).PricefeedKeeper)
			version, err := pricefeedModule.OnChanOpenInit(
				suite.chainA.GetContext(),
				channel.Ordering,
				channel.GetConnectionHops(),
				path.EndpointA.ChannelConfig.PortID,
				path.EndpointA.ChannelID,
				chanCap,
				counterparty,
				channel.GetVersion(),
			)

			// check result
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(types.Version, version)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(version, "")
			}
		})
	}
}

func (suite *PricefeedTestSuite) TestOnChanOpenTry() {
	var (
		channel             *channeltypes.Channel
		chanCap             *capabilitytypes.Capability
		path                *ibctesting.Path
		counterparty        channeltypes.Counterparty
		counterpartyVersion string
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success", func() {}, true,
		},
		{
			"max channels reached", func() {
				path.EndpointA.ChannelID = channeltypes.FormatChannelIdentifier(math.MaxUint32 + 1)
			}, false,
		},
		{
			"capability already claimed", func() {
				err := testingapp.GetTestingApp(suite.chainA).ScopedPricefeedKeeper.ClaimCapability(
					suite.chainA.GetContext(),
					chanCap,
					host.ChannelCapabilityPath(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID),
				)
				suite.Require().NoError(err)
			}, false,
		},
		{
			"invalid order - ORDERED", func() {
				channel.Ordering = channeltypes.ORDERED
			}, false,
		},
		{
			"invalid port ID", func() {
				path.EndpointA.ChannelConfig.PortID = ibctesting.MockPort
			}, false,
		},
		{
			"invalid counterparty version", func() {
				counterpartyVersion = "version"
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			// setup
			suite.SetupTest() // reset
			path = NewPricefeedPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)
			path.EndpointA.ChannelID = ibctesting.FirstChannelID

			counterparty = channeltypes.NewCounterparty(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID)
			channel = &channeltypes.Channel{
				State:          channeltypes.TRYOPEN,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{path.EndpointA.ConnectionID},
				Version:        types.Version,
			}
			counterpartyVersion = types.Version

			// get module
			module, _, err := suite.chainA.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(
				suite.chainA.GetContext(),
				types.PortID,
			)
			suite.Require().NoError(err)

			chanCap, err = suite.chainA.App.GetScopedIBCKeeper().
				NewCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(types.PortID, path.EndpointA.ChannelID))
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			// modify test data
			tc.malleate()

			// call function
			version, err := cbs.OnChanOpenTry(
				suite.chainA.GetContext(),
				channel.Ordering,
				channel.GetConnectionHops(),
				path.EndpointA.ChannelConfig.PortID,
				path.EndpointA.ChannelID,
				chanCap,
				channel.Counterparty,
				counterpartyVersion,
			)

			// check result
			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(types.Version, version)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal("", version)
			}
		})
	}
}

func (suite *PricefeedTestSuite) TestOnChanOpenAck() {
	var counterpartyVersion string

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success", func() {}, true,
		},
		{
			"invalid counterparty version", func() {
				counterpartyVersion = "version"
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			// setup
			suite.SetupTest() // reset
			path := NewPricefeedPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)
			path.EndpointA.ChannelID = ibctesting.FirstChannelID
			counterpartyVersion = types.Version

			// get module
			module, _, err := suite.chainA.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(
				suite.chainA.GetContext(),
				types.PortID,
			)
			suite.Require().NoError(err)

			// get route
			cbs, ok := suite.chainA.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			// modify test data
			tc.malleate()

			// call function
			err = cbs.OnChanOpenAck(
				suite.chainA.GetContext(),
				path.EndpointA.ChannelConfig.PortID,
				path.EndpointA.ChannelID,
				path.EndpointA.Counterparty.ChannelID,
				counterpartyVersion,
			)

			// check result
			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *PricefeedTestSuite) TestOnRecvPacket() {
	var packetData []byte
	testCases := []struct {
		name          string
		malleate      func()
		expAckSuccess bool
	}{
		{
			"success", func() {}, true,
		},
		{
			"fails - cannot unmarshal packet data", func() {
				packetData = []byte("invalid data")
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			// setup
			suite.SetupTest() // reset
			path := NewPricefeedPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)

			// prepare packet
			msg := &bandtypes.OracleResponsePacketData{
				ClientID:      "test",
				RequestID:     1,
				AnsCount:      1,
				RequestTime:   1000000000,
				ResolveTime:   1000000000,
				ResolveStatus: bandtypes.RESOLVE_STATUS_SUCCESS,
				Result:        []byte{},
			}
			packetData = msg.GetBytes()

			// modify test data
			tc.malleate()

			packet := channeltypes.NewPacket(
				packetData,
				uint64(1),
				path.EndpointA.ChannelConfig.PortID,
				path.EndpointA.ChannelID,
				path.EndpointB.ChannelConfig.PortID,
				path.EndpointB.ChannelID,
				clienttypes.NewHeight(0, 100),
				0,
			)

			// prepare expected ack
			expectedAck := channeltypes.NewResultAcknowledgement(nil)

			// get module
			module, _, err := suite.chainA.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(
				suite.chainA.GetContext(),
				path.EndpointA.ChannelConfig.PortID,
			)
			suite.Require().NoError(err)

			// get route
			cbs, ok := suite.chainA.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			// call recv packet
			ack := cbs.OnRecvPacket(suite.chainA.GetContext(), packet, nil)

			// check result
			if tc.expAckSuccess {
				suite.Require().True(ack.Success())
				suite.Require().Equal(expectedAck, ack)
			} else {
				suite.Require().False(ack.Success())
			}
		})
	}
}

func (suite *PricefeedTestSuite) TestOnAcknowledgementPacket() {
	var ackData []byte
	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success", func() {}, true,
		},
		{
			"fails - cannot unmarshal packet acknowledgement", func() {
				ackData = []byte{}
			}, false,
		},
		{
			"fails - cannot unmarshal packet data", func() {
				ackData = channeltypes.NewResultAcknowledgement([]byte{}).Acknowledgement()
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			// setup
			suite.SetupTest() // reset
			path := NewPricefeedPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)

			// prepare packet
			packet := channeltypes.NewPacket(
				[]byte("empty packet data"),
				suite.chainA.SenderAccount.GetSequence(),
				path.EndpointA.ChannelConfig.PortID,
				path.EndpointA.ChannelID,
				path.EndpointB.ChannelConfig.PortID,
				path.EndpointB.ChannelID,
				clienttypes.NewHeight(0, 100),
				0,
			)

			// prepare ack
			ack := &bandtypes.OracleRequestPacketAcknowledgement{
				RequestID: 1,
			}
			ackData = channeltypes.NewResultAcknowledgement(ack.GetBytes()).Acknowledgement()

			// modify test data
			tc.malleate()

			// get module
			module, _, err := suite.chainA.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(
				suite.chainA.GetContext(),
				path.EndpointA.ChannelConfig.PortID,
			)
			suite.Require().NoError(err)

			// get route
			cbs, ok := suite.chainB.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			// call ack
			err = cbs.OnAcknowledgementPacket(suite.chainA.GetContext(), packet, ackData, nil)

			// check result
			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
