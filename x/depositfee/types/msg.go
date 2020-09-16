package types

import (

	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"

)
//----------------------------------------------------------------------------------------------------------------------------
type MsgOpenFee struct {
	Owner AccountID `json:"owner" yaml:"owner"`
}

func NewMsgOpenFee(owner AccountID) MsgOpenFee {
	return MsgOpenFee{Owner: owner}
}

// Route should return the name of the module
func (msg MsgOpenFee) Route() string { return RouterKey }

func (msg MsgOpenFee) Type() chainTypes.Name { return chainTypes.MustName("openfee") }

func (msg MsgOpenFee) Sender() AccountID {
	return msg.Owner
}

func (msg MsgOpenFee) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.Owner.Empty() {
		return ErrEmptyOwnerAccount
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgPrestoreFee struct {
	Owner AccountID `json:"owner" yaml:"owner"`
	Amount Coin`json:"amount" yaml:"amount"`
}

func NewMsgPrestoreFee(owner AccountID,amount Coin) MsgPrestoreFee {
	return MsgPrestoreFee{Owner: owner,Amount:amount}
}

// Route should return the name of the module
func (msg MsgPrestoreFee) Route() string { return RouterKey }

func (msg MsgPrestoreFee) Type() chainTypes.Name { return chainTypes.MustName("prestorefee") }

func (msg MsgPrestoreFee) Sender() AccountID {
	return msg.Owner
}

func (msg MsgPrestoreFee) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.Owner.Empty() {
		return ErrEmptyOwnerAccount
	}

	if !msg.Amount.Amount.IsPositive() {
		return ErrBadAmount
	}
	return nil
}
//----------------------------------------------------------------------------------------------------------------------------
type MsgClaimFee struct {
	Owner AccountID `json:"owner" yaml:"owner"`
	Amount Coin`json:"amount" yaml:"amount"`
}

func NewMsgClaimFee(owner AccountID,amount Coin) MsgClaimFee {
	return MsgClaimFee{Owner: owner,Amount:amount}
}

// Route should return the name of the module
func (msg MsgClaimFee) Route() string { return RouterKey }

func (msg MsgClaimFee) Type() chainTypes.Name { return chainTypes.MustName("claimfee") }

func (msg MsgClaimFee) Sender() AccountID {
	return msg.Owner
}

func (msg MsgClaimFee) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures either empty or valid
	if msg.Owner.Empty() {
		return ErrEmptyOwnerAccount
	}

	if !msg.Amount.Amount.IsPositive() {
		return ErrBadAmount
	}
	return nil
}