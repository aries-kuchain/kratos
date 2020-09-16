package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/depositfee/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"

)

func (k Keeper) GetFeeInfo(ctx sdk.Context, owner AccountID) (feeInfo types.FeeInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeInfoKey(owner)
	value := store.Get(key)
	if value == nil {
		return feeInfo, false
	}

	feeInfo = types.MustUnmarshalFeeInfo(k.cdc, value)
	return feeInfo, true
}

func (k Keeper) SetFeeInfo(ctx sdk.Context, feeInfo types.FeeInfo) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalFeeInfo(k.cdc, feeInfo)
	store.Set(types.GetFeeInfoKey(feeInfo.Owner), b)
}

func (k Keeper) NewFeeInfo(ctx sdk.Context, owner AccountID) (err error) {
	if k.ValidatorAccount(ctx,owner) == false {
		return types.ErrUnKnowAccount
	}

	_,found := k.GetFeeInfo(ctx,owner)
	if found {
		return nil
	}

	feeInfo := types.NewFeeInfo(owner)
	k.SetFeeInfo(ctx,feeInfo)
	return nil
}

func (k Keeper) PreStoreFee(ctx sdk.Context, owner AccountID,amount Coin) (totalPreStoreFee sdk.Int,err error){
	feeInfo,found := k.GetFeeInfo(ctx,owner)
	if !found {
		return sdk.ZeroInt(),types.ErrFeeInfoNotExist
	}

	err = k.supplyKeeper.ModuleCoinsToPower(ctx, types.ModuleName, chainTypes.NewCoins(amount))
	if err != nil {
		return sdk.ZeroInt(), err
	}

	feeInfo.PrestoreFee = feeInfo.PrestoreFee.Add(amount.Amount)
	k.SetFeeInfo(ctx,feeInfo)
	return feeInfo.PrestoreFee,nil
}