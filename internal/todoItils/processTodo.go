package todoItils

import (
	"github.com/ivadey/debrix-cli/internal/dbUtils"
	"regexp"
	"strings"
)

func processTodo(text string, pattern *regexp.Regexp) dbUtils.TodoItem {
	todoLocation := pattern.FindIndex([]byte(text))
	res := text[todoLocation[1]:]

	var scope string
	res, scope = extractScope(res)

	var annotations TodoAnnotations
	res, annotations = extractAnnotations(res)

	if strings.HasPrefix(res, ":") {
		res = strings.Replace(res, ":", "", 1)
		res = strings.TrimSpace(res)
	}

	return dbUtils.TodoItem{
		RelativePath: "",
		FileName:     "",
		Line:         0,
		Task:         res,
		Scope:        scope,
		Assignee:     annotations.Assignee,
		Due:          annotations.Due,
		Reminder:     annotations.Reminder,

		IsCompleted: false,
	}
}

func extractScope(todoComment string) (string, string) {
	scope := ""
	res := string(scopePattern.ReplaceAllFunc([]byte(todoComment), func(entry []byte) []byte {
		if scope == "" {
			scope = string(entry)[1 : len(entry)-1]
			return []byte("")
		}
		return entry
	}))

	return res, scope
}

func extractAnnotations(todoComment string) (string, TodoAnnotations) {
	assignee := ""
	due := ""
	reminder := ""

	filteredText := string(annotationPattern.ReplaceAllFunc([]byte(todoComment), func(entry []byte) []byte {
		match := annotationPattern.FindStringSubmatch(string(entry))
		name, value := match[1], match[2]

		if name == "assign" {
			assignee = value
			return []byte("")
		} else if name == "due" {
			due = value // todo: validate value (must be date in format YYYY-MM-DD)
			return []byte("")
		} else if name == "remind" {
			reminder = value // todo: validate value (must be date in format YYYY-MM-DD)
			return []byte("")
		}

		return entry
	}))

	return filteredText, TodoAnnotations{assignee, due, reminder}
}

type TodoAnnotations struct {
	Assignee string
	Due      string
	Reminder string
}

var scopePattern = regexp.MustCompile(`^(\([^)]+\))`)
var annotationPattern = regexp.MustCompile(`\[!(\w+):([^]]+)]`)
