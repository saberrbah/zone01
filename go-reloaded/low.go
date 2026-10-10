
package main

import (
	"strconv"
	"strings"
)

func convertLow(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" || strings.HasPrefix(words[i], "(low,") {
			count := 1
			markerLen := 1

			if words[i] != "(low)" {
				if i+1 >= len(words) {
					continue
				}

				n := strings.TrimSuffix(words[i+1], ")")
				var err error
				count, err = strconv.Atoi(n)

				if err != nil || count < 1 {
					continue
				}

				markerLen = 2
			}

			if count > i {
				continue
			}

			for j := i - count; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}

			words = append(words[:i], words[i+markerLen:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}
