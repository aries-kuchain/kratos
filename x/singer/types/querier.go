package types

import ()

const (
	QuerySingerInfo = "singer_info"
)

type QuerySingerInfoParams struct {
	SingerAccount AccountID `json:"deposit_id" yaml:"deposit_id"`
}

func NewQueryDepositParams(singerAccount AccountID) QuerySingerInfoParams {
	return QuerySingerInfoParams{
		SingerAccount: singerAccount,
	}
}
