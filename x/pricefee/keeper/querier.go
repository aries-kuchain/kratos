package keeper

import (
	//	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/KuChainNetwork/kuchain/x/pricefee/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryFeeInfo:
			return queryDepositInfo(ctx, path[1:], req, keeper)

		default:
			return nil, nil //sdk.ErrUnknownRequest("unknown bank query endpoint")
		}
	}
}

func queryDepositInfo(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) (res []byte, err error) {
	var params types.QueryFeeInfoParams
	err = k.cdc.UnmarshalJSON(req.Data, &params)

	ctx.Logger().Debug("queryFeeInfo:", params, err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	feeInfo, found := k.GetFeeInfo(ctx, params.OwnerAccount)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "feeinfo  do not find")
	}

	bz, err := codec.MarshalJSONIndent(k.cdc, feeInfo)
	ctx.Logger().Debug("queryValidatorOutstandingRewards:", bz, "err:", err)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
