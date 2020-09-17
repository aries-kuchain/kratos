package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

)

func (k Keeper) MinAccessAmount(ctx sdk.Context) (res sdk.Int) {
	k.paramstore.Get(ctx, types.KeyMinAccessAmount, &res)
	return res
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MinAccessAmount(ctx),
	)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}