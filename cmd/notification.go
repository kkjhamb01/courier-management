package cmd

import (
	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/spf13/cobra"
)

var NotificationCmd = &cobra.Command{
	Use:   service.Notification,
	Short: "Notification Registration Service",
	Run: func(cmd *cobra.Command, args []string) {
		service.FilterRegisteredServices(func(s service.Service) bool {
			if s.Name() != service.Notification {
				return true
			}

			return false
		})
	},
}
