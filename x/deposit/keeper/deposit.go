package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"fmt"
)

func (k Keeper) GetDepositInfo(ctx sdk.Context, depositID string) (depositInfo types.DepositInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDepositInfoKey(depositID)
	value := store.Get(key)
	if value == nil {
		return depositInfo, false
	}

	depositInfo = types.MustUnmarshalDepositInfo(k.cdc, value)
	return depositInfo, true
}

func (k Keeper) SetDepositInfo(ctx sdk.Context, depositInfo types.DepositInfo) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalDepositInfo(k.cdc, depositInfo)
	store.Set(types.GetDepositInfoKey(depositInfo.DepositID), b)
}

func (k Keeper) NewDepositInfo(ctx sdk.Context, ownerAccount AccountID, asset Coin) (depositID string, err error) {
	depositId := fmt.Sprintf("%s-%s-%s", ownerAccount.String(), asset.String(), ctx.BlockHeader().Time.Format("2006-01-02-15:04:05"))

	_, found := k.GetDepositInfo(ctx, depositId)
	if found {
		return depositId, types.ErrDepositAlreadyExist
	}

	legalCoin, found := k.GetLegalCoin(ctx, asset)
	if !found {
		return depositId, types.ErrLegalCoinNotExist
	}

	if legalCoin.Status != types.Permint {
		return depositId, types.ErrLegalCoinNotExist
	}

	//扣除费用   计算费用稍后
	_,err = k.pricefeeKeeper.LockFee(ctx,ownerAccount,asset.Amount)
	if err != nil {
		return depositId,err
	}
	//最低抵押以及签署人数
	depositInfo := types.NewDepositInfo(depositId, ownerAccount, asset)
	pickedSingers,err := k.singerKeeper.PickSinger(ctx,depositId,asset.Amount,3)

	if err != nil {
		return depositId,err
	}

	depositInfo.SetSingers(pickedSingers)
	k.SetDepositInfo(ctx, depositInfo)
	return depositId, nil
}

func (k Keeper) GetAllDepositInfo(ctx sdk.Context) (depositInfos []types.DepositInfo) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DepositInfoKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		depositInfo := types.MustUnmarshalDepositInfo(k.cdc, iterator.Value())
		depositInfos = append(depositInfos, depositInfo)
	}

	return depositInfos
}
