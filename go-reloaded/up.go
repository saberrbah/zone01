package main

import (
	"strconv"
	"strings"
)

func convertUp(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		// Handle the exact (up) marker.
		if words[i] == "(up)" {
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}

			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		// Handle (up, number), including numbers separated by spaces.
		if words[i] == "(up," {
			// Check that the number exists.
			if i+1 >= len(words) {
				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}

			numberText := strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(numberText)

			// Remove both marker parts if the count is invalid,
			// zero, negative, or too large for int.
			if err != nil || count <= 0 {
				words = append(words[:i], words[i+2:]...)
				i--
				continue
			}

			// Apply uppercase to the available previous words.
			if count > i {
				count = i
			}

			for j := i - count; j < i; j++ {
				words[j] = strings.ToUpper(words[j])
			}

			// Remove the marker and its number.
			words = append(words[:i], words[i+2:]...)
			i--
			continue
		}

		// Handle (up,3) when there are no spaces.
		if strings.HasPrefix(words[i], "(up,") &&
			strings.HasSuffix(words[i], ")") {
			numberText := strings.TrimSuffix(strings.TrimPrefix(words[i], "(up,"), ")")
			count, err := strconv.Atoi(numberText)

			if err != nil || count <= 0 {
				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}

			if count > i {
				count = i
			}

			for j := i - count; j < i; j++ {
				words[j] = strings.ToUpper(words[j])
			}

			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}