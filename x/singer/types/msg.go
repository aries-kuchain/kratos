package types

import (
	//	"encoding/json"
	//	sdk "github.com/cosmos/cosmos-sdk/types"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
)

type MsgRegisterSinger struct {
	SingerAccount AccountID `json:"singer_account" yaml:"singer_account"`
}

func NewRegisterSinger(singerAccount AccountID) MsgRegisterSinger {
	return MsgRegisterSinger{SingerAccount: singerAccount}
}

// Route should return the name of the module
func (msg MsgRegisterSinger) Route() string { return RouterKey }

func (msg MsgRegisterSinger) Type() chainTypes.Name { return chainTypes.MustName("registersinger") }

func (msg MsgRegisterSinger) Sender() AccountID {
	return msg.SingerAccount
}

func (msg MsgRegisterSinger) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SingerAccount.Empty() {
		return ErrEmptySingerAccount
	}
	return nil
}
//--------------------------------------------------------------------------------------------------------------
type MsgPayAccess struct {
	SingerAccount AccountID `json:"singer_account" yaml:"singer_account"`
	Amount        Coin      `json:"amount" yaml:"amount"`
}

func NewMsgPayAccess(singerAccount AccountID, amount Coin) MsgPayAccess {
	return MsgPayAccess{SingerAccount: singerAccount, Amount: amount}
}

// Route should return the name of the module
func (msg MsgPayAccess) Route() string { return RouterKey }

func (msg MsgPayAccess) Type() chainTypes.Name { return chainTypes.MustName("payaccess") }

func (msg MsgPayAccess) Sender() AccountID {
	return msg.SingerAccount
}

func (msg MsgPayAccess) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SingerAccount.Empty() {
		return ErrEmptySingerAccount
	}

	if !msg.Amount.Amount.IsPositive() {
		return ErrBadAccessAmount
	}
	return nil
}
//---------------------------------------------------------------------------------------------------------------------------
type MsgActiveSinger struct {
	SystemAccount AccountID `json:"system_account" yaml:"system_account"`
	SingerAccount AccountID `json:"singer_account" yaml:"singer_account"`
}

func NewMsgActiveSinger(systemAccount, singerAccount AccountID) MsgActiveSinger {
	return MsgActiveSinger{SystemAccount: systemAccount, SingerAccount: singerAccount}
}

// Route should return the name of the module
func (msg MsgActiveSinger) Route() string { return RouterKey }

func (msg MsgActiveSinger) Type() chainTypes.Name { return chainTypes.MustName("activesinger") }

func (msg MsgActiveSinger) Sender() AccountID {
	return msg.SystemAccount
}

func (msg MsgActiveSinger) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SystemAccount.Empty() {
		return ErrEmptySystemAccount
	}
	if msg.SingerAccount.Empty() {
		return ErrEmptySingerAccount
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgPayBTCMortgate struct {
	SingerAccount AccountID `json:"singer_account" yaml:"singer_account"`
	Amount        Coin      `json:"amount" yaml:"amount"`
}

func NewMsgPayBTCMortgate(singerAccount AccountID, amount Coin) MsgPayBTCMortgate {
	return MsgPayBTCMortgate{SingerAccount: singerAccount, Amount: amount}
}

// Route should return the name of the module
func (msg MsgPayBTCMortgate) Route() string { return RouterKey }

func (msg MsgPayBTCMortgate) Type() chainTypes.Name { return chainTypes.MustName("paybtcmortgage") }

func (msg MsgPayBTCMortgate) Sender() AccountID {
	return msg.SingerAccount
}

func (msg MsgPayBTCMortgate) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SingerAccount.Empty() {
		return ErrEmptySingerAccount
	}

	if !msg.Amount.Amount.IsPositive() {
		return ErrBadAccessAmount
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgClaimBTCMortgate struct {
	SingerAccount AccountID `json:"singer_account" yaml:"singer_account"`
	Amount        Coin      `json:"amount" yaml:"amount"`
}

func NewMsgClaimBTCMortgate(singerAccount AccountID, amount Coin) MsgClaimBTCMortgate {
	return MsgClaimBTCMortgate{SingerAccount: singerAccount, Amount: amount}
}

// Route should return the name of the module
func (msg MsgClaimBTCMortgate) Route() string { return RouterKey }

func (msg MsgClaimBTCMortgate) Type() chainTypes.Name { return chainTypes.MustName("claimbtcmortgage") }

func (msg MsgClaimBTCMortgate) Sender() AccountID {
	return msg.SingerAccount
}

func (msg MsgClaimBTCMortgate) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SingerAccount.Empty() {
		return ErrEmptySingerAccount
	}

	if !msg.Amount.Amount.IsPositive() {
		return ErrBadAccessAmount
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------

type MsgClaimAccess struct {
	SingerAccount AccountID `json:"singer_account" yaml:"singer_account"`
}

func NewMsgClaimAccess(singerAccount AccountID) MsgClaimAccess {
	return MsgClaimAccess{SingerAccount: singerAccount}
}

// Route should return the name of the module
func (msg MsgClaimAccess) Route() string { return RouterKey }

func (msg MsgClaimAccess) Type() chainTypes.Name { return chainTypes.MustName("claimaccess") }

func (msg MsgClaimAccess) Sender() AccountID {
	return msg.SingerAccount
}

func (msg MsgClaimAccess) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SingerAccount.Empty() {
		return ErrEmptySingerAccount
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgLogoutSinger struct {
	SingerAccount AccountID `json:"singer_account" yaml:"singer_account"`
}

func NewMsgLogoutSinger(singerAccount AccountID) MsgLogoutSinger {
	return MsgLogoutSinger{SingerAccount: singerAccount}
}

// Route should return the name of the module
func (msg MsgLogoutSinger) Route() string { return RouterKey }

func (msg MsgLogoutSinger) Type() chainTypes.Name { return chainTypes.MustName("logoutsinger") }

func (msg MsgLogoutSinger) Sender() AccountID {
	return msg.SingerAccount
}

func (msg MsgLogoutSinger) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SingerAccount.Empty() {
		return ErrEmptySingerAccount
	}
	return nil
}

//----------------------------------------------------------------------------------------------------------------------------
type MsgSetBtcAddress struct {
	SingerAccount AccountID `json:"singer_account" yaml:"singer_account"`
	DepoistID string `json:"deposit_id" yaml:"deposit_id"`
	BtcAddress []byte `json:"btc_address" yaml:"btc_address"`
}

func NewMsgSetBtcAddress(singerAccount AccountID,depositID string,btcAddress []byte) MsgSetBtcAddress {
	return MsgSetBtcAddress{
		SingerAccount: singerAccount,
		DepoistID:depositID,
		BtcAddress:btcAddress,
	}
}

// Route should return the name of the module
func (msg MsgSetBtcAddress) Route() string { return RouterKey }

func (msg MsgSetBtcAddress) Type() chainTypes.Name { return chainTypes.MustName("setbtcaddress") }

func (msg MsgSetBtcAddress) Sender() AccountID {
	return msg.SingerAccount
}

func (msg MsgSetBtcAddress) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.SingerAccount.Empty() {
		return ErrEmptySingerAccount
	}
	return nil
}