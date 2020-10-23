package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetSpvInfo(ctx sdk.Context, depositID string, ownerAccount AccountID) (spvInfo singerTypes.SpvInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDepositSingerSpvKey(depositID, ownerAccount)
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
	store.Set(types.GetDepositSingerSpvKey(spvInfo.DepositID, spvInfo.SpvSubmiter), b)
}

func (k Keeper) NewSpvInfo(ctx sdk.Context, spvInfo singerTypes.SpvInfo) (err error) {
	depositInfo, found := k.GetDepositInfo(ctx, spvInfo.DepositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if depositInfo.Status != types.AddressReady {
		return types.ErrStatusNotAddressReady
	}

	if !depositInfo.Owner.Eq(spvInfo.SpvSubmiter) {
		return types.ErrSpvSubmiterNotOwner
	}

	k.SetSpvInfo(ctx, spvInfo)
	depositInfo.Status = types.DepositSpvReady
	k.SetDepositInfo(ctx, depositInfo)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSunmitSpv,
			sdk.NewAttribute(types.AttributeKeyDepositID,depositInfo.DepositID),
			sdk.NewAttribute(types.AttributeKeyOwner, depositInfo.Owner.String()),
		),
	})
	return k.singerKeeper.SetSpvReady(ctx, spvInfo.DepositID)
}

func (k Keeper) GetDepositSpv(ctx sdk.Context, depositID string)(spvInfo singerTypes.SpvInfo,found bool) {

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetDepositSpvKey(depositID))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		red := singerTypes.MustUnmarshalSpvInfo(k.cdc, iterator.Value())
		return red,true
	}
	return spvInfo,false
}

func (k Keeper) SetCashOut(ctx sdk.Context, depositID string) (err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if depositInfo.Status == types.CashOut {
		return nil
	}

	if depositInfo.Status != types.Cashing {
		return types.ErrStatusNotCashing
	}

	depositInfo.Status = types.CashOut
	k.SetDepositInfo(ctx, depositInfo)

	return nil
}
