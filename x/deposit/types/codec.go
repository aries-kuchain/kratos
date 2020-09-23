package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(KuMsgCreateDeposit{}, "kuchain/KuMsgCreateDeposit", nil)
	cdc.RegisterConcrete(KuMsgCreateLegalCoin{}, "kuchain/KuMsgCreateLegalCoin", nil)
	cdc.RegisterConcrete(KuMsgPermintLegalCoin{}, "kuchain/KuMsgPermintLegalCoin", nil)
	cdc.RegisterConcrete(KuMsgProhibitLegalCoin{}, "kuchain/KuMsgProhibitLegalCoin", nil)
	cdc.RegisterConcrete(KuMsgSubmitSpv{}, "deposit/KuMsgSubmitSpv", nil)
	cdc.RegisterConcrete(KuMsgTransferDeposit{}, "deposit/KuMsgTransferDeposit", nil)
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
