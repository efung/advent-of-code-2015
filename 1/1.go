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
	var tok rune
	var floor int
	var pos int
	var firstbasementpos int
	firstbasementpos = 0
	for tok != scanner.EOF {
		pos++
		tok = s.Scan()
		if s.TokenText() == "(" {
			floor++
		} else if s.TokenText() == ")" {
			floor--
		}
		if firstbasementpos == 0 && floor == -1 {
			firstbasementpos = pos
		}
	}

	fmt.Printf("Final floor: %d\n", floor)
	fmt.Printf("Basement floor instruction position: %d\n", firstbasementpos)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
