package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

func (k Keeper) MortgageRage(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyMortgageRate, &res)
	return res
}

func (k Keeper) DepositFeeRate(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyDepositFeeRate, &res)
	return res
}

func (k Keeper) ClaimFeeRate(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyClaimFeeRate, &res)
	return res
}

func (k Keeper) Threshold(ctx sdk.Context) (res int) {
	k.paramstore.Get(ctx, types.KeyThreshold, &res)
	return res
}

func (k Keeper) WaitTime(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyWaitTime, &res)
	return res
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MortgageRage(ctx),
		k.DepositFeeRate(ctx),
		k.ClaimFeeRate(ctx),
		k.Threshold(ctx),
		k.WaitTime(ctx),
	)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}