package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	yaml "gopkg.in/yaml.v2"
	"time"
)

type PriceInfo struct {
	Base       Coin
	Quote      Coin
	RecordTime time.Time
	Remark     string
}

func NewPriceInfo(base, quote Coin, remark string) PriceInfo {
	return PriceInfo{
		Base:       base,
		Quote:      quote,
		RecordTime: time.Now().UTC(),
		Remark:     remark,
	}
}

// return the redelegation
func MustMarshalPriceInfo(cdc *codec.Codec, priceInfo PriceInfo) []byte {
	return cdc.MustMarshalBinaryBare(&priceInfo)
}

// unmarshal a redelegation from a store value
func MustUnmarshalPriceInfo(cdc *codec.Codec, value []byte) PriceInfo {
	priceInfo, err := UnmarshalPriceInfo(cdc, value)
	if err != nil {
		panic(err)
	}
	return priceInfo
}

// unmarshal a redelegation from a store value
func UnmarshalPriceInfo(cdc *codec.Codec, value []byte) (v PriceInfo, err error) {
	err = cdc.UnmarshalBinaryBare(value, &v)
	return v, err
}

// String implements the Stringer interface for a PriceInfo object.
func (v PriceInfo) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}
