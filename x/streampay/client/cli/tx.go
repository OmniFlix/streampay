package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/OmniFlix/streampay/v2/x/streampay/types"
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
		GetCmdClaimStreamedAmount(),
	)

	return streamPaymentTxCmd
}

// GetCmdStreamSend implements the stream-send command
func GetCmdStreamSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stream-send",
		Short: "creates a stream payment",
		Long: "creates a stream payment with the given recipient and amount.\n" +
			"stream payment will start immediately aftet the successful execution of this command.\n" +
			"types of stream-payments:\n" +
			"1. continuous: stream payment will be sent continuously until the end time is reached\n" +
			"2. delayed: stream payment will be sent after the end time is reached (use --delayed flag)\n" +
			"3. periodic: stream payment will be sent periodically until the end time is reached" +
			" (use --stream-periods-file)\n\n" +
			"example stream-periods-file:\n" +
			"[\n" +
			"	{\n" +
			"		\"amount\": 100000,\n" +
			"		\"duration\": 60\n" +
			"	},\n" +
			"	{\n" +
			"		\"amount\": 100000,\n" +
			"		\"duration\": 60\n" +
			"	},\n" +
			"]\n\n" +
			"cancellation of stream payments:\n" +
			"using --cancellable flag, you can create a cancellable stream payment\n",
		Example: fmt.Sprintf(
			"$ %s tx streampay stream-send [recipient] [amount] --payment-fee [amount] --duration <duration>"+
				" --chain-id <chain-id> --from <sender> --fees <fees>\n\n"+
				"delayed payment:\n"+
				"$ %s tx streampay stream-send [recipient] [amount] --payment-fee [amount]--duration <stream-duration> --delayed"+
				" --chain-id <chain-id> --from <sender> --fees <fees>\n\n"+
				"periodic payment:\n"+
				"$ %s tx streampay stream-send [recipient] [amount]  --payment-fee [amount] --stream-periods-file <stream-periods-file>"+
				" --chain-id <chain-id> --from <sender> --fees <fees>\n\n",
			version.AppName, version.AppName, version.AppName,
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
			durationStr, err := cmd.Flags().GetString(FlagDuration)
			if err != nil {
				return err
			}
			duration, err := time.ParseDuration(durationStr)
			if err != nil {
				return err
			}
			if duration <= 0 {
				return fmt.Errorf("duration should be a positive value")
			}
			delayed, err := cmd.Flags().GetBool(FlagDelayed)
			if err != nil {
				return err
			}
			streamPeriodsFile, err := cmd.Flags().GetString(FlagStreamPeriodsFile)
			if err != nil {
				return err
			}
			_type := types.TypeContinuous
			var periods []*types.Period
			if delayed {
				_type = types.TypeDelayed
			} else if streamPeriodsFile != "" {
				_type = types.TypePeriodic
				periods, err = parsePeriods(streamPeriodsFile)
				if err != nil {
					return err
				}
			} else {
				periods = nil
			}
			cancellable, err := cmd.Flags().GetBool(FlagCancellable)
			if err != nil {
				return err
			}
			paymentFeeStr, err := cmd.Flags().GetString(FlagPaymentFee)
			if err != nil {
				return err
			}
			paymentFee, err := sdk.ParseCoinNormalized(paymentFeeStr)
			if err != nil {
				return err
			}

			msg := types.NewMsgStreamSend(
				sender.String(),
				recipient.String(),
				amount,
				_type,
				duration,
				periods,
				cancellable,
				paymentFee,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsStreamSend)
	_ = cmd.MarkFlagRequired(FlagDuration)
	_ = cmd.MarkFlagRequired(FlagPaymentFee)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdStopStream implements the stop-stream command
func GetCmdStopStream() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop-stream",
		Short: "stops a stream payment",
		Long: "stops a stream payment with the given stream id if stream-payment is cancellable.\n" +
			"if stream-payment is not cancellable, this command will fail.\n" +
			"if payment is in progress only remaining amount will be refunded to the sender.\n" +
			"streamed amount can be claimed using claim command.\n",
		Example: fmt.Sprintf(
			"$ %s tx streampay stop-stream [stream-id] --chain-id <chain-id> --from <sender> --fees <fees>",
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

func GetCmdClaimStreamedAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim",
		Short: "claim streamed amount",
		Long:  "claim streamed amount for the given stream id.\n",
		Example: fmt.Sprintf(
			"$ %s tx streampay claim [stream-id] --chain-id <chain-id> --from <sender> --fees <fees>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			claimer := clientCtx.GetFromAddress()
			msg := types.NewMsgClaimStreamedAmount(args[0], claimer.String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func parsePeriods(filePath string) ([]*types.Period, error) {
	var periods []*types.Period
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &periods)
	if err != nil {
		return nil, err
	}
	return periods, nil
}
