package utils

import (
	"os"
	"path/filepath"
	"slices"
)

func GetFilesToInspect(atPath string, exclude []string) []string {
	res := make([]string, 0)

	entries, _ := os.ReadDir(atPath)
	for _, entry := range entries {
		if slices.Contains(exclude, entry.Name()) {
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
