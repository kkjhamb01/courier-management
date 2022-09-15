package cmd

import (
	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/spf13/cobra"
)

var OfferingCmd = &cobra.Command{
	Use:   service.Offering,
	Short: "Offering Service",
	Run: func(cmd *cobra.Command, args []string) {
		service.FilterRegisteredServices(func(s service.Service) bool {
			if s.Name() != service.Offering {
				return true
			}

			return false
		})
	},
}
