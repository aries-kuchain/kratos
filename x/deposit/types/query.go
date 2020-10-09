package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	yaml "gopkg.in/yaml.v2"
)


type QueryDepositMortgageRatioResponse struct {
	DepositID string `json:"deposit_id" yaml:"deposit_id"`
	BaseMortgageRatio   sdk.Int                    `json:"base_mortgage_ratio" yaml:"base_mortgage_ratio"`
	CurrentMortgageRatio   sdk.Int                    `json:"current_mortgage_ratio" yaml:"current_mortgage_ratio"`
	PunishMortgageRatio   sdk.Int                    `json:"punish_mortgage_ratio" yaml:"punish_mortgage_ratio"`
}

func NewQueryDepositMortgageRatioResponse(depositID string,baseRatio,currentRatio,punishRatio sdk.Int) QueryDepositMortgageRatioResponse {
	return QueryDepositMortgageRatioResponse{DepositID: depositID, BaseMortgageRatio: baseRatio,CurrentMortgageRatio:currentRatio,PunishMortgageRatio:punishRatio}
}

func (v QueryDepositMortgageRatioResponse) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}