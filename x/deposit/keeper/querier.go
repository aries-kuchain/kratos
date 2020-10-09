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

	err,baseRatio,currentRatio,punishRatio := k.GetMortgageRatio(ctx, params.DepositID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	depositMortgageRatio := types.NewQueryDepositMortgageRatioResponse(params.DepositID,baseRatio,currentRatio,punishRatio)

	bz, err := codec.MarshalJSONIndent(k.cdc, depositMortgageRatio)
	ctx.Logger().Debug("queryValidatorOutstandingRewards:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

