package singer

import (
	"github.com/KuChainNetwork/kuchain/chain/constants"
	"github.com/KuChainNetwork/kuchain/chain/msg"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/singer/external"
	"github.com/KuChainNetwork/kuchain/x/singer/keeper"
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"fmt"
)

func NewHandler(k keeper.Keeper) msg.Handler {
	return func(ctx chainTypes.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.KuMsgRegisterSinger:
			return handleKuMsgRegisterSinger(ctx, k, msg)
		case types.KuMsgPayAccess:
			return handleKuMsgPayAccess(ctx, k, msg)
		case types.KuMsgActiveSinger:
			return handleKuMsgActiveSinger(ctx, k, msg)
		case types.KuMsgBTCMortgage:
			return handleKuMsgBTCMortgage(ctx, k, msg)
		case types.KuMsgClaimBTCMortgate:
			return handleKuMsgClaimBTCMortgate(ctx, k, msg)
		case types.KuMsgClaimAccess:
			return handleKuMsgClaimAccess(ctx, k, msg)
		case types.KuMsgLogoutSinger:
			return handleKuMsgLogoutSinger(ctx, k, msg)
		case types.KuMsgMsgSetBtcAddress:
			return handleKuMsgMsgSetBtcAddress(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}

func handleKuMsgRegisterSinger(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgRegisterSinger) (*sdk.Result, error) {
	msgData := types.MsgRegisterSinger{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.SingerAccount)

	sdkCtx := ctx.Context()

	if _, found := k.GetSingerInfo(sdkCtx, msgData.SingerAccount); found {
		return nil, types.ErrSingerExists
	}

	if !k.ValidatorAccount(sdkCtx, msgData.SingerAccount) {
		return nil, types.ErrUnKnowAccount
	}
	
	k.NewSingerInfo(sdkCtx, msgData.SingerAccount)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgPayAccess(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgPayAccess) (*sdk.Result, error) {
	msgData := types.MsgPayAccess{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgPayAccess data unmarshal error")
	}

	ctx.RequireAuth(msgData.SingerAccount)

	sdkCtx := ctx.Context()

	if _, found := k.GetSingerInfo(sdkCtx, msgData.SingerAccount); !found {
		return nil, types.ErrSingerNotExists
	}

	if msgData.Amount.Denom != external.DefaultBondDenom {
		return nil, types.ErrBadDenom
	}

	_, err := k.SingerAddAccess(sdkCtx, msgData.SingerAccount, msgData.Amount)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgActiveSinger(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgActiveSinger) (*sdk.Result, error) {
	msgData := types.MsgActiveSinger{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgActiveSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.SystemAccount)

	sdkCtx := ctx.Context()

	if _, found := k.GetSingerInfo(sdkCtx, msgData.SingerAccount); !found {
		return nil, types.ErrSingerNotExists
	}

	name, ok := msgData.SystemAccount.ToName()
	if ok && constants.IsSystemAccount(name) {
		return nil, types.ErrNotSystemAccount
	}

	if err := k.ActiveSingerInfo(sdkCtx, msgData.SingerAccount); err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgBTCMortgage(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgBTCMortgage) (*sdk.Result, error) {
	msgData := types.MsgPayBTCMortgate{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgPayBTCMortgate data unmarshal error")
	}

	ctx.RequireAuth(msgData.SingerAccount)

	sdkCtx := ctx.Context()

	if _, found := k.GetSingerInfo(sdkCtx, msgData.SingerAccount); !found {
		return nil, types.ErrSingerNotExists
	}

	if msgData.Amount.Denom != external.DefaultBondDenom {
		return nil, types.ErrBadDenom
	}

	_, err := k.SingerAddBTCMortgate(sdkCtx, msgData.SingerAccount, msgData.Amount)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}


func handleKuMsgClaimBTCMortgate(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgClaimBTCMortgate) (*sdk.Result, error) {
	msgData := types.MsgClaimBTCMortgate{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgClaimBTCMortgate data unmarshal error")
	}

	ctx.RequireAuth(msgData.SingerAccount)

	sdkCtx := ctx.Context()

	if _, found := k.GetSingerInfo(sdkCtx, msgData.SingerAccount); !found {
		return nil, types.ErrSingerNotExists
	}

	if msgData.Amount.Denom != external.DefaultBondDenom {
		return nil, types.ErrBadDenom
	}

	_, err := k.SingerClaimBTCMortgate(sdkCtx, msgData.SingerAccount, msgData.Amount)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgClaimAccess(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgClaimAccess) (*sdk.Result, error) {
	msgData := types.MsgClaimAccess{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.SingerAccount)

	sdkCtx := ctx.Context()

	if _, found := k.GetSingerInfo(sdkCtx, msgData.SingerAccount); !found {
		return nil, types.ErrSingerNotExists
	}

	_,err := k.SingerClaimAccess(sdkCtx,msgData.SingerAccount)
	if  err != nil {
		return nil,err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgLogoutSinger(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgLogoutSinger) (*sdk.Result, error) {
	msgData := types.MsgLogoutSinger{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.SingerAccount)

	sdkCtx := ctx.Context()

	if _, found := k.GetSingerInfo(sdkCtx, msgData.SingerAccount); !found {
		return nil, types.ErrSingerNotExists
	}

	err := k.SingerLogoutAccess(sdkCtx,msgData.SingerAccount)
	if  err != nil {
		return nil,err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgMsgSetBtcAddress(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgMsgSetBtcAddress) (*sdk.Result, error) {
	msgData := types.MsgSetBtcAddress{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.SingerAccount)

	sdkCtx := ctx.Context()

	if _, found := k.GetSingerInfo(sdkCtx, msgData.SingerAccount); !found {
		return nil, types.ErrSingerNotExists
	}

	err := k.NewDepositBtcAddress(sdkCtx,msgData.DepoistID,msgData.SingerAccount,msgData.BtcAddress)
	if  err != nil {
		return nil,err
	}

	if k.CheckBtcAddressReady(sdkCtx,msgData.DepoistID) {
		err = k.SetBtcAddressReady(sdkCtx,msgData.DepoistID,msgData.BtcAddress)
		if err != nil {
			return nil,err
		}
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}