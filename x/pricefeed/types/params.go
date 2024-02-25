package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	DefaultAskCount       = uint64(16)
	DefaultMinCount       = uint64(10)
	DefaultMinDsCount     = uint64(3)
	DefaultPrepareGasBase = uint64(3000)
	DefaultPrepareGasEach = uint64(600)
	DefaultExecuteGasBase = uint64(70000)
	DefaultExecuteGasEach = uint64(7500)
	DefaultSourceChannel  = NotSet
)

var (
	DefaultFeeLimit = sdk.NewCoins(sdk.NewInt64Coin("uband", 1000000))
)

var _ paramtypes.ParamSet = (*Params)(nil)

// NewParams creates a new Params instance
func NewParams(
	askCount, minCount, minDsCount, prepareGasBase, prepareGasEach, executeGasBase, executeGasEach uint64,
	sourceChannel string,
	feeLimit sdk.Coins,
) Params {
	return Params{
		AskCount:       askCount,
		MinCount:       minCount,
		MinDsCount:     minDsCount,
		PrepareGasBase: prepareGasBase,
		PrepareGasEach: prepareGasEach,
		ExecuteGasBase: executeGasBase,
		ExecuteGasEach: executeGasEach,
		SourceChannel:  sourceChannel,
		FeeLimit:       feeLimit,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultAskCount,
		DefaultMinCount,
		DefaultMinDsCount,
		DefaultPrepareGasBase,
		DefaultPrepareGasEach,
		DefaultExecuteGasBase,
		DefaultExecuteGasEach,
		DefaultSourceChannel,
		DefaultFeeLimit,
	)
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateUint64("ask count", true)(p.AskCount); err != nil {
		return err
	}
	if err := validateUint64("min count", true)(p.MinCount); err != nil {
		return err
	}
	if err := validateUint64("min ds count", true)(p.MinDsCount); err != nil {
		return err
	}
	if err := validateUint64("prepare gas base", true)(p.PrepareGasBase); err != nil {
		return err
	}
	if err := validateUint64("prepare gas each", true)(p.PrepareGasEach); err != nil {
		return err
	}
	if err := validateUint64("execute gas base", true)(p.ExecuteGasBase); err != nil {
		return err
	}
	if err := validateUint64("execute gas each", true)(p.ExecuteGasEach); err != nil {
		return err
	}
	if err := validateString("source channel")(p.SourceChannel); err != nil {
		return err
	}
	if err := validateFeeLimit(p.FeeLimit); err != nil {
		return err
	}

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

func validateString(name string) func(interface{}) error {
	return func(i interface{}) error {
		_, ok := i.(string)
		if !ok {
			return fmt.Errorf("%s must be string: %T", name, i)
		}
		return nil
	}
}

func validateFeeLimit(i interface{}) error {
	_, ok := i.(sdk.Coins)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "type: %T, expected sdk.Coins", i)
	}
	return nil
}
