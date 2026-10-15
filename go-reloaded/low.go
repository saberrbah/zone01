package main

import (
	"strconv"
	"strings"
)

func convertLow(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		punctuation := ""

		// Handle punctuation attached to (low).
		if strings.HasPrefix(words[i], "(low)") {
			punctuation = strings.TrimPrefix(words[i], "(low)")
			words[i] = "(low)"
		}

		if words[i] != "(low)" && !strings.HasPrefix(words[i], "(low,") {
			continue
		}

		count := 1
		markerLen := 1

		if words[i] != "(low)" {
			if i+1 >= len(words) {
				continue
			}

			numberToken := words[i+1]
			closeIndex := strings.Index(numberToken, ")")
			if closeIndex == -1 {
				continue
			}

			n := numberToken[:closeIndex]
			punctuation += numberToken[closeIndex+1:]

			var err error
			count, err = strconv.Atoi(n)
			if err != nil || count < 1 {
				continue
			}

			markerLen = 2
		}

		// Not enough words before the marker: use the ones that exist.
		if count > i {
			count = i
		}

		for j := i - count; j < i; j++ {
			words[j] = strings.ToLower(words[j])
		}

		words = append(words[:i], words[i+markerLen:]...)

		if punctuation != "" && i > 0 {
			words[i-1] += punctuation
		}

		i--
	}

	return strings.Join(words, " ")
}