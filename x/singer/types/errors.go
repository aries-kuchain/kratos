package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEmptySingerAccount  = sdkerrors.Register(ModuleName, 2, "empty singer account")
	ErrSingerExists        = sdkerrors.Register(ModuleName, 3, "singer has already exist")
	ErrUnKnowAccount       = sdkerrors.Register(ModuleName, 4, "singer is not a known account")
	ErrBadAccessAmount     = sdkerrors.Register(ModuleName, 5, "invalid access amount")
	ErrSingerNotExists     = sdkerrors.Register(ModuleName, 6, "singer does not  exist")
	ErrBadDenom            = sdkerrors.Register(ModuleName, 7, "invalid coin denomination")
	ErrEmptySystemAccount  = sdkerrors.Register(ModuleName, 8, "empty system account")
	ErrNotSystemAccount    = sdkerrors.Register(ModuleName, 9, "operate account is not  system account")
	ErrSingerAlreadyActive = sdkerrors.Register(ModuleName, 10, "singer already active")
	ErrInsufficientAccessAsset = sdkerrors.Register(ModuleName, 11, "singer Access Asset is insufficient")
	ErrMortgageNotEnough = sdkerrors.Register(ModuleName, 12, "singer mortgage is not enough")
	ErrAccessIsZero = sdkerrors.Register(ModuleName, 13, "singer access is zero")
	ErrSingerStatusLock = sdkerrors.Register(ModuleName, 14, "singer type is lock")
	ErrDepositAlreadyExist = sdkerrors.Register(ModuleName, 15, "deposit already exist")
	ErrSingerStatusNotActive = sdkerrors.Register(ModuleName, 16, "deposit status is not active")
	ErrNotEnoughSingers = sdkerrors.Register(ModuleName, 17, "Singer is not enough")
	ErrNotDepositSInger = sdkerrors.Register(ModuleName, 18, "Singer is not in the deposit singer group")
	ErrDepositNotExist = sdkerrors.Register(ModuleName, 19, "deposit not exist")
	ErrDepositStatusNotOpen = sdkerrors.Register(ModuleName, 20, "deposit status is not open")
	ErrDepositStatusNotAddressReady = sdkerrors.Register(ModuleName, 21, "deposit status is not AddressReady")
	ErrDepositStatusNotSpvReady = sdkerrors.Register(ModuleName, 22, "deposit status is not SPVReady")
	ErrDepositStatusNotActive = sdkerrors.Register(ModuleName, 23, "deposit status is not Active")
)
