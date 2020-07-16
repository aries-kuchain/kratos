package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(StdTx{}, "kuchain/Tx", nil)
	cdc.RegisterInterface((*KuMsgData)(nil), nil)
}

// module wide codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

// Cdc get codec for types
func Cdc() *codec.Codec {
	return ModuleCdc
}