package textprocessor

import (
	"regexp"
	"strconv"
)

// ConvertNumbers преобразует числа с тегами (hex) и (bin) в десятичные.
func ConvertNumbers(input string) string {
	for {
		changed := false

		// Обработка (hex)
		hexRe := regexp.MustCompile(`(\S+)\s*\(hex\)`)
		input = hexRe.ReplaceAllStringFunc(input, func(match string) string {
			parts := hexRe.FindStringSubmatch(match)
			if len(parts) < 2 {
				return match
			}
			// Проверяем, является ли слово валидным шестнадцатеричным числом
			if !regexp.MustCompile(`^[0-9a-fA-F]+$`).MatchString(parts[1]) {
				return parts[1] // Возвращаем слово без тега, если невалидное
			}
			decimal, err := strconv.ParseInt(parts[1], 16, 64)
			if err != nil {
				return parts[1]
			}
			changed = true
			return strconv.FormatInt(decimal, 10)
		})

		// Обработка (bin)
		binRe := regexp.MustCompile(`(\S+)\s*\(bin\)`)
		input = binRe.ReplaceAllStringFunc(input, func(match string) string {
			parts := binRe.FindStringSubmatch(match)
			if len(parts) < 2 {
				return match
			}
			// Проверяем, является ли слово валидным двоичным числом
			if !regexp.MustCompile(`^[01]+$`).MatchString(parts[1]) {
				return parts[1]
			}
			decimal, err := strconv.ParseInt(parts[1], 2, 64)
			if err != nil {
				return parts[1]
			}
			changed = true
			return strconv.FormatInt(decimal, 10)
		})

		if !changed {
			break
		}
	}
	return input
}
