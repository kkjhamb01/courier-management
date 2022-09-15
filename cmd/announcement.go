package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.artin.ai/backend/courier-management/common/service"
)

var AnnouncementCmd = &cobra.Command{
	Use:   service.Announcement,
	Short: "Announcement Service",
	Run: func(cmd *cobra.Command, args []string) {
		service.FilterRegisteredServices(func(s service.Service) bool {
			if s.Name() != service.Announcement {
				return true
			}

			return false
		})
	},
}
