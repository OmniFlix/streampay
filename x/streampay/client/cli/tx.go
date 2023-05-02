package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"strconv"
	"time"

	"github.com/OmniFlix/streampay/x/streampay/types"
	"github.com/cosmos/cosmos-sdk/client"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	streamPaymentTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	streamPaymentTxCmd.AddCommand(
		GetCmdStreamSend(),
		GetCmdStopStream(),
	)

	return streamPaymentTxCmd
}

// GetCmdStreamSend implements the stream-send command
func GetCmdStreamSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "stream-send",
		Long: "creates a stream payment",
		Example: fmt.Sprintf(
			"$ %s tx streampay stream-send [recipient] [amount] --end-time <end-timestamp> ",
			version.AppName,
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()
			recipient, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			endTime, err := cmd.Flags().GetString(FlagEndTime)
			if err != nil {
				return err
			}
			if endTime == "" {
				return fmt.Errorf("endtime is required")
			}
			endTimestamp, err := strconv.ParseInt(endTime, 10, 64)
			if err != nil {
				return err
			}
			etm := time.Unix(endTimestamp, 0)
			if etm.Unix() <= time.Now().Unix() {
				return fmt.Errorf("endtime should be in future")
			}
			delayed, err := cmd.Flags().GetBool(FlagDelayed)
			if err != nil {
				return err
			}
			_type := types.TypeContinuous
			if delayed {
				_type = types.TypeDelayed
			}

			msg := types.NewMsgStreamSend(sender.String(), recipient.String(), amount, _type, etm)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsStreamSend)
	_ = cmd.MarkFlagRequired(FlagEndTime)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdStopStream implements the stop-stream command
func GetCmdStopStream() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "stop-stream",
		Long: "stops a stream payment",
		Example: fmt.Sprintf(
			"$ %s tx streampay stop-stream [stream-id]",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := clientCtx.GetFromAddress()
			msg := types.NewMsgStopStream(args[0], sender.String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
