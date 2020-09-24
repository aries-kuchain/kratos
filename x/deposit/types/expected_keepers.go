// noalias
// DONTCOVER
package types

import (
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"
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
	Create(ctx sdk.Context, creator, symbol chainTypes.Name, maxSupply chainTypes.Coin, canIssue, canLock, canBurn bool, issue2Height int64, initSupply chainTypes.Coin, desc []byte) error
	Issue(ctx sdk.Context, creator, symbol chainTypes.Name, amount chainTypes.Coin) error
	SpendableCoins(ctx sdk.Context, addr chainTypes.AccountID) chainTypes.Coins
}

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
	LockFee(ctx sdk.Context, owner AccountID,amount sdk.Int) (totalPreStoreFee sdk.Int,err error)
	TransferFee(ctx sdk.Context, from,to AccountID,amount sdk.Int) (totalPreStoreFee sdk.Int,err error)
	UnLockFee(ctx sdk.Context, owner AccountID,amount sdk.Int) (totalPreStoreFee sdk.Int,err error)
}

type SingerKeeper interface {
	PickSinger(ctx sdk.Context,depositID string,minStake sdk.Int,threshold int) (pickedSingerInfo  singerTypes.SingerInfos,err error)
	SetSpvReady(ctx sdk.Context, depositID string) (err error)
	SetClaimAddress(ctx sdk.Context, depositID string,claimAddress []byte) (err error) 
	FinishDeposit(ctx sdk.Context, depositID string) (err error) 
}

