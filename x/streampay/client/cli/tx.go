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
			feeStr, err := cmd.Flags().GetString(FlagStreamPaymentFee)
			if err != nil {
				return err
			}
			fee, err := sdk.ParseCoinNormalized(feeStr)
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
				fee,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsStreamSend)
	_ = cmd.MarkFlagRequired(FlagDuration)
	_ = cmd.MarkFlagRequired(FlagStreamPaymentFee)
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

func GetCmdClaimStreamedAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "claim",
		Long: "claim streamed amount",
		Example: fmt.Sprintf(
			"$ %s tx streampay claim [stream-id]",
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
