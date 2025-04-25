package cmd

import (
	"fmt"
	"github.com/ivadey/debrix-cli/internal/mdGenerator"
	"github.com/ivadey/debrix-cli/internal/todos"
	"github.com/ivadey/debrix-cli/internal/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate TODO.md from sources",
	Run: func(cmd *cobra.Command, args []string) {
		workDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		config := utils.GetConfig()
		filesToParse := utils.GetFilesToInspect(workDir, config.Exclude)

		total := len(filesToParse)
		collected := 0
		fmt.Printf("Total amount of files to inspect: %v\n", len(filesToParse))

		todosInfo := make([]todos.TodoInfo, 0)

		var todoPattern = regexp.MustCompile("(?i)(" + strings.Join(config.Pattern, "|") + ")")
		for _, filePath := range filesToParse {
			collected++
			utils.RenderProgressBar(collected, total, 80)

			collectedTodos := todos.Collect(workDir, filePath, todoPattern)

			todosInfo = append(todosInfo, collectedTodos...)
		}
		fmt.Println("")

		var res string = mdGenerator.Generate(todosInfo, config)

		err = os.WriteFile(filepath.Join(workDir, config.OutFile), []byte(res), 0644)
		if err != nil {
			panic(err)
		}
	},
}
