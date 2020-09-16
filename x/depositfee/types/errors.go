package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEmptyOwnerAccount  = sdkerrors.Register(ModuleName, 2, "empty owner account")
	ErrFeeInfoExist  = sdkerrors.Register(ModuleName, 3, "fee info already exist")
	ErrUnKnowAccount       = sdkerrors.Register(ModuleName, 4, "singer is not a known account")
	ErrBadAmount        = sdkerrors.Register(ModuleName, 5, "amount is inValid")
	ErrFeeInfoNotExist  = sdkerrors.Register(ModuleName, 6, "fee does not exist")

)