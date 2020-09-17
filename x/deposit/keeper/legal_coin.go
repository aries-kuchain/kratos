package keeper

import (
	// chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/KuChainNetwork/kuchain/chain/constants"
	// "github.com/KuChainNetwork/kuchain/x/singer/external"
//	"fmt"
)

func (k Keeper) GetLegalCoin(ctx sdk.Context, asset Coin) (legalCoin types.LegalCoin, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetLegalCoinKey(asset)
	value := store.Get(key)
	if value == nil {
		return legalCoin, false
	}

	legalCoin = types.MustUnmarshalLegalCoin(k.cdc, value)
	return legalCoin, true
}

func (k Keeper) SetLegalCoin(ctx sdk.Context,legalCoin types.LegalCoin) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalLegalCoin(k.cdc, legalCoin)
	store.Set(types.GetLegalCoinKey(legalCoin.Asset), b)
}

func (k Keeper) CreateLegalCoin(ctx sdk.Context,systemAccount AccountID,asset Coin) (err error) {
	if _,found := k.GetLegalCoin(ctx,asset);found {
		return types.ErrLegalCoinAlreadyExist
	}
	name, ok := systemAccount.ToName()
	if ok && constants.IsSystemAccount(name) {
		return types.ErrNotSystemAccount
	}
	legalCoin := types.NewLegalCoin(asset)
	k.SetLegalCoin(ctx,legalCoin)
	return nil
}

func (k Keeper) ProhibitLegalCoin(ctx sdk.Context,systemAccount AccountID,asset Coin) (err error) {
	legalCoin,found := k.GetLegalCoin(ctx,asset)
	if !found {
		return types.ErrLegalCoinNotExist
	}

	name, ok := systemAccount.ToName()
	if ok && constants.IsSystemAccount(name) {
		return types.ErrNotSystemAccount
	}
	legalCoin.Status = types.Prohibit
	k.SetLegalCoin(ctx,legalCoin)
	return nil
}

func (k Keeper) PermintLegalCoin(ctx sdk.Context,systemAccount AccountID,asset Coin) (err error) {
	legalCoin,found := k.GetLegalCoin(ctx,asset)
	if !found {
		return types.ErrLegalCoinNotExist
	}

	name, ok := systemAccount.ToName()
	if ok && constants.IsSystemAccount(name) {
		return types.ErrNotSystemAccount
	}
	legalCoin.Status = types.Permint
	k.SetLegalCoin(ctx,legalCoin)
	return nil
}
