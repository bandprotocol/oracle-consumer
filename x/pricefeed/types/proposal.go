package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	// UpdateSymbolRequest defines the type for a UpdateSymbolRequestProposal
	UpdateSymbolRequest = "UpdateSymbolRequest"
)

// Assert UpdateSymbolRequestProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &UpdateSymbolRequestProposal{}

func init() {
	govtypes.RegisterProposalType(UpdateSymbolRequest)
}

func NewUpdateSymbolRequestProposal(title, description string, symbols Symbols) *UpdateSymbolRequestProposal {
	return &UpdateSymbolRequestProposal{title, description, symbols.Symbols}
}

// String implements the Stringer interface.
func (usrp UpdateSymbolRequestProposal) String() string {
	return "ABV"
}

// GetTitle returns the title of a update symbol request proposal.
func (usrp *UpdateSymbolRequestProposal) GetTitle() string { return usrp.Title }

// GetDescription returns the description of a update symbol request proposal.
func (usrp *UpdateSymbolRequestProposal) GetDescription() string { return usrp.Description }

// ProposalRoute returns the routing key of a update symbol request proposal.
func (usrp *UpdateSymbolRequestProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a update symbol request proposal.
func (usrp *UpdateSymbolRequestProposal) ProposalType() string { return UpdateSymbolRequest }

// ValidateBasic validates the update symbol request proposal.
func (usrp *UpdateSymbolRequestProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(usrp)
	if err != nil {
		return err
	}

	return nil
}
