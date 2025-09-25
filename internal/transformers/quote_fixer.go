package textprocessor

import (
	"regexp"
	"strings"
)

// FixQuotes исправляет одинарные и двойные кавычки, сохраняя апострофы.
func FixQuotes(input string) string {
	// Одинарные кавычки: сохраняем пробелы только внутри кавычек, не склеиваем слова
	reSingle := regexp.MustCompile(`'\s*([^']*?)\s*'`)
	input = reSingle.ReplaceAllStringFunc(input, func(match string) string {
		parts := reSingle.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}
		trimmed := strings.Trim(parts[1], " ")
		return "'" + trimmed + "'"
	})

	// Двойные кавычки: аналогично
	reDouble := regexp.MustCompile(`"\s*([^\"]*?)\s*"`)
	input = reDouble.ReplaceAllStringFunc(input, func(match string) string {
		parts := reDouble.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}
		trimmed := strings.Trim(parts[1], " ")
		return "\"" + trimmed + "\""
	})

	return input
}
