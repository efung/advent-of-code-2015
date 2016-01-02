package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var R map[string]*Reindeer = make(map[string]*Reindeer)
var Pts map[string]int = make(map[string]int)

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
	for scanner.Scan() {
		line := scanner.Text()
		res := regex.FindStringSubmatch(line)
		speed, err := strconv.Atoi(res[2])
		check(err)
		power_dur, err := strconv.Atoi(res[3])
		check(err)
		rest_dur, err := strconv.Atoi(res[4])
		check(err)

		R[res[1]] = &Reindeer{res[1], speed, power_dur, rest_dur}
	}

	for t := 1; t <= race_length; t++ {
		lead_distance := 0
		var best_reindeers []string
		for _, r := range R {
			d := r.Travelled(t)
			if d > lead_distance {
				lead_distance = d
				best_reindeers = []string{r.name}
			} else if d == lead_distance {
				best_reindeers = append(best_reindeers, r.name)
			}
		}
		for _, winner := range best_reindeers {
			Pts[winner]++
		}
	}

	high_score := 0
	for name, pts := range Pts {
		fmt.Printf("%s has %d pts\n", name, pts)
		if pts > high_score {
			high_score = pts
		}
	}

	fmt.Printf("\nHigh score: %d\n", high_score)
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
