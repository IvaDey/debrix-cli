package mdGenerator

import (
	"fmt"
	"github.com/ivadey/debrix-cli/internal/todos"
	"github.com/ivadey/debrix-cli/internal/utils"
	"strings"
)

func generateTodoItem(lang string, todoInfo todos.TodoInfo, config *utils.Config) string {
	localizedLabels := getLabels(lang)

	res := fmt.Sprintf(
		"- [ ] [%s:%d](%s) %s",
		todoInfo.FileName,
		todoInfo.Line,
		utils.GenerateLink(todoInfo, config),
		todoInfo.Task,
	)
	annotations := make([]string, 0)

	if todoInfo.Due != "" {
		annotations = append(annotations, fmt.Sprintf("⏰ %s: %v", localizedLabels.deadline, todoInfo.Due))
	}

	if todoInfo.Reminder != "" {
		annotations = append(annotations, fmt.Sprintf("🔔 %s: %v", localizedLabels.reminder, todoInfo.Reminder))
	}

	if todoInfo.Assignee != "" {
		annotations = append(annotations, fmt.Sprintf("👤 %s: %v", localizedLabels.assignee, todoInfo.Assignee))
	}

	if len(annotations) > 0 {
		res += "<br>\n  " + strings.Join(annotations, " | ")
	}

	return res
}
