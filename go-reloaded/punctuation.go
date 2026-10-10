
package main

import "strings"

func fixPunctuation(text string) string {
	input := []byte(text)
	output := make([]byte, 0, len(input))

	for i := 0; i < len(input); {
		if strings.ContainsRune(".,!?:;", rune(input[i])) {
			// Remove spaces before punctuation.
			for len(output) > 0 && output[len(output)-1] == ' ' {
				output = output[:len(output)-1]
			}

			// Keep consecutive punctuation together.
			for i < len(input) &&
				strings.ContainsRune(".,!?:;", rune(input[i])) {
				output = append(output, input[i])
				i++
			}

			// Remove spaces after punctuation.
			for i < len(input) && input[i] == ' ' {
				i++
			}

			// Add one space before the next word,
			// but not before another punctuation mark.
			if i < len(input) &&
				!strings.ContainsRune(".,!?:;", rune(input[i])) {
				output = append(output, ' ')
			}
		} else {
			if input[i] == ' ' {
				if len(output) > 0 && output[len(output)-1] != ' ' {
					output = append(output, ' ')
				}
			} else {
				output = append(output, input[i])
			}
			i++
		}
	}

	return strings.TrimRight(string(output), " ")
}
