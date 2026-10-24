package main

import (
	"strconv"
	"strings"
	"unicode"
)

func convertCap(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		count := 1
		markerLen := 1

		// Handle the exact (cap) marker.
		if words[i] == "(cap)" {
			count = 1
		} else if words[i] == "(cap," {
			// Handle (cap, number) when the number is separated by a space.
			if i+1 >= len(words) {
				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}

			numberText := strings.TrimSuffix(words[i+1], ")")
			countValue, err := strconv.Atoi(numberText)

			// Remove both marker parts for invalid, zero, negative,
			// or overflowing counts.
			if err != nil || countValue <= 0 {
				words = append(words[:i], words[i+2:]...)
				i--
				continue
			}

			count = countValue
			markerLen = 2
		} else if strings.HasPrefix(words[i], "(cap,") &&
			strings.HasSuffix(words[i], ")") {
			// Handle (cap,number) without a space.
			numberText := strings.TrimSuffix(
				strings.TrimPrefix(words[i], "(cap,"),
				")",
			)
			countValue, err := strconv.Atoi(numberText)

			// Remove the marker for invalid counts.
			if err != nil || countValue <= 0 {
				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}

			count = countValue
		} else {
			continue
		}

		// Use only the available previous words.
		if count > i {
			count = i
		}

		// Capitalize the selected words.
		for j := i - count; j < i; j++ {
			runes := []rune(strings.ToLower(words[j]))

			// Skip words starting with punctuation or quotes.
			if len(runes) == 0 ||
				(!unicode.IsLetter(runes[0]) && !unicode.IsDigit(runes[0])) {
				continue
			}

			for k, r := range runes {
				if unicode.IsLetter(r) {
					runes[k] = unicode.ToUpper(r)
					break
				}
				if unicode.IsDigit(r) {
					break
				}
			}

			words[j] = string(runes)
		}

		// Remove the marker and, if necessary, its count.
		words = append(words[:i], words[i+markerLen:]...)
		i--
	}

	return strings.Join(words, " ")
}