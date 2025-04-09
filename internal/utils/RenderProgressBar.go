package utils

import (
	"fmt"
)

func RenderProgressBar(current, total int, width int) {
	percent := float64(current) / float64(total)
	filled := int(percent * float64(width))
	empty := width - filled

	bar := fmt.Sprintf(
		"\r[%s%s] %3.0f%%",
		repeat("█", filled),
		repeat("·", empty),
		percent*100,
	)

	fmt.Print(bar)
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
