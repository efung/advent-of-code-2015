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
	santax := 0
	santay := 0
	robox := 0
	roboy := 0
	housemap := map[string]int{}
	delivered(housemap, santax, santay)
	delivered(housemap, robox, roboy)
	for tok != scanner.EOF {
		tok = s.Scan()
		if tok != scanner.EOF {
			switch s.TokenText() {
			case "^":
				santay++
			case "v":
				santay--
			case ">":
				santax++
			case "<":
				santax--
			}
			delivered(housemap, santax, santay)
		}
		tok = s.Scan()
		if tok != scanner.EOF {
			switch s.TokenText() {
			case "^":
				roboy++
			case "v":
				roboy--
			case ">":
				robox++
			case "<":
				robox--
			}
			delivered(housemap, robox, roboy)
		}
	}

	fmt.Printf("Houses with presents (Santa and Robo-Santa): %d\n", len(housemap))
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
