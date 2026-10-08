package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertHex(words []string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {
			number, err := strconv.ParseInt(words[i-1], 16, 64)

			if err != nil {
				fmt.Println("Error converting hexadecimal:", err)
				return words
			}

			words[i-1] = strconv.FormatInt(number, 10)
			words = append(words[:i], words[i+1:]...)
		}
	}

	return words
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(data)
	words := strings.Fields(text)

	words = convertHex(words)

	text = strings.Join(words, " ")
	data = []byte(text)

	err = os.WriteFile(outputFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
