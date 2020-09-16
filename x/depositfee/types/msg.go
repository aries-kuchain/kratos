package types

import (
	"encoding/json"
	//"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgSetStore{}
)

// MsgSetName defines a SetName message
type MsgSetStore struct {
	Name  string
	Value string
	Owner sdk.AccAddress
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetStore(name string, value string, owner sdk.AccAddress) MsgSetStore {
	return MsgSetStore{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgSetStore) Route() string { return "easystore" }

// Type should return the action
func (msg MsgSetStore) Type() string { return "set_store" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetStore) ValidateBasic() error {
	if msg.Owner.Empty() {
		// return fmt.Errorf(
		// 	"Owner is empty",
		// )
		return nil //sdkerrors.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		// return fmt.Errorf(
		// 	"name or value is empty",
		// )
		return nil //sdkerrors.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetStore) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgSetStore) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
