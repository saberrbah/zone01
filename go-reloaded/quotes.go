
package main

import "strings"

func fixQuotes(text string) string {
	runes := []rune(text)
	var result strings.Builder
	open := false

	for i := 0; i < len(runes); i++ {
		if runes[i] == '\'' {
			if !open {
				// Opening quote: skip spaces after it.
				result.WriteRune('\'')
				for i+1 < len(runes) && runes[i+1] == ' ' {
					i++
				}
				open = true
			} else {
				// Closing quote: remove spaces before it.
				current := result.String()
				result.Reset()
				result.WriteString(strings.TrimRight(current, " "))
				result.WriteRune('\'')
				open = false
			}
		} else {
			result.WriteRune(runes[i])
		}
	}

	return result.String()
}
