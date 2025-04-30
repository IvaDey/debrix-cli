package utils

import (
	"fmt"
	ignore "github.com/sabhiram/go-gitignore"
	"os"
	"path/filepath"
)

func GetFilesToInspect(atPath string, exclude []string) []string {
	res := make([]string, 0)

	entries, err := os.ReadDir(atPath)
	if err != nil {
		panic(fmt.Sprintf("error reading dir %s: %s", atPath, err.Error()))
	}

	matcher := ignore.CompileIgnoreLines(exclude...)
	workDir, _ := os.Getwd()

	for _, entry := range entries {
		relativePath, _ := filepath.Rel(workDir, filepath.Join(atPath, entry.Name()))
		if matcher.MatchesPath(relativePath) {
			continue
		}

		if entry.IsDir() {
			res = append(res, GetFilesToInspect(filepath.Join(atPath, entry.Name()), exclude)...)
		} else {
			res = append(res, filepath.Join(atPath, entry.Name()))
		}
	}

	return res
}
