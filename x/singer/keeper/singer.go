package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"

)

func (k Keeper) GetSingerInfo(ctx sdk.Context, singerAccount AccountID) (singerInfo types.SingerInfo,found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetSingerInfoKey(singerAccount)
	value := store.Get(key)
	if value == nil {
		return singerInfo, false
	}

	singerInfo = types.MustUnmarshalSingerInfo(k.cdc, value)
	return singerInfo, true
}

func (k Keeper) SetSingerInfo(ctx sdk.Context, singerInfo types.SingerInfo) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalSingerInfo(k.cdc, singerInfo)
	store.Set(types.GetSingerInfoKey(singerInfo.SingerAccount), b)
}

func (k Keeper) SingerAddAccess(ctx sdk.Context,singerAccount AccountID,amount Coin) (totalAccess sdk.Int,err error) {
	singer,found := k.GetSingerInfo(ctx,singerAccount)

	if !found {
		return sdk.ZeroInt(),types.ErrSingerNotExists
	}

		//将coin转换成coinpower
	err = k.supplyKeeper.ModuleCoinsToPower(ctx, types.ModuleName, chainTypes.NewCoins(amount))
	if err != nil {
		return sdk.ZeroInt(), err
	}

	singer.AccessAsset = singer.AccessAsset.Add(amount.Amount)
	k.SetSingerInfo(ctx,singer)
	return singer.AccessAsset,nil
}
