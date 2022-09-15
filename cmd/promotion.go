package cmd

import (
	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/spf13/cobra"
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
