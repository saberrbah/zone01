package main

import (
	"strconv"
	"strings"
)

func convertBin(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		// Only recognize the exact (bin) marker.
		if words[i] != "(bin)" {
			continue
		}

		// If a previous word exists, try to convert it from binary to decimal.
		if i > 0 {
			number, err := strconv.ParseInt(words[i-1], 2, 64)

			if err == nil {
				words[i-1] = strconv.FormatInt(number, 10)
			}
		}

		// Delete the marker, even if there is no previous word.
		words = append(words[:i], words[i+1:]...)
		i--
	}

	return strings.Join(words, " ")
}