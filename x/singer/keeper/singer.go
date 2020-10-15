package keeper

import (
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/KuChainNetwork/kuchain/x/singer/external"
	"sort"
)

func (k Keeper) GetSingerInfo(ctx sdk.Context, singerAccount AccountID) (singerInfo types.SingerInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetSingerInfoKey(singerAccount)
	value := store.Get(key)
	if value == nil {
		return singerInfo, false
	}

	singerInfo = types.MustUnmarshalSingerInfo(k.cdc, value)
	return singerInfo, true
}

func (k Keeper) GetAllSingerInfo(ctx sdk.Context) (singerInfos []types.SingerInfo) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SingerInfoKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		singerInfo := types.MustUnmarshalSingerInfo(k.cdc, iterator.Value())
		singerInfos = append(singerInfos, singerInfo)
	}

	return singerInfos
}

func (k Keeper) SetSingerInfo(ctx sdk.Context, singerInfo types.SingerInfo) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalSingerInfo(k.cdc, singerInfo)
	store.Set(types.GetSingerInfoKey(singerInfo.SingerAccount), b)
}

func (k Keeper) NewSingerInfo(ctx sdk.Context, singerAccount AccountID) (err error) {
	newSingerInfo := types.NewSingerInfo(singerAccount)
	k.SetSingerInfo(ctx, newSingerInfo)
	err = k.pricefeeKeeper.NewFeeInfo(ctx,singerAccount)
	return err
}

func (k Keeper) SingerAddAccess(ctx sdk.Context, singerAccount AccountID, amount Coin) (totalAccess sdk.Int, err error) {
	singer, found := k.GetSingerInfo(ctx, singerAccount)

	if !found {
		return sdk.ZeroInt(), types.ErrSingerNotExists
	}

	err = k.supplyKeeper.ModuleCoinsToPower(ctx, types.ModuleName, chainTypes.NewCoins(amount))
	if err != nil {
		return sdk.ZeroInt(), err
	}

	singer.AccessAsset = singer.AccessAsset.Add(amount.Amount)
	k.SetSingerInfo(ctx, singer)
	return singer.AccessAsset, nil
}

func (k Keeper) ActiveSingerInfo(ctx sdk.Context, singerAccount AccountID) error {
	singer, found := k.GetSingerInfo(ctx, singerAccount)

	if !found {
		return types.ErrSingerNotExists
	}

	if singer.Status == types.Active {
		return types.ErrSingerAlreadyActive
	}

	if singer.AccessAsset.LT(k.MinAccessAmount(ctx)) {
		return types.ErrInsufficientAccessAsset
	}
	singer.Status = types.Active
	k.SetSingerInfo(ctx, singer)
	return nil
}

func (k Keeper) SingerAddBTCMortgate(ctx sdk.Context, singerAccount AccountID, amount Coin) (totalAccess sdk.Int, err error) {
	singer, found := k.GetSingerInfo(ctx, singerAccount)

	if !found {
		return sdk.ZeroInt(), types.ErrSingerNotExists
	}

	err = k.supplyKeeper.ModuleCoinsToPower(ctx, types.ModuleName, chainTypes.NewCoins(amount))
	if err != nil {
		return sdk.ZeroInt(), err
	}

	singer.SignatureMortgage = singer.SignatureMortgage.Add(amount.Amount)
	k.SetSingerInfo(ctx, singer)
	return singer.SignatureMortgage, nil
}

func (k Keeper) SingerClaimBTCMortgate(ctx sdk.Context, singerAccount AccountID, amount Coin) (totalAccess sdk.Int, err error) {
	singer, found := k.GetSingerInfo(ctx, singerAccount)

	if !found {
		return sdk.ZeroInt(), types.ErrSingerNotExists
	}

	if singer.SignatureMortgage.LT(amount.Amount) {
		return sdk.ZeroInt(), types.ErrMortgageNotEnough
	}

	//状态判断
	if singer.Status == types.Lock {
		return sdk.ZeroInt(), types.ErrSingerStatusNotActive
	}

	err = k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, singerAccount, chainTypes.NewCoins(amount))
	if err != nil {
		return sdk.ZeroInt(), err
	}

	singer.SignatureMortgage = singer.SignatureMortgage.Sub(amount.Amount)
	k.SetSingerInfo(ctx, singer)
	return singer.SignatureMortgage, nil
}

func (k Keeper) SingerClaimAccess(ctx sdk.Context, singerAccount AccountID) (totalAccess sdk.Int, err error) {
	singer, found := k.GetSingerInfo(ctx, singerAccount)

	if !found {
		return sdk.ZeroInt(), types.ErrSingerNotExists
	}

	if singer.AccessAsset.LTE(k.MinAccessAmount(ctx)) {
		return sdk.ZeroInt(),types.ErrInsufficientAccessAsset
	}

	amount := chainTypes.NewCoin( external.DefaultBondDenom,singer.AccessAsset.Sub(k.MinAccessAmount(ctx)))

	err = k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, singerAccount, chainTypes.NewCoins(amount))
	if err != nil {
		return sdk.ZeroInt(), err
	}

	singer.AccessAsset = singer.AccessAsset.Sub(amount.Amount)
	k.SetSingerInfo(ctx, singer)
	return singer.AccessAsset, nil
}


func (k Keeper) SingerLogoutAccess(ctx sdk.Context, singerAccount AccountID) ( err error) {
	singer, found := k.GetSingerInfo(ctx, singerAccount)

	if !found {
		return  types.ErrSingerNotExists
	}

	if singer.Status == types.Lock {//TODO
		return  types.ErrSingerStatusNotActive
	}

	if singer.AccessAsset.IsZero() {
		return types.ErrAccessIsZero
	}

	if !singer.LockMortgage.IsZero() {
		return types.ErrSingerIsBusy
	}

	amount := chainTypes.NewCoin( external.DefaultBondDenom,singer.AccessAsset)

	err = k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, singerAccount, chainTypes.NewCoins(amount))
	if err != nil {
		return  err
	}

	singer.AccessAsset = sdk.ZeroInt()
	singer.Status = types.InActive
	k.SetSingerInfo(ctx, singer)
	return nil
}

func (k Keeper) lockSinger(ctx sdk.Context, singers []AccountID) (err error) {
	for _, singer := range singers {
		singerInfo,found := k.GetSingerInfo(ctx,singer)
		if !found {
			return types.ErrSingerNotExists
		}
		singerInfo.Status = types.Lock
		k.SetSingerInfo(ctx,singerInfo)
	}
	return nil
}

func (k Keeper) lockMortgage(ctx sdk.Context,singerInfors types.SingerInfos,lockAmount sdk.Int) (err error) {
	for _, singerInfo := range singerInfors {
		if singerInfo.Status != types.Active {
			return types.ErrSingerStatusNotActive
		}
		if  singerInfo.SignatureMortgage.LT(lockAmount) {
			return types.ErrMortgageNotEnough
		}
		singerInfo.SignatureMortgage = singerInfo.SignatureMortgage.Sub(lockAmount)
		singerInfo.LockMortgage = singerInfo.LockMortgage.Add(lockAmount)
		k.SetSingerInfo(ctx,singerInfo)
	}
	return nil
}

func (k Keeper) PickSinger(ctx sdk.Context,depositID string,minStake sdk.Int,threshold int) (pickedSingerInfo  types.SingerInfos,err error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SingerInfoKey)
	defer iterator.Close()
	var singerInfos types.SingerInfos

	for ; iterator.Valid(); iterator.Next() {
		singerInfo := types.MustUnmarshalSingerInfo(k.cdc, iterator.Value())
		if singerInfo.Status == types.Active && singerInfo.SignatureMortgage.GTE(minStake) {
			singerInfos = append(singerInfos, singerInfo)
		}
	}

	sort.Sort(singerInfos)

	if len(singerInfos) < threshold {
		return nil,types.ErrNotEnoughSingers
	}

	for i:=0;i<threshold; i=i+1 {
		pickedSingerInfo = append(pickedSingerInfo,singerInfos[i])
	}

	k.NewDepositInfo(ctx,depositID,threshold,pickedSingerInfo,minStake)
	err = k.lockMortgage(ctx,pickedSingerInfo,minStake)
	if err != nil {
		return nil,err
	}

	return pickedSingerInfo,nil
}

func (k Keeper) unlockSinger(ctx sdk.Context,singers []AccountID) (err error) {
	for _, singer := range singers {
		singerInfo,found := k.GetSingerInfo(ctx,singer)
		if !found {
			return types.ErrSingerNotExists
		}
		if singerInfo.Status != types.Lock {
			return types.ErrSingerStatusNotLock
		}
		singerInfo.Status = types.Active
		k.SetSingerInfo(ctx,singerInfo)
	}
	return nil
}

func (k Keeper) unlockMortgage(ctx sdk.Context,singers []AccountID,unlockAmount sdk.Int) (err error) {
	for _, singer := range singers {
		singerInfo,found := k.GetSingerInfo(ctx,singer)
		if !found {
			return types.ErrSingerNotExists
		}
		if singerInfo.LockMortgage.LT(unlockAmount) {
			return types.ErrMortgageNotEnough
		}
		singerInfo.LockMortgage = singerInfo.LockMortgage.Sub(unlockAmount)
		singerInfo.SignatureMortgage = singerInfo.SignatureMortgage.Add(unlockAmount)
		k.SetSingerInfo(ctx,singerInfo)
	}
	return nil
}

func  (k Keeper) punishSinger(ctx sdk.Context,singers []AccountID,punishAmount sdk.Int) (err error) {
	for _, singer := range singers {
		singerInfo,found := k.GetSingerInfo(ctx,singer)
		if !found {
			return types.ErrSingerNotExists
		}
		if singerInfo.LockMortgage.LT(punishAmount) {
			return types.ErrMortgageNotEnough
		}

		singerInfo.LockMortgage = singerInfo.LockMortgage.Sub(punishAmount)
		k.SetSingerInfo(ctx,singerInfo)
	}
	return nil
}