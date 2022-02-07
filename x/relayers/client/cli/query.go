package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group relayers queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdRelayersByDenom())
	cmd.AddCommand(CmdShowThreshold())
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdRelayersByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relayers-by-denom [denom]",
		Short: "Query relayers_by_denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryRelayersByDenomRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.RelayersByDenom(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowThreshold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-threshold [denom]",
		Short: "shows a threshold",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetThresholdRequest{
				Denom: args[0],
			}
			res, err := queryClient.Threshold(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
