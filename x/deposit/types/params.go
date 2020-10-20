package types

import (
	"fmt"
	"github.com/KuChainNetwork/kuchain/x/singer/external"
	"github.com/cosmos/cosmos-sdk/codec"
	yaml "gopkg.in/yaml.v2"
	"time"
)

// Default parameter namespace
const (
	DefaultParamspace = ModuleName

	DefaultMortgageRate     = 150
	DefaultDepositFeeRate   = 5
	DefaultClaimFeeRate     = 5
	DefaultThreshold        = 3
	DefaultLackMortgageRate = 110

	DefaultWaitTime         time.Duration = time.Second * 10 //3 days
	DefaultDepositLifeCycle time.Duration = time.Second * 10 //0.5 years

)

var (
	KeyMortgageRate     = []byte("mortgagerate")
	KeyDepositFeeRate   = []byte("depositfeerate")
	KeyClaimFeeRate     = []byte("claimfeerate")
	KeyThreshold        = []byte("threshold")
	KeyWaitTime         = []byte("waittime")
	KeyLackMortgageRate = []byte("lackmortgagerate")
	KeyDepositLifeCycle = []byte("depositlifecycle")
)

type Params struct {
	MortgageRate     int64         `json:"mortgate_rage" yaml:"mortgate_rage"`
	DepositFeeRate   int64         `json:"deposit_fee_rate" yaml:"deposit_fee_rate"`
	ClaimFeeRate     int64         `json:"claim_fee_rate" yaml:"claim_fee_rate"`
	Threshold        int           `json:"threshold" yaml:"threshold"`
	WaitTime         time.Duration `json:"wait_time" yaml:"wait_time"`
	LackMortgageRate int64         `json:"lack_mortgate_rage" yaml:"lack_mortgate_rage"`
	DepositLifeCycle time.Duration `json:"deposit_life_cycle" yaml:"deposit_life_cycle"`
}

func NewParams(
	mortgageRage int64,
	depositFeeRate int64,
	claimFeeRate int64,
	threshold int,
	waitTime time.Duration,
	lackMortgageRage int64,
	depositLifeCycle time.Duration,
) Params {
	return Params{
		MortgageRate:     mortgageRage,
		DepositFeeRate:   depositFeeRate,
		ClaimFeeRate:     claimFeeRate,
		Threshold:        threshold,
		WaitTime:         waitTime,
		LackMortgageRate: lackMortgageRage,
		DepositLifeCycle: depositLifeCycle,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultMortgageRate,
		DefaultDepositFeeRate,
		DefaultClaimFeeRate,
		DefaultThreshold,
		DefaultWaitTime,
		DefaultLackMortgageRate,
		DefaultDepositLifeCycle,
	)
}

// ParamKeyTable for slashing module
func ParamKeyTable() external.ParamsKeyTable {
	return external.ParamsNewKeyTable().RegisterParamSet(&Params{})
}

func (p *Params) ParamSetPairs() external.ParamsSetPairs {
	return external.ParamsSetPairs{
		external.NewParamSetPair(KeyMortgageRate, &p.MortgageRate, validateMortgageRate),
		external.NewParamSetPair(KeyDepositFeeRate, &p.DepositFeeRate, validateFeeRate),
		external.NewParamSetPair(KeyClaimFeeRate, &p.ClaimFeeRate, validateFeeRate),
		external.NewParamSetPair(KeyThreshold, &p.Threshold, validateThreshold),
		external.NewParamSetPair(KeyWaitTime, &p.WaitTime, validateWaitTime),
		external.NewParamSetPair(KeyLackMortgageRate, &p.LackMortgageRate, validateLackMortgageRate),
		external.NewParamSetPair(KeyDepositLifeCycle, &p.DepositLifeCycle, validateWaitTime),
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
	if err := validateMortgageRate(p.MortgageRate); err != nil {
		return err
	}

	if err := validateFeeRate(p.DepositFeeRate); err != nil {
		return err
	}

	if err := validateFeeRate(p.ClaimFeeRate); err != nil {
		return err
	}
	return nil
}

func validateMortgageRate(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 100 {
		return fmt.Errorf("mortgage rate must be greater 100: %d", v)
	}

	return nil
}

func validateLackMortgageRate(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("mortgage rate must be greater 0: %d", v)
	}

	return nil
}

func validateFeeRate(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("fee rate must be :positive %d", v)
	}

	return nil
}

func validateThreshold(i interface{}) error {
	v, ok := i.(int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("Threshold must be :positive %d", v)
	}

	return nil
}

func validateWaitTime(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unbonding time must be positive: %d", v)
	}

	return nil
}
