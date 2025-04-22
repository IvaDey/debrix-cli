package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the Version number of Debrix tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
