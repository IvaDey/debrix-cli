package mdGenerator

import (
	"fmt"
	"github.com/ivadey/debrix-cli/internal/todos"
	"strings"
)

func generateTable(lang string, todosInfo []todos.TodoInfo) string {
	localizedLabels := getLabels(lang)

	res := fmt.Sprintf(
		"| 📝 TODO | ⏳ %s | 🔔 %s | 👤 %s | 📄 %s |\n",
		localizedLabels.deadline,
		localizedLabels.reminder,
		localizedLabels.assignee,
		localizedLabels.file,
	)
	res += "|---|---|---|---|---|\n"

	for _, todoInfo := range todosInfo {
		todo := strings.ReplaceAll(todoInfo.Task, "\n", "<br>")
		todo = strings.ReplaceAll(todo, "|", "\\|")

		deadline := todoInfo.Due
		if deadline == "" {
			deadline = "-"
		}

		reminder := todoInfo.Reminder
		if reminder == "" {
			reminder = "-"
		}

		assignee := todoInfo.Assignee
		if assignee == "" {
			assignee = "-"
		}

		file := fmt.Sprintf(
			"[%s:%d](%s#L%d)",
			todoInfo.FileName,
			todoInfo.Line,
			todoInfo.RelativePath,
			todoInfo.Line,
		)

		res += fmt.Sprintf(
			"|%s|%s|%s|%s|%s|\n",
			strings.ReplaceAll(todoInfo.Task, "\n", "<br>"),
			deadline,
			reminder,
			assignee,
			file,
		)
	}

	return res
}
