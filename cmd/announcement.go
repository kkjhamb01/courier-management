package cmd

import (
	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/spf13/cobra"
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
