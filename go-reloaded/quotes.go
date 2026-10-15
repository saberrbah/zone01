package main

import (
	"strings"
	"unicode"
)

func fixQuotes(text string) string {
	runes := []rune(text)
	var result strings.Builder
	open := false

	for i := 0; i < len(runes); i++ {
		if runes[i] != '\'' {
			result.WriteRune(runes[i])
			continue
		}

		// Preserve apostrophes inside words, like don't and I'm.
		if i > 0 && i+1 < len(runes) &&
			(unicode.IsLetter(runes[i-1]) || unicode.IsDigit(runes[i-1])) &&
			(unicode.IsLetter(runes[i+1]) || unicode.IsDigit(runes[i+1])) {
			result.WriteRune('\'')
			continue
		}

		if !open {
			// Opening quote: remove spaces immediately inside it.
			result.WriteRune('\'')
			for i+1 < len(runes) && runes[i+1] == ' ' {
				i++
			}
			open = true
		} else {
			// Closing quote: remove spaces immediately before it.
			current := strings.TrimRight(result.String(), " ")
			result.Reset()
			result.WriteString(current)
			result.WriteRune('\'')
			open = false

			// Prevent the closing quote from joining two words.
			if i+1 < len(runes) &&
				(unicode.IsLetter(runes[i+1]) || unicode.IsDigit(runes[i+1]) || runes[i+1] == '\'') {
				result.WriteRune(' ')
			}
		}
	}

	return result.String()
}