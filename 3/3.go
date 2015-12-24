package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var tok rune
	x := 0
	y := 0
	var housemap = map[string]int{}
	delivered(housemap, x, y)
	for tok != scanner.EOF {
		tok = s.Scan()
		if tok != scanner.EOF {
			switch s.TokenText() {
			case "^":
				y++
			case "v":
				y--
			case ">":
				x++
			case "<":
				x--
			}
			delivered(housemap, x, y)
		}
	}

	fmt.Printf("Houses with presents: %d\n", len(housemap))
}

func delivered(housemap map[string]int, x int, y int) {
	var key string
	key = strconv.Itoa(x) + "X" + strconv.Itoa(y)
	housemap[key] = housemap[key] + 1
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
