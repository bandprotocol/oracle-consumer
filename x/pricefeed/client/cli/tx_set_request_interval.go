package cli

import (
	"fmt"
	"strconv"

	"consumer/x/pricefeed/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSetRequestInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-request-interval [symbol] [oracle-script-id] [block-interval]",
		Short: "Broadcast message SetRequestInterval",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOracleScriptId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("unable to parse request ID: %w", err)
			}

			argBlockInterval, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return fmt.Errorf("unable to parse request ID: %w", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetRequestInterval(
				args[0],
				argOracleScriptId,
				argBlockInterval,
				clientCtx.GetFromAddress(),
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
