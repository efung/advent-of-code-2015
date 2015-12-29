package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}

	seq := os.Args[1]
	iters, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%d: %s\n", 0, seq)
	for i := 1; i <= iters; i++ {
		seq = LookAndSay(seq)
		//fmt.Printf("%d: %s\n", i, seq)
	}
	fmt.Printf("After %d iters, len(sequence): %d\n", iters, len(seq))
}

func LookAndSay(input string) string {
	runLength := 1
	var Char rune
	prevChar := rune(input[0])
	output := make([]string, 2*len(input))
	for _, Char = range input[1:] {
		if Char != prevChar {
			output = append(output, strconv.Itoa(runLength), string(prevChar))
			runLength = 1
		} else {
			runLength++
		}
		prevChar = Char
	}
	output = append(output, strconv.Itoa(runLength), string(prevChar))
	return strings.Join(output, "")
}
