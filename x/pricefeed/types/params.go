package types

import (
	fmt "fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	DefaultAskCount    = uint64(16)
	DefaultMinCount    = uint64(10)
	DefaultMinDsCount  = uint64(3)
	DefaultPrepareGasA = uint64(20000)
	DefaultPrepareGasB = uint64(20000)
	DefaultExecuteGasA = uint64(20000)
	DefaultExecuteGasB = uint64(20000)
)

var (
	KeyAskCount    = []byte("AskCount")
	KeyMinCount    = []byte("MinCount")
	KeyMinDsCount  = []byte("MinDsCount")
	KeyPrepareGasA = []byte("PrepareGasA")
	KeyPrepareGasB = []byte("PrepareGasB")
	KeyExecuteGasA = []byte("ExecuteGasA")
	KeyExecuteGasB = []byte("ExecuteGasB")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(askCount, minCount, minDsCount, prepareGasA, prepareGasB, executeGasA, executeGasB uint64) Params {
	return Params{
		AskCount:    askCount,
		MinCount:    minCount,
		MinDsCount:  minDsCount,
		PrepareGasA: prepareGasA,
		PrepareGasB: prepareGasB,
		ExecuteGasA: executeGasA,
		ExecuteGasB: executeGasB,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultAskCount, DefaultMinCount, DefaultMinDsCount, DefaultPrepareGasA, DefaultPrepareGasB, DefaultExecuteGasA, DefaultExecuteGasB)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAskCount, &p.AskCount, validateUint64("ask count", true)),
		paramtypes.NewParamSetPair(KeyMinCount, &p.MinCount, validateUint64("min count", true)),
		paramtypes.NewParamSetPair(KeyMinDsCount, &p.MinDsCount, validateUint64("min ds count", true)),
		paramtypes.NewParamSetPair(KeyPrepareGasA, &p.PrepareGasA, validateUint64("prepare gas a", true)),
		paramtypes.NewParamSetPair(KeyPrepareGasB, &p.PrepareGasB, validateUint64("prepare gas b", true)),
		paramtypes.NewParamSetPair(KeyExecuteGasA, &p.ExecuteGasA, validateUint64("execute gas a", true)),
		paramtypes.NewParamSetPair(KeyExecuteGasB, &p.ExecuteGasB, validateUint64("execute gas b", true)),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateUint64(name string, positiveOnly bool) func(interface{}) error {
	return func(i interface{}) error {
		v, ok := i.(uint64)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
		if v <= 0 && positiveOnly {
			return fmt.Errorf("%s must be positive: %d", name, v)
		}
		return nil
	}
}
