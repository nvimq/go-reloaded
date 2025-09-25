package textprocessor

import (
	"regexp"
	"strings"
)

// FixPunctuation форматирует пунктуацию согласно требованиям.
func FixPunctuation(input string) string {
	// Шаг 1: Нормализуем избыточную пунктуацию
	input = regexp.MustCompile(`\.{2,}`).ReplaceAllString(input, "...") // Многоточия: любое количество точек → ...
	input = regexp.MustCompile(`!{2,}`).ReplaceAllString(input, "!!!")  // Восклицательные знаки → !!!
	input = regexp.MustCompile(`\?{2,}`).ReplaceAllString(input, "?")   // Вопросительные знаки → ?
	input = regexp.MustCompile(`:{2,}`).ReplaceAllString(input, ":")    // Двоеточия → :
	input = regexp.MustCompile(`;{2,}`).ReplaceAllString(input, ";")    // Точки с запятой → ;
	input = regexp.MustCompile(`,{2,}`).ReplaceAllString(input, ",")    // Запятые → ,

	// Шаг 2: Обрабатываем комбинации !? или ?!
	input = regexp.MustCompile(`(!+\?+|\?+!+)`).ReplaceAllString(input, "!?") // !? или ?! → !?

	// Шаг 3: Удаляем пробелы перед пунктуацией и добавляем пробел после
	input = regexp.MustCompile(`\s*([.,!?:;])\s*`).ReplaceAllString(input, "$1 ")

	// Шаг 4: Убедимся, что многоточия примыкают к предыдущему слову
	input = regexp.MustCompile(`(\S)\s*\.\.\.\s*`).ReplaceAllString(input, "$1... ")

	// Шаг 5: Добавляем пробел после пунктуации, если его нет (кроме закрывающих скобок)
	input = regexp.MustCompile(`([.,!?:;])([^\s)])`).ReplaceAllString(input, "$1 $2")

	// Шаг 6: Обрабатываем двоеточия отдельно
	input = regexp.MustCompile(`(\w):(\w)`).ReplaceAllString(input, "$1: $2")

	// Шаг 7: Удаляем лишние пробелы
	input = strings.Join(strings.Fields(input), " ")

	return strings.TrimSpace(input)
}
