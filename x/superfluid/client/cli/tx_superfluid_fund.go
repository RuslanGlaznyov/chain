package cli

import (
	"github.com/spf13/cast"
	"strconv"

	"github.com/KYVENetwork/chain/x/superfluid/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSuperfluidFund() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fund [pool-id] [val-addr] [amount]",
		Short: "Broadcast message superfluid-fund",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPoolId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			argValAddr := cast.ToString(args[1])

			argAmount, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSuperfluidFund(
				clientCtx.GetFromAddress().String(),
				argPoolId,
				argValAddr,
				argAmount,
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
