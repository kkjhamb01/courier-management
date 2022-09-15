package cmd

import (
	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/spf13/cobra"
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
