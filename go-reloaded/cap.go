package main

import (
	"strconv"
	"strings"
	"unicode"
)

func convertCap(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		punctuation := ""

		// Handle punctuation attached to (cap).
		if strings.HasPrefix(words[i], "(cap)") {
			punctuation = strings.TrimPrefix(words[i], "(cap)")
			words[i] = "(cap)"
		}

		if words[i] != "(cap)" && !strings.HasPrefix(words[i], "(cap,") {
			continue
		}

		count := 1
		markerLen := 1

		if words[i] != "(cap)" {
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
			runes := []rune(strings.ToLower(words[j]))

			if len(runes) > 0 {
				runes[0] = unicode.ToUpper(runes[0])
				words[j] = string(runes)
			}
		}

		words = append(words[:i], words[i+markerLen:]...)

		if punctuation != "" && i > 0 {
			words[i-1] += punctuation
		}

		i--
	}

	return strings.Join(words, " ")
}