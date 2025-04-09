package mdGenerator

import (
	"fmt"
	"github.com/ivadey/debrix-cli/internal/todos"
	"strings"
)

func generateTodoItem(lang string, todoInfo todos.TodoInfo) string {
	localizedLabels := getLabels(lang)

	res := fmt.Sprintf(
		"- [ ] [%s:%d](%s#L%d) %s",
		todoInfo.FileName,
		todoInfo.Line,
		todoInfo.RelativePath,
		todoInfo.Line,
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
