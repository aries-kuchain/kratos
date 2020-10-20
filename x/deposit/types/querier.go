package types

import ()

const (
	QueryDepositMortgageRatioParams = "deposit_mortgage_ratio"
	QueryAllDeposit                 = "all_deposit"
	QueryDepositInfo                = "deposit_info"
	QueryUserDeposit                = "user_deposit"
	QueryCashReadyDeposit                = "cashready_deposit"

)

type QueryDepositParams struct {
	DepositID string `json:"deposit_id" yaml:"deposit_id"`
}

func NewQueryDepositParams(depositID string) QueryDepositParams {
	return QueryDepositParams{
		DepositID: depositID,
	}
}

//---------------------------------------------------------------------------------------------------------------------------
type QueryUserDepositParams struct {
	OwerAccount AccountID `json:"owner_account" yaml:"owner_account"`
}

func NewQueryUserDepositParams(ownerAccount AccountID) QueryUserDepositParams {
	return QueryUserDepositParams{
		OwerAccount: ownerAccount,
	}
}
