
package types

import ()

const (
	QueryDepositMortgageRatioParams                      = "deposit_mortgage_ratio"

)

type QueryDepositParams struct {
	DepositID string `json:"deposit_id" yaml:"deposit_id"`
}

func NewQueryDepositParams(depositID string) QueryDepositParams {
	return QueryDepositParams{
		DepositID: depositID,
	}
}