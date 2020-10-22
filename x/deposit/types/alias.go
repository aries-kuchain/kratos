package types

import (
	"github.com/KuChainNetwork/kuchain/chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	AccountID  = types.AccountID
	Dec        = sdk.Dec
	Coin       = types.Coin
	Coins      = types.Coins
	Name       = types.Name
)

const (
	AccIDStoreKeyLen = types.AccIDStoreKeyLen
)

var (
	NewAccountIDFromByte = types.NewAccountIDFromByte
	NewCoin              = types.NewCoin
	NewCoins             = types.NewCoins
)
