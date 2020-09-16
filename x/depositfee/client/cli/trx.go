package cli

import (
	"github.com/spf13/cobra"

	types "github.com/KuChainNetwork/kuchain/x/depositfee/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Auth transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetCmdOpenFee(cdc),
		GetCmdPreStoreFee(cdc),
	)
	return txCmd
}

// GetTxCmd returns the transaction commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Auth transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetCmdQueryFeeInfo(queryRoute, cdc),
	)
	return txCmd
}
