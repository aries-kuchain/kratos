package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	yaml "gopkg.in/yaml.v2"
)


type QueryDepositMortgageRatioResponse struct {
	DepositID string `json:"deposit_id" yaml:"deposit_id"`
	BaseMortgageRatio   sdk.Int                    `json:"base_mortgage_ratio" yaml:"base_mortgage_ratio"`
}

func NewQueryDepositMortgageRatioResponse(depositID string,baseRatio sdk.Int) QueryDepositMortgageRatioResponse {
	return QueryDepositMortgageRatioResponse{DepositID: depositID, BaseMortgageRatio: baseRatio}
}

func (v QueryDepositMortgageRatioResponse) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}