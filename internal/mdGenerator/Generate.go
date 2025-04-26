package mdGenerator

import (
	"fmt"
	"github.com/ivadey/debrix-cli/internal/dbUtils"
	"github.com/ivadey/debrix-cli/internal/utils"
	"strings"
)

func Generate(todoItems []dbUtils.TodoItem, config *utils.Config) string {
	lang := config.Language
	layout := config.Layout

	localizedLabels := getLabels(lang)

	var res string
	opened := len(todoItems)
	completed := 0
	withDueDate := 0
	withReminder := 0
	withAssignee := 0
	unscopedTodos := make([]dbUtils.TodoItem, 0)
	scopedTodos := make(map[string][]dbUtils.TodoItem)

	for _, todoItem := range todoItems {
		if todoItem.IsCompleted {
			opened--
			completed++
		}

		if todoItem.Due != "" {
			withDueDate += 1
		}

		if todoItem.Reminder != "" {
			withReminder += 1
		}

		if todoItem.Assignee != "" {
			withAssignee += 1
		}

		if todoItem.Scope == "" {
			unscopedTodos = append(unscopedTodos, todoItem)
		} else {
			scopedTodos[todoItem.Scope] = append(scopedTodos[todoItem.Scope], todoItem)
		}
	}

	header := generateHeader(lang, overview{
		Total:        len(todoItems),
		Opened:       opened,
		Done:         completed,
		WithDeadline: withDueDate,
		WithReminder: withReminder,
		WithAssignee: withAssignee,
	})

	if len(unscopedTodos) > 0 {
		res += fmt.Sprintf(
			"## 🗂 %s\n\n%s\n\n",
			localizedLabels.general,
			generateTodosLayout(lang, unscopedTodos, layout, config),
		)
	}
	for scope, items := range scopedTodos {
		res += fmt.Sprintf(
			"## 🧩 %s: "+scope+"\n\n%s\n\n",
			localizedLabels.scope,
			generateTodosLayout(lang, items, layout, config),
		)
	}

	return header + "\n\n" + res
}

func generateTodosLayout(lang string, todosInfo []dbUtils.TodoItem, layout string, config *utils.Config) string {
	switch layout {
	case "table":
		return generateTable(lang, todosInfo, config)
	default:
		items := make([]string, 0)
		for _, todoItem := range todosInfo {
			items = append(items, generateTodoItem(lang, todoItem, config))
		}
		return strings.Join(items, "\n")
	}
}
