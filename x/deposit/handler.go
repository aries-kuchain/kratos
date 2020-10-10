package deposit

import (
	"github.com/KuChainNetwork/kuchain/chain/msg"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/deposit/keeper"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/KuChainNetwork/kuchain/chain/constants"
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
		case types.KuMsgDepositClaimCoin:
			return handleKuMsgDepositClaimCoin(ctx, k, msg)
		case types.KuMsgFinishDeposit:
			return handleKuMsgFinishDeposit(ctx, k, msg)
		case types.KuMsgWaitTimeout:
			return handleKuMsgWaitTimeout(ctx, k, msg)
		case types.KuMsgReportWrongSpv:
			return handleKuMsgReportWrongSpv(ctx, k, msg)
		case types.KuMsgJudgeDepositSpv:
			return handleKuMsgJudgeDepositSpv(ctx, k, msg)
		case types.KuMsgClaimAberrant:
			return handleKuMsgClaimAberrant(ctx, k, msg)
		case types.KuMsgClaimMortgage:
			return handleKuMsgClaimMortgage(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}

// Handle a message to buy name
func handleKuKuMsgCreateDeposit(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgCreateDeposit) (*sdk.Result, error) {
	msgData := types.MsgCreateDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgCreateDeposit data unmarshal error")
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
		return nil, sdkerrors.Wrapf(err, "msg MsgProhibitLegalCoin data unmarshal error")
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
		return nil, sdkerrors.Wrapf(err, "msg MsgCreateLegalCoin data unmarshal error")
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
		return nil, sdkerrors.Wrapf(err, "msg MsgProhibitLegalCoin data unmarshal error")
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
		return nil, sdkerrors.Wrapf(err, "msg MsgSubmitSpv data unmarshal error")
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
		return nil, sdkerrors.Wrapf(err, "msg MsgTransferDeposit data unmarshal error")
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
		return nil, sdkerrors.Wrapf(err, "msg MsgDepositToCoin data unmarshal error")
	}

	ctx.RequireAuth(msgData.Owner)

	sdkCtx := ctx.Context()

	if err := keeper.DepositToCoin(sdkCtx, msgData.DepositID,msgData.Owner); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgDepositClaimCoin(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgDepositClaimCoin) (*sdk.Result, error) {
	msgData := types.MsgDepositClaimCoin{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgDepositClaimCoin data unmarshal error")
	}

	ctx.RequireAuth(msgData.Owner)

	sdkCtx := ctx.Context()

	if err := keeper.ClaimDeposit(sdkCtx, msgData.DepositID,msgData.Owner,msgData.Asset,msgData.ClaimAddress); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgFinishDeposit(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgFinishDeposit) (*sdk.Result, error) {
	msgData := types.MsgFinishDeposit{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgFinishDeposit data unmarshal error")
	}

	ctx.RequireAuth(msgData.Owner)

	sdkCtx := ctx.Context()

	if err := keeper.FinishDeposit(sdkCtx, msgData.DepositID,msgData.Owner); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgWaitTimeout(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgWaitTimeout) (*sdk.Result, error) {
	msgData := types.MsgWaitTimeout{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgWaitTimeout data unmarshal error")
	}

	ctx.RequireAuth(msgData.Owner)

	sdkCtx := ctx.Context()

	if err := keeper.WaitTimeOut(sdkCtx, msgData.DepositID,msgData.Owner); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgReportWrongSpv(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgReportWrongSpv) (*sdk.Result, error) {
	msgData := types.MsgReportWrongSpv{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgReportWrongSpv data unmarshal error")
	}

	ctx.RequireAuth(msgData.Owner)

	sdkCtx := ctx.Context()

	if err := keeper.ReportWrongSingerSpv(sdkCtx, msgData.DepositID,msgData.Owner); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgJudgeDepositSpv(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgJudgeDepositSpv) (*sdk.Result, error) {
	msgData := types.MsgJudgeDepositSpv{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgJudgeDepositSpv data unmarshal error")
	}

	name, ok := msgData.SystemAccount.ToName()
	if ok && constants.IsSystemAccount(name) {
		return nil, types.ErrNotSystemAccount
	}

	ctx.RequireAuth(msgData.SystemAccount)

	sdkCtx := ctx.Context()

	if err := keeper.JudgeSpvRight(sdkCtx, msgData.DepositID,msgData.SystemAccount,msgData.SpvIsRight,msgData.FeeToSinger); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgClaimAberrant(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgClaimAberrant) (*sdk.Result, error) {
	msgData := types.MsgClaimAberrant{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgClaimAberrant data unmarshal error")
	}

	ctx.RequireAuth(msgData.ClaimAccount)

	sdkCtx := ctx.Context()

	if err := keeper.ClaimAberrantDeposit(sdkCtx, msgData.DepositID,msgData.ClaimAccount); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleKuMsgClaimMortgage(ctx chainTypes.Context, keeper keeper.Keeper, msg types.KuMsgClaimMortgage) (*sdk.Result, error) {
	msgData := types.MsgClaimMortgage{}
	if err := msg.UnmarshalData(Cdc(), &msgData); err != nil {
		return nil, sdkerrors.Wrapf(err, "msg MsgClaimMortgage data unmarshal error")
	}

	ctx.RequireAuth(msgData.ClaimAccount)

	sdkCtx := ctx.Context()

	if err := keeper.ClaimMortgageDeposit(sdkCtx, msgData.DepositID,msgData.ClaimAccount); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}