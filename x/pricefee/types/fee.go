package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	yaml "gopkg.in/yaml.v2"
)

type FeeInfo struct {
	Owner AccountID
	PrestoreFee sdk.Int
    LockedFee  sdk.Int
}

func NewFeeInfo (owner AccountID) FeeInfo {
	return FeeInfo {
		Owner:owner,
		PrestoreFee:sdk.ZeroInt(),
		LockedFee:sdk.ZeroInt(),
	}
}


// return the redelegation
func MustMarshalFeeInfo(cdc *codec.Codec, feeInfo FeeInfo) []byte {
	return cdc.MustMarshalBinaryBare(&feeInfo)
}

// unmarshal a redelegation from a store value
func MustUnmarshalFeeInfo(cdc *codec.Codec, value []byte) FeeInfo {
	feeInfo, err := UnmarshalFeeInfo(cdc, value)
	if err != nil {
		panic(err)
	}
	return feeInfo
}

// unmarshal a redelegation from a store value
func UnmarshalFeeInfo(cdc *codec.Codec, value []byte) (v FeeInfo, err error) {
	err = cdc.UnmarshalBinaryBare(value, &v)
	return v, err
}

// String implements the Stringer interface for a FeeInfo object.
func (v FeeInfo) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}
