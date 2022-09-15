package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.artin.ai/backend/courier-management/common/service"
)

var pricing = &cobra.Command{
	Use:   service.Pricing,
	Short: "Pricing Service",
	Run: func(cmd *cobra.Command, args []string) {
		service.FilterRegisteredServices(func(s service.Service) bool {
			if s.Name() != service.Pricing {
				return true
			}

			return false
		})
	},
}
