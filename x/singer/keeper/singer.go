package keeper

import (
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/KuChainNetwork/kuchain/x/singer/external"

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

func (k Keeper) SetSingerInfo(ctx sdk.Context, singerInfo types.SingerInfo) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalSingerInfo(k.cdc, singerInfo)
	store.Set(types.GetSingerInfoKey(singerInfo.SingerAccount), b)
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

	if singer.Status != types.InActive {
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
		return sdk.ZeroInt(), types.ErrMortgageNotEnough
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

	if singer.AccessAsset.LT(k.MinAccessAmount(ctx)) {
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

	if singer.Status == types.Lock {
		return  types.ErrMortgageNotEnough
	}

	if singer.AccessAsset.IsZero() {
		return types.ErrAccessIsZero
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