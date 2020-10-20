package types

import ()

const (
	QueryFeeInfo = "fee_info"
)

type QueryFeeInfoParams struct {
	OwnerAccount AccountID `json:"singer_account" yaml:"singer_account"`
}

func NewQuerySingerInfoParams(ownerAccount AccountID) QueryFeeInfoParams {
	return QueryFeeInfoParams{
		OwnerAccount: ownerAccount,
	}
}

//-----------------------------------------------------------------------------------------------------------------------
