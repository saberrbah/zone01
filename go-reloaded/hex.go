package main

import (
	"strconv"
	"strings"
)

func convertHex(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if strings.TrimRight(words[i], ".,!?;:") == "(hex)" {
			if i == 0 {
				continue
			}
			number, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				continue
			}

			// Keep punctuation attached to the marker, like (hex).
			punctuation := strings.TrimPrefix(words[i], "(hex)")

			words[i-1] = strconv.FormatInt(number, 10) + punctuation
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}