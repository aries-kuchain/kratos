package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEmptyOwnerAccount  = sdkerrors.Register(ModuleName, 2, "empty owner account")
	ErrBadAmount  = sdkerrors.Register(ModuleName, 3, "invalid amount")
	ErrDepositAlreadyExist  = sdkerrors.Register(ModuleName, 4, "Deposit Already exist")
	ErrLegalCoinAlreadyExist  = sdkerrors.Register(ModuleName, 5, "Legal Coin Already exist")
	ErrNotSystemAccount    = sdkerrors.Register(ModuleName, 6, "operate account is not  system account")
	ErrLegalCoinNotExist  = sdkerrors.Register(ModuleName, 7, "Legal Coin does not exist")
	ErrAssetSymbolError =  sdkerrors.Register(ModuleName, 8, "coin denom should be equal")
)