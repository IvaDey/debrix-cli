package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "debrix",
	Short: "Debrix – CLI tool to manage TODOs in code",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		if err != nil {
			fmt.Println("Error:", err)
		}
		os.Exit(1)
	}
}
