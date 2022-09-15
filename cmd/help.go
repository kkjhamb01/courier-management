package cmd

import (
	"os"
	"strings"

	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/spf13/cobra"
)

var helpCommand = &cobra.Command{
	Use:   "help",
	Short: "How to use Courier Manager commands",
	Long: `
	To start all the services run ` + os.Args[0] + ` start
	
	To start a specific service run ` + os.Args[0] + ` start --service $service_name

	To start multiple services rune ` + os.Args[0] + ` start --service $service_name1,$service_name2
	
	possible values for $service_name: ` + strings.Join(service.AllAvailableServices(), ","),
}
