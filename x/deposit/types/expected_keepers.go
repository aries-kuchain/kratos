// noalias
// DONTCOVER
package types

import (
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/account/exported"
	//"github.com/KuChainNetwork/kuchain/x/slashing/external"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper expected account keeper
type AccountKeeper interface {
	chainTypes.AccountAuther
	GetAccount(sdk.Context, chainTypes.AccountID) exported.Account
	IterateAccounts(ctx sdk.Context, process func(exported.Account) (stop bool))
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	chainTypes.AssetTransfer

	SpendableCoins(ctx sdk.Context, addr chainTypes.AccountID) chainTypes.Coins
}
