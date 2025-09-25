package textprocessor

import (
	"strings"
)

// FixArticles исправляет артикли a/an перед словами.
func FixArticles(input string) string {
	silentHWords := map[string]bool{
		"honest":    true,
		"heir":      true,
		"honorific": true,
		"honor":     true,
		"herb":      true,
		"hour":      true,
		"homage":    true,
	}
	// Исключения, где "a" несмотря на гласную (юзер, юнит, юниформ, юниверсити, южуал)
	aExceptions := map[string]bool{
		"university": true,
		"uniform":    true,
		"user":       true,
		"usual":      true,
		"unit":       true,
	}
	conjunctions := map[string]bool{
		"and": true,
		"or":  true,
		"but": true,
		"for": true,
		"yet": true,
		"so":  true,
	}
	vowels := "aeiouAEIOU"

	words := strings.Fields(input)
	result := make([]string, 0, len(words))

	for i := 0; i < len(words); i++ {
		currentWord := words[i]
		currentLower := strings.ToLower(currentWord)
		isArticle := currentLower == "a" || currentLower == "an"

		if isArticle && i+1 < len(words) {
			nextWord := words[i+1]
			nextStripped := strings.ToLower(strings.TrimLeft(nextWord, "'\""))
			isValidWord := len(nextStripped) > 0 && (nextStripped[0] >= 'a' && nextStripped[0] <= 'z')

			if isValidWord && !conjunctions[nextStripped] {
				useAn := false
				if silentHWords[nextStripped] {
					useAn = true
				} else if aExceptions[nextStripped] {
					useAn = false
				} else if strings.ContainsAny(string(nextStripped[0]), vowels) {
					useAn = true
				}

				if useAn {
					if currentLower == "a" {
						if isUpper(currentWord) {
							result = append(result, "An")
						} else {
							result = append(result, "an")
						}
					} else {
						result = append(result, currentWord)
					}
				} else {
					if currentLower == "an" {
						if isUpper(currentWord) {
							result = append(result, "A")
						} else {
							result = append(result, "a")
						}
					} else {
						result = append(result, currentWord)
					}
				}
				continue
			}
		}
		result = append(result, currentWord)
	}

	return strings.Join(result, " ")
}

func isUpper(word string) bool {
	if len(word) == 0 {
		return false
	}
	return word[0] >= 'A' && word[0] <= 'Z'
}
