package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var Guests map[string]*Guest = make(map[string]*Guest)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)

	regex := regexp.MustCompilePOSIX(`([[:alpha:]]+) would (gain|lose) ([[:digit:]]+) happiness units by sitting next to ([[:alpha:]]+)\.`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res := regex.FindStringSubmatch(line)

		me := res[1]
		them := res[4]
		gainlose := res[2]
		deltaHappiness, err := strconv.Atoi(res[3])
		check(err)

		if Guests[me] == nil {
			Guests[me] = NewGuest(me)
		}
		sign := 1
		if gainlose == "gain" {
			sign = 1
		} else if gainlose == "lose" {
			sign = -1
		}
		Guests[me].happiness[them] = sign * deltaHappiness
	}

	includeSelf := false
	if includeSelf {
		selfGuest := NewGuest("self")
		for k, _ := range Guests {
			selfGuest.happiness[k] = 0
			Guests[k].happiness[selfGuest.name] = 0
		}
		Guests["self"] = selfGuest
	}

	// Because the table is circular, we want to avoid generating permutations that are
	// circular permutations of each other. We do this by fixing one element, then
	// generating permutations of the remaining ones.
	p := make([]string, len(Guests))
	i := 0
	for k, _ := range Guests {
		p[i] = k
		i++
	}

	iter := PermuteIter(p[1:])
	max_happiness := 0
	for sign := iter(); sign != 0; sign = iter() {
		h := TableHappiness(p)
		if h > max_happiness {
			fmt.Printf("%v = %d\n", p, h)
			max_happiness = h
		}
	}

	fmt.Printf("max(happiness): %d\n", max_happiness)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Guest struct {
	name      string
	happiness map[string]int
}

func NewGuest(guestName string) *Guest {
	return &Guest{name: guestName, happiness: make(map[string]int)}
}

func (me *Guest) SeatedNextTo(guestName string) int {
	return me.happiness[guestName]
}

func TableHappiness(names []string) int {
	happiness := 0
	for i, guest := range names {
		var left, right int
		if i == 0 {
			left = len(names) - 1
			right = i + 1
		} else if i == len(names)-1 {
			left = i - 1
			right = 0
		} else {
			left = i - 1
			right = i + 1
		}

		happiness += Guests[guest].SeatedNextTo(names[left]) + Guests[guest].SeatedNextTo(names[right])
	}
	return happiness
}
