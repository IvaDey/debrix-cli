package main

import (
	"fmt"
	"github.com/ivadey/debrix-cli/internal/mdGenerator"
	"github.com/ivadey/debrix-cli/internal/todos"
	"github.com/ivadey/debrix-cli/internal/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := utils.GetConfig(workDir)
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
}
