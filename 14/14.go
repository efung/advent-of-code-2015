package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)
	race_length, err := strconv.Atoi(os.Args[2])
	check(err)

	regex := regexp.MustCompilePOSIX(`([[:alpha:]]+) can fly ([[:digit:]]+) km/s for ([[:digit:]]+) seconds, but then must rest for ([[:digit:]]+) seconds\.`)
	scanner := bufio.NewScanner(file)
	bestDistance := 0
	for scanner.Scan() {
		line := scanner.Text()
		res := regex.FindStringSubmatch(line)
		speed, err := strconv.Atoi(res[2])
		check(err)
		power_dur, err := strconv.Atoi(res[3])
		check(err)
		rest_dur, err := strconv.Atoi(res[4])
		check(err)

		r := &Reindeer{res[1], speed, power_dur, rest_dur}
		d := r.Travelled(race_length)
		fmt.Printf("%s travels %d km\n", r.name, d)
		if d > bestDistance {
			bestDistance = d
		}
	}
	fmt.Printf("\nBest distance after %d s: %d km\n", race_length, bestDistance)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Reindeer struct {
	name      string
	speed     int
	power_dur int
	rest_dur  int
}

func (r *Reindeer) Travelled(length int) int {
	period := r.power_dur + r.rest_dur
	d1 := (length / period) * r.power_dur * r.speed
	rem := length % period
	d2 := min(rem, r.power_dur) * r.speed
	return d1 + d2
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
