package types

import (
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"

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

	return nil
}