package cli

import (
	"github.com/spf13/cobra"

	types "github.com/KuChainNetwork/kuchain/x/deposit/types"
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
		GetCmdPermintLegalCoin(cdc),
		GetCmdProhibitLegalCoin(cdc),
		GetCmdCreateDeposit(cdc),
		GetCmdCreateCoin(cdc),
		GetCmdSubmitSpv(cdc),
		GetCmdTransferDeposit(cdc),
		GetCmdDepositToCoin(cdc),
		GetCmdSetAddress(cdc),
		GetCmdFinishDeposit(cdc),
		GetCmdDepostWaitTimeOut(cdc),
		GetReportSingerWrongSpv(cdc),
		GetJudgeSpv(cdc),
		GetCmdClaimAberrant(cdc),
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
		GetCmdQueryLegalCoin(queryRoute, cdc),
		GetCmdQueryDeposit(queryRoute, cdc),
		GetCmdQueryAllDeposit(queryRoute, cdc),
		GetCmdQueryAllLegalCoin(queryRoute, cdc),
		GetCmdQueryDepositSpv(queryRoute, cdc),
		GetCmdQueryDepositMortgageRatio(queryRoute, cdc),
	)
	return txCmd
}
