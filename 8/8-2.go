package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)

	scanner := bufio.NewScanner(file)
	code_chars := 0
	encoded_chars := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			code_chars++
			switch c {
			case '\\':
				encoded_chars += 2
			case '"':
				encoded_chars += 2
			default:
				encoded_chars++
			}
		}
		encoded_chars += 2 // surrounding double quotes
	}

	fmt.Printf("Encoded characters: %d\n", encoded_chars)
	fmt.Printf("Code characters: %d\n", code_chars)
	fmt.Printf("Difference: %d\n", encoded_chars-code_chars)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
