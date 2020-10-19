package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

func (k Keeper) MinAccessAmount(ctx sdk.Context) (res sdk.Int) {
	k.paramstore.Get(ctx, types.KeyMinAccessAmount, &res)
	return res
}

func (k Keeper) WaitTime(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyWaitTime, &res)
	return res
}

func (k Keeper) PunishRate(ctx sdk.Context) (res int) {
	k.paramstore.Get(ctx, types.KeyPunishRage, &res)
	return res
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MinAccessAmount(ctx),
		k.WaitTime(ctx),
		k.PunishRate(ctx),
	)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
