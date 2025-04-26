package cmd

import (
	"github.com/ivadey/debrix-cli/internal/dbUtils"
	"github.com/ivadey/debrix-cli/internal/mdGenerator"
	"github.com/ivadey/debrix-cli/internal/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(cleanupCmd)
}

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Remove all completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		dbUtils.RemoveCompleted()
		config := utils.GetConfig()

		storedData := *dbUtils.FetchAll()
		todoItems := make([]dbUtils.TodoItem, len(storedData))
		for index, item := range storedData {
			todoItems[index] = item.TodoItem
		}
		res := mdGenerator.Generate(todoItems, config)

		err = os.WriteFile(filepath.Join(workDir, config.OutFile), []byte(res), 0644)
		if err != nil {
			panic(err)
		}
	},
}
