package todos

import (
	"github.com/ivadey/debrix-cli/internal/parser"
	"path/filepath"
	"regexp"
	"strings"
)

func Collect(basePath string, path string, pattern *regexp.Regexp) []TodoInfo {
	comments, err := parser.GetComments(path)
	if err != nil {
		return nil
	}

	todos := make([]TodoInfo, 0)

	for _, comment := range comments {
		for index, line := range strings.Split(comment.Text, "\n") {
			if pattern.MatchString(line) {
				todoInfo := processTodo(line, pattern)
				todoInfo.RelativePath, _ = filepath.Rel(basePath, path)
				todoInfo.FileName = filepath.Base(path)
				todoInfo.Line = comment.Line + uint32(index)

				todos = append(todos, todoInfo)
			}
		}
	}

	return todos
}
