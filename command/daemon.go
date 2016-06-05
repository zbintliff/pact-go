package command

import (
	"github.com/pact-foundation/pact-go/daemon"
	"github.com/spf13/cobra"
)

var port int
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Creates a daemon for the Pact DSLs to communicate with",
	Long:  `Creates a daemon for the Pact DSLs to communicate with`,
	Run: func(cmd *cobra.Command, args []string) {
		setLogLevel(verbose, logLevel)

		svc := &daemon.PactMockService{}
		svc.Setup()
		daemon.NewDaemon(svc).StartDaemon(port)
	},
}

func init() {
	daemonCmd.LocalFlags().IntVarP(&port, "port", "p", 6666, "Local daemon port")
	RootCmd.AddCommand(daemonCmd)
}