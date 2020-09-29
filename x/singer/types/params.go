package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"fmt"
	"github.com/KuChainNetwork/kuchain/x/singer/external"
	yaml "gopkg.in/yaml.v2"
	"github.com/cosmos/cosmos-sdk/codec"
	"time"
)

// Default parameter namespace
const (
	DefaultParamspace = ModuleName
	DefaultMinAccessAmountByPower int64 = 1
	DefaultWaitTime time.Duration = time.Second * 10
	DefaultPunishRate int = 10
)

var (
	KeyMinAccessAmount     = []byte("MinAccessAmount")
	KeyWaitTime     = []byte("waittime")
	KeyPunishRage     = []byte("punishrate")

)

type Params struct {
	MinAccessAmount sdk.Int `json:"min_access_amount" yaml:"min_access_amount"`
	WaitTime  time.Duration `json:"wait_time" yaml:"wait_time"`
	PunishRage int `json:"punish_rate" yaml:"punish_rate"`
}

func NewParams(
	accessAmount sdk.Int,
	waitTime time.Duration,
	punishRate int,
) Params {
	return Params{
		MinAccessAmount:accessAmount,
		WaitTime:waitTime,
		PunishRage:punishRate,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		external.TokensFromConsensusPower(DefaultMinAccessAmountByPower),
		DefaultWaitTime,
		DefaultPunishRate,
	)
}

// ParamKeyTable for slashing module
func ParamKeyTable() external.ParamsKeyTable {
	return external.ParamsNewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() external.ParamsSetPairs {
	return external.ParamsSetPairs{
		external.NewParamSetPair(KeyMinAccessAmount, &p.MinAccessAmount, validateMinAccessAmount),
		external.NewParamSetPair(KeyWaitTime, &p.WaitTime, validateWaitTime),
		external.NewParamSetPair(KeyPunishRage, &p.PunishRage, validatePunishRate),

	}
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// unmarshal the current staking params value from store key or panic
func MustUnmarshalParams(cdc *codec.Codec, value []byte) Params {
	params, err := UnmarshalParams(cdc, value)
	if err != nil {
		panic(err)
	}
	return params
}

// unmarshal the current staking params value from store key
func UnmarshalParams(cdc *codec.Codec, value []byte) (params Params, err error) {
	err = cdc.UnmarshalBinaryBare(value, &params)
	if err != nil {
		return
	}
	return
}

// validate a set of params
func (p Params) Validate() error {
	if err := validateMinAccessAmount(p.MinAccessAmount); err != nil {
		return err
	}

	return nil
}


func validateMinAccessAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("min Access Amount must be positive: %d", v)
	}

	return nil
}

func validateWaitTime(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("wait time must be positive: %d", v)
	}

	return nil
}

func validatePunishRate(i interface{}) error {
	v, ok := i.(int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("punish rate must be positive: %d", v)
	}

	return nil
}
