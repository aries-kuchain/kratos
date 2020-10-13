package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/pricefee/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/KuChainNetwork/kuchain/chain/constants"
)

func (k Keeper) GetPriceInfo(ctx sdk.Context, base,quote Coin) (priceInfo types.PriceInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPriceInfoKey(base,quote)
	value := store.Get(key)
	if value == nil {
		return priceInfo, false
	}

	priceInfo = types.MustUnmarshalPriceInfo(k.cdc, value)
	return priceInfo, true
}

func (k Keeper) SetPriceInfo(ctx sdk.Context, priceInfo types.PriceInfo) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalPriceInfo(k.cdc, priceInfo)
	store.Set(types.GetPriceInfoKey(priceInfo.Base,priceInfo.Quote), b)
}

func (k Keeper) SetPrice(ctx sdk.Context, msgSetPrice types.MsgSetPrice) (err error){
	name, ok := msgSetPrice.SystemAccount.ToName()
	if !ok {
		return types.ErrSystemNotAddress
	}
	if !constants.IsSystemAccount(name) {
		return types.ErrNotSystemAccount
	}
	priceInfo := types.NewPriceInfo(msgSetPrice.Base,msgSetPrice.Quote,msgSetPrice.Remark)
	k.SetPriceInfo(ctx,priceInfo)
	return nil
}