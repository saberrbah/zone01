package main

import "strings"

const punctuationMarks = ".,!?:;"

func processText(text string) string {
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		// Keep the indentation of the line.
		trimmed := strings.TrimLeft(line, " \t")
		indent := line[:len(line)-len(trimmed)]

		// Normalize markers before processing them.
		trimmed = normalizeMarkers(trimmed)
		
		trimmed = convertHex(trimmed)
		trimmed = convertBin(trimmed)
		trimmed = convertUp(trimmed)
		trimmed = convertLow(trimmed)
		trimmed = convertCap(trimmed)

		// Fix punctuation, quotes, and articles.
		trimmed = fixPunctuation(trimmed)
		trimmed = fixQuotes(trimmed)
		trimmed = fixArticles(trimmed)

		// A line starting with punctuation has no indentation.
		if trimmed != "" && strings.ContainsRune(punctuationMarks, rune(trimmed[0])) {
			indent = ""
		}

		lines[i] = indent + trimmed
	}

	// Attach punctuation-only lines to the previous line.
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