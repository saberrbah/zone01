package main

import (
	"strings"
)

func processText(text string) string {
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		line = convertHex(line)
		line = convertBin(line)
		line = convertUp(line)
		line = convertLow(line)
		line = convertCap(line)
		lines[i] = line
	}

	return strings.Join(lines, "\n")
}
