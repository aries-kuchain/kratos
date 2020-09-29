package cli

import (
	"bufio"
	"github.com/KuChainNetwork/kuchain/chain/client/txutil"
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/KuChainNetwork/kuchain/chain/hexutil"

	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramscutils  "github.com/KuChainNetwork/kuchain/x/deposit/client/utils"

)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {

	singerTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Singer transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	singerTxCmd.AddCommand(flags.PostCommands(
		GetCmdRegisterSinger(cdc),
		GetCmdPayAccess(cdc),
		GetCmdActiveSinger(cdc),
		GetCmdPayMortgage(cdc),
		GetCmdClaimMortgage(cdc),
		GetCmdClaimAccess(cdc),
		GetCmdLogoutSinger(cdc),
		GetCmdSetAddress(cdc),
		GetCmdActiveDeposit(cdc),
		GetCmdSubmitSpv(cdc),
		GetCmdDepostWaitTimeOut(cdc),
		GetCmdReportWrongSpv(cdc),
	)...)

	return singerTxCmd
}

func GetCmdRegisterSinger(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-singer [singer-account]",
		Short: "register to be a new singer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", singerAccount)
			}

			msg := types.NewKuMsgRegisterSinger(authAccAddress, singerAccount)
			cliCtx = cliCtx.WithFromAccount(singerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdPayAccess(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pay-access [singer-account] [amount]",
		Short: "register to be a new singer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			amount, err := chainTypes.ParseCoin(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "amount parse error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", singerAccount)
			}

			msg := types.NewKuMsgPayAccess(authAccAddress, singerAccount, amount)
			cliCtx = cliCtx.WithFromAccount(singerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdActiveSinger(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "active-singer [system-account] [singer-account]",
		Short: "register to be a new singer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			systemAccount, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, systemAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", singerAccount)
			}

			msg := types.NewKuMsgActiveSinger(authAccAddress, systemAccount, singerAccount)
			cliCtx = cliCtx.WithFromAccount(singerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}


func GetCmdPayMortgage(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pay-mortgage [singer-account] [amount]",
		Short: "register to be a new singer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			amount, err := chainTypes.ParseCoin(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "amount parse error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", singerAccount)
			}


			msg := types.NewKuMsgBTCMortgage(authAccAddress, singerAccount, amount)
			cliCtx = cliCtx.WithFromAccount(singerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}


func GetCmdClaimMortgage(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-mortgage [singer-account] [amount]",
		Short: "register to be a new singer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			amount, err := chainTypes.ParseCoin(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "amount parse error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", singerAccount)
			}

			msg := types.NewKuMsgClaimBTCMortgate(authAccAddress, singerAccount, amount)
			cliCtx = cliCtx.WithFromAccount(singerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}


func GetCmdClaimAccess(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-access [singer-account]",
		Short: "register to be a new singer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", singerAccount)
			}

			msg := types.NewKuMsgClaimAccess(authAccAddress, singerAccount)
			cliCtx = cliCtx.WithFromAccount(singerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}


func GetCmdLogoutSinger(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout-singer [singer-account]",
		Short: "register to be a new singer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", singerAccount)
			}

			msg := types.NewKuMsgLogoutSinger(authAccAddress, singerAccount)
			cliCtx = cliCtx.WithFromAccount(singerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSetAddress(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-address [deposit-id] [singer-account] [btc-address]",
		Short: "register to be a new singer",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", singerAccount)
			}

			msg := types.NewKuMsgMsgSetBtcAddress(authAccAddress, singerAccount,args[0],args[2])
			cliCtx = cliCtx.WithFromAccount(singerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdActiveDeposit(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "active-deposit [deposit-id] [singer-account]",
		Short: "active a deposit",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", singerAccount)
			}

			msg := types.NewKuMsgActiveDeposit(authAccAddress, singerAccount,args[0])
			cliCtx = cliCtx.WithFromAccount(singerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
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

			spvInfo := types.NewSpvInfo(args[1],spvSubmiter,hexutil.MustDecode(paramsSpv.Version),hexutil.MustDecode(paramsSpv.TxInputVector),hexutil.MustDecode(paramsSpv.TxOutputVector),
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

	return cmd
}

func GetCmdDepostWaitTimeOut(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "singer-wait-timeout [depositID] [singer]",
		Short: "notify depisit time out ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", singerAccount)
			}

			msg := types.NewKuMsgWaitTimeout(authAccAddress, args[0],singerAccount)
			cliCtx = cliCtx.WithFromAccount(singerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdReportWrongSpv(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signer-report-wrong-spv [depositID] [singer]",
		Short: "report deposit spv is wrong",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, singerAccount)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", singerAccount)
			}

			msg := types.NewKuMsgReportSpvWrong(authAccAddress, args[0],singerAccount)
			cliCtx = cliCtx.WithFromAccount(singerAccount)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[1])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}