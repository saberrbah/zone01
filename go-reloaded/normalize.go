package main

import "regexp"

// markerPattern matches a marker with any spacing inside the parentheses,
// like "( up )", "(up,3)" or "(  low  ,  2 )".
var markerPattern = regexp.MustCompile(`\(\s*(hex|bin|up|low|cap)\s*(?:,\s*(-?\d+)\s*)?\)`)

// normalizeMarkers rewrites every marker to its standard form:
// "(up)" or "(up, 3)".
func normalizeMarkers(text string) string {
	return markerPattern.ReplaceAllStringFunc(text, func(match string) string {
		parts := markerPattern.FindStringSubmatch(match)

		if parts[2] == "" {
			return "(" + parts[1] + ")"
		}

		return "(" + parts[1] + ", " + parts[2] + ")"
	})
}
