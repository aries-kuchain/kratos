package cli

import (
	"strings"

	"github.com/spf13/cobra"

//	"github.com/KuChainNetwork/kuchain/x/depositfee/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	client "github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	authtxb "github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

// GetBroadcastCommand returns the tx broadcast command.
func GeteasystoreCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store [name] [value]",
		Short: "store a value on chain",
		Long: strings.TrimSpace(`just store a value for a simple test
`),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			inBuf := cmd.InOrStdin()
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := authtxb.NewTxBuilderFromCLI(inBuf).WithTxEncoder(client.GetTxEncoder(cdc))

//			msgStoreData := types.NewMsgSetStore(args[0], args[1], cliCtx.GetFromAddress())
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{})
		},
	}

	return flags.PostCommands(cmd)[0]
}
