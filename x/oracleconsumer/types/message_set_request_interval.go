package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetRequestInterval = "set_request_interval"

var _ sdk.Msg = &MsgCreateRequestInterval{}

func NewMsgSetRequestInterval(oracleScriptId uint64, calldata []byte, blockInterval uint64, askCount uint64, minCount uint64, feeLimit sdk.Coins, prepareGas uint64, executeGas uint64, sender string) *MsgCreateRequestInterval {
	return &MsgCreateRequestInterval{
		OracleScriptId: oracleScriptId,
		Calldata:       calldata,
		BlockInterval:  blockInterval,
		AskCount:       askCount,
		MinCount:       minCount,
		FeeLimit:       feeLimit,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
		Sender:         sender,
	}
}

func (msg *MsgCreateRequestInterval) Route() string {
	return RouterKey
}

func (msg *MsgCreateRequestInterval) Type() string {
	return TypeMsgSetRequestInterval
}

func (msg *MsgCreateRequestInterval) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRequestInterval) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRequestInterval) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
