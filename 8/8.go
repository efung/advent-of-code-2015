package main

import (
	"bufio"
	"fmt"
	"os"
	"text/scanner"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)

	var s scanner.Scanner
	s.Init(bufio.NewReader(file))
	s.Mode = scanner.ScanChars
	tok := s.Next()
	string_chars := 0
	code_chars := 0
	for tok != scanner.EOF {
		if tok == '\n' {
			tok = s.Next()
			continue
		}
		code_chars++
		switch tok {
		case '\\':
			escape_tok := s.Next()
			code_chars++
			switch escape_tok {
			case '\\', '"':
				string_chars++
			case 'x':
				_ = s.Next()
				code_chars++
				_ = s.Next()
				code_chars++
				string_chars++
			}
		case '"':
			break
		default:
			string_chars++
		}
		tok = s.Next()
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
