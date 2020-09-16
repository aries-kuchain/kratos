package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/depositfee/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

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