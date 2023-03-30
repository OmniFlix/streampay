package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/OmniFlix/payment-stream/x/streampay/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group marketplace queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(
		GetCmdQueryPaymentStreams(),
	)

	return cmd
}

// GetCmdQueryPaymentStreams implements the query payment streams command.
func GetCmdQueryPaymentStreams() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "stream-payments ",
		Long:    "Query stream payments.",
		Example: fmt.Sprintf("$ %s query streampay stream-payments", version.AppName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.StreamingPayments(
				context.Background(),
				&types.QueryStreamPaymentsRequest{
					Pagination: pageReq,
				})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryStreamPayment implements the query stream payment command.
func GetCmdQueryStreamPayment() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "stream-payment",
		Long:    "Query stream payment.",
		Example: fmt.Sprintf("$ %s query steampay stream-payment <id>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())

			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			if err != nil {
				return err
			}

			res, err := queryClient.StreamingPayment(
				context.Background(),
				&types.QueryStreamPaymentRequest{
					Id: args[0],
				})

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}