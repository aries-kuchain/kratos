package types

import (
	// "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
)

const (
	ModuleName      = "pricefee"
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
	FeeInfoKey = []byte{0x11} // prefix for each key to a singerInfo index
	PriceInfoKey = []byte{0x21} // prefix for each key to a singerInfo index

)

const (
	AccountIDlen = sdk.AddrLen + 1
)

func GetFeeInfoKey(owner AccountID) []byte {
	return append(FeeInfoKey, owner.StoreKey()...)
}

func GetPriceInfoKey(base Coin,quote Coin) []byte {
	tmp := append(FeeInfoKey, []byte(base.Denom)...)
	return append(tmp,[]byte(quote.Denom)...)
}