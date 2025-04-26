package mdGenerator

import (
	"fmt"
	"github.com/ivadey/debrix-cli/internal/dbUtils"
	"github.com/ivadey/debrix-cli/internal/todoItils"
	"github.com/ivadey/debrix-cli/internal/utils"
	"strings"
)

func generateTable(lang string, todosInfo []dbUtils.TodoItem, config *utils.Config) string {
	localizedLabels := getLabels(lang)

	res := fmt.Sprintf(
		"| 📝 TODO | ⏳ %s | 🔔 %s | 👤 %s | 📄 %s |\n",
		localizedLabels.deadline,
		localizedLabels.reminder,
		localizedLabels.assignee,
		localizedLabels.file,
	)
	res += "|---|---|---|---|---|\n"

	for _, todoItem := range todosInfo {
		todo := strings.ReplaceAll(todoItem.Task, "\n", "<br>")
		todo = strings.ReplaceAll(todo, "|", "\\|")

		deadline := todoItem.Due
		if deadline == "" {
			deadline = "-"
		} else if todoItem.IsCompleted {
			deadline = fmt.Sprintf("~~%s~~", deadline)
		}

		reminder := todoItem.Reminder
		if reminder == "" {
			reminder = "-"
		} else if todoItem.IsCompleted {
			reminder = fmt.Sprintf("~~%s~~", reminder)
		}

		assignee := todoItem.Assignee
		if assignee == "" {
			assignee = "-"
		} else if todoItem.IsCompleted {
			assignee = fmt.Sprintf("~~%s~~", assignee)
		}

		file := fmt.Sprintf(
			"[%s:%d](%s)",
			todoItem.FileName,
			todoItem.Line,
			todoItils.GenerateLink(todoItem, config),
		)
		if todoItem.IsCompleted {
			file = fmt.Sprintf("~~%s~~", file)
		}

		task := strings.ReplaceAll(todoItem.Task, "\n", "<br>")
		if todoItem.IsCompleted {
			task = fmt.Sprintf("~~%s~~", task)
		}

		res += fmt.Sprintf(
			"|%s|%s|%s|%s|%s|\n",
			task,
			deadline,
			reminder,
			assignee,
			file,
		)
	}

	return res
}
