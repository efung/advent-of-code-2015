package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)

	nice := 0
	nice2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if AtLeastThreeVowels(word) && AtLeastOneDoubleLetter(word) && HasNoForbiddenPairs(word) {
			nice++
		}
		if HasRepeatingPair(word) && HasRepeatingLetterWithIntervening(word) {
			nice2++
		}
	}

	fmt.Printf("Nice words: %d\n", nice)
	fmt.Printf("Revised nice words: %d\n", nice2)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func AtLeastThreeVowels(word string) bool {
	matched, err := regexp.MatchString(".*[aeiou].*[aeiou].*[aeiou].*", word)
	if err != nil {
		fmt.Println(err.Error())
	}
	return matched
}

func AtLeastOneDoubleLetter(word string) bool {
	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			return true
		}
	}
	return false
}

func HasNoForbiddenPairs(word string) bool {
	matched, err := regexp.MatchString("ab|cd|pq|xy", word)
	if err != nil {
		fmt.Println(err.Error())
	}
	return !matched
}

func HasRepeatingPair(word string) bool {
	for i := 0; i < len(word)-3; i++ {
		pair := word[i : i+2]
		for j := i + 2; j < len(word)-1; j++ {
			candidate := word[j : j+2]
			if pair == candidate {
				return true
			}
		}
	}
	return false
}

func HasRepeatingLetterWithIntervening(word string) bool {
	for i := 0; i < len(word)-2; i++ {
		c1 := word[i]
		c3 := word[i+2]
		if c1 == c3 {
			return true
		}
	}
	return false
}
