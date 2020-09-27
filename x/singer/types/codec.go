package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(KuMsgRegisterSinger{}, "singer/KuMsgRegisterSinger", nil)
	cdc.RegisterConcrete(KuMsgPayAccess{}, "singer/KuMsgPayAccess", nil)
	cdc.RegisterConcrete(KuMsgActiveSinger{}, "singer/KuMsgActiveSinger", nil)
	cdc.RegisterConcrete(KuMsgBTCMortgage{}, "singer/KuMsgBTCMortgage", nil)
	cdc.RegisterConcrete(KuMsgClaimBTCMortgate{}, "singer/KuMsgClaimBTCMortgate", nil)
	cdc.RegisterConcrete(KuMsgClaimAccess{}, "singer/KuMsgClaimAccess", nil)
	cdc.RegisterConcrete(KuMsgLogoutSinger{}, "singer/KuMsgLogoutSinger", nil)
	cdc.RegisterConcrete(KuMsgMsgSetBtcAddress{}, "singer/KuMsgMsgSetBtcAddress", nil)
	cdc.RegisterConcrete(KuMsgActiveDeposit{}, "singer/KuMsgActiveDeposit", nil)
	cdc.RegisterConcrete(KuMsgSubmitSpv{}, "singer/KuMsgSubmitSpv", nil)
	cdc.RegisterConcrete(KuMsgWaitTimeout{}, "singer/KuMsgWaitTimeout", nil)
	cdc.RegisterConcrete(KuMsgReportSpvWrong{}, "singer/KuMsgReportSpvWrong", nil)
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
