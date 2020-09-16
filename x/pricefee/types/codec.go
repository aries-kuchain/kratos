package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(KuMsgOpenFee{}, "kuchain/KuMsgOpenFee", nil)
	cdc.RegisterConcrete(KuMsgPrestoreFee{}, "kuchain/KuMsgPrestoreFee", nil)
	cdc.RegisterConcrete(KuMsgClaimFee{}, "kuchain/KuMsgClaimFee", nil)
	cdc.RegisterConcrete(KuMsgSetPrice{}, "kuchain/KuMsgSetPrice", nil)

}

var (
	ModuleCdc = codec.New()
)

// Cdc get codec for types
func Cdc() *codec.Codec {
	return ModuleCdc
}

func init() {
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
