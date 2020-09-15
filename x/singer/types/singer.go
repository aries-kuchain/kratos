package types

import (
	//	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	yaml "gopkg.in/yaml.v2"
	//	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
)

type SingerStatus int32

// staking constants
const (
	InActive SingerStatus = 1
	Active   SingerStatus = 2
	Lock     SingerStatus = 3

	SingerTypeInActive = "InActive"
	SingerTypeActive   = "Active"
	SingerTypeLock     = "Lock"
)

// Equal compares two BondStatus instances
func (b SingerStatus) Equal(b2 SingerStatus) bool {
	return byte(b) == byte(b2)
}

// String implements the Stringer interface for BondStatus.
func (b SingerStatus) String() string {
	switch b {
	case InActive:
		return SingerTypeInActive

	case Active:
		return SingerTypeActive

	case Lock:
		return SingerTypeLock

	default:
		panic("invalid singer status")
	}
}

type SingerInfo struct {
	SingerAccount     AccountID    `json:"singer_account" yaml:"singer_account"`
	AccessAsset       sdk.Int      `json:"access_asset" yaml:"access_asset"`
	Status            SingerStatus `json:"status" yaml:"status"`
	SignatureMortgage sdk.Int      `json:"signature_mortgage" yaml:"signature_mortgage"`
}

func NewSingerInfo(singerName AccountID) SingerInfo {
	return SingerInfo{
		SingerAccount:     singerName,
		AccessAsset:       sdk.ZeroInt(),
		Status:            InActive,
		SignatureMortgage: sdk.ZeroInt(),
	}
}

// return the redelegation
func MustMarshalSingerInfo(cdc *codec.Codec, singer SingerInfo) []byte {
	return cdc.MustMarshalBinaryBare(&singer)
}

// unmarshal a redelegation from a store value
func MustUnmarshalSingerInfo(cdc *codec.Codec, value []byte) SingerInfo {
	singer, err := UnmarshalSingerInfo(cdc, value)
	if err != nil {
		panic(err)
	}
	return singer
}

// unmarshal a redelegation from a store value
func UnmarshalSingerInfo(cdc *codec.Codec, value []byte) (v SingerInfo, err error) {
	err = cdc.UnmarshalBinaryBare(value, &v)
	return v, err
}

// String implements the Stringer interface for a SingerInfo object.
func (v SingerInfo) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}
