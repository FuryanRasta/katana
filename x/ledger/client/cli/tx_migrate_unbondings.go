package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var _ = strconv.Itoa(0)
var FlagUnbondings = "unbondings"

func CmdMigrateUnbondings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate-unbondings [denom]",
		Short: "Migrate unbondings",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Broadcast message migrate init with an unbonding list which can be given through a JSON file.

Example:
$ %s tx ledger migrate-unbondings uratom  --unbondings="path/to/unbondings.json" --from admin

Where unbondings.json could be like this:
[
    {
        "denom": "uratom",
        "pool": "cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75",
        "era": 2,
        "unbondings": [
            {
                "amount": "1000000",
                "recipient": "cosmos1cad0efr25faywnjp8qp36l8zlqa2sgz0jwn0hl"
            },
            {
                "amount": "2000000",
                "recipient": "cosmos13mwxtgrljf9d5r72sc28496ua4lsga0jvmqz8x"
            }
        ]
    },
    {
        "denom": "uratom",
        "pool": "cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75",
        "era": 3,
        "unbondings": [
            {
                "amount": "1000000",
                "recipient": "cosmos1cad0efr25faywnjp8qp36l8zlqa2sgz0jwn0hl"
            },
            {
                "amount": "2000000",
                "recipient": "cosmos13mwxtgrljf9d5r72sc28496ua4lsga0jvmqz8x"
            }
        ]
    }
]
`, version.AppName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			unbondings, err := parseUnbondingsFlags(cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgMigrateUnbondings(
				clientCtx.GetFromAddress().String(),
				args[0],
				unbondings,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagUnbondings, "", "unbondings file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func parseUnbondingsFlags(fs *pflag.FlagSet) ([]*types.PoolUnbond, error) {
	ud := make([]*types.PoolUnbond, 0)
	udFile, err := fs.GetString(FlagUnbondings)
	if err != nil {
		return nil, err
	}
	if len(udFile) == 0 {
		return nil, fmt.Errorf("unbondings json file not give")
	}

	contents, err := os.ReadFile(udFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &ud)
	if err != nil {
		return nil, err
	}

	return ud, nil
}