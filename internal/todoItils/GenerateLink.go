package todoItils

import (
	"github.com/ivadey/debrix-cli/internal/dbUtils"
	"github.com/ivadey/debrix-cli/internal/utils"
	"regexp"
	"strconv"
)

func GenerateLink(todoItem dbUtils.TodoItem, config *utils.Config) string {
	res := string(filePathPattern.ReplaceAllFunc([]byte(config.LinkTemplate), func(entry []byte) []byte {
		return []byte(todoItem.RelativePath)
	}))

	res = string(lineNumberPattern.ReplaceAllFunc([]byte(res), func(entry []byte) []byte {
		return []byte(strconv.Itoa(int(todoItem.Line)))
	}))

	return res
}

var filePathPattern = regexp.MustCompile("{{filePath}}")
var lineNumberPattern = regexp.MustCompile("{{lineNumber}}")
