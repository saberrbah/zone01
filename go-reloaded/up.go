// package main

// import (
// 	"strconv"
// 	"strings"
// )

// func convertUp(text string) string {
// 	words := strings.Fields(text)

// 	for i := 0; i < len(words); i++ {
// 		if words[i] == "(up)" || strings.HasPrefix(words[i], "(up,") {
// 			count := 1
// 			markerLen := 1

// 			if words[i] != "(up)" {
// 				if i+1 >= len(words) {
// 					continue
// 				}

// 				n := strings.TrimSuffix(words[i+1], ")")
// 				var err error
// 				count, err = strconv.Atoi(n)

// 				if err != nil || count < 1 {
// 					continue
// 				}

// 				markerLen = 2
// 			}

// 			if count > i {
// 				continue
// 			}

// 			for j := i - count; j < i; j++ {
// 				words[j] = strings.ToUpper(words[j])
// 			}

// 			words = append(words[:i], words[i+markerLen:]...)
// 			i--
// 		}
// 	}

// 	return strings.Join(words, " ")
// }


package main

import (
	"strconv"
	"strings"
)

func convertUp(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		// Handle punctuation attached to (up).
		punctuation := ""
		if strings.HasPrefix(words[i], "(up)") {
			punctuation = strings.TrimPrefix(words[i], "(up)")
			words[i] = "(up)"
		}

		if words[i] != "(up)" && !strings.HasPrefix(words[i], "(up,") {
			continue
		}

		count := 1
		markerLen := 1

		if words[i] != "(up)" {
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
			words[j] = strings.ToUpper(words[j])
		}

		words = append(words[:i], words[i+markerLen:]...)

		if punctuation != "" && i > 0 {
			words[i-1] += punctuation
		}

		i--
	}

	return strings.Join(words, " ")
}
