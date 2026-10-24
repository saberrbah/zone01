package main

import (
	"fmt"
	"os"
	"strings"
)

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

	if !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println("Error: output file must be a .txt file")
		return
	}

	text := string(data)
	text = processText(text)


	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
