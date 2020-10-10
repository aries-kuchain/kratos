package cli

import (
	"bufio"
	"github.com/KuChainNetwork/kuchain/chain/client/txutil"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/chain/hexutil"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramscutils  "github.com/KuChainNetwork/kuchain/x/deposit/client/utils"
	singerTypes	"github.com/KuChainNetwork/kuchain/x/singer/types"


)

func GetCmdPermintLegalCoin(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "permint-legalcoin [systemAccount] [asset]",
		Short: "Permint a legal coin ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			systemAccount, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			asset, err := chainTypes.ParseCoin(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "amount parse error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, systemAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", systemAccount)
			}

			msg := types.NewKuMsgPermintLegalCoin(authAccAddress, systemAccount, asset)
			cliCtx = cliCtx.WithFromAccount(systemAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetCmdProhibitLegalCoin(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prohibit-legalcoin [systemAccount] [asset]",
		Short: "prohibit a legal coin ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			systemAccount, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			asset, err := chainTypes.ParseCoin(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "amount parse error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, systemAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", systemAccount)
			}

			msg := types.NewKuMsgProhibitLegalCoin(authAccAddress, systemAccount, asset)
			cliCtx = cliCtx.WithFromAccount(systemAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

// GetCmdCreateLegalCoin returns the tx broadcast command.
func GetCmdCreateDeposit(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-deposit [owner] [amount]",
		Short: "create a deposit ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			owner, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			amount, err := chainTypes.ParseCoin(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "amount parse error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, owner)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", owner)
			}

			msg := types.NewKuMsgCreateDeposit(authAccAddress, owner, amount)
			cliCtx = cliCtx.WithFromAccount(owner)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetCmdCreateCoin(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-legalcoin [systemAccount] [asset] [symbol]",
		Short: "create a legal coin ",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			systemAccount, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			asset, err := chainTypes.ParseCoin(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "amount parse error")
			}

			symbol := chainTypes.MustName(args[2])

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, systemAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", systemAccount)
			}

			msg := types.NewKuMsgCreateLegalCoin(authAccAddress, systemAccount, asset, symbol)
			cliCtx = cliCtx.WithFromAccount(systemAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetCmdSubmitSpv(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-spv [spv-submiter] [depositID] [spv-file]",
		Short: "submit a spv for deposit ",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			spvSubmiter, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			paramsSpv, err := paramscutils.ParseParamsSpvJSON(cdc, args[2])
			if err != nil {
				return err
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, spvSubmiter)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", spvSubmiter)
			}

			spvInfo := singerTypes.NewSpvInfo(args[1],spvSubmiter,hexutil.MustDecode(paramsSpv.Version),hexutil.MustDecode(paramsSpv.TxInputVector),hexutil.MustDecode(paramsSpv.TxOutputVector),
				hexutil.MustDecode(paramsSpv.TxLockTime),hexutil.MustDecode(paramsSpv.MerkleProof),hexutil.MustDecode(paramsSpv.BitcoinHeaders),paramsSpv.FundingOutputIndex,paramsSpv.TxIndexInBlock,
			)
			msg := types.NewKuMsgSubmitSpv(authAccAddress, spvInfo)
			cliCtx = cliCtx.WithFromAccount(spvSubmiter)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetCmdTransferDeposit(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-deposit [depositID] [from] [to] [memo]",
		Short: "create a legal coin ",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			fromAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			toAccount, err := chainTypes.NewAccountIDFromStr(args[2])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, fromAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", fromAccount)
			}

			msg := types.NewKuMsgTransferDeposit(authAccAddress, args[0],fromAccount, toAccount, args[3])
			cliCtx = cliCtx.WithFromAccount(fromAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetCmdDepositToCoin(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-to-coin [depositID] [owner]",
		Short: "trans deposit to coin ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			ownerAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

		
			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, ownerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", ownerAccount)
			}

			msg := types.NewKuMsgDepositToCoin(authAccAddress, args[0],ownerAccount)
			cliCtx = cliCtx.WithFromAccount(ownerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}


func GetCmdSetAddress(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-deposit [deposit-id] [owner-account] [amount] [btc-address]",
		Short: "claim a deposit ",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			ownerAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, ownerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", ownerAccount)
			}

			asset, err := chainTypes.ParseCoin(args[2])
			if err != nil {
				return sdkerrors.Wrap(err, "amount parse error")
			}

			msg := types.NewKuMsgDepositClaimCoin(authAccAddress,args[0], ownerAccount,asset,args[3])
			cliCtx = cliCtx.WithFromAccount(ownerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetCmdFinishDeposit(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "finish-deposit [depositID] [owner]",
		Short: "finish a deposit ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			ownerAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

		
			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, ownerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", ownerAccount)
			}

			msg := types.NewKuMsgFinishDeposit(authAccAddress, args[0],ownerAccount)
			cliCtx = cliCtx.WithFromAccount(ownerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetCmdDepostWaitTimeOut(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-wait-timeout [depositID] [owner]",
		Short: "notify depisit time out ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			ownerAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, ownerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", ownerAccount)
			}

			msg := types.NewKuMsgWaitTimeout(authAccAddress, args[0],ownerAccount)
			cliCtx = cliCtx.WithFromAccount(ownerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetReportSingerWrongSpv(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-wrong-spv [depositID] [owner]",
		Short: "report singer spv is wrong ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			ownerAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, ownerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", ownerAccount)
			}

			msg := types.NewKuMsgReportWrongSpv(authAccAddress, args[0],ownerAccount)
			cliCtx = cliCtx.WithFromAccount(ownerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetJudgeSpv(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "judge-spv [depositID] [systemAccount] [spvIsRight] [feeToSinger]",
		Short: "judge spv is right ",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			systemAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, systemAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", systemAccount)
			}

			spvIsRight := args[2] == "1"
			feeToSinger := args[3] == "1"

			msg := types.NewKuMsgJudgeDepositSpv(authAccAddress, args[0],systemAccount,spvIsRight,feeToSinger)
			cliCtx = cliCtx.WithFromAccount(systemAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetCmdClaimAberrant(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-aberrant [deposit-id] [claim-account] [amount] ",
		Short: "claim a aberrant deposit ",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			claimAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, claimAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", claimAccount)
			}

			asset, err := chainTypes.ParseCoin(args[2])
			if err != nil {
				return sdkerrors.Wrap(err, "amount parse error")
			}

			msg := types.NewKuMsgClaimAberrant(authAccAddress,args[0], claimAccount,asset)
			cliCtx = cliCtx.WithFromAccount(claimAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetCmdClaimMortgage(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-mortgage-deposit [deposit-id] [claim-account] [amount] ",
		Short: "claim a lack mortgage deposit ",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			claimAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, claimAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", claimAccount)
			}

			asset, err := chainTypes.ParseCoin(args[2])
			if err != nil {
				return sdkerrors.Wrap(err, "amount parse error")
			}

			msg := types.NewKuMsgClaimMortgage(authAccAddress,args[0], claimAccount,asset)
			cliCtx = cliCtx.WithFromAccount(claimAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}

func GetCmdCashReadyDeposit(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cashready-deposit [deposit-id] [operator-account] ",
		Short: "CashReady an active  deposit ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			operatorAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, operatorAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", operatorAccount)
			}

			msg := types.NewKuMsgCashReadyDeposit(authAccAddress,args[0], operatorAccount)
			cliCtx = cliCtx.WithFromAccount(operatorAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}