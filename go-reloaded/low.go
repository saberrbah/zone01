
package main

import (
	"strings"
)

func convertLow(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" {
			if i == 0 {
				continue
			}

			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}