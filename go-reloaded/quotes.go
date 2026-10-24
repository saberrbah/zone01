package main

import "strings"

func fixQuotes(text string) string {
	var result strings.Builder

	for i := 0; i < len(text); i++ {
		if text[i] != '\'' {
			result.WriteByte(text[i])
			continue
		}

		// A quote before a letter or digit can be part of a word, like '90s.
		if i+1 < len(text) && isWordChar(text[i+1]) {
			result.WriteByte(text[i])
			continue
		}

		// Find the next possible closing quote.
		end := -1
		for j := i + 1; j < len(text); j++ {
			if text[j] != '\'' {
				continue
			}

			// A quote followed by a word character may be an apostrophe.
			if j+1 < len(text) && isWordChar(text[j+1]) {
				continue
			}

			end = j
			break
		}

		// If no closing quote exists, preserve the remaining text.
		if end == -1 {
			result.WriteString(text[i:])
			break
		}

		// Remove whitespace only inside the confirmed quote pair.
		inside := strings.TrimSpace(text[i+1 : end])
		result.WriteByte('\'')
		result.WriteString(inside)
		result.WriteByte('\'')

		i = end
	}

	return result.String()
}

func isWordChar(c byte) bool {
	return c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z' ||
		c >= '0' && c <= '9' ||
		c == '_'
}