package keeper

import (
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
)

func (k Keeper) GetSpvInfo(ctx sdk.Context, depositID string,singerAccount AccountID) (spvInfo singerTypes.SpvInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDepositSingerSpvKey(depositID,singerAccount)
	value := store.Get(key)
	if value == nil {
		return spvInfo, false
	}

	spvInfo = singerTypes.MustUnmarshalSpvInfo(k.cdc, value)
	return spvInfo, true
}

func (k Keeper) SetSpvInfo(ctx sdk.Context, spvInfo singerTypes.SpvInfo) {
	store := ctx.KVStore(k.storeKey)
	b := singerTypes.MustMarshalSpvInfo(k.cdc, spvInfo)
	store.Set(types.GetDepositSingerSpvKey(spvInfo.DepositID,spvInfo.SpvSubmiter), b)
}

func (k Keeper) NewSpvInfo(ctx sdk.Context, spvInfo singerTypes.SpvInfo) (err error ) {
	depositInfo,found := k.GetDepositInfo(ctx,spvInfo.DepositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if depositInfo.Status != types.AddressReady {
		return types.ErrStatusNotAddressReady
	}

	if !depositInfo.Owner.Eq(spvInfo.SpvSubmiter) {
		return types.ErrSpvSubmiterNotOwner
	}

	k.SetSpvInfo(ctx,spvInfo)
	depositInfo.Status = types.DepositSpvReady
	k.SetDepositInfo(ctx,depositInfo)
	return k.singerKeeper.SetSpvReady(ctx,spvInfo.DepositID)
}