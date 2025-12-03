package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//Declare flags and parse command-line arguments (name, value, usage)
	countBytes := flag.Bool("c", false, "count bytes")
	countLines := flag.Bool("l", false, "count lines")
	countWords := flag.Bool("w", false, "count words")
	countChars := flag.Bool("m", false, "count characters")

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: go run main.go [flags] <filename>")
		os.Exit(1)
	}

	filename := args[0]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// If no flags are provided, count all
	if !*countBytes && !*countLines && !*countWords && !*countChars {
		*countBytes = true
		*countLines = true
		*countWords = true
	}

	result := ""

	if *countLines {
		result += fmt.Sprintf("%d ", countLinesFunc(data))
	}
	if *countWords {
		result += fmt.Sprintf("%d ", countWordsFunc(data))
	}
	if *countChars {
		result += fmt.Sprintf("%d ", countCharsFunc(data))
	}
	if *countBytes {
		result += fmt.Sprintf("%d ", len(data))
	}

	result += filename
	fmt.Println(result)
}

func countLinesFunc(data []byte) int {
	count := 0
	for _, b := range data {
		if b == '\n' {
			count++
		}
	}
	return count
}

func countWordsFunc(data []byte) int {
	count := 0
	inWord := false
	for _, b := range data {
		if b == ' ' || b == '\n' || b == '\t' {
			inWord = false
		} else if !inWord {
			count++
			inWord = true
		}
	}
	return count
}

func countCharsFunc(data []byte) int {
	return len([]rune(string(data)))
}
