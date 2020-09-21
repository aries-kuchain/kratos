package types

import (
	//	"fmt"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName      = "singer"
	StoreKey        = ModuleName
	RouterKey       = ModuleName
	keyCommunityTax = ModuleName
	QueryValue      = ModuleName
)

var (
	ModuleAccountName = chainTypes.MustName(ModuleName)
	ModuleAccountID   = chainTypes.NewAccountIDFromName(ModuleAccountName)
)

var (
	SingerInfoKey = []byte{0x11} // prefix for each key to a singerInfo index
	DepositInfoKey = []byte{0x21} 
	DepositBtvAddressKey = []byte{0x31} 

)

const (
	AccountIDlen = sdk.AddrLen + 1
)

func GetSingerInfoKey(singerAccount AccountID) []byte {
	return append(SingerInfoKey, singerAccount.StoreKey()...)
}

func GetDepositInfoKey(depositID string) []byte {
	return append(DepositInfoKey, []byte(depositID)...)
}

func GetDepositBtcAddressDepositKey(depositID string) []byte {
	return append(DepositBtvAddressKey, []byte(depositID)...)
}

func GetDepositBtcAddressKey(depositID string,singerAccount AccountID) []byte {
	return append(GetDepositBtcAddressDepositKey(depositID), singerAccount.StoreKey()...)
}