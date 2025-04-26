package todoItils

import (
	"github.com/ivadey/debrix-cli/internal/dbUtils"
	"github.com/ivadey/debrix-cli/internal/parser"
	"github.com/ivadey/debrix-cli/internal/utils"
	"path/filepath"
	"regexp"
	"strings"
)

func Collect(basePath string, path string, pattern *regexp.Regexp) []dbUtils.TodoItem {
	comments, err := parser.GetComments(path)
	if err != nil {
		return nil
	}

	todos := make([]dbUtils.TodoItem, 0)

	db := dbUtils.OpenDb()
	var storedItems []dbUtils.StoredTodo
	db.Where("file_name = ? and is_completed = false", filepath.Base(path)).Find(&storedItems)

	for _, comment := range comments {
		for index, line := range strings.Split(comment.Text, "\n") {
			if pattern.MatchString(line) {
				todoItem := processTodo(line, pattern)
				todoItem.RelativePath, _ = filepath.Rel(basePath, path)
				todoItem.FileName = filepath.Base(path)
				todoItem.Line = comment.Line + uint32(index)

				author, _ := utils.GetAuthorForLine(path, todoItem.Line)
				todoItem.Author = author

				matchedIndex := findIndex(todoItem.Task, &storedItems)
				if matchedIndex == -1 {
					db.Create(&dbUtils.StoredTodo{TodoItem: todoItem})
				} else {
					storedItem := storedItems[matchedIndex]
					changes := getChanges(todoItem, storedItem)

					if len(changes) > 0 {
						db.Model(&storedItem).Updates(changes)
					}
					removeAt(&storedItems, matchedIndex)
				}

				todos = append(todos, todoItem)
			}
		}
	}

	for _, item := range storedItems {
		db.Model(&item).Update("is_completed", true)
	}

	return todos
}

func getChanges(todoItem dbUtils.TodoItem, storedItem dbUtils.StoredTodo) map[string]interface{} {
	changes := make(map[string]interface{})

	if storedItem.Line != todoItem.Line {
		changes["line"] = todoItem.Line
	}
	if storedItem.Scope != todoItem.Scope {
		changes["scope"] = todoItem.Scope
	}
	if storedItem.Due != todoItem.Due {
		changes["due"] = todoItem.Due
	}
	if storedItem.Reminder != todoItem.Reminder {
		changes["reminder"] = todoItem.Reminder
	}
	if storedItem.Assignee != todoItem.Assignee {
		changes["assignee"] = todoItem.Assignee
	}

	return changes
}

func findIndex(text string, items *[]dbUtils.StoredTodo) int {
	for index, item := range *items {
		if item.Task == text {
			return index
		}
	}

	return -1
}

func removeAt(items *[]dbUtils.StoredTodo, index int) {
	if index < 0 || index >= len(*items) {
		return
	}

	*items = append((*items)[:index], (*items)[index+1:]...)
}
