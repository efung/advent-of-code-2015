package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	password := os.Args[1]
	next := []rune(password)
	for {
		next = Increment(next)
		if HasStraight(next) && HasTwoPairs(next) {
			fmt.Printf("Next password: %s\n", string(next))
			break
		}
	}
}

func Increment(input []rune) []rune {
	output := make([]rune, 8)
	forbidden, i := HasForbiddenChars(input)
	if forbidden {
		for c := 0; c < len(input); c++ {
			if c < i {
				output[c] = input[c]
			} else if c == i {
				output[c] = input[c] + 1
			} else {
				output[c] = 'a'
			}
		}
	} else {
		output = input
		for c := len(input) - 1; c >= 0; c-- {
			done := true
			var c_incr rune
			switch input[c] {
			case 'z':
				c_incr = 'a'
				done = false
			case 'h':
				c_incr = 'j'
			case 'n':
				c_incr = 'p'
			case 'k':
				c_incr = 'm'
			default:
				c_incr = input[c] + 1
			}
			output[c] = c_incr
			if done {
				break
			}
		}
	}
	return output
}

func HasForbiddenChars(input []rune) (bool, int) {
	for i, c := range input {
		if c == 'i' || c == 'o' || c == 'l' {
			return true, i
		}
	}

	return false, 0
}

func HasStraight(input []rune) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i] <= 'x' {
			if input[i+1] == input[i]+1 && input[i+2] == input[i]+2 {
				return true
			}
		}
	}
	return false
}

func HasTwoPairs(input []rune) bool {
	var pair1, pair2 bool
	i := 0
	for ; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			pair1 = true
			break
		}
	}
	for j := i + 2; j < len(input)-1; j++ {
		if input[j] == input[j+1] {
			pair2 = true
			break
		}
	}
	return pair1 && pair2
}
