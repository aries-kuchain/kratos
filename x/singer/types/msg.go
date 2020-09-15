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
	return MsgRegisterSinger{SingerAccount:singerAccount}
}

// Route should return the name of the module
func (msg MsgRegisterSinger) Route() string { return RouterKey }

// Type should return the action
//func (msg MsgRegisterSinger) Type() string { return "register_singer" }

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

type MsgPayAccess struct {
	SingerAccount AccountID `json:"singer_account" yaml:"singer_account"`
	Amount           Coin      `json:"amount" yaml:"amount"`
}

func NewMsgPayAccess(singerAccount AccountID,amount Coin) MsgPayAccess {
	return MsgPayAccess{SingerAccount:singerAccount,Amount:amount}
}


// Route should return the name of the module
func (msg MsgPayAccess) Route() string { return RouterKey }

// Type should return the action
//func (msg MsgRegisterSinger) Type() string { return "register_singer" }

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