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
		Use:   "set-request-interval [os-id] [calldata] [block-interval] [ask-count] [min-count] [fee-limit] [prepare-gas] [execute-gas]",
		Short: "Broadcast message SetRequestInterval",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOsId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("unable to parse request ID: %w", err)
			}

			argCalldata, err := cmd.Flags().GetBytesHex(args[1])
			if err != nil {
				return err
			}

			argBlockInterval, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return fmt.Errorf("unable to parse request ID: %w", err)
			}

			argAskCount, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return fmt.Errorf("unable to parse request ID: %w", err)
			}

			argMinCount, err := strconv.ParseUint(args[4], 10, 64)
			if err != nil {
				return fmt.Errorf("unable to parse request ID: %w", err)
			}

			argFeeLimit, err := sdk.ParseCoinsNormalized(args[5])
			if err != nil {
				return err
			}

			argPrepareGas, err := strconv.ParseUint(args[6], 10, 64)
			if err != nil {
				return fmt.Errorf("unable to parse request ID: %w", err)
			}

			argExecuteGas, err := strconv.ParseUint(args[7], 10, 64)
			if err != nil {
				return fmt.Errorf("unable to parse request ID: %w", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetRequestInterval(
				argOsId,
				argCalldata,
				argBlockInterval,
				argAskCount,
				argMinCount,
				argFeeLimit,
				argPrepareGas,
				argExecuteGas,
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
