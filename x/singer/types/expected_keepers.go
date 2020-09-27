// noalias
// DONTCOVER
package types

import (
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/account/exported"
	supplyexported "github.com/KuChainNetwork/kuchain/x/supply/exported"
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

// SupplyKeeper defines the expected supply Keeper (noalias)
type SupplyKeeper interface {
	GetSupply(ctx sdk.Context) supplyexported.SupplyI

	InitModuleAccount(ctx sdk.Context, moduleName string) error
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) supplyexported.ModuleAccountI

	// TODO remove with genesis 2-phases refactor https://github.com/cosmos/cosmos-sdk/issues/2862
	SetModuleAccount(sdk.Context, supplyexported.ModuleAccountI)
	ModuleCoinsToPower(ctx sdk.Context, recipientModule string, amt Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr AccountID, amt Coins) error

	BurnCoins(ctx sdk.Context, name chainTypes.AccountID, amt Coins) error
}

type PriceFeeKeeper interface {
	NewFeeInfo(ctx sdk.Context, owner AccountID) (err error)
}

type DepositKeeper interface {
	SetDepositBtcAddress(ctx sdk.Context,depositID string,btcAddress []byte)(err error)
	ActiveDeposit(ctx sdk.Context,depositID string) (err error)
	SetCashOut(ctx sdk.Context, depositID string)  (err error )
	AberrantDeposit(ctx sdk.Context,depositID string) (err error)
	ExternalCloseDeposit(ctx sdk.Context,depositID string) (err error)
	SetWrongDepositSpv(ctx sdk.Context,depositID string) (err error) 
}