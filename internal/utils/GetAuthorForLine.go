package utils

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func GetAuthorForLine(file string, forLine uint32) (string, error) {
	cmd := exec.Command("git", "blame", "-L", fmt.Sprintf("%d,%d", forLine, forLine), "--porcelain", file)

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "0000000000000000000000000000000000000000") {
			return getCurrentGitUser()
		}

		if strings.HasPrefix(line, "author ") {
			author := strings.TrimPrefix(line, "author ")
			return author, nil
		}
	}
	return "", fmt.Errorf("author not found")
}

func getCurrentGitUser() (string, error) {
	cmd := exec.Command("git", "config", "user.name")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
