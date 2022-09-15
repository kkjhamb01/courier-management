package cmd

import (
	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/spf13/cobra"
)

var delivery = &cobra.Command{
	Use:   service.Delivery,
	Short: "Delivery Service",
	Run: func(cmd *cobra.Command, args []string) {
		service.FilterRegisteredServices(func(s service.Service) bool {
			if s.Name() != service.Delivery {
				return true
			}

			return false
		})
	},
}
