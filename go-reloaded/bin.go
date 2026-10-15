package main

import (
	"strconv"
	"strings"
)

func convertBin(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if strings.TrimRight(words[i], ".,!?;:") == "(bin)" {
			if i == 0 {
				continue
			}
			number, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				continue
			}

			words[i-1] = strconv.FormatInt(number, 10)
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}
