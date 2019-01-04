package cmd

import (
	"github.com/spf13/cobra"
)

var numCmd = &cobra.Command{
	Use:   "num",
	Short: "All the numeric related commands",
}

func init() {
	rootCmd.AddCommand(numCmd)
}
