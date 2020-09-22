package cli

import (
	//	"fmt"
	//"github.com/KuChainNetwork/kuchain/chain/client/flags"
	chainTypes "github.com/KuChainNetwork/kuchain/chain/types"
	"github.com/KuChainNetwork/kuchain/x/deposit/types"
	singerTypes "github.com/KuChainNetwork/kuchain/x/singer/types"
	//"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	//"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	//"strings"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// GetCmdResolveName queries information about a name
func GetCmdQueryLegalCoin(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "legal-coin [asset]",
		Short: "Query a legal coin information",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			asset, err := chainTypes.ParseCoin(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "amount parse error")
			}

			res, _, err := cliCtx.QueryStore(types.GetLegalCoinKey(asset), storeName)
			if err != nil {
				return err
			}

			legalCoin, err := types.UnmarshalLegalCoin(types.Cdc(), res)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(legalCoin)
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

func GetCmdQueryAllLegalCoin(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "all-legalcoin",
		Short: "Query all deposit ",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			resKVs, _, err := cliCtx.QuerySubspace(types.LegalDepositCoinKey, storeName)
			if err != nil {
				return err
			}

			var legalCoins []types.LegalCoin
			for _, kv := range resKVs {
				legalcoin, err := types.UnmarshalLegalCoin(types.Cdc(), kv.Value)
				if err != nil {
					return err
				}

				legalCoins = append(legalCoins, legalcoin)
			}

			return cliCtx.PrintOutput(legalCoins)
		},
	}
}


func GetCmdQueryDepositSpv(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deposit-spv [depositID]",
		Short: "Query a deposit spv info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			resKVs, _, err := cliCtx.QuerySubspace(types.GetDepositSpvKey(args[0]), storeName)
			if err != nil {
				return err
			}

			var spvInfos []singerTypes.SpvInfo
			for _, kv := range resKVs {
				spvInfo, err := singerTypes.UnmarshalSpvInfo(types.Cdc(), kv.Value)
				if err != nil {
					return err
				}

				spvInfos = append(spvInfos, spvInfo)
			}

			return cliCtx.PrintOutput(spvInfos)
		},
	}
}