package types

import (
	"github.com/KuChainNetwork/kuchain/chain/msg"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//	"github.com/tendermint/tendermint/crypto"
)

var (
	RouterKeyName = chainTypes.MustName(RouterKey)
)

type KuMsgRegisterSinger struct {
	chainTypes.KuMsg
}

func NewKuMsgRegisterSinger(auth sdk.AccAddress, singerAccount AccountID) KuMsgRegisterSinger {
	return KuMsgRegisterSinger{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgRegisterSinger{
				SingerAccount: singerAccount,
			}),
		),
	}
}

func (msg KuMsgRegisterSinger) ValidateBasic() error {
	msgData := MsgRegisterSinger{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgPayAccess struct {
	chainTypes.KuMsg
}

func NewKuMsgPayAccess(auth sdk.AccAddress, singerAccount AccountID, amount Coin) KuMsgPayAccess {
	return KuMsgPayAccess{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithTransfer(singerAccount, ModuleAccountID, chainTypes.Coins{amount}),
			msg.WithData(Cdc(), &MsgPayAccess{
				SingerAccount: singerAccount,
				Amount:        amount,
			}),
		),
	}
}

func (msg KuMsgPayAccess) ValidateBasic() error {
	if err := msg.KuMsg.ValidateTransfer(); err != nil {
		return err
	}

	msgData := MsgPayAccess{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	if err := msg.KuMsg.ValidateTransferRequire(ModuleAccountID, chainTypes.NewCoins(msgData.Amount)); err != nil {
		return chainTypes.ErrKuMsgInconsistentAmount
	}
	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgActiveSinger struct {
	chainTypes.KuMsg
}

func NewKuMsgActiveSinger(auth sdk.AccAddress, systemAccount, singerAccount AccountID) KuMsgActiveSinger {
	return KuMsgActiveSinger{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgActiveSinger{
				SystemAccount: systemAccount,
				SingerAccount: singerAccount,
			}),
		),
	}
}

func (msg KuMsgActiveSinger) ValidateBasic() error {
	msgData := MsgActiveSinger{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	return msgData.ValidateBasic()
}

//--------------------------------------------------------------------------------------------------------------------------
type KuMsgBTCMortgage struct {
	chainTypes.KuMsg
}

func NewKuMsgBTCMortgage(auth sdk.AccAddress, singerAccount AccountID, amount Coin) KuMsgBTCMortgage {
	return KuMsgBTCMortgage{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithTransfer(singerAccount, ModuleAccountID, chainTypes.Coins{amount}),
			msg.WithData(Cdc(), &MsgPayBTCMortgate{
				SingerAccount: singerAccount,
				Amount:        amount,
			}),
		),
	}
}

func (msg KuMsgBTCMortgage) ValidateBasic() error {
	if err := msg.KuMsg.ValidateTransfer(); err != nil {
		return err
	}

	msgData := MsgPayBTCMortgate{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	if err := msg.KuMsg.ValidateTransferRequire(ModuleAccountID, chainTypes.NewCoins(msgData.Amount)); err != nil {
		return chainTypes.ErrKuMsgInconsistentAmount
	}
	return msgData.ValidateBasic()
}

//--------------------------------------------------------------------------------------------------------------------------
type KuMsgClaimBTCMortgate struct {
	chainTypes.KuMsg
}

func NewKuMsgClaimBTCMortgate(auth sdk.AccAddress, singerAccount AccountID, amount Coin) KuMsgClaimBTCMortgate {
	return KuMsgClaimBTCMortgate{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgClaimBTCMortgate{
				SingerAccount: singerAccount,
				Amount:        amount,
			}),
		),
	}
}

func (msg KuMsgClaimBTCMortgate) ValidateBasic() error {
	msgData := MsgClaimBTCMortgate{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	return msgData.ValidateBasic()
}

//--------------------------------------------------------------------------------------------------------------------------

type KuMsgClaimAccess struct {
	chainTypes.KuMsg
}

func NewKuMsgClaimAccess(auth sdk.AccAddress, singerAccount AccountID) KuMsgClaimAccess {
	return KuMsgClaimAccess{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgClaimAccess{
				SingerAccount: singerAccount,
			}),
		),
	}
}

func (msg KuMsgClaimAccess) ValidateBasic() error {
	msgData := MsgClaimAccess{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	return msgData.ValidateBasic()
}

//--------------------------------------------------------------------------------------------------------------------------
type KuMsgLogoutSinger struct {
	chainTypes.KuMsg
}

func NewKuMsgLogoutSinger(auth sdk.AccAddress, singerAccount AccountID) KuMsgLogoutSinger {
	return KuMsgLogoutSinger{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgLogoutSinger{
				SingerAccount: singerAccount,
			}),
		),
	}
}

func (msg KuMsgLogoutSinger) ValidateBasic() error {
	msgData := MsgLogoutSinger{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	return msgData.ValidateBasic()
}

//--------------------------------------------------------------------------------------------------------------------------
type KuMsgMsgSetBtcAddress struct {
	chainTypes.KuMsg
}

func NewKuMsgMsgSetBtcAddress(auth sdk.AccAddress, singerAccount AccountID, depositID string, btcAddress string) KuMsgMsgSetBtcAddress {
	return KuMsgMsgSetBtcAddress{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgSetBtcAddress{
				SingerAccount: singerAccount,
				DepoistID:     depositID,
				BtcAddress:    btcAddress,
			}),
		),
	}
}

func (msg KuMsgMsgSetBtcAddress) ValidateBasic() error {
	msgData := MsgSetBtcAddress{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	return msgData.ValidateBasic()
}

//--------------------------------------------------------------------------------------------------------------------------
type KuMsgActiveDeposit struct {
	chainTypes.KuMsg
}

func NewKuMsgActiveDeposit(auth sdk.AccAddress, singerAccount AccountID, depositID string) KuMsgActiveDeposit {
	return KuMsgActiveDeposit{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgActiveDeposit{
				SingerAccount: singerAccount,
				DepositID:     depositID,
			}),
		),
	}
}

func (msg KuMsgActiveDeposit) ValidateBasic() error {
	msgData := MsgActiveDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}
	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgSubmitSpv struct {
	chainTypes.KuMsg
}

func NewKuMsgSubmitSpv(auth sdk.AccAddress, spvInfo SpvInfo) KuMsgSubmitSpv {
	return KuMsgSubmitSpv{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgSubmitSpv{
				SpvInfo: spvInfo,
			}),
		),
	}
}

func (msg KuMsgSubmitSpv) ValidateBasic() error {
	msgData := MsgSubmitSpv{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgWaitTimeout struct {
	chainTypes.KuMsg
}

func NewKuMsgWaitTimeout(auth sdk.AccAddress, depositID string, singerAccount AccountID) KuMsgWaitTimeout {
	return KuMsgWaitTimeout{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgWaitTimeout{
				DepositID:     depositID,
				SingerAccount: singerAccount,
			}),
		),
	}
}

func (msg KuMsgWaitTimeout) ValidateBasic() error {

	msgData := MsgWaitTimeout{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgReportSpvWrong struct {
	chainTypes.KuMsg
}

func NewKuMsgReportSpvWrong(auth sdk.AccAddress, depositID string, singerAccount AccountID) KuMsgReportSpvWrong {
	return KuMsgReportSpvWrong{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgReportSpvWrong{
				DepositID:     depositID,
				SingerAccount: singerAccount,
			}),
		),
	}
}

func (msg KuMsgReportSpvWrong) ValidateBasic() error {

	msgData := MsgReportSpvWrong{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}
