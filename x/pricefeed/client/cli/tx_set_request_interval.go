package cli

import (
	"fmt"
	"strconv"

	"consumer/x/pricefeed/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSetRequestInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-request-interval [symbol] [oracle-script-id] [fee-limit] [block-interval]",
		Short: "Broadcast message SetRequestInterval",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOracleScriptId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("unable to parse request ID: %w", err)
			}

			argFeeLimit, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			argBlockInterval, err := strconv.ParseUint(args[3], 10, 64)
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
				argFeeLimit,
				argBlockInterval,
				clientCtx.GetFromAddress().String(),
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
