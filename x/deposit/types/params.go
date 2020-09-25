package types

import (
	"fmt"
	"github.com/KuChainNetwork/kuchain/x/singer/external"
	yaml "gopkg.in/yaml.v2"
	"github.com/cosmos/cosmos-sdk/codec"
)

// Default parameter namespace
const (
	DefaultParamspace = ModuleName

	DefaultMortgageRate = 150
	DefaultDepositFeeRate = 5
	DefaultClaimFeeRate = 5
	DefaultThreshold = 3
)

var (
	KeyMortgageRate     = []byte("mortgagerate")
	KeyDepositFeeRate     = []byte("depositfeerate")
	KeyClaimFeeRate     = []byte("claimfeerate")
	KeyThreshold     = []byte("threshold")
)


type Params struct {
	MortgageRate int64 `json:"mortgate_rage" yaml:"mortgate_rage"`
	DepositFeeRate int64 `json:"deposit_fee_rate" yaml:"deposit_fee_rate"`
	ClaimFeeRate int64  `json:"claim_fee_rate" yaml:"claim_fee_rate"`
	Threshold int `json:"threshold" yaml:"threshold"`
}

func NewParams(
	mortgageRage int64,
	depositFeeRate int64,
	claimFeeRate int64,
	threshold int,
) Params {
	return Params{
		MortgageRate:mortgageRage,
		DepositFeeRate:depositFeeRate,
		ClaimFeeRate:claimFeeRate,
		Threshold:threshold,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultMortgageRate,
		DefaultDepositFeeRate,
		DefaultClaimFeeRate,
		DefaultThreshold,
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