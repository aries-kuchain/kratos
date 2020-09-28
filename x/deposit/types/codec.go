package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(KuMsgCreateDeposit{}, "deposit/KuMsgCreateDeposit", nil)
	cdc.RegisterConcrete(KuMsgCreateLegalCoin{}, "deposit/KuMsgCreateLegalCoin", nil)
	cdc.RegisterConcrete(KuMsgPermintLegalCoin{}, "deposit/KuMsgPermintLegalCoin", nil)
	cdc.RegisterConcrete(KuMsgProhibitLegalCoin{}, "deposit/KuMsgProhibitLegalCoin", nil)
	cdc.RegisterConcrete(KuMsgSubmitSpv{}, "deposit/KuMsgSubmitSpv", nil)
	cdc.RegisterConcrete(KuMsgTransferDeposit{}, "deposit/KuMsgTransferDeposit", nil)
	cdc.RegisterConcrete(KuMsgDepositToCoin{}, "deposit/KuMsgDepositToCoin", nil)
	cdc.RegisterConcrete(KuMsgDepositClaimCoin{}, "deposit/KuMsgDepositClaimCoin", nil)
	cdc.RegisterConcrete(KuMsgFinishDeposit{}, "deposit/KuMsgFinishDeposit", nil)
	cdc.RegisterConcrete(KuMsgWaitTimeout{}, "deposit/KuMsgWaitTimeout", nil)
	cdc.RegisterConcrete(KuMsgReportWrongSpv{}, "deposit/KuMsgReportWrongSpv", nil)
	cdc.RegisterConcrete(KuMsgJudgeDepositSpv{}, "deposit/KuMsgJudgeDepositSpv", nil)

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
