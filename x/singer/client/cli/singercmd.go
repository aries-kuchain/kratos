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

	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
