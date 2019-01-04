package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var oneCmd = &cobra.Command{
	Use:   "one",
	Short: "first subcommand in numerics",
	Run:   runOne,
}

func init() {
	numCmd.AddCommand(oneCmd)
}

func runOne(cmd *cobra.Command, args []string) {
	fmt.Println("inside one, config value: ", viper.GetString("config"))
}
