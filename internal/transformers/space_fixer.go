package textprocessor

import (
	"regexp"
	"strings"
)

// FixSpaces нормализует пробелы вокруг скобок.
func FixSpaces(input string) string {
	input = regexp.MustCompile(`\s*\(\s*`).ReplaceAllString(input, " (")
	input = regexp.MustCompile(`\s*\)\s*`).ReplaceAllString(input, ") ")
	return strings.TrimSpace(input)
}
