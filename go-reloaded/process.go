package main

import "strings"

const punctuationMarks = ".,!?:;"

func processText(text string) string {
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		// Keep the indentation of the line.
		trimmed := strings.TrimLeft(line, " \t")
		indent := line[:len(line)-len(trimmed)]

		// Punctuation and quotes first, so markers like (up, 2)
		// count real words only. Markers are then applied in order.
		trimmed = fixPunctuation(trimmed)
		trimmed = fixQuotes(trimmed)
		trimmed = convertMarkers(trimmed)
		trimmed = fixArticles(trimmed)

		// A line starting with punctuation has no indentation.
		if trimmed != "" && strings.ContainsRune(punctuationMarks, rune(trimmed[0])) {
			indent = ""
		}

		lines[i] = indent + trimmed
	}

	// Punctuation at the start of a line belongs to the last word before it.
	result := []string{}
	for _, line := range lines {
		if len(result) > 0 &&
			line != "" &&
			strings.ContainsRune(punctuationMarks, rune(line[0])) &&
			strings.TrimSpace(result[len(result)-1]) != "" {
			result[len(result)-1] += line
			continue
		}
		result = append(result, line)
	}

	return strings.Join(result, "\n")
}