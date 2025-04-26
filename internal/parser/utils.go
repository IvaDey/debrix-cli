package parser

import (
	"fmt"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/bash"
	"github.com/smacker/go-tree-sitter/css"
	"github.com/smacker/go-tree-sitter/html"
	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/smacker/go-tree-sitter/lua"
	"github.com/smacker/go-tree-sitter/php"
	"github.com/smacker/go-tree-sitter/python"
	"github.com/smacker/go-tree-sitter/ruby"
	"github.com/smacker/go-tree-sitter/rust"
	"github.com/smacker/go-tree-sitter/yaml"
	"strings"
)

func cleanComment(node *sitter.Node, source []byte) string {
	comment := node.Content(source)

	if strings.HasPrefix(comment, "//") {
		comment = comment[2:]
	} else if strings.HasPrefix(comment, "/*") {
		comment = comment[2 : len(comment)-2]
	} else if strings.HasPrefix(comment, "#") {
		comment = comment[1:]
	} else if strings.HasPrefix(comment, "<!--") {
		comment = comment[4 : len(comment)-3]
	} else if strings.HasPrefix(comment, "--[[") {
		comment = comment[4 : len(comment)-4]
	} else if strings.HasPrefix(comment, "--") {
		comment = comment[2:]
	} else if strings.HasPrefix(comment, "=begin") {
		comment = comment[6 : len(comment)-5]
	}

	return strings.TrimSpace(comment)
}

func getLanguage(fileExt string) (*sitter.Language, error) {
	switch fileExt {
	// JavaScript, TypeScript, Go, Swift, Objective-C, Dart, java, kotlin, c/c++, c#
	case ".js", ".mjs", ".cjs", ".ts", ".mts", ".cts", ".jsx", ".tsx", ".go", ".swift", ".dart",
		".c", ".cpp", ".h", ".hpp", ".cc", ".hh", ".m", ".java", ".kt", ".kts", ".cs", ".csx",
		".scala":
		return javascript.GetLanguage(), nil

	// Python, Elixir, Perl
	case ".py", ".ex", ".exs", ".pl", ".pm", ".t":
		return python.GetLanguage(), nil

	// HTML
	case ".html", ".htm", ".xml":
		return html.GetLanguage(), nil

	// CSS
	case ".css":
		return css.GetLanguage(), nil

	// 	Bash
	case ".sh":
		return bash.GetLanguage(), nil

	// 	Lua
	case ".lua":
		return lua.GetLanguage(), nil

	// 	PHP
	case ".php", ".phptml", ".phps", ".php4", ".php5":
		return php.GetLanguage(), nil

	// 	YAML
	case ".yml":
		return yaml.GetLanguage(), nil

	// 	Ruby
	case ".rb", ".ru", ".ruby":
		return ruby.GetLanguage(), nil

	// 	Rust
	case ".rs":
		return rust.GetLanguage(), nil

	default:
		return nil, fmt.Errorf("Files with extension '%s' not yet supported ", fileExt)
	}
}
