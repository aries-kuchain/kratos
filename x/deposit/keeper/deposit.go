package keeper

import (
	// chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/KuChainNetwork/kuchain/x/singer/external"
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

func (k Keeper) NewDepositInfo(ctx sdk.Context,ownerAccount AccountID,asset Coin) (err error){
	depositId := fmt.Sprintf("%s-%s-%s",ownerAccount.String(),asset.String(),ctx.BlockHeader().Time.Format("2006-01-02 15:04:05") )

	_,found := k.GetDepositInfo(ctx,depositId)
	if  found {
		return types.ErrDepositAlreadyExist
	}
//判断代币，收取费用
	depositInfo := types.NewDepositInfo(depositId,ownerAccount,asset)
	k.SetDepositInfo(ctx,depositInfo)
	return nil
}