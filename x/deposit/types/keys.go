package types

import (
	// "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
)

const (
	ModuleName      = "deposit"
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
	DepositInfoKey = []byte{0x11} // prefix for each key to a singerInfo index
	legalDepositCoinKey = []byte{0x21}
)

const (
	AccountIDlen = sdk.AddrLen + 1
)

func GetDepositInfoKey(depositID string) []byte {
	return append(DepositInfoKey, []byte(depositID)...)
}

func GetLegalCoinKey(asset Coin) []byte{
	return append(legalDepositCoinKey, []byte(asset.Denom)...)
} 


