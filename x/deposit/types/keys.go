package types

import (
	// "fmt"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	QuerierRoute      = ModuleName
)

var (
	DepositInfoKey      = []byte{0x11} // prefix for each key to a singerInfo index
	LegalDepositCoinKey = []byte{0x21}
	DepositSpvKey       = []byte{0x31}
	DepositGradeKey       = []byte{0x41}
)

const (
	AccountIDlen = sdk.AddrLen + 1
)

func GetDepositInfoKey(depositID string) []byte {
	return append(DepositInfoKey, []byte(depositID)...)
}

func GetLegalCoinKey(asset Coin) []byte {
	return append(LegalDepositCoinKey, []byte(asset.Denom)...)
}

func GetDepositSpvKey(depositID string) []byte {
	return append(DepositSpvKey, []byte(depositID)...)
}

func GetDepositSingerSpvKey(depositID string, singerAccount AccountID) []byte {
	return append(GetDepositSpvKey(depositID), singerAccount.StoreKey()...)
}

func GetDepositGradeKey(asset Coin) []byte {
	json,err := Coins{asset}.MarshalJSON()
	if err != nil {
		return DepositGradeKey
	}
	return append(DepositGradeKey, json...)
}
