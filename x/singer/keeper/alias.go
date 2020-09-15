package keeper

import (
	"github.com/KuChainNetwork/kuchain/chain/types"
)

type (
	AccountID = types.AccountID
	KuMsg     = types.KuMsg
	Name      = types.Name
	Coins     = types.Coins
	Coin     = types.Coin
)

var (
	MustName                = types.MustName
	NewCoins                = types.NewCoins
	NewCoin                 = types.NewCoin
	NewAccountIDFromAccAdd  = types.NewAccountIDFromAccAdd
	NewAccountIDFromConsAdd = types.NewAccountIDFromConsAdd
	NewAccountIDFromByte    = types.NewAccountIDFromByte
	NewInt                  = types.NewInt
)
