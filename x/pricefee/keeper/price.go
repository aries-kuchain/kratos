package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/pricefee/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	// "github.com/KuChainNetwork/kuchain/x/pricefee/external"

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

func (k Keeper) SetpriceInfo(ctx sdk.Context, priceInfo types.PriceInfo) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalPriceInfo(k.cdc, priceInfo)
	store.Set(types.GetPriceInfoKey(priceInfo.Base,priceInfo.Quote), b)
}