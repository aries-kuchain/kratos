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