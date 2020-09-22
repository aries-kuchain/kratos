package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	yaml "gopkg.in/yaml.v2"
	"fmt"
//	"encoding/hex"
)

type DepositStatus uint32

const (
	Open DepositStatus = 1
	AddressReady DepositStatus = 2
	Close DepositStatus = 3

	DepositStatusOpen = "open"
	DepositStatusAddressReady = "addressReady"
	DepositStatusClose = "close"
)

type DepositInfo struct {
	DepositID string `json:"deposit_id" yaml:"deposit_id"`
	Threshold int `json:"threshold" yaml:"threshold"`
	Singers []AccountID `json:"singers" yaml:"singers"`
	minStake sdk.Int  
	Status DepositStatus `json:"status" yaml:"status"`
}


func NewDepositInfo(depositID string,threshold int,minStake sdk.Int) DepositInfo {
	return DepositInfo{
		DepositID:     depositID,
		Threshold:       threshold,
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

func (v *DepositInfo) SetSingers(singers SingerInfos) {
		for _,singerInfo := range singers {
		v.Singers = append(v.Singers ,singerInfo.SingerAccount)
	}
} 

func (v DepositInfo) CheckSinger(singerAccount AccountID) bool {
	for _,depositSinger := range v.Singers {
		if depositSinger.Eq(singerAccount) {
			return true
		}
	}
	return false
}

//--------------------------------------------------------------------------------------------------
type DepositBtcAddress struct {
	DepositID string
	Singer AccountID
	BtcAddress []byte
}

func NewDepositBtcAddress(depositID string,singer AccountID,btcAddress []byte) DepositBtcAddress {
	return DepositBtcAddress{
		DepositID:     depositID,
		Singer:       singer,
		BtcAddress: btcAddress,
	}
}

// return the redelegation
func MustMarshalDepositBtcAddress(cdc *codec.Codec, depositBtcAddress DepositBtcAddress) []byte {
	return cdc.MustMarshalBinaryBare(&depositBtcAddress)
}

// unmarshal a redelegation from a store value
func MustUnmarshalDepositBtcAddress(cdc *codec.Codec, value []byte) DepositBtcAddress {
	depositBtcAddress, err := UnmarshalDepositBtcAddress(cdc, value)
	if err != nil {
		panic(err)
	}
	return depositBtcAddress
}

// unmarshal a redelegation from a store value
func UnmarshalDepositBtcAddress(cdc *codec.Codec, value []byte) (v DepositBtcAddress, err error) {
	err = cdc.UnmarshalBinaryBare(value, &v)
	return v, err
}

// // String implements the Stringer interface for a SingerInfo object.
// func (v DepositBtcAddress) String() string {
// 	out, _ := yaml.Marshal(v)
// 	return string(out)
// }

func (v DepositBtcAddress) String() string {
	return fmt.Sprintf(`DepositID:%s\n
	Singer:%s\n
	BtcAddress:%x\n`,v.DepositID,v.Singer.String(),v.BtcAddress)
}