package types

import ()

const (
	QuerySingerInfo = "singer_info"
	QueryAllDeposit = "all_deposit"
	QueryDepositInfo = "deposit_info"
	QuerySingerDeposit = "singer_deposit"
)

type QuerySingerInfoParams struct {
	SingerAccount AccountID `json:"singer_account" yaml:"singer_account"`
}

func NewQuerySingerInfoParams(singerAccount AccountID) QuerySingerInfoParams {
	return QuerySingerInfoParams{
		SingerAccount: singerAccount,
	}
}
//-----------------------------------------------------------------------------------------------------------------------
type QueryDepositInfoParams struct {
	DepositID string `json:"deposit_id" yaml:"deposit_id"`
}

func NewQueryDepositParams(depositID string) QueryDepositInfoParams {
	return QueryDepositInfoParams{
		DepositID: depositID,
	}
}
//-----------------------------------------------------------------------------------------------------------------------
