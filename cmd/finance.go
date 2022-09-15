package cmd

import (
	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/spf13/cobra"
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
