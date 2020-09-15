package singer

import (

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/KuChainNetwork/kuchain/chain/msg"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/singer/keeper"
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	"github.com/KuChainNetwork/kuchain/x/singer/external"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(k keeper.Keeper) msg.Handler {
	return func(ctx chainTypes.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.KuMsgRegisterSinger:
			return handleKuMsgRegisterSinger(ctx, k, msg)
		case types.KuMsgPayAccess:
			return handleKuMsgPayAccess(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}


func handleKuMsgRegisterSinger(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgRegisterSinger) (*sdk.Result, error) {
	msgData := types.MsgRegisterSinger{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg CreateValidator data unmarshal error")
	}

	ctx.RequireAuth(msgData.SingerAccount)

	sdkCtx := ctx.Context()
	
	if _,found := k.GetSingerInfo(sdkCtx,msgData.SingerAccount);found {
		return nil, types.ErrSingerExists
	}
		
	if !k.ValidatorAccount(sdkCtx, msgData.SingerAccount) {
		return nil, types.ErrUnKnowAccount
	}

	newSingerInfo := types.NewSingerInfo(msgData.SingerAccount)

	k.SetSingerInfo(sdkCtx,newSingerInfo)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgPayAccess(ctx chainTypes.Context, k keeper.Keeper, msg types.KuMsgPayAccess) (*sdk.Result, error) {
	msgData := types.MsgPayAccess{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg CreateValidator data unmarshal error")
	}

	ctx.RequireAuth(msgData.SingerAccount)

	sdkCtx := ctx.Context()
	
	if _,found := k.GetSingerInfo(sdkCtx,msgData.SingerAccount);!found {
		return nil, types.ErrSingerNotExists
	}

	if msgData.Amount.Denom != external.DefaultBondDenom {
		return nil, types.ErrBadDenom
	}

	_,err := k.SingerAddAccess(sdkCtx,msgData.SingerAccount,msgData.Amount)
	if err != nil {
		return nil,err
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
