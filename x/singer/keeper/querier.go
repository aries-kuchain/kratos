package keeper

import (
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QuerySingerInfo:
			return querySingerInfo(ctx, path[1:], req, keeper)
		case types.QueryAllDeposit:
			return queryAllDeposit(ctx, path[1:], req, keeper)
		default:
			return nil, nil //sdk.ErrUnknownRequest("unknown bank query endpoint")
		}
	}
}

type QueryResResolve struct {
	Value string `json:"value"`
}

// nolint: unparam
func querySingerInfo(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	var params types.QuerySingerInfoParams
	err = k.cdc.UnmarshalJSON(req.Data, &params)

	ctx.Logger().Debug("querySingerInfo:", params, err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	singerInfo, found := k.GetSingerInfo(ctx, params.SingerAccount)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "singer  do not find")
	}

	singerInfoResponse := types.NewQueryDepositMortgageRatioResponse(singerInfo.SingerAccount, singerInfo.AccessAsset, singerInfo.Status, singerInfo.SignatureMortgage, singerInfo.LockMortgage)

	bz, err := codec.MarshalJSONIndent(k.cdc, singerInfoResponse)
	ctx.Logger().Debug("queryValidatorOutstandingRewards:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
func queryAllDeposit(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	depositInfos := k.GetAllDeposit(ctx)

	allDepositInfoResponse := types.NewQueryAllDepositResponse(depositInfos)
	bz, err := codec.MarshalJSONIndent(k.cdc, allDepositInfoResponse)
	ctx.Logger().Debug("queryValidatorOutstandingRewards:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}