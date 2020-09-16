package types

import (

	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"

)
//----------------------------------------------------------------------------------------------------------------------------
type MsgOpenFee struct {
	Owner AccountID `json:"singer_account" yaml:"singer_account"`
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