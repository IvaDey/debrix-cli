package mdGenerator

import (
	"fmt"
	"github.com/ivadey/debrix-cli/internal/dbUtils"
	"github.com/ivadey/debrix-cli/internal/todoItils"
	"github.com/ivadey/debrix-cli/internal/utils"
	"strings"
)

func generateTodoItem(lang string, todoItem dbUtils.TodoItem, config *utils.Config) string {
	localizedLabels := getLabels(lang)

	mark := " "
	if todoItem.IsCompleted {
		mark = "x"
	}

	res := fmt.Sprintf(
		"- [%s] [%s:%d](%s) %s",
		mark,
		todoItem.FileName,
		todoItem.Line,
		todoItils.GenerateLink(todoItem, config),
		utils.EscapeForMarkdown(todoItem.Task),
	)
	annotations := make([]string, 0)

	if todoItem.Due != "" {
		annotations = append(annotations, fmt.Sprintf("⏰ %s: %v", localizedLabels.deadline, utils.EscapeForMarkdown(todoItem.Due)))
	}

	if todoItem.Reminder != "" {
		annotations = append(annotations, fmt.Sprintf("🔔 %s: %v", localizedLabels.reminder, utils.EscapeForMarkdown(todoItem.Reminder)))
	}

	if todoItem.Assignee != "" {
		annotations = append(annotations, fmt.Sprintf("👤 %s: %v", localizedLabels.assignee, utils.EscapeForMarkdown(todoItem.Assignee)))
	}

	if len(annotations) > 0 {
		res += "<br>\n  " + strings.Join(annotations, " | ")
	}

	return res
}
