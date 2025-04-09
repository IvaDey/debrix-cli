package todos

import (
	"regexp"
	"strings"
)

func processTodo(text string, pattern *regexp.Regexp) TodoInfo {
	todoLocation := pattern.FindIndex([]byte(text))
	res := text[todoLocation[1]:]

	scope := ""
	res = string(scopePattern.ReplaceAllFunc([]byte(res), func(entry []byte) []byte {
		if scope == "" {
			scope = string(entry)[1 : len(entry)-1]
			return []byte("")
		}
		return entry
	}))

	assignee := ""
	due := ""
	reminder := ""
	res = string(annotationPattern.ReplaceAllFunc([]byte(res), func(entry []byte) []byte {
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

	if strings.HasPrefix(res, ":") {
		res = strings.Replace(res, ":", "", 1)
		res = strings.TrimSpace(res)
	}

	return TodoInfo{
		RelativePath: "",
		FileName:     "",
		Line:         0,
		Task:         res,
		Scope:        scope,
		Assignee:     assignee,
		Due:          due,
		Reminder:     reminder,
	}
}

type TodoInfo struct {
	RelativePath string
	FileName     string
	Line         uint32
	Task         string
	Scope        string
	Assignee     string
	Due          string
	Reminder     string
}

var scopePattern = regexp.MustCompile(`^(\([^)]+\))`)
var annotationPattern = regexp.MustCompile(`\[!(\w+):([^]]+)]`)
