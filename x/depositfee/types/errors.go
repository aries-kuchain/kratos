package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEmptyOwnerAccount  = sdkerrors.Register(ModuleName, 2, "empty owner account")
)