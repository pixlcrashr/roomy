package cmd

import (
	"os"

	_ "github.com/lib/pq"
	"github.com/pixlcrashr/roomy/pkg/cfg"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	config  *cfg.Config
)

var rootCmd = &cobra.Command{
	Use:   "roomy",
	Short: "Room and place reservation management system",
	Long:  `Roomy is a room and place reservation management system for organizations.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		config, err = cfg.Load(cfgFile)
		if err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default: ./config.yaml)")
}
