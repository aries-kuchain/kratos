package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	yaml "gopkg.in/yaml.v2"
)

type DepositStatus uint32

const (
	Open DepositStatus = 1
	Close DepositStatus = 2

	DepositStatusOpen = "open"
	DepositStatusClose = "close"
)

type DepositInfo struct {
	DepositID string
	Threshold int
	Singers SingerInfos
	minStake sdk.Int
	Status DepositStatus
}


func NewDepositInfo(depositID string,threshold int,singer SingerInfos,minStake sdk.Int) DepositInfo {
	return DepositInfo{
		DepositID:     depositID,
		Threshold:       threshold,
		Singers:            singer,
		minStake: minStake,
		Status:Open,
	}
}

// return the redelegation
func MustMarshalDepositInfo(cdc *codec.Codec, depositInfo DepositInfo) []byte {
	return cdc.MustMarshalBinaryBare(&depositInfo)
}

// unmarshal a redelegation from a store value
func MustUnmarshalDepositInfo(cdc *codec.Codec, value []byte) DepositInfo {
	depositInfo, err := UnmarshalDepositInfo(cdc, value)
	if err != nil {
		panic(err)
	}
	return depositInfo
}

// unmarshal a redelegation from a store value
func UnmarshalDepositInfo(cdc *codec.Codec, value []byte) (v DepositInfo, err error) {
	err = cdc.UnmarshalBinaryBare(value, &v)
	return v, err
}

// String implements the Stringer interface for a SingerInfo object.
func (v DepositInfo) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}