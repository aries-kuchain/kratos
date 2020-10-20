package types

import (
	"github.com/KuChainNetwork/kuchain/chain/msg"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	RouterKeyName = chainTypes.MustName(RouterKey)
)

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgCreateDeposit struct {
	chainTypes.KuMsg
}

func NewKuMsgCreateDeposit(auth sdk.AccAddress, ownerAccountID AccountID, amount Coin) KuMsgCreateDeposit {
	return KuMsgCreateDeposit{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgCreateDeposit{
				OwnerAccount: ownerAccountID,
				Amount:       amount,
			}),
		),
	}
}

func (msg KuMsgCreateDeposit) ValidateBasic() error {
	msgData := MsgCreateDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgCreateLegalCoin struct {
	chainTypes.KuMsg
}

func NewKuMsgCreateLegalCoin(auth sdk.AccAddress, systemAccountID AccountID, amount Coin, symbol chainTypes.Name) KuMsgCreateLegalCoin {
	return KuMsgCreateLegalCoin{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgCreateLegalCoin{
				SystemAccount: systemAccountID,
				MaxSupply:     amount,
				Symbol:        symbol,
			}),
		),
	}
}

func (msg KuMsgCreateLegalCoin) ValidateBasic() error {
	msgData := MsgCreateLegalCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgProhibitLegalCoin struct {
	chainTypes.KuMsg
}

func NewKuMsgProhibitLegalCoin(auth sdk.AccAddress, systemAccountID AccountID, amount Coin) KuMsgProhibitLegalCoin {
	return KuMsgProhibitLegalCoin{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgProhibitLegalCoin{
				SystemAccount: systemAccountID,
				Amount:        amount,
			}),
		),
	}
}

func (msg KuMsgProhibitLegalCoin) ValidateBasic() error {
	msgData := MsgProhibitLegalCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgPermintLegalCoin struct {
	chainTypes.KuMsg
}

func NewKuMsgPermintLegalCoin(auth sdk.AccAddress, systemAccountID AccountID, amount Coin) KuMsgPermintLegalCoin {
	return KuMsgPermintLegalCoin{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgProhibitLegalCoin{
				SystemAccount: systemAccountID,
				Amount:        amount,
			}),
		),
	}
}

func (msg KuMsgPermintLegalCoin) ValidateBasic() error {
	msgData := MsgPermintLegalCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgSubmitSpv struct {
	chainTypes.KuMsg
}

func NewKuMsgSubmitSpv(auth sdk.AccAddress, spvInfo singerTypes.SpvInfo) KuMsgSubmitSpv {
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
type KuMsgTransferDeposit struct {
	chainTypes.KuMsg
}

func NewKuMsgTransferDeposit(auth sdk.AccAddress, depositID string, from, to AccountID, memo string) KuMsgTransferDeposit {
	return KuMsgTransferDeposit{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgTransferDeposit{
				DepositID: depositID,
				From:      from,
				To:        to,
				Memo:      memo,
			}),
		),
	}
}

func (msg KuMsgTransferDeposit) ValidateBasic() error {
	msgData := MsgTransferDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgDepositToCoin struct {
	chainTypes.KuMsg
}

func NewKuMsgDepositToCoin(auth sdk.AccAddress, depositID string, owner AccountID) KuMsgDepositToCoin {
	return KuMsgDepositToCoin{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgDepositToCoin{
				DepositID: depositID,
				Owner:     owner,
			}),
		),
	}
}

func (msg KuMsgDepositToCoin) ValidateBasic() error {
	msgData := MsgDepositToCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgDepositClaimCoin struct {
	chainTypes.KuMsg
}

func NewKuMsgDepositClaimCoin(auth sdk.AccAddress, depositID string, owner AccountID, asset Coin, claimAddress string) KuMsgDepositClaimCoin {
	return KuMsgDepositClaimCoin{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithTransfer(owner, ModuleAccountID, chainTypes.Coins{asset}),
			msg.WithData(Cdc(), &MsgDepositClaimCoin{
				DepositID:    depositID,
				Owner:        owner,
				Asset:        asset,
				ClaimAddress: claimAddress,
			}),
		),
	}
}

func (msg KuMsgDepositClaimCoin) ValidateBasic() error {
	if err := msg.KuMsg.ValidateTransfer(); err != nil {
		return err
	}

	msgData := MsgDepositClaimCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	if err := msg.KuMsg.ValidateTransferRequire(ModuleAccountID, chainTypes.NewCoins(msgData.Asset)); err != nil {
		return chainTypes.ErrKuMsgInconsistentAmount
	}
	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgFinishDeposit struct {
	chainTypes.KuMsg
}

func NewKuMsgFinishDeposit(auth sdk.AccAddress, depositID string, owner AccountID) KuMsgFinishDeposit {
	return KuMsgFinishDeposit{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgFinishDeposit{
				DepositID: depositID,
				Owner:     owner,
			}),
		),
	}
}

func (msg KuMsgFinishDeposit) ValidateBasic() error {

	msgData := MsgFinishDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgWaitTimeout struct {
	chainTypes.KuMsg
}

func NewKuMsgWaitTimeout(auth sdk.AccAddress, depositID string, owner AccountID) KuMsgWaitTimeout {
	return KuMsgWaitTimeout{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgWaitTimeout{
				DepositID: depositID,
				Owner:     owner,
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
type KuMsgReportWrongSpv struct {
	chainTypes.KuMsg
}

func NewKuMsgReportWrongSpv(auth sdk.AccAddress, depositID string, owner AccountID) KuMsgReportWrongSpv {
	return KuMsgReportWrongSpv{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgReportWrongSpv{
				DepositID: depositID,
				Owner:     owner,
			}),
		),
	}
}

func (msg KuMsgReportWrongSpv) ValidateBasic() error {

	msgData := MsgReportWrongSpv{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgJudgeDepositSpv struct {
	chainTypes.KuMsg
}

func NewKuMsgJudgeDepositSpv(auth sdk.AccAddress, depositID string, systemAccount AccountID, spvIsRight bool, feeToSinger bool) KuMsgJudgeDepositSpv {
	return KuMsgJudgeDepositSpv{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgJudgeDepositSpv{
				DepositID:     depositID,
				SystemAccount: systemAccount,
				SpvIsRight:    spvIsRight,
				FeeToSinger:   feeToSinger,
			}),
		),
	}
}

func (msg KuMsgJudgeDepositSpv) ValidateBasic() error {

	msgData := MsgJudgeDepositSpv{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgClaimAberrant struct {
	chainTypes.KuMsg
}

func NewKuMsgClaimAberrant(auth sdk.AccAddress, depositID string, claimAccount AccountID, amount Coin) KuMsgClaimAberrant {
	return KuMsgClaimAberrant{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithTransfer(claimAccount, ModuleAccountID, chainTypes.Coins{amount}),
			msg.WithData(Cdc(), &MsgClaimAberrant{
				DepositID:    depositID,
				ClaimAccount: claimAccount,
				Amount:       amount,
			}),
		),
	}
}

func (msg KuMsgClaimAberrant) ValidateBasic() error {
	if err := msg.KuMsg.ValidateTransfer(); err != nil {
		return err
	}

	msgData := MsgClaimAberrant{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	if err := msg.KuMsg.ValidateTransferRequire(ModuleAccountID, chainTypes.NewCoins(msgData.Amount)); err != nil {
		return chainTypes.ErrKuMsgInconsistentAmount
	}
	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgClaimMortgage struct {
	chainTypes.KuMsg
}

func NewKuMsgClaimMortgage(auth sdk.AccAddress, depositID string, claimAccount AccountID, amount Coin) KuMsgClaimMortgage {
	return KuMsgClaimMortgage{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithTransfer(claimAccount, ModuleAccountID, chainTypes.Coins{amount}),
			msg.WithData(Cdc(), &MsgClaimMortgage{
				DepositID:    depositID,
				ClaimAccount: claimAccount,
				Amount:       amount,
			}),
		),
	}
}

func (msg KuMsgClaimMortgage) ValidateBasic() error {
	if err := msg.KuMsg.ValidateTransfer(); err != nil {
		return err
	}

	msgData := MsgClaimMortgage{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	if err := msg.KuMsg.ValidateTransferRequire(ModuleAccountID, chainTypes.NewCoins(msgData.Amount)); err != nil {
		return chainTypes.ErrKuMsgInconsistentAmount
	}
	return msgData.ValidateBasic()
}

//---------------------------------------------------------------------------------------------------------------------------------------------
type KuMsgCashReadyDeposit struct {
	chainTypes.KuMsg
}

func NewKuMsgCashReadyDeposit(auth sdk.AccAddress, depositID string, operator AccountID) KuMsgCashReadyDeposit {
	return KuMsgCashReadyDeposit{
		*msg.MustNewKuMsg(
			RouterKeyName,
			msg.WithAuth(auth),
			msg.WithData(Cdc(), &MsgCashReadyDeposit{
				DepositID: depositID,
				Operator:  operator,
			}),
		),
	}
}

func (msg KuMsgCashReadyDeposit) ValidateBasic() error {

	msgData := MsgCashReadyDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return err
	}

	return msgData.ValidateBasic()
}
