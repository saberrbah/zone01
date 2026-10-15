package main

import (
	"strconv"
	"strings"
)

func convertUp(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		prefix := ""
		punctuation := ""

		// Handle punctuation before (up), like !(up)!
		for len(words[i]) > 0 &&
			strings.ContainsRune(".,!?:;", rune(words[i][0])) {
			prefix += string(words[i][0])
			words[i] = words[i][1:]
		}

		// Handle punctuation after (up), like (up)!
		if strings.HasPrefix(words[i], "(up)") {
			punctuation = strings.TrimPrefix(words[i], "(up)")
			words[i] = "(up)"
		}

		if words[i] != "(up)" && !strings.HasPrefix(words[i], "(up,") {
			if prefix != "" {
				words[i] = prefix + words[i]
			}
			continue
		}

		count := 1
		markerLen := 1

		if words[i] != "(up)" {
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
			words[j] = strings.ToUpper(words[j])
		}

		words = append(words[:i], words[i+markerLen:]...)

		if i > 0 {
			words[i-1] += punctuation

			if prefix != "" {
				// Avoid adding punctuation twice.
				if !strings.HasSuffix(words[i-1], prefix) {
					words[i-1] += prefix
				}
			}
		}

		i--
	}

	return strings.Join(words, " ")
}