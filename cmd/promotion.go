package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.artin.ai/backend/courier-management/common/service"
)

var PromotionCmd = &cobra.Command{
	Use:   service.Promotion,
	Short: "Promotion Service",
	Run: func(cmd *cobra.Command, args []string) {
		service.FilterRegisteredServices(func(s service.Service) bool {
			if s.Name() != service.Promotion {
				return true
			}

			return false
		})
	},
}
