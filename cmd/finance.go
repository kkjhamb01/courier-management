package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.artin.ai/backend/courier-management/common/service"
)

var FinanceCmd = &cobra.Command{
	Use:   service.Finance,
	Short: "Finance Service",
	Run: func(cmd *cobra.Command, args []string) {
		service.FilterRegisteredServices(func(s service.Service) bool {
			if s.Name() != service.Finance {
				return true
			}

			return false
		})
	},
}
