package main

import "strings"

func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		// Check whether the current word is "a", ignoring case.
		if strings.ToLower(words[i]) != "a" {
			continue
		}

		nextWord := strings.ToLower(words[i+1])

		// Leave the article unchanged if the next word starts with a quote or punctuation.
		if nextWord == "" || !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", rune(nextWord[0])) {
			continue
		}

		// Change "a" to "an" before a vowel or h.
		if strings.ContainsRune("aeiouh", rune(nextWord[0])) {
			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}
	}

	return strings.Join(words, " ")
}