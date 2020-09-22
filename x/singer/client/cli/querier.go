package cli

import (
	"fmt"
	"github.com/KuChainNetwork/kuchain/chain/client/flags"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/singer/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"strings"
)

// GetCmdResolveName queries information about a name

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	singerQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the singer module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	singerQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQuerySinger(queryRoute, cdc),
		GetCmdQueryAllSinger(queryRoute, cdc),
		GetCmdQueryAllDeposit(queryRoute, cdc),
		GetCmdQueryDeposit(queryRoute, cdc),
		GetCmdQueryDepositBtcAddress(queryRoute, cdc),
	)...)

	return singerQueryCmd

}

// GetCmdQueryValidator implements the validator query command.
func GetCmdQuerySinger(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "singer [singer-account]",
		Short: "Query a singer",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details about an individual singer.

Example:
$ %s query singer singer jack
`,
				version.ClientName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			singerAccount, err := chainTypes.NewAccountIDFromStr(args[0])
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryStore(types.GetSingerInfoKey(singerAccount), storeName)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return fmt.Errorf("no validator found with address %s", singerAccount)
			}

			singer, err := types.UnmarshalSingerInfo(types.Cdc(), res)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(singer)
		},
	}
}

func GetCmdQueryAllSinger(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "all-singer",
		Short: "Query all singer",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details about an individual singer.

Example:
$ %s query singer singer jack
`,
				version.ClientName,
			),
		),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			resKVs, _, err := cliCtx.QuerySubspace(types.SingerInfoKey, storeName)
			if err != nil {
				return err
			}

			var singerInfos types.SingerInfos
			for _, kv := range resKVs {
				singerInfo, err := types.UnmarshalSingerInfo(types.Cdc(), kv.Value)
				if err != nil {
					return err
				}

				singerInfos = append(singerInfos, singerInfo)
			}

			return cliCtx.PrintOutput(singerInfos)
		},
	}
}

func GetCmdQueryAllDeposit(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "all-deposit",
		Short: "Query all deposit ",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			resKVs, _, err := cliCtx.QuerySubspace(types.DepositInfoKey, storeName)
			if err != nil {
				return err
			}

			var depositInfos []types.DepositInfo
			for _, kv := range resKVs {
				depositInfo, err := types.UnmarshalDepositInfo(types.Cdc(), kv.Value)
				if err != nil {
					return err
				}

				depositInfos = append(depositInfos, depositInfo)
			}

			return cliCtx.PrintOutput(depositInfos)
		},
	}
}

func GetCmdQueryDeposit(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deposit [depositID]",
		Short: "Query a deposit information",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryStore(types.GetDepositInfoKey(args[0]), storeName)
			if err != nil {
				return err
			}

			depositInfo, err := types.UnmarshalDepositInfo(types.Cdc(), res)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(depositInfo)
		},
	}
}


func GetCmdQueryDepositBtcAddress(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deposit-btcaddress [depositID]",
		Short: "Query deposit btc address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			resKVs, _, err := cliCtx.QuerySubspace(types.GetDepositBtcAddressDepositKey(args[0]), storeName)
			if err != nil {
				return err
			}

			var depositBtcAddress []types.DepositBtcAddress
			for _, kv := range resKVs {
				btcAddress, err := types.UnmarshalDepositBtcAddress(types.Cdc(), kv.Value)
				if err != nil {
					return err
				}

				depositBtcAddress = append(depositBtcAddress, btcAddress)
			}

			return cliCtx.PrintOutput(depositBtcAddress)
		},
	}
}