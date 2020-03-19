// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

const defaultRes = "Whatever."

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if endsWithQuestionMark(remark) && isAllCaps(remark) && !hasNoLetters(remark) {
		return "Calm down, I know what I'm doing!"
	}
	if isAllCaps(remark) && !hasNoLetters(remark) {
		return "Whoa, chill out!"
	}
	if endsWithQuestionMark(remark) {
		return "Sure."
	}
	if hasNoLetters(remark) {
		return defaultRes
	}
	return defaultRes
}

func endsWithQuestionMark(remark string) bool {
	r := remark[len(remark)-1]
	return string(r) == "?"
}

func isAllCaps(remark string) bool {
	for i := 0; i < len(remark); i++ {
		if unicode.IsLetter(rune(remark[i])) {
			if !unicode.IsUpper(rune(remark[i])) {
				return false
			}
		}
	}
	return true
}

func hasNoLetters(remark string) bool {
	for i := 0; i < len(remark); i++ {
		if unicode.IsLetter(rune(remark[i])) {
			return false
		}
	}
	return true
}
