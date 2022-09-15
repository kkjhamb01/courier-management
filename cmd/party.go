package cmd

import (
	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/spf13/cobra"
)

var PartyCmd = &cobra.Command{
	Use:   service.Party,
	Short: "Party Service",
	Run: func(cmd *cobra.Command, args []string) {
		service.FilterRegisteredServices(func(s service.Service) bool {
			if s.Name() != service.Party {
				return true
			}

			return false
		})
	},
}
