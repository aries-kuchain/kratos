package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(KuMsgRegisterSinger{}, "kuchain/KuMsgRegisterSinger", nil)
	cdc.RegisterConcrete(KuMsgPayAccess{}, "kuchain/KuMsgPayAccess", nil)
	cdc.RegisterConcrete(KuMsgActiveSinger{}, "kuchain/KuMsgActiveSinger", nil)
	cdc.RegisterConcrete(KuMsgBTCMortgage{}, "kuchain/KuMsgBTCMortgage", nil)
	cdc.RegisterConcrete(KuMsgClaimBTCMortgate{}, "kuchain/KuMsgClaimBTCMortgate", nil)
	cdc.RegisterConcrete(KuMsgClaimAccess{}, "kuchain/KuMsgClaimAccess", nil)
	cdc.RegisterConcrete(KuMsgLogoutSinger{}, "kuchain/KuMsgLogoutSinger", nil)
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