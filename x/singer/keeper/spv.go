package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetSpvInfo(ctx sdk.Context, depositID string,singerAccount AccountID) (spvInfo types.SpvInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDepositSingerSpvKey(depositID,singerAccount)
	value := store.Get(key)
	if value == nil {
		return spvInfo, false
	}

	spvInfo = types.MustUnmarshalSpvInfo(k.cdc, value)
	return spvInfo, true
}

func (k Keeper) SetSpvInfo(ctx sdk.Context, spvInfo types.SpvInfo) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalSpvInfo(k.cdc, spvInfo)
	store.Set(types.GetDepositSingerSpvKey(spvInfo.DepositID,spvInfo.SpvSubmiter), b)
}

func (k Keeper) GetDepositActiveInfo(ctx sdk.Context, depositID string,singerAccount AccountID) (depositActiveInfo types.DepositActiveInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDepositSingerActiveKey(depositID,singerAccount)
	value := store.Get(key)
	if value == nil {
		return depositActiveInfo, false
	}

	depositActiveInfo = types.MustUnmarshalDepositActiveInfo(k.cdc, value)
	return depositActiveInfo, true
}

func (k Keeper) SetDepositActiveInfo(ctx sdk.Context, depositActiveInfo types.DepositActiveInfo) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalDepositActiveInfo(k.cdc, depositActiveInfo)
	store.Set(types.GetDepositSingerActiveKey(depositActiveInfo.DepositID,depositActiveInfo.Singer), b)
}

func  (k Keeper) NewClaimSpvInfo(ctx sdk.Context, spvInfo types.SpvInfo) (err error) {
	depositInfo,found := k.GetDepositInfo(ctx,spvInfo.DepositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if  depositInfo.Status != types.Cashing {
		if depositInfo.Status != types.CashOut {
			return types.ErrDepositStatusNotSpvReady
		}
	}

	if !depositInfo.CheckSinger(spvInfo.SpvSubmiter) {
		return types.ErrNotDepositSInger
	}
	
	k.SetSpvInfo(ctx,spvInfo)
	depositInfo.Status = types.CashOut
	k.SetDepositInfo(ctx,depositInfo)

	return k.depositKeeper.SetCashOut(ctx,spvInfo.DepositID)
}