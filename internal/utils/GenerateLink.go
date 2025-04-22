package utils

import (
	"github.com/ivadey/debrix-cli/internal/todos"
	"regexp"
	"strconv"
)

func GenerateLink(todoInfo todos.TodoInfo, config *Config) string {
	res := string(filePathPattern.ReplaceAllFunc([]byte(config.LinkTemplate), func(entry []byte) []byte {
		return []byte(todoInfo.RelativePath)
	}))

	res = string(lineNumberPattern.ReplaceAllFunc([]byte(res), func(entry []byte) []byte {
		return []byte(strconv.Itoa(int(todoInfo.Line)))
	}))

	return res
}

var filePathPattern = regexp.MustCompile("{{filePath}}")
var lineNumberPattern = regexp.MustCompile("{{lineNumber}}")
