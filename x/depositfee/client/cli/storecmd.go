package cli

import (
	"bufio"
	"github.com/KuChainNetwork/kuchain/chain/client/txutil"
	"github.com/KuChainNetwork/kuchain/x/depositfee/types"
//	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"strings"

	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// GetBroadcastCommand returns the tx broadcast command.
func GetCmdOpenFee(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open-fee [owner] ",
		Short: "open a fee record",
		Long: strings.TrimSpace(`open a fee record for deposit
`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := txutil.NewTxBuilderFromCLI(inBuf).WithTxEncoder(txutil.GetTxEncoder(cdc))
			cliCtx := txutil.NewKuCLICtxByBuf(cdc, inBuf)

			owner, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "validator account id error")
			}

			authAccAddress, err := txutil.QueryAccountAuth(cliCtx, owner)
			if err != nil {
				return sdkerrors.Wrapf(err, "query account %s auth error", owner)
			}

			msg := types.NewKuMsgRegisterSinger(authAccAddress, owner)
			cliCtx = cliCtx.WithFromAccount(owner)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return flags.PostCommands(cmd)[0]
}


func GetCmdPreStoreFee(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prestore-fee [owner-account] [amount]",
		Short: "prestore some fee",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
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

			msg := types.NewKuMsgPrestoreFee(authAccAddress, owner, amount)
			cliCtx = cliCtx.WithFromAccount(owner)
			if txBldr.FeePayer().Empty() {
				txBldr = txBldr.WithPayer(args[0])
			}
			return txutil.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	return flags.PostCommands(cmd)[0]

//	return cmd
}