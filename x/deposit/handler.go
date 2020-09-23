package deposit

import (
	"github.com/KuChainNetwork/kuchain/chain/msg"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/deposit/keeper"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(k keeper.Keeper) msg.Handler {
	return func(ctx chainTypes.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.KuMsgCreateDeposit:
			return handleKuKuMsgCreateDeposit(ctx, k, msg)
		case types.KuMsgCreateLegalCoin:
			return handleKuMsgCreateLegalCoin(ctx, k, msg)
		case types.KuMsgProhibitLegalCoin:
			return handleKuMsgProhibitLegalCoin(ctx, k, msg)
		case types.KuMsgPermintLegalCoin:
			return handleKuMsgPermintLegalCoin(ctx, k, msg)
		case types.KuMsgSubmitSpv:
			return handleKuMsgSubmitSpv(ctx, k, msg)
		case types.KuMsgTransferDeposit:
			return handleKuMsgTransferDeposit(ctx, k, msg)
		case types.KuMsgDepositToCoin:
			return handleKuMsgDepositToCoin(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}

// Handle a message to buy name
func handleKuKuMsgCreateDeposit(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgCreateDeposit) (*sdk.Result, error) {
	msgData := types.MsgCreateDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.OwnerAccount)

	sdkCtx := ctx.Context()

	_, err := keeper.NewDepositInfo(sdkCtx, msgData.OwnerAccount, msgData.Amount)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// Handle a message to buy name
func handleKuMsgProhibitLegalCoin(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgProhibitLegalCoin) (*sdk.Result, error) {
	msgData := types.MsgProhibitLegalCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.SystemAccount)

	sdkCtx := ctx.Context()

	if err := keeper.ProhibitLegalCoin(sdkCtx, msgData.SystemAccount, msgData.Amount); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil

}

// Handle a message to buy name
func handleKuMsgCreateLegalCoin(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgCreateLegalCoin) (*sdk.Result, error) {
	msgData := types.MsgCreateLegalCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.SystemAccount)

	sdkCtx := ctx.Context()

	if err := keeper.CreateLegalCoin(sdkCtx, msgData.SystemAccount, msgData.MaxSupply, msgData.Symbol); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgPermintLegalCoin(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgPermintLegalCoin) (*sdk.Result, error) {
	msgData := types.MsgProhibitLegalCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.SystemAccount)

	sdkCtx := ctx.Context()

	if err := keeper.PermintLegalCoin(sdkCtx, msgData.SystemAccount, msgData.Amount); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgSubmitSpv(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgSubmitSpv) (*sdk.Result, error) {
	msgData := types.MsgSubmitSpv{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.SpvInfo.SpvSubmiter)

	sdkCtx := ctx.Context()

	if err := keeper.NewSpvInfo(sdkCtx, msgData.SpvInfo); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgTransferDeposit(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgTransferDeposit) (*sdk.Result, error) {
	msgData := types.MsgTransferDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.From)

	sdkCtx := ctx.Context()

	if err := keeper.TransferDeposit(sdkCtx, msgData.DepositID,msgData.From,msgData.To); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgDepositToCoin(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgDepositToCoin) (*sdk.Result, error) {
	msgData := types.MsgDepositToCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgRegisterSinger data unmarshal error")
	}

	ctx.RequireAuth(msgData.Owner)

	sdkCtx := ctx.Context()

	if err := keeper.DepositToCoin(sdkCtx, msgData.DepositID,msgData.Owner); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}