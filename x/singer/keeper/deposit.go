package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"bytes"
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


func  (k Keeper) NewDepositBtcAddress(ctx sdk.Context, depositID string,singer AccountID,btcAddress []byte) (err error){
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

	var temBtcAddress []byte

	for _,singer := range depositInfo.Singers {
		btcAddress,found := k.GetDepositBtcAddress(ctx,depositID,singer)
		if !found {
			return false
		}
		if len(temBtcAddress) != 0 &&  !bytes.Equal(temBtcAddress,btcAddress.BtcAddress) {
			return false
		}
		temBtcAddress = []byte("")
		temBtcAddress = append(temBtcAddress,btcAddress.BtcAddress...)
	}
	return true
}

func (k Keeper) SetBtcAddressReady(ctx sdk.Context, depositID string,btcAddress []byte)(err error) {
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

	if  depositInfo.Status != types.SPVReady {
		return types.ErrDepositStatusNotSpvReady
	}

	depositInfo.Status = types.DepositActive
	k.SetDepositInfo(ctx,depositInfo)

	return nil
}

func (k Keeper)  SetClaimAddress(ctx sdk.Context, depositID string,claimAddress []byte) (err error) {
	depositInfo,found := k.GetDepositInfo(ctx,depositID)
	if !found {
		return types.ErrDepositNotExist
	}

	if  depositInfo.Status != types.DepositActive {
		return types.ErrDepositStatusNotActive
	}

	depositInfo.Status = types.Cashing
	depositInfo.ClaimAddress = append(depositInfo.ClaimAddress,claimAddress...)
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