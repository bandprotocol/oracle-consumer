/*
NOTE: Usage of x/params to manage parameters is deprecated in favor of x/gov
controlled execution of MsgUpdateParams messages. These types remains solely
for migration purposes and will be removed in a future release.
*/
package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyAskCount       = []byte("AskCount")
	KeyMinCount       = []byte("MinCount")
	KeyMinDsCount     = []byte("MinDsCount")
	KeyPrepareGasBase = []byte("PrepareGasBase")
	KeyPrepareGasEach = []byte("PrepareGasEach")
	KeyExecuteGasBase = []byte("ExecuteGasBase")
	KeyExecuteGasEach = []byte("ExecuteGasEach")
	KeySourceChannel  = []byte("SourceChannel")
	KeyFeeLimit       = []byte("FeeLimit")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAskCount, &p.AskCount, validateUint64("ask count", true)),
		paramtypes.NewParamSetPair(KeyMinCount, &p.MinCount, validateUint64("min count", true)),
		paramtypes.NewParamSetPair(KeyMinDsCount, &p.MinDsCount, validateUint64("min ds count", true)),
		paramtypes.NewParamSetPair(KeyPrepareGasBase, &p.PrepareGasBase, validateUint64("prepare gas base", true)),
		paramtypes.NewParamSetPair(KeyPrepareGasEach, &p.PrepareGasEach, validateUint64("prepare gas each", true)),
		paramtypes.NewParamSetPair(KeyExecuteGasBase, &p.ExecuteGasBase, validateUint64("execute gas base", true)),
		paramtypes.NewParamSetPair(KeyExecuteGasEach, &p.ExecuteGasEach, validateUint64("execute gas each", true)),
		paramtypes.NewParamSetPair(KeySourceChannel, &p.SourceChannel, validateString("source channel")),
		paramtypes.NewParamSetPair(KeyFeeLimit, &p.FeeLimit, validateFeeLimit),
	}
}
