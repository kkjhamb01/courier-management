package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.artin.ai/backend/courier-management/common/service"
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
