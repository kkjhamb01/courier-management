package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	_ "github.com/kkjhamb01/courier-management/announcement"
	_ "github.com/kkjhamb01/courier-management/delivery"
	_ "github.com/kkjhamb01/courier-management/finance"
	_ "github.com/kkjhamb01/courier-management/notification"
	_ "github.com/kkjhamb01/courier-management/offering"
	_ "github.com/kkjhamb01/courier-management/party"
	_ "github.com/kkjhamb01/courier-management/pricing"
	_ "github.com/kkjhamb01/courier-management/promotion"
	_ "github.com/kkjhamb01/courier-management/rating"
	_ "github.com/kkjhamb01/courier-management/uaa"
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
