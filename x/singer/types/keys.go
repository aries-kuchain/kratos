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
	keyCommunityTax = "just for test"
	QueryValue      = ModuleName
)

var (
	ModuleAccountName = chainTypes.MustName(ModuleName)
	ModuleAccountID   = chainTypes.NewAccountIDFromName(ModuleAccountName)
)

var (
	SingerInfoKey = []byte{0x11} // prefix for each key to a singerInfo index
)

const (
	AccountIDlen = sdk.AddrLen + 1
)

func GetSingerInfoKey(singerAccount AccountID) []byte {
	return append(SingerInfoKey, singerAccount.StoreKey()...)
}
