package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	yaml "gopkg.in/yaml.v2"
)

type QuerySingerInfoResponse struct {
	SingerAccount     AccountID    `json:"singer_account" yaml:"singer_account"`
	AccessAsset       sdk.Int      `json:"access_asset" yaml:"access_asset"`
	Status            SingerStatus `json:"status" yaml:"status"`
	SignatureMortgage sdk.Int      `json:"signature_mortgage" yaml:"signature_mortgage"`
	LockMortgage      sdk.Int      `json:"locked_morgage" yaml:"locked_morgage"`
}

func NewQueryDepositMortgageRatioResponse(singerName AccountID, accessAsset sdk.Int, status SingerStatus, signatureMortgage, lockMortgage sdk.Int) QuerySingerInfoResponse {
	return QuerySingerInfoResponse{
		SingerAccount:     singerName,
		AccessAsset:       accessAsset,
		Status:            status,
		SignatureMortgage: signatureMortgage,
		LockMortgage:      lockMortgage,
	}
}

func (v QuerySingerInfoResponse) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

type QueryAllDepositResponse struct {
	DepositInfos []DepositInfo  `json:"all_deposit_info" yaml:"all_deposit_info"`
}

func NewQueryAllDepositResponse(depositInfos []DepositInfo) QueryAllDepositResponse{
	var result QueryAllDepositResponse
	for _,depositInfo := range depositInfos {
		result.DepositInfos = append(result.DepositInfos,depositInfo)
	}
	return result
}

func NewQuerySingerDepositResponse(depositInfos []DepositInfo,singerAccount AccountID) QueryAllDepositResponse{
	var result QueryAllDepositResponse
	for _,depositInfo := range depositInfos {
		for _,depositSinger := range depositInfo.Singers {
			if depositSinger.Eq(singerAccount) &&  depositInfo.Status != Close {
				result.DepositInfos = append(result.DepositInfos,depositInfo)
			}
		}
		
	}
	return result
}


func (v QueryAllDepositResponse) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}