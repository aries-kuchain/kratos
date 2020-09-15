package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


var (
	ErrEmptySingerAccount              = sdkerrors.Register(ModuleName, 2, "empty singer account")
	ErrSingerExists								= sdkerrors.Register(ModuleName, 3, "singer has already exist")
	ErrUnKnowAccount                   = sdkerrors.Register(ModuleName, 4, "singer is not a known account")
	ErrBadAccessAmount                   = sdkerrors.Register(ModuleName, 5, "invalid access amount")
	ErrSingerNotExists								= sdkerrors.Register(ModuleName, 6, "singer does not  exist")
	ErrBadDenom                        = sdkerrors.Register(ModuleName, 7, "invalid coin denomination")

	
)