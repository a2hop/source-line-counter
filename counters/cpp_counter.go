package counters

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type CppCounter struct{}

func (c *CppCounter) Name() string {
	return "C/C++"
}

func (c *CppCounter) ShouldCount(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".cpp" || ext == ".cc" || ext == ".cxx" ||
		ext == ".hpp" || ext == ".h" || ext == ".hxx" || ext == ".c"
}

func (c *CppCounter) CountLines(filePath string, includeWhitespace bool) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	inMultiLineComment := false

	for scanner.Scan() {
		if includeWhitespace {
			lineCount++
		} else {
			line := strings.TrimSpace(scanner.Text())

			// Handle multi-line comments
			if inMultiLineComment {
				if strings.Contains(line, "*/") {
					inMultiLineComment = false
				}
				continue
			}

			if strings.HasPrefix(line, "/*") {
				inMultiLineComment = true
				continue
			}

			// Skip empty lines and single-line comments
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
