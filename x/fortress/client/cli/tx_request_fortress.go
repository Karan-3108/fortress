package cli

import (
	"strconv"

	"github.com/Karan-3108/Fortress/x/fortress/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRequestFortress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-fortress [amount] [fee] [collateral] [deadline]",
		Short: "Broadcast message request-fortress",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argFee := args[1]
			argCollateral := args[2]
			argDeadline := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestFortress(
				clientCtx.GetFromAddress().String(),
				argAmount,
				argFee,
				argCollateral,
				argDeadline,
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
