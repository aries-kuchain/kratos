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
		case types.QueryDepositInfo:
			return queryDepositInfo(ctx, path[1:], req, keeper)
		case types.QuerySingerDeposit:
			return querySingerDeposit(ctx, path[1:], req, keeper)
		case types.QueryDepositAddress:
			return queryDepositAddress(ctx, path[1:], req, keeper)
		case types.QueryDepositSpv:
			return queryDepositSpv(ctx, path[1:], req, keeper)
		case types.QueryParameters:
			return queryParameters(ctx, path[1:], req, keeper)
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

func queryDepositInfo(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	var params types.QueryDepositInfoParams
	err = k.cdc.UnmarshalJSON(req.Data, &params)

	ctx.Logger().Debug("querySingerInfo:", params, err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	depositInfo, found := k.GetDepositInfo(ctx, params.DepositID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "singer  do not find")
	}

	bz, err := codec.MarshalJSONIndent(k.cdc, depositInfo)
	ctx.Logger().Debug("queryValidatorOutstandingRewards:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func querySingerDeposit(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	var params types.QuerySingerInfoParams
	err = k.cdc.UnmarshalJSON(req.Data, &params)

	ctx.Logger().Debug("querySingerInfo:", params, err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	depositInfos := k.GetAllDeposit(ctx)
	allDepositInfoResponse := types.NewQuerySingerDepositResponse(depositInfos, params.SingerAccount)

	bz, err := codec.MarshalJSONIndent(k.cdc, allDepositInfoResponse)
	ctx.Logger().Debug("queryValidatorOutstandingRewards:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryDepositAddress(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	var params types.QueryDepositInfoParams
	err = k.cdc.UnmarshalJSON(req.Data, &params)

	ctx.Logger().Debug("queryDepositAddress:", params, err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	depositAddress := k.GetDepositAddress(ctx, params.DepositID)

	bz, err := codec.MarshalJSONIndent(k.cdc, depositAddress)
	ctx.Logger().Debug("queryDepositAddress:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryDepositSpv(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	var params types.QueryDepositInfoParams
	err = k.cdc.UnmarshalJSON(req.Data, &params)

	ctx.Logger().Debug("queryDepositAddress:", params, err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	depositSpv := k.GetDepositClaimSpv(ctx, params.DepositID)

	bz, err := codec.MarshalJSONIndent(k.cdc, depositSpv)
	ctx.Logger().Debug("queryDepositAddress:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryParameters(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	params := k.GetParams(ctx)

	res, err = codec.MarshalJSONIndent(k.cdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}