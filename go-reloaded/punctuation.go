package main

import (
	"strings"
)

func fixPunctuation(text string) string {
	var result strings.Builder

	for i := 0; i < len(text); i++ {
		// If the character is punctuation, remove spaces before it.
		if strings.ContainsRune(".,!?:;", rune(text[i])) {
			current := result.String()
			result.Reset()
			result.WriteString(strings.TrimRight(current, " "))

			// Write consecutive punctuation marks.
			for i < len(text) && strings.ContainsRune(".,!?:;", rune(text[i])) {
				result.WriteByte(text[i])
				i++
			}

			// Skip spaces after punctuation.
			for i < len(text) && text[i] == ' ' {
				i++
			}

			// Add one space after punctuation if text continues.
			if i < len(text) {
				result.WriteByte(' ')
			}

			i--
		} else {
			result.WriteByte(text[i])
		}
	}

	return strings.TrimRight(result.String(), " ")
}