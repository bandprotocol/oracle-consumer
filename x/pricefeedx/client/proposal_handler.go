package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/bandprotocol/oracle-consumer/x/pricefeedx/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitUpdateSymbolRequestProposalTxCmd)
