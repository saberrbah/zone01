
package main

import (
	"strings"
)

func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			nextWord := strings.ToLower(words[i+1])

			nextWord = strings.TrimLeft(nextWord, "\"'([{")

			if nextWord != "" && strings.ContainsRune("aeiouh", rune(nextWord[0])) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return strings.Join(words, " ")
}
