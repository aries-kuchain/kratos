package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"fmt"
	"github.com/KuChainNetwork/kuchain/x/deposit/external"

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
	depositInfo.DepositChangeTime = ctx.BlockHeader().Time
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
	threshold := k.Threshold(ctx)
	quoteAmount := k.CalQuoteAmount(ctx,asset,Coin{Denom:external.DefaultBondDenom,Amount:sdk.NewInt(1)})
	feeAmount := quoteAmount.MulRaw(k.DepositFeeRate(ctx)).QuoRaw(int64(100*threshold)).MulRaw(int64(threshold))
	_,err = k.pricefeeKeeper.LockFee(ctx,ownerAccount,feeAmount)
	if err != nil {
		return depositId,err
	}
	depositInfo := types.NewDepositInfo(depositId, ownerAccount, asset)
	eachMortgage := quoteAmount.MulRaw(k.MortgageRage(ctx)).QuoRaw(int64(100*threshold))
	pickedSingers,err := k.singerKeeper.PickSinger(ctx,depositId,eachMortgage,threshold)

	if err != nil {
		return depositId,err
	}
	depositInfo.CurrentFee = feeAmount
	depositInfo.TotalFee = feeAmount
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

func (k Keeper) SetDepositBtcAddress(ctx sdk.Context,depositID string,btcAddress []byte)(err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if depositInfo.Status != types.SingerReady {
		return types.ErrStatusNotSingerReady
	}

	depositInfo.DepositAddress = append(depositInfo.DepositAddress,btcAddress...)
	depositInfo.Status = types.AddressReady
	k.SetDepositInfo(ctx,depositInfo)
	return nil
}

func (k Keeper) ActiveDeposit(ctx sdk.Context,depositID string) (err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if depositInfo.Status != types.DepositSpvReady {
		return types.ErrDepositNotExist
	}

	threshold := len(depositInfo.Singers)

	for _,singerAccount := range depositInfo.Singers {
		_,err := k.pricefeeKeeper.TransferFee(ctx,depositInfo.Owner,singerAccount,depositInfo.CurrentFee.QuoRaw(int64(threshold)))
		if err != nil {
			return err
		}
	}
	depositInfo.CurrentFee = sdk.ZeroInt()
	depositInfo.Status = types.Active
	k.SetDepositInfo(ctx,depositInfo)
	return nil
}

func (k Keeper) TransferDeposit(ctx sdk.Context,depositID string,from,to AccountID) (err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if depositInfo.Status != types.Active {
		return types.ErrStatusNotActive
	}

	if !depositInfo.Owner.Eq(from) {
		return types.ErrNotOwnerAccount
	}

	depositInfo.Owner = to
	k.SetDepositInfo(ctx,depositInfo)

	return nil
}

func (k Keeper) DepositToCoin(ctx sdk.Context,depositID string,owner AccountID) (err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if depositInfo.Status != types.Active {
		return types.ErrStatusNotActive
	}

	if !depositInfo.Owner.Eq(owner) {
		return types.ErrNotOwnerAccount
	}

	depositInfo.Status = types.CashReady
	k.SetDepositInfo(ctx,depositInfo)

	legalCoin,found := k.GetLegalCoin(ctx,depositInfo.Asset)
	if !found {
		return types.ErrLegalCoinNotExist
	}

	err = k.bankKeeper.Issue(ctx,types.ModuleAccountName,legalCoin.Symbol,depositInfo.Asset)
	if err != nil {
		return err
	}
	err = k.supplyKeeper.ModuleCoinsToPower(ctx,types.ModuleName,Coins{depositInfo.Asset})
	if err != nil {
		return err
	}
	return k.supplyKeeper.SendCoinsFromModuleToAccount(ctx,types.ModuleName,owner,Coins{depositInfo.Asset})
}

func (k Keeper) ClaimDeposit(ctx sdk.Context,depositID string,owner AccountID,asset Coin,claimAddress []byte) (err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if depositInfo.Status != types.CashReady {
		return types.ErrStatusNotCashReady
	}

	if !depositInfo.Asset.IsEqual(asset) {
		return types.ErrCoinNotEqual
	}
	threshold := len(depositInfo.Singers)
	quoteAmount := k.CalQuoteAmount(ctx,asset,Coin{Denom:external.DefaultBondDenom,Amount:sdk.NewInt(1)})
	feeAmount := quoteAmount.MulRaw(k.ClaimFeeRate(ctx)).QuoRaw(int64(100*threshold)).MulRaw(int64(threshold))
	_,err = k.pricefeeKeeper.LockFee(ctx,owner,feeAmount)
	if err != nil {
		return err
	}

	depositInfo.WithDrawAddress = append(depositInfo.WithDrawAddress,claimAddress...)
	depositInfo.Status = types.Cashing
	depositInfo.Owner = owner
	depositInfo.CurrentFee = feeAmount
	depositInfo.TotalFee = depositInfo.TotalFee.Add(feeAmount)
	k.SetDepositInfo(ctx,depositInfo)
	return k.singerKeeper.SetClaimAddress(ctx,depositID,claimAddress)
}

func (k Keeper) FinishDeposit(ctx sdk.Context,depositID string,owner AccountID)(err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if depositInfo.Status != types.CashOut {
		return types.ErrStatusNotCashOut
	}

	if !depositInfo.Owner.Eq(owner) {
		return types.ErrNotOwnerAccount
	}
	threshold := len(depositInfo.Singers)

	for _,singerAccount := range depositInfo.Singers {
		_,err := k.pricefeeKeeper.TransferFee(ctx,depositInfo.Owner,singerAccount,depositInfo.CurrentFee.QuoRaw(int64(threshold)))
		if err != nil {
			return err
		}
		_,err = k.pricefeeKeeper.UnLockFee(ctx,singerAccount,depositInfo.TotalFee.QuoRaw(3))
		if err != nil {
			return err
		}
	}

	depositInfo.Status = types.Finish
	k.SetDepositInfo(ctx,depositInfo)
	//销毁代币
	err = k.supplyKeeper.BurnCoins(ctx,types.ModuleAccountID,Coins{depositInfo.Asset})
	if err != nil {
		return err
	}

	return k.singerKeeper.FinishDeposit(ctx,depositID)
}

func (k Keeper) CalQuoteAmount(ctx sdk.Context,base,quote Coin) sdk.Int {
	priceInfo,found := k.pricefeeKeeper.GetPriceInfo(ctx,base,quote)
	if !found {
		return sdk.ZeroInt()
	}

	return base.Amount.Mul(priceInfo.Quote.Amount).Quo(priceInfo.Base.Amount)
}

func (k Keeper) WaitTimeOut(ctx sdk.Context,depositID string,owner AccountID) (err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}
	
	if !depositInfo.Owner.Eq(owner) {
		return types.ErrNotOwnerAccount
	}

	if depositInfo.DepositChangeTime.Add(k.WaitTime(ctx)).After(ctx.BlockHeader().Time) {
		return types.ErrNotReachWaitTime
	}

	if depositInfo.Status == types.SingerReady {
		_,err = k.pricefeeKeeper.UnLockFee(ctx,depositInfo.Owner,depositInfo.CurrentFee)
		if err != nil {
			return err
		}

		depositInfo.Status = types.Aberrant
		k.SetDepositInfo(ctx,depositInfo)
		return k.singerKeeper.AberrantDeposit(ctx,depositID)
	}

	if depositInfo.Status == types.DepositSpvReady {
		depositInfo.Status = types.Active
		k.SetDepositInfo(ctx,depositInfo)

		threshold := len(depositInfo.Singers)

		for _,singerAccount := range depositInfo.Singers {
			_,err := k.pricefeeKeeper.TransferFee(ctx,depositInfo.Owner,singerAccount,depositInfo.CurrentFee.QuoRaw(int64(threshold)))
			if err != nil {
				return err
			}
		}
		return k.singerKeeper.ActiveSingerDeposit(ctx,depositID)
	}

	if depositInfo.Status == types.Cashing {
		err = k.supplyKeeper.ModuleCoinsToPower(ctx,types.ModuleName,Coins{depositInfo.Asset})
		if err != nil {
			return err
		}
		err = k.supplyKeeper.SendCoinsFromModuleToAccount(ctx,types.ModuleName,owner,Coins{depositInfo.Asset})
		if err != nil {
			return err
		}
		_,err = k.pricefeeKeeper.UnLockFee(ctx,depositInfo.Owner,depositInfo.CurrentFee)
		if err != nil {
			return err
		}
		depositInfo.Status = types.Aberrant
		k.SetDepositInfo(ctx,depositInfo)
		return k.singerKeeper.AberrantDeposit(ctx,depositID)
	}

	return  types.ErrNotWaitStatus
}

func (k Keeper) AberrantDeposit(ctx sdk.Context,depositID string) (err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if depositInfo.Status != types.AddressReady {
		return types.ErrStatusNotAddressReady
	}

	threshold := len(depositInfo.Singers)

	for _,singerAccount := range depositInfo.Singers {
		_,err := k.pricefeeKeeper.TransferFee(ctx,depositInfo.Owner,singerAccount,depositInfo.CurrentFee.QuoRaw(int64(threshold)))
		if err != nil {
			return err
		}
		_,err = k.pricefeeKeeper.UnLockFee(ctx,singerAccount,depositInfo.TotalFee.QuoRaw(3))
		if err != nil {
			return err
		}
	}

	depositInfo.Status = types.Aberrant
	k.SetDepositInfo(ctx,depositInfo)
	return nil
}

func  (k Keeper) ExternalCloseDeposit(ctx sdk.Context,depositID string) (err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}
		if depositInfo.Status != types.CashOut {
		return types.ErrStatusNotCashOut
	}

	threshold := len(depositInfo.Singers)

	for _,singerAccount := range depositInfo.Singers {
		_,err := k.pricefeeKeeper.TransferFee(ctx,depositInfo.Owner,singerAccount,depositInfo.CurrentFee.QuoRaw(int64(threshold)))
		if err != nil {
			return err
		}
		_,err = k.pricefeeKeeper.UnLockFee(ctx,singerAccount,depositInfo.TotalFee.QuoRaw(3))
		if err != nil {
			return err
		}
	}

	depositInfo.Status = types.Finish
	k.SetDepositInfo(ctx,depositInfo)
	//销毁代币
	return k.supplyKeeper.BurnCoins(ctx,types.ModuleAccountID,Coins{depositInfo.Asset})
}