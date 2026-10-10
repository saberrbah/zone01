
package main

import (
	"strings"
)

func convertCap(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" {
			if i == 0 {
				continue
			}

			word := strings.ToLower(words[i-1])
			words[i-1] = strings.ToUpper(word[:1]) + word[1:]

			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}