package cmd

import (
	"github.com/felicianotech/ubuntu-remote-server/api"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	portFl uint16

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Runs the server",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

			app := api.App{}
			app.Initialize()

			log.Info("Server initialized.")
			log.Info("Starting the server...")

			app.Run(":2004")
		},
	}
)

func init() {
	runCmd.Flags().Uint16Var(&portFl, "port", 2004, "port to listen to")

	rootCmd.AddCommand(runCmd)
}
