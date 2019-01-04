package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "cobrademo",
	Short: "Demo of the persistent flag not being attached multiple times",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(cobraInit)
	// add config flag
	rootCmd.PersistentFlags().String(
		"config",
		os.ExpandEnv("$HOME/.config"),
		"Path to config file")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

func cobraInit() {
	viper.AutomaticEnv()
}
