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
	string_chars := 0
	escape := false
	hex_escape := false
	seen_one_hex := false
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			switch c {
			case ' ', '\t', '\n':
			default:
				code_chars++
			}

			switch c {
			case '\\':
				if escape {
					string_chars++
					escape = false
				} else {
					escape = true
				}
			case '"':
				if escape {
					escape = false
					string_chars++
				}
			case 'x':
				if escape {
					hex_escape = true
					escape = false
				} else {
					string_chars++
				}
			case 'a', 'b', 'c', 'd', 'e', 'f', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if hex_escape {
					seen_one_hex = true
					hex_escape = false
				} else if seen_one_hex {
					string_chars++
					seen_one_hex = false
				} else {
					string_chars++
				}
			default:
				string_chars++
			}
		}
	}

	fmt.Printf("Code characters: %d\n", code_chars)
	fmt.Printf("String characters: %d\n", string_chars)
	fmt.Printf("Difference: %d\n", code_chars-string_chars)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
