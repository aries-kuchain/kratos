package types

import (
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"
	"github.com/KuChainNetwork/kuchain/chain/hexutil"

)

//--------------------------------------------------------------------------------------------------------------
type MsgCreateDeposit struct {
	OwnerAccount AccountID `json:"owner_account" yaml:"owner_account"`
	Amount       Coin      `json:"amount" yaml:"amount"`
}

func NewMsgCreateDeposit(ownerAccount AccountID, amount Coin) MsgCreateDeposit {
	return MsgCreateDeposit{OwnerAccount: ownerAccount, Amount: amount}
}

// Route should return the name of the module
func (msg MsgCreateDeposit) Route() string { return RouterKey }

func (msg MsgCreateDeposit) Type() chainTypes.Name { return chainTypes.MustName("createdeposit") }

func (msg MsgCreateDeposit) Sender() AccountID {
	return msg.OwnerAccount
}

func (msg MsgCreateDeposit) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.OwnerAccount.Empty() {
		return ErrEmptyOwnerAccount
	}

	if !msg.Amount.Amount.IsPositive() {
		return ErrBadAmount
	}
	return nil
}

//----------------------------------------------------------------------------------------------------------------------------
type MsgProhibitLegalCoin struct {
	SystemAccount AccountID `json:"owner_account" yaml:"owner_account"`
	Amount        Coin      `json:"amount" yaml:"amount"`
}

func NewMsgProhibitLegalCoin(systemAccount AccountID, amount Coin) MsgProhibitLegalCoin {
	return MsgProhibitLegalCoin{SystemAccount: systemAccount, Amount: amount}
}

// Route should return the name of the module
func (msg MsgProhibitLegalCoin) Route() string { return RouterKey }

func (msg MsgProhibitLegalCoin) Type() chainTypes.Name {
	return chainTypes.MustName("prohibitlegalcoin")
}

func (msg MsgProhibitLegalCoin) Sender() AccountID {
	return msg.SystemAccount
}

func (msg MsgProhibitLegalCoin) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SystemAccount.Empty() {
		return ErrEmptyOwnerAccount
	}

	if !msg.Amount.Amount.IsPositive() {
		return ErrBadAmount
	}
	return nil
}

//----------------------------------------------------------------------------------------------------------------------------
type MsgPermintLegalCoin struct {
	SystemAccount AccountID `json:"owner_account" yaml:"owner_account"`
	Amount        Coin      `json:"amount" yaml:"amount"`
}

func NewMsgPermintLegalCoin(systemAccount AccountID, amount Coin) MsgPermintLegalCoin {
	return MsgPermintLegalCoin{SystemAccount: systemAccount, Amount: amount}
}

// Route should return the name of the module
func (msg MsgPermintLegalCoin) Route() string { return RouterKey }

func (msg MsgPermintLegalCoin) Type() chainTypes.Name { return chainTypes.MustName("permintlegalcoin") }

func (msg MsgPermintLegalCoin) Sender() AccountID {
	return msg.SystemAccount
}

func (msg MsgPermintLegalCoin) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SystemAccount.Empty() {
		return ErrEmptyOwnerAccount
	}

	if !msg.Amount.Amount.IsPositive() {
		return ErrBadAmount
	}
	return nil
}

//----------------------------------------------------------------------------------------------------------------------------
type MsgCreateLegalCoin struct {
	SystemAccount AccountID       `json:"owner_account" yaml:"owner_account"`
	MaxSupply     Coin            `json:"amount" yaml:"amount"`
	Symbol        chainTypes.Name `json:"symbol" yaml:"symbol"`
}

func NewMsgCreateLegalCoin(systemAccount AccountID, amount Coin) MsgCreateLegalCoin {
	return MsgCreateLegalCoin{SystemAccount: systemAccount, MaxSupply: amount}
}

// Route should return the name of the module
func (msg MsgCreateLegalCoin) Route() string { return RouterKey }

func (msg MsgCreateLegalCoin) Type() chainTypes.Name { return chainTypes.MustName("createlegalcoin") }

func (msg MsgCreateLegalCoin) Sender() AccountID {
	return msg.SystemAccount
}

func (msg MsgCreateLegalCoin) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SystemAccount.Empty() {
		return ErrEmptyOwnerAccount
	}

	if !msg.MaxSupply.Amount.IsPositive() {
		return ErrBadAmount
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgSubmitSpv struct {
	singerTypes.SpvInfo
}

func NewMsgSubmitSpv(spvInfo singerTypes.SpvInfo) MsgSubmitSpv {
	return MsgSubmitSpv{SpvInfo:spvInfo}
}

// Route should return the name of the module
func (msg MsgSubmitSpv) Route() string { return RouterKey }

func (msg MsgSubmitSpv) Type() chainTypes.Name { return chainTypes.MustName("depositsubmitspv") }

func (msg MsgSubmitSpv) Sender() AccountID {
	return msg.SpvInfo.SpvSubmiter
}

func (msg MsgSubmitSpv) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SpvInfo.SpvSubmiter.Empty() {
		return ErrEmptyOwnerAccount
	}

	if len(msg.SpvInfo.DepositID) == 0 {
		return ErrEmptyDepositID
	}

	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgTransferDeposit struct {
	DepositID string
	From AccountID
	To AccountID
	Memo string
}

func NewMsgTransferDeposit(depositID string,from,to AccountID,memo string) MsgTransferDeposit {
	return MsgTransferDeposit{
		DepositID:depositID,
		From:from,
		To:to,
		Memo:memo,
	}
}

// Route should return the name of the module
func (msg MsgTransferDeposit) Route() string { return RouterKey }

func (msg MsgTransferDeposit) Type() chainTypes.Name { return chainTypes.MustName("transferdeposit") }

func (msg MsgTransferDeposit) Sender() AccountID {
	return msg.From
}

func (msg MsgTransferDeposit) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.From.Empty() {
		return ErrEmptyOwnerAccount
	}

	if msg.To.Empty() {
		return ErrEmptyOwnerAccount
	}

	if len(msg.DepositID) == 0 {
		return ErrEmptyDepositID
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgDepositToCoin struct {
	DepositID string
	Owner AccountID
}

func NewMsgDepositToCoin(depositID string,owner AccountID) MsgDepositToCoin {
	return MsgDepositToCoin{
		DepositID:depositID,
		Owner:owner,
	}
}

// Route should return the name of the module
func (msg MsgDepositToCoin) Route() string { return RouterKey }

func (msg MsgDepositToCoin) Type() chainTypes.Name { return chainTypes.MustName("deposittocoin") }

func (msg MsgDepositToCoin) Sender() AccountID {
	return msg.Owner
}

func (msg MsgDepositToCoin) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.Owner.Empty() {
		return ErrEmptyOwnerAccount
	}

	if len(msg.DepositID) == 0 {
		return ErrEmptyDepositID
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgDepositClaimCoin struct {
	DepositID string
	Owner AccountID
	Asset Coin
	ClaimAddress string
}

func NewMsgDepositClaimCoin(depositID string,owner AccountID,asset Coin,claimAddress string) MsgDepositClaimCoin {
	return MsgDepositClaimCoin{
		DepositID:depositID,
		Owner:owner,
		Asset:asset,
		ClaimAddress:claimAddress,
	}
}

// Route should return the name of the module
func (msg MsgDepositClaimCoin) Route() string { return RouterKey }

func (msg MsgDepositClaimCoin) Type() chainTypes.Name { return chainTypes.MustName("deposittocoin") }

func (msg MsgDepositClaimCoin) Sender() AccountID {
	return msg.Owner
}

func (msg MsgDepositClaimCoin) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.Owner.Empty() {
		return ErrEmptyOwnerAccount
	}

	if len(msg.DepositID) == 0 {
		return ErrEmptyDepositID
	}

	if !msg.Asset.IsPositive() {
		return ErrBadAmount
	}

	if !hexutil.IsValidBtcAddress(msg.ClaimAddress) {
		return ErrWrongBtcAddress
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgFinishDeposit struct {
	DepositID string
	Owner AccountID
}

func NewMsgFinishDeposit(depositID string,owner AccountID,asset Coin,claimAddress []byte) MsgFinishDeposit {
	return MsgFinishDeposit{
		DepositID:depositID,
		Owner:owner,
	}
}

// Route should return the name of the module
func (msg MsgFinishDeposit) Route() string { return RouterKey }

func (msg MsgFinishDeposit) Type() chainTypes.Name { return chainTypes.MustName("finishdepsit") }

func (msg MsgFinishDeposit) Sender() AccountID {
	return msg.Owner
}

func (msg MsgFinishDeposit) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.Owner.Empty() {
		return ErrEmptyOwnerAccount
	}

	if len(msg.DepositID) == 0 {
		return ErrEmptyDepositID
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgWaitTimeout struct {
	DepositID string
	Owner AccountID
}

func NewMsgWaitTimeout(depositID string,owner AccountID,asset Coin) MsgWaitTimeout {
	return MsgWaitTimeout{
		DepositID:depositID,
		Owner:owner,
	}
}

// Route should return the name of the module
func (msg MsgWaitTimeout) Route() string { return RouterKey }

func (msg MsgWaitTimeout) Type() chainTypes.Name { return chainTypes.MustName("singertimeout") }

func (msg MsgWaitTimeout) Sender() AccountID {
	return msg.Owner
}

func (msg MsgWaitTimeout) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.Owner.Empty() {
		return ErrEmptyOwnerAccount
	}

	if len(msg.DepositID) == 0 {
		return ErrEmptyDepositID
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgReportWrongSpv struct {
	DepositID string
	Owner AccountID
}

func NewMsgReportWrongSpv(depositID string,owner AccountID,asset Coin ) MsgReportWrongSpv {
	return MsgReportWrongSpv{
		DepositID:depositID,
		Owner:owner,
	}
}

// Route should return the name of the module
func (msg MsgReportWrongSpv) Route() string { return RouterKey }

func (msg MsgReportWrongSpv) Type() chainTypes.Name { return chainTypes.MustName("reportwrongspv") }

func (msg MsgReportWrongSpv) Sender() AccountID {
	return msg.Owner
}

func (msg MsgReportWrongSpv) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.Owner.Empty() {
		return ErrEmptyOwnerAccount
	}

	if len(msg.DepositID) == 0 {
		return ErrEmptyDepositID
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgJudgeDepositSpv struct {
	DepositID string
	SystemAccount AccountID
	SpvIsRight	bool
	FeeToSinger bool
}

func NewMsgJudgeDepositSpv(depositID string,systemAccount AccountID,asset Coin,spvIsRight bool,feeToSinger bool) MsgJudgeDepositSpv {
	return MsgJudgeDepositSpv{
		DepositID:depositID,
		SystemAccount:systemAccount,
		SpvIsRight:spvIsRight,
		FeeToSinger:feeToSinger,
	}
}

// Route should return the name of the module
func (msg MsgJudgeDepositSpv) Route() string { return RouterKey }

func (msg MsgJudgeDepositSpv) Type() chainTypes.Name { return chainTypes.MustName("judgedepositspv") }

func (msg MsgJudgeDepositSpv) Sender() AccountID {
	return msg.SystemAccount
}

func (msg MsgJudgeDepositSpv) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SystemAccount.Empty() {
		return ErrEmptyOwnerAccount
	}

	if len(msg.DepositID) == 0 {
		return ErrEmptyDepositID
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgClaimAberrant struct {
	DepositID string
	ClaimAccount AccountID
	Amount	Coin
}

func NewMsgClaimAberrant(depositID string,claimAccount AccountID,amount Coin) MsgClaimAberrant {
	return MsgClaimAberrant{
		DepositID:depositID,
		ClaimAccount:claimAccount,
		Amount:amount,
	}
}

// Route should return the name of the module
func (msg MsgClaimAberrant) Route() string { return RouterKey }

func (msg MsgClaimAberrant) Type() chainTypes.Name { return chainTypes.MustName("claimaberrant") }

func (msg MsgClaimAberrant) Sender() AccountID {
	return msg.ClaimAccount
}

func (msg MsgClaimAberrant) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.ClaimAccount.Empty() {
		return ErrEmptyOwnerAccount
	}

	if len(msg.DepositID) == 0 {
		return ErrEmptyDepositID
	}
	
	if !msg.Amount.Amount.IsPositive() {
		return ErrBadAmount
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgClaimMortgage struct {
	DepositID string
	ClaimAccount AccountID
	Amount	Coin
}

func NewMsgClaimMortgage(depositID string,claimAccount AccountID,amount Coin) MsgClaimMortgage {
	return MsgClaimMortgage{
		DepositID:depositID,
		ClaimAccount:claimAccount,
		Amount:amount,
	}
}

// Route should return the name of the module
func (msg MsgClaimMortgage) Route() string { return RouterKey }

func (msg MsgClaimMortgage) Type() chainTypes.Name { return chainTypes.MustName("claimmortgage") }

func (msg MsgClaimMortgage) Sender() AccountID {
	return msg.ClaimAccount
}

func (msg MsgClaimMortgage) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.ClaimAccount.Empty() {
		return ErrEmptyOwnerAccount
	}

	if len(msg.DepositID) == 0 {
		return ErrEmptyDepositID
	}
	
	if !msg.Amount.Amount.IsPositive() {
		return ErrBadAmount
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgCashReadyDeposit struct {
	DepositID string
	Operator  AccountID
}

func NewMsgCashReadyDeposit(depositID string,operator AccountID) MsgCashReadyDeposit {
	return MsgCashReadyDeposit{
		DepositID:depositID,
		Operator:operator,
	}
}

// Route should return the name of the module
func (msg MsgCashReadyDeposit) Route() string { return RouterKey }

func (msg MsgCashReadyDeposit) Type() chainTypes.Name { return chainTypes.MustName("cashreadydeposit") }

func (msg MsgCashReadyDeposit) Sender() AccountID {
	return msg.Operator
}

func (msg MsgCashReadyDeposit) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid

	if len(msg.DepositID) == 0 {
		return ErrEmptyDepositID
	}
	return nil
}