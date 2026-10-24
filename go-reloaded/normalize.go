package main

import "regexp"

// This regex finds markers like (up), ( up ), (cap, 2), and ( cap, 2 ).
// It allows spaces inside the parentheses and an optional number.
var markerPattern = regexp.MustCompile(`\(\s*(hex|bin|up|low|cap)\s*(?:,\s*(-?\d+)\s*)?\)`)

// normalizeMarkers puts every valid marker into a standard format.
func normalizeMarkers(text string) string {

	// Find every marker in the text and replace it with a normalized version.
	return markerPattern.ReplaceAllStringFunc(text, func(match string) string {

		// Extract the marker name and optional number from the matched text.
		parts := markerPattern.FindStringSubmatch(match)

		// Start building the marker with an opening parenthesis and its name.
		marker := "(" + parts[1]

		// If the marker has a number, add a comma, a space, and that number.
		if parts[2] != "" {
			marker += ", " + parts[2]
		}

		// Add the closing parenthesis and return the normalized marker.
		return marker + ")"
	})
}