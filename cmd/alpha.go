package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var alphaCmd = &cobra.Command{
	Use:   "alpha",
	Short: "All the alphabet related commands",
	Run:   runAlpha,
}

func init() {
	rootCmd.AddCommand(alphaCmd)
}

func runAlpha(cmd *cobra.Command, args []string) {
	fmt.Println("inside alpha, config value: ", viper.GetString("config"))
}
