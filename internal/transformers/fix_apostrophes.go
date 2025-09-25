package textprocessor

import (
	"regexp"
)

// FixApostrophes исправляет апострофы в сокращениях и притяжательных формах.
func FixApostrophes(input string) string {
	// Обрабатываем сокращения (can' t → can't)
	re1 := regexp.MustCompile(`(\w)\s*'\s*([a-zA-Z])`)
	input = re1.ReplaceAllString(input, "$1'$2")

	// Обрабатываем притяжательные формы (systers ' → systers')
	re2 := regexp.MustCompile(`(\w+)\s*'\s*(\b|$)`)
	input = re2.ReplaceAllString(input, "$1'$2")

	return input
}
