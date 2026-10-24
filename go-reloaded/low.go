package main

import (
	"strconv"
	"strings"
)

func convertLow(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		// Handle the exact (low) marker.
		if words[i] == "(low)" {
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}

			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		// Handle (low, number) when the number is separated by a space.
		if words[i] == "(low," {
			// Remove an incomplete marker.
			if i+1 >= len(words) {
				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}

			numberText := strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(numberText)

			// Remove both marker parts for invalid, zero, negative,
			// or overflowing counts.
			if err != nil || count <= 0 {
				words = append(words[:i], words[i+2:]...)
				i--
				continue
			}

			// Use only the available previous words.
			if count > i {
				count = i
			}

			for j := i - count; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}

			// Remove the marker and its number.
			words = append(words[:i], words[i+2:]...)
			i--
			continue
		}

		// Handle (low,number) without a space.
		if strings.HasPrefix(words[i], "(low,") &&
			strings.HasSuffix(words[i], ")") {
			numberText := strings.TrimSuffix(
				strings.TrimPrefix(words[i], "(low,"),
				")",
			)
			count, err := strconv.Atoi(numberText)

			// Remove the marker if the count is invalid.
			if err != nil || count <= 0 {
				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}

			// Use only the available previous words.
			if count > i {
				count = i
			}

			for j := i - count; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}

			// Remove the marker.
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}