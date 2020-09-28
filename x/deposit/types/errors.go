package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEmptyOwnerAccount     = sdkerrors.Register(ModuleName, 2, "empty owner account")
	ErrBadAmount             = sdkerrors.Register(ModuleName, 3, "invalid amount")
	ErrDepositAlreadyExist   = sdkerrors.Register(ModuleName, 4, "Deposit Already exist")
	ErrLegalCoinAlreadyExist = sdkerrors.Register(ModuleName, 5, "Legal Coin Already exist")
	ErrNotSystemAccount      = sdkerrors.Register(ModuleName, 6, "operate account is not  system account")
	ErrLegalCoinNotExist     = sdkerrors.Register(ModuleName, 7, "Legal Coin does not exist")
	ErrAssetSymbolError      = sdkerrors.Register(ModuleName, 8, "coin denom should be equal")
	ErrDepositNotExist   = sdkerrors.Register(ModuleName, 9, "Deposit does not exist")
	ErrStatusNotSingerReady   = sdkerrors.Register(ModuleName, 10, "Deposit status is not SingerReady")
	ErrStatusNotAddressReady   = sdkerrors.Register(ModuleName, 11, "Deposit status is not AddressReady")
	ErrSpvSubmiterNotOwner   = sdkerrors.Register(ModuleName, 12, "spv submiter is not the deposit owner")
	ErrNotOwnerAccount   = sdkerrors.Register(ModuleName, 13, "operator account is not deposit owner account")
	ErrStatusNotActive   = sdkerrors.Register(ModuleName, 14, "Deposit status is not Active")
	ErrCoinNotEqual   = sdkerrors.Register(ModuleName, 15, "Claim coin does not equal deposit coin")
	ErrStatusNotCashReady   = sdkerrors.Register(ModuleName, 16, "Deposit status is not CashReady")
	ErrStatusNotCashing   = sdkerrors.Register(ModuleName, 17, "Deposit status is not Cashing")
	ErrStatusNotCashOut   = sdkerrors.Register(ModuleName, 18, "Deposit status is not CashOut")
	ErrPriceNotFound   = sdkerrors.Register(ModuleName, 19, "Deposit status is not CashOut")
	ErrNotReachWaitTime  = sdkerrors.Register(ModuleName, 20, "Do not Reach Wait Time")
	ErrNotWaitStatus  = sdkerrors.Register(ModuleName, 21, "Deposit Status is not a wait for other operator status")
	ErrNotJudgeStatus  = sdkerrors.Register(ModuleName, 22, "Deposit Status is not a judge status")

)
