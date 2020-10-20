package cli

import (
	"fmt"
	//"github.com/KuChainNetwork/kuchain/chain/client/flags"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/pricefee/types"
	//	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
$ %s query pricefee fee jack
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

// GetCmdQueryValidator implements the validator query command.
func GetCmdQueryPriceInfo(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "price [base] [quote ]",
		Short: "Query a fee info",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details about an individual fee.

Example:
$ %s query pricefee fee jack
`,
				version.ClientName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			base, err := chainTypes.ParseCoin(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "base parse error")
			}

			quote, err := chainTypes.ParseCoin(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "quote parse error")
			}

			res, _, err := cliCtx.QueryStore(types.GetPriceInfoKey(base, quote), storeName)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return fmt.Errorf("no price found with base %s,quote %s", base, quote)

			}

			priceInfo, err := types.UnmarshalPriceInfo(types.Cdc(), res)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(priceInfo)
		},
	}
}
