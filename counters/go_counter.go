package counters

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type GoCounter struct{}

func (g *GoCounter) Name() string {
	return "Go"
}

func (g *GoCounter) ShouldCount(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".go"
}

func (g *GoCounter) CountLines(filePath string, includeWhitespace bool) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if includeWhitespace {
			lineCount++
		} else {
			line := strings.TrimSpace(scanner.Text())
			// Skip empty lines and comments
			if line != "" && !strings.HasPrefix(line, "//") {
				lineCount++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}
