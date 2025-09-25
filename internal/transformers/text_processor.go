package textprocessor

import (
	"strings"
)

// ProcessText выполняет полную обработку текста, повторяя процесс 5 раз для каждой строки.
func ProcessText(input string) string {
	lines := strings.Split(input, "\n")
	processedLines := make([]string, 0, len(lines))

	for _, line := range lines {
		processedLine := line
		processedLine = FixSpaces(processedLine)      // Нормализуем пробелы
		processedLine = ConvertNumbers(processedLine) // Обрабатываем числа (hex, bin)
		processedLine = ModifyCase(processedLine)     // Обрабатываем теги (up, low, cap)
		processedLine = FixQuotes(processedLine)      // Исправляем кавычки
		processedLine = FixApostrophes(processedLine) // Исправляем апострофы
		processedLine = FixPunctuation(processedLine) // Форматируем пунктуацию
		processedLine = FixArticles(processedLine)    // Корректируем a/an
		processedLines = append(processedLines, processedLine)
	}

	return strings.Join(processedLines, "\n")
}
