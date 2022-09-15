package cmd

import (
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/spf13/cobra"
)

var startServiceNames string

// to capture interrupt signal and stop the running service(s)
var signals chan os.Signal

func init() {
	StartCmd.AddCommand(OfferingCmd)
	StartCmd.AddCommand(UaaCmd)
	StartCmd.AddCommand(PartyCmd)
	StartCmd.AddCommand(FinanceCmd)
	StartCmd.AddCommand(NotificationCmd)
	StartCmd.AddCommand(delivery)
	StartCmd.AddCommand(pricing)
	StartCmd.AddCommand(RatingCmd)
	StartCmd.AddCommand(PromotionCmd)
	StartCmd.AddCommand(AnnouncementCmd)
	StartCmd.Flags().StringVarP(&startServiceNames, "names", "n", "", "To start multiple services separated by a comma")

	// stop the running service(s) on interrupt signal
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals
		logger.Info("Interrupt signal received")
		service.StopAllRegisteredServices()
		signal.Stop(signals)
		os.Exit(0)
	}()
}

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start one or more services of the courier management application",
	Example: os.Args[0] + ` start --names offering,apigateway
` + os.Args[0] + ` start offering`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
		logger.InitLogger()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// filter registered services by names flag
		service.FilterRegisteredServices(func(s service.Service) bool {
			names := strings.Split(startServiceNames, ",")

			for _, name := range names {
				if s.Name() == name {
					return false
				}
			}

			return true
		})
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if len(service.RegisteredServices) == 0 {
			logger.Warning("Wrong service name(s) have been entered, no service is running")
			return
		}

		err := sentry.Init(sentry.ClientOptions{
			Dsn: config.Tracing().DSN,
		})
		if err != nil {
			logger.Fatalf("sentry.Init: %s", err)
		}
		defer sentry.Flush(2 * time.Second)

		service.StartAllRegisteredServices(ctx)
	},
}
