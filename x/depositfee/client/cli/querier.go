package cli

import (
	"fmt"
	//"github.com/KuChainNetwork/kuchain/chain/client/flags"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/depositfee/types"
//	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"strings"
)

// GetCmdQueryValidator implements the validator query command.
func GetCmdQueryFeeInfo(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "fee [owner-account]",
		Short: "Query a fee info",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details about an individual fee.

Example:
$ %s query depositfee fee jack
`,
				version.ClientName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			owner, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryStore(types.GetFeeInfoKey(owner), storeName)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return fmt.Errorf("no validator found with address %s", owner)
			}

			singer, err := types.UnmarshalFeeInfo(types.Cdc(), res)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(singer)
		},
	}
}
