package cmd

import (
	"fmt"
	"os"

	"github.com/pixlcrashr/roomy/pkg/api"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Roomy server",
	Long:  `Start the Roomy HTTP server to serve the API and web interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := api.NewServer(nil)

		fmt.Printf("Starting server on %s\n", config.Server.Address)
		if err := server.Listen(config.Server.Address); err != nil {
			fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
