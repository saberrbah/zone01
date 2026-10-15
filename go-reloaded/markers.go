package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

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

// convertMarkers applies (hex), (bin), (up), (low) and (cap) markers
// from left to right, so every marker only sees the words before it.
func convertMarkers(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		name, count, size, rest, ok := parseMarker(words, i)
		if !ok {
			continue
		}

		if i > 0 {
			switch name {
			case "hex":
				words[i-1] = convertBase(words[i-1], 16)
			case "bin":
				words[i-1] = convertBase(words[i-1], 2)
			default:
				start := i - count
				if start < 0 {
					start = 0
				}
				for j := start; j < i; j++ {
					words[j] = transform(name, words[j])
				}
			}

			// Keep punctuation that was attached to the marker.
			words[i-1] += rest
		}

		// Remove the marker (one word, or two for "(up, 2)").
		words = append(words[:i], words[i+size:]...)
		i--
	}

	return strings.Join(words, " ")
}

// parseMarker checks whether words[i] starts a marker. It returns the marker
// name, how many words it applies to, how many words the marker uses, and any
// punctuation stuck to its closing parenthesis.
func parseMarker(words []string, i int) (name string, count, size int, rest string, ok bool) {
	word := words[i]

	for _, n := range []string{"hex", "bin", "up", "low", "cap"} {
		simple := "(" + n + ")"
		if strings.HasPrefix(word, simple) {
			return n, 1, 1, strings.TrimPrefix(word, simple), true
		}

		if n == "hex" || n == "bin" || word != "("+n+"," || i+1 >= len(words) {
			continue
		}

		token := words[i+1]
		end := strings.Index(token, ")")
		if end == -1 {
			continue
		}

		c, err := strconv.Atoi(token[:end])
		if err != nil {
			continue
		}

		return n, c, 2, token[end+1:], true
	}

	return "", 0, 0, "", false
}

// convertBase turns a number in the given base into decimal.
// If the word is not a valid number, it is left unchanged.
func convertBase(word string, base int) string {
	number, err := strconv.ParseInt(word, base, 64)
	if err != nil {
		return word
	}

	return strconv.FormatInt(number, 10)
}

func transform(name, word string) string {
	switch name {
	case "up":
		return strings.ToUpper(word)
	case "low":
		return strings.ToLower(word)
	}

	return capitalize(word)
}

// capitalize lowercases the word, then uppercases its first letter
// (skipping leading punctuation such as a quote).
func capitalize(word string) string {
	runes := []rune(strings.ToLower(word))

	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			break
		}
		if unicode.IsDigit(r) {
			break
		}
	}

	return string(runes)
}