package main

import "strings"

func fixPunctuation(text string) string {
	input := []byte(text)
	output := make([]byte, 0, len(input))

	for i := 0; i < len(input); {
		// Keep URLs unchanged.
		if strings.HasPrefix(string(input[i:]), "https://") ||
			strings.HasPrefix(string(input[i:]), "http://") {
			for i < len(input) &&
				input[i] != ' ' &&
				input[i] != '\n' &&
				input[i] != '\t' {
				output = append(output, input[i])
				i++
			}
			continue
		}

		// Handle punctuation.
		if strings.ContainsRune(".,!?:;", rune(input[i])) {
			// Remove spaces before punctuation.
			for len(output) > 0 && output[len(output)-1] == ' ' {
				output = output[:len(output)-1]
			}

			// Copy punctuation, including consecutive marks.
			for i < len(input) &&
				strings.ContainsRune(".,!?:;", rune(input[i])) {
				output = append(output, input[i])
				i++
			}

			// Skip spaces after punctuation.
			for i < len(input) && input[i] == ' ' {
				i++
			}

			// Add one space before the next word,
			// except at the end of the line.
			if i < len(input) {
				output = append(output, ' ')
			}
			continue
		}

		// Normalize repeated spaces.
		if input[i] == ' ' {
			if len(output) > 0 && output[len(output)-1] != ' ' {
				output = append(output, ' ')
			}
		} else {
			output = append(output, input[i])
		}

		i++
	}

	return strings.TrimRight(string(output), " ")
}