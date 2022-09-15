package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.artin.ai/backend/courier-management/common/service"
)

var UaaCmd = &cobra.Command{
	Use:   service.Uaa,
	Short: "UAA Service",
	Run: func(cmd *cobra.Command, args []string) {
		service.FilterRegisteredServices(func(s service.Service) bool {
			if s.Name() != service.Uaa {
				return true
			}

			return false
		})
	},
}
