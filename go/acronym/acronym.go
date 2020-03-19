// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// this is failing with
// acronym_test.go:11: Acronym test [Something - I made up from thin air], expected [SIMUFTA], actual [SIMUFTA]
// wtf?
// maybe it's expecting the actual object, and i'm creating a new struct?

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	// purify input to get around weird rules like keeping the first letter after a hyphen
	// i'd like to not have custom logic here, but unicode.IsPunct fails at strings like "Halley's Comet"
	for i := 0; i < len(s); i++ {
		if s[i] == '?' || s[i] == '-' {
			s = s[:i] + " " + s[(i+1):]
		}
	}

	words := strings.Split(s, " ")
	firstLetters := make([]rune, 0)

	for i := 0; i < len(words); i++ {
		words[i] = strings.TrimSpace(words[i])

		firstLetters = append(firstLetters, firstLetter(words[i]))
	}

	return string(firstLetters)
}

func firstLetter(s string) rune {
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			return unicode.ToUpper(rune(s[i]))
		}
	}
	return 0
}
