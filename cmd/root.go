package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"

	_ "gitlab.artin.ai/backend/courier-management/announcement"
	_ "gitlab.artin.ai/backend/courier-management/delivery"
	_ "gitlab.artin.ai/backend/courier-management/finance"
	_ "gitlab.artin.ai/backend/courier-management/notification"
	_ "gitlab.artin.ai/backend/courier-management/offering"
	_ "gitlab.artin.ai/backend/courier-management/party"
	_ "gitlab.artin.ai/backend/courier-management/pricing"
	_ "gitlab.artin.ai/backend/courier-management/promotion"
	_ "gitlab.artin.ai/backend/courier-management/rating"
	_ "gitlab.artin.ai/backend/courier-management/uaa"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
)

func init() {
	RootCmd.SetHelpCommand(helpCommand)
	RootCmd.AddCommand(StartCmd)
}

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Courier Manager",
	Run: func(cmd *cobra.Command, args []string) {
		helpCommand.Help()
	},
}

func Execute() {
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	if err := RootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
