package mdGenerator

import (
	"fmt"
	"github.com/ivadey/debrix-cli/internal/todos"
	"strings"
)

func Generate(lang string, todosInfo []todos.TodoInfo, layout string) string {
	localizedLabels := getLabels(lang)

	var res string
	withDueDate := 0
	withReminder := 0
	withAssignee := 0
	unscopedTodos := make([]todos.TodoInfo, 0)
	scopedTodos := make(map[string][]todos.TodoInfo)

	for _, todoInfo := range todosInfo {
		if todoInfo.Due != "" {
			withDueDate += 1
		}

		if todoInfo.Reminder != "" {
			withReminder += 1
		}

		if todoInfo.Assignee != "" {
			withAssignee += 1
		}

		if todoInfo.Scope == "" {
			unscopedTodos = append(unscopedTodos, todoInfo)
		} else {
			scopedTodos[todoInfo.Scope] = append(scopedTodos[todoInfo.Scope], todoInfo)
		}
	}

	header := generateHeader(lang, overview{
		Total:        len(todosInfo),
		Opened:       len(todosInfo),
		Done:         0,
		WithDeadline: withDueDate,
		WithReminder: withReminder,
		WithAssignee: withAssignee,
	})

	if len(unscopedTodos) > 0 {
		res += fmt.Sprintf(
			"## 🗂 %s\n\n%s\n\n",
			localizedLabels.general,
			generateTodosLayout(lang, unscopedTodos, layout),
		)
	}
	for scope, todoItems := range scopedTodos {
		res += fmt.Sprintf(
			"## 🧩 %s: "+scope+"\n\n%s\n\n",
			localizedLabels.scope,
			generateTodosLayout(lang, todoItems, layout),
		)
	}

	return header + "\n\n" + res
}

func generateTodosLayout(lang string, todosInfo []todos.TodoInfo, layout string) string {
	switch layout {
	case "table":
		return generateTable(lang, todosInfo)
	default:
		items := make([]string, 0)
		for _, todoInfo := range todosInfo {
			items = append(items, generateTodoItem(lang, todoInfo))
		}
		return strings.Join(items, "\n")
	}
}
