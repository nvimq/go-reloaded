package textprocessor

import (
	"regexp"
	"strconv"
	"strings"
)

// ModifyCase изменяет регистр слов по тегам (up, low, cap).
func ModifyCase(input string) string {
	re := regexp.MustCompile(`([^\(\)\n]*?)\s*\(\s*(up|low|cap)(?:,\s*(\S+))?\s*\)`)
	for re.MatchString(input) {
		match := re.FindStringSubmatchIndex(input)
		if match == nil {
			break
		}

		text := input[match[2]:match[3]]
		operation := strings.TrimSpace(input[match[4]:match[5]])
		count := 1

		if match[6] != -1 {
			countStr := input[match[6]:match[7]]
			// Проверяем, является ли параметр числом
			if n, err := strconv.Atoi(countStr); err == nil && n > 0 {
				count = n
			} else {
				// Если параметр невалидный (например, "10(bin)"), оставляем текст без изменений
				count = 1
			}
		}

		// Рекурсивно обрабатываем текст внутри тега
		text = ModifyCase(text)
		words := strings.Fields(text)
		if len(words) == 0 {
			input = input[:match[0]] + input[match[1]:]
			continue
		}

		if count > len(words) {
			count = len(words)
		}

		startWord := len(words) - count
		if startWord < 0 {
			startWord = 0
		}

		for i := startWord; i < len(words); i++ {
			switch operation {
			case "up":
				words[i] = strings.ToUpper(words[i])
			case "low":
				words[i] = strings.ToLower(words[i])
			case "cap":
				words[i] = Capitalize(words[i])
			}
		}

		modifiedText := strings.Join(words, " ")
		input = input[:match[0]] + modifiedText + input[match[1]:]
	}
	return strings.TrimSpace(input)
}

// Capitalize делает первую букву слова заглавной, остальные — строчными.
func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	if len(word) == 1 {
		return strings.ToUpper(word)
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
