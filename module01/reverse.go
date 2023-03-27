package module01

import (
	"strings"
)

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {

	var rWord = []rune(word)
	var sbReverse = strings.Builder{}

	for i := len(rWord) - 1; i >= 0; i-- {
		sbReverse.WriteRune(rWord[i])
	}

	return sbReverse.String()
}
