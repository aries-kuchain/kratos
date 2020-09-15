package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


var (
	ErrEmptySingerAccount              = sdkerrors.Register(ModuleName, 2, "empty singer account")
	ErrSingerExists								= sdkerrors.Register(ModuleName, 3, "singer has already exist")
	ErrUnKnowAccount                   = sdkerrors.Register(ModuleName, 4, "singer is not a known account")
)