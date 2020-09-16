package pricefee

import (
	//	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdk "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/KuChainNetwork/kuchain/chain/msg"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/pricefee/keeper"
	"github.com/KuChainNetwork/kuchain/x/pricefee/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(k keeper.Keeper) msg.Handler {
	return func(ctx chainTypes.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.KuMsgOpenFee:
			return handleKuMsgOpenFee(ctx, k, msg)
		case types.KuMsgPrestoreFee:
			return handleKuMsgPrestoreFee(ctx, k, msg)
		case types.KuMsgClaimFee:
			return handleKuMsgClaimFee(ctx, k, msg)
		case types.KuMsgSetPrice:
			return handleKuMsgSetPrice(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}


func handleKuMsgOpenFee(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgOpenFee) (*sdk.Result, error) {
	msgData := types.MsgOpenFee{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.Owner)

	sdkCtx := ctx.Context()

	if _, found := k.GetFeeInfo(sdkCtx, msgData.Owner); found {
		return nil, types.ErrFeeInfoExist
	}

	err := k.NewFeeInfo(sdkCtx,msgData.Owner)
	if err != nil {
		return nil,err
	} 

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}


func handleKuMsgPrestoreFee(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgPrestoreFee) (*sdk.Result, error) {
	msgData := types.MsgPrestoreFee{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.Owner)

	sdkCtx := ctx.Context()

	_,err := k.PreStoreFee(sdkCtx,msgData.Owner,msgData.Amount)
	if err != nil {
		return nil,err
	} 

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgClaimFee(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgClaimFee) (*sdk.Result, error) {
	msgData := types.MsgClaimFee{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.Owner)

	sdkCtx := ctx.Context()

	_,err := k.ClaimFee(sdkCtx,msgData.Owner,msgData.Amount)
	if err != nil {
		return nil,err
	} 

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgSetPrice(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgSetPrice) (*sdk.Result, error) {
	msgData := types.MsgSetPrice{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.SystemAccount)

	sdkCtx := ctx.Context()

	err := k.SetPrice(sdkCtx,msgData)
	if err != nil {
		return nil,err
	} 

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}