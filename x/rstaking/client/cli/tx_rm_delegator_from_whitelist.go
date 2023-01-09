package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/warmage-sports/katana/x/rstaking/types"
)

var _ = strconv.Itoa(0)

func CmdRmDelegatorFromWhitelist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm-delegator-from-whitelist [del-address]",
		Short: "Remove delegator from whitelist",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDelAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRmDelegatorFromWhitelist(
				clientCtx.GetFromAddress().String(),
				argDelAddress,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
