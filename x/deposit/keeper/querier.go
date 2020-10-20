package keeper

import (
	//	"encoding/json"

	//	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryDepositMortgageRatioParams:
			return queryQueryDepositMortgageRatio(ctx, path[1:], req, k)
		case types.QueryAllDeposit:
			return queryAllDeposit(ctx, path[1:], req, k)
		case types.QueryDepositInfo:
			return queryDepositInfo(ctx, path[1:], req, k)
		case types.QueryUserDeposit:
			return queryUserDeposit(ctx, path[1:], req, k)
		case types.QueryCashReadyDeposit:
			return queryCashReadyDeposit(ctx, path[1:], req, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func queryQueryDepositMortgageRatio(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryDepositParams
	err := k.cdc.UnmarshalJSON(req.Data, &params)

	ctx.Logger().Debug("queryValidatorOutstandingRewards:", params, err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	err, baseRatio := k.GetMortgageRatio(ctx, params.DepositID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	depositMortgageRatio := types.NewQueryDepositMortgageRatioResponse(params.DepositID, baseRatio)

	bz, err := codec.MarshalJSONIndent(k.cdc, depositMortgageRatio)
	ctx.Logger().Debug("queryValidatorOutstandingRewards:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryAllDeposit(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	depositInfos := k.GetAllDepositInfo(ctx)
	bz, err := codec.MarshalJSONIndent(k.cdc, depositInfos)
	ctx.Logger().Debug("queryAllDeposit:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryDepositInfo(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	var params types.QueryDepositParams
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

func queryUserDeposit(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	var params types.QueryUserDepositParams
	err = k.cdc.UnmarshalJSON(req.Data, &params)

	ctx.Logger().Debug("querySingerInfo:", params, err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	depositInfos := k.GetAllDepositInfo(ctx)

	reponse := types.NewQueryAllDepositWithOwnerResponse(depositInfos, params.OwerAccount)
	bz, err := codec.MarshalJSONIndent(k.cdc, reponse)
	ctx.Logger().Debug("queryValidatorOutstandingRewards:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryCashReadyDeposit(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	depositInfos := k.GetAllDepositInfo(ctx)
	reponse := types.NewQueryAllDepositWithCashReadyResponse(depositInfos)
	bz, err := codec.MarshalJSONIndent(k.cdc, reponse)
	ctx.Logger().Debug("queryAllDeposit:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
