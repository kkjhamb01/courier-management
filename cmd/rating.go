package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.artin.ai/backend/courier-management/common/service"
)

var RatingCmd = &cobra.Command{
	Use:   service.Rating,
	Short: "Rating Service",
	Run: func(cmd *cobra.Command, args []string) {
		service.FilterRegisteredServices(func(s service.Service) bool {
			if s.Name() != service.Rating {
				return true
			}

			return false
		})
	},
}
