package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/spellshape/network/x/participation/types"
)

func CmdParticipate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "participate [auction-id] [tier-id]",
		Short: "Add as a participant for the auction with a specified tier benefit",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAuctionID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argTierID, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgParticipate(
				clientCtx.GetFromAddress().String(),
				argAuctionID,
				argTierID,
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
