package utils

import (
	"strings"
)

func EscapeForMarkdown(text string) string {
	replacer := strings.NewReplacer(
		`*`, `\*`,
		`_`, `\_`,
		`#`, `\#`,
		`>`, `\>`,
		`-`, `\-`,
		`+`, `\+`,
		`=`, `\=`,
		`|`, `\|`,
		`~`, `\~`,
		"`", "\\`",
		`[`, `\[`,
		`]`, `\]`,
		`(`, `\(`,
		`)`, `\)`,
		`{`, `\{`,
		`}`, `\}`,
		`<`, `&lt;`,
		`>`, `&gt;`,
		`\`, `\\`,
	)
	return replacer.Replace(text)
}
