package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/singer/external"
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

func (k Keeper) NewDepositInfo(ctx sdk.Context, depositID string,threshold int,singer types.SingerInfos,minStake sdk.Int) (err error) {
	_,found := k.GetDepositInfo(ctx,depositID)
	if found {
		return types.ErrDepositAlreadyExist
	}
	depositInfo := types.NewDepositInfo(depositID,threshold,minStake)
	depositInfo.SetSingers(singer)
	k.SetDepositInfo(ctx,depositInfo)
	return nil
}

func (k Keeper) GetDepositBtcAddress(ctx sdk.Context, depositID string,singer AccountID) (depositBtcAddress types.DepositBtcAddress, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDepositBtcAddressKey(depositID,singer)
	value := store.Get(key)
	if value == nil {
		return depositBtcAddress, false
	}

	depositBtcAddress = types.MustUnmarshalDepositBtcAddress(k.cdc, value)
	return depositBtcAddress, true
}

func (k Keeper) SetDepositBtcAddress(ctx sdk.Context, depositBtcAddress types.DepositBtcAddress) {
	store := ctx.KVStore(k.storeKey)
	b := types.MustMarshalDepositBtcAddress(k.cdc, depositBtcAddress)
	store.Set(types.GetDepositBtcAddressKey(depositBtcAddress.DepositID,depositBtcAddress.Singer), b)
}


func  (k Keeper) NewDepositBtcAddress(ctx sdk.Context, depositID string,singer AccountID,btcAddress string) (err error){
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if !depositInfo.CheckSinger(singer) {
		return types.ErrNotDepositSInger
	}

	if depositInfo.Status != types.Open {
		return types.ErrDepositStatusNotOpen
	}

	depoistBtcAddress := types.NewDepositBtcAddress(depositID,singer,btcAddress)
	k.SetDepositBtcAddress(ctx,depoistBtcAddress)

	return nil
}

func  (k Keeper) CheckBtcAddressReady(ctx sdk.Context, depositID string) bool {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return false
	}

	if  depositInfo.Status != types.Open {
		return true
	}

	var temBtcAddress string

	for _,singer := range depositInfo.Singers {
		btcAddress,found := k.GetDepositBtcAddress(ctx,depositID,singer)
		if !found {
			return false
		}
		if len(temBtcAddress) != 0 && temBtcAddress != btcAddress.BtcAddress  {
			return false
		}
		temBtcAddress = btcAddress.BtcAddress
	}
	return true
}

func (k Keeper) SetBtcAddressReady(ctx sdk.Context, depositID string,btcAddress string)(err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if  depositInfo.Status != types.Open {
		return types.ErrDepositStatusNotOpen
	}

	depositInfo.Status = types.AddressReady
	k.SetDepositInfo(ctx,depositInfo)

	return k.depositKeeper.SetDepositBtcAddress(ctx,depositID,btcAddress)
}

func (k Keeper) SetSpvReady(ctx sdk.Context, depositID string) (err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if  depositInfo.Status != types.AddressReady {
		return types.ErrDepositStatusNotAddressReady
	}

	depositInfo.Status = types.SPVReady
	k.SetDepositInfo(ctx,depositInfo)
	return nil
}

func (k Keeper) SetActiveDeposit(ctx sdk.Context, depositID string,singerAccount AccountID) (err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if  depositInfo.Status != types.SPVReady {
		return types.ErrDepositStatusNotSpvReady
	}

	if !depositInfo.CheckSinger(singerAccount) {
		return types.ErrNotDepositSInger
	}
	
	depositActiveInfo := types.NewDepositActiveInfo(depositID,singerAccount)
	k.SetDepositActiveInfo(ctx,depositActiveInfo)

	return nil
}

func  (k Keeper)  CheckActiveReady(ctx sdk.Context, depositID string) (bool) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return false
	}

	if  depositInfo.Status < types.SPVReady {
		return false
	}

	if  depositInfo.Status > types.SPVReady {
		return true
	}

	for _,singer := range depositInfo.Singers {
		_,found := k.GetDepositActiveInfo(ctx,depositID,singer)
		if !found {
			return false
		}
	}

	return true
}

func (k Keeper)  ActiveDeposit(ctx sdk.Context, depositID string) (err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if  depositInfo.Status != types.SPVReady {
		return types.ErrDepositStatusNotSpvReady
	}

	depositInfo.Status = types.DepositActive
	k.SetDepositInfo(ctx,depositInfo)

	return k.depositKeeper.ActiveDeposit(ctx,depositID) 
}

func (k Keeper) ActiveSingerDeposit(ctx sdk.Context, depositID string) (err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	depositInfo.Status = types.DepositActive
	k.SetDepositInfo(ctx,depositInfo)

	return nil
}

func (k Keeper)  SetClaimAddress(ctx sdk.Context, depositID string,claimAddress string) (err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if  depositInfo.Status != types.DepositActive {
		return types.ErrDepositStatusNotActive
	}

	depositInfo.Status = types.Cashing
	depositInfo.ClaimAddress = claimAddress
	k.SetDepositInfo(ctx,depositInfo)
	return nil
}

func (k Keeper) FinishDeposit(ctx sdk.Context, depositID string) (err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if  depositInfo.Status != types.CashOut {
		return types.ErrDepositStatusNotCashOut
	}

	depositInfo.Status = types.Close
	k.SetDepositInfo(ctx,depositInfo)

	return k.unlockSinger(ctx,depositInfo.Singers)
}

func (k Keeper) AberrantDeposit(ctx sdk.Context, depositID string)(err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	depositInfo.Status = types.Aberrant
	k.SetDepositInfo(ctx,depositInfo)
	return nil
}

func (k Keeper) WaitTimeOut(ctx sdk.Context,depositID string,singerAccount AccountID) (err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}
	
	if !depositInfo.CheckSinger(singerAccount) {
		return types.ErrNotDepositSInger
	}

	if depositInfo.DepositChangeTime.Add(k.WaitTime(ctx)).After(ctx.BlockHeader().Time) {
		return types.ErrNotReachWaitTime
	}

	if depositInfo.Status == types.AddressReady {
		k.depositKeeper.AberrantDeposit(ctx,depositID)
		depositInfo.Status = types.Close
		k.SetDepositInfo(ctx,depositInfo)

		return k.unlockSinger(ctx,depositInfo.Singers)
	}

	if depositInfo.Status == types.CashOut {
		k.depositKeeper.ExternalCloseDeposit(ctx,depositID)
		depositInfo.Status = types.Close
		k.SetDepositInfo(ctx,depositInfo)

		return k.unlockSinger(ctx,depositInfo.Singers)
	}
	return types.ErrNotWaitStatus
}

func (k Keeper) ReportSpvWrong(ctx sdk.Context,depositID string,singerAccount AccountID) (err error) {
	depositInfo, found := k.GetDepositInfo(ctx, depositID)
	if !found {
		return types.ErrDepositNotExist
	}
	
	if !depositInfo.CheckSinger(singerAccount) {
		return types.ErrNotDepositSInger
	}

	if  depositInfo.Status != types.SPVReady {
		return types.ErrDepositStatusNotSpvReady
	}

	depositInfo.Status = types.WrongDepositSPV
	k.SetDepositInfo(ctx,depositInfo)

	return k.depositKeeper.SetWrongDepositSpv(ctx,depositID)
}

func (k Keeper) SetWrongSingerSpv(ctx sdk.Context, depositID string)(err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if  depositInfo.Status != types.CashOut {
		return types.ErrDepositStatusNotCashOut
	}

	depositInfo.Status = types.WrongSingerSPV
	k.SetDepositInfo(ctx,depositInfo)
	return nil
}

func (k Keeper) AberrantFinishDeposit(ctx sdk.Context, depositID string)(err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	depositInfo.Status = types.Close
	k.SetDepositInfo(ctx,depositInfo)
	return k.unlockSinger(ctx,depositInfo.Singers)
}

func (k Keeper) FinishDepositPunishSinger(ctx sdk.Context, depositID string,owner AccountID)(err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}
	//punish Singer
	punishRate := k.PunishRate(ctx)
	minStake :=  depositInfo.GetMinStake()
	punishAmount := minStake.MulRaw(int64(punishRate)).QuoRaw(100)
	k.punishSinger(ctx,depositInfo.Singers,punishAmount)
	//transfer coins to deposit owner
	amount := chainTypes.NewCoin( external.DefaultBondDenom,punishAmount.MulRaw(int64(len(depositInfo.Singers))))
	err = k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, chainTypes.NewCoins(amount))
	depositInfo.Status = types.Close
	k.SetDepositInfo(ctx,depositInfo)
	return nil
}

func (k Keeper) FinishAberrantDeposit(ctx sdk.Context, depositID string,claimAccount AccountID)(err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}
	if  depositInfo.Status != types.Aberrant {
		return types.ErrDepositStatusNotAberrant
	}
	//transfer singers mortgage to claimAccount  how much?
	minStake :=  depositInfo.GetMinStake()
	k.punishSinger(ctx,depositInfo.Singers,minStake)
	amount := chainTypes.NewCoin( external.DefaultBondDenom,minStake.MulRaw(int64(len(depositInfo.Singers))))
	err = k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, claimAccount, chainTypes.NewCoins(amount))
	depositInfo.Status = types.Close
	k.SetDepositInfo(ctx,depositInfo)
	return nil
}

func (k Keeper) GetMortgageRatio(ctx sdk.Context, depositID string,baseMortgage sdk.Int) (err error,baseRate,currentRate sdk.Int) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist,sdk.ZeroInt(),sdk.ZeroInt()
	}

	minStake :=  depositInfo.GetMinStake()
	threshold := len(depositInfo.Singers)
	baseRate = minStake.MulRaw(int64(threshold*100)).Quo(baseMortgage)

	currentStake := sdk.ZeroInt()
	for _,singerAccount := range depositInfo.Singers {
		singerInfo,found := k.GetSingerInfo(ctx,singerAccount)
		if !found {
			return types.ErrSingerNotExists,baseRate,sdk.ZeroInt()
		}
		currentStake = currentStake.Add(singerInfo.SignatureMortgage)
	}

	currentRate = currentStake.MulRaw(int64(100)).Quo(baseMortgage)
	return nil,baseRate,currentRate

}