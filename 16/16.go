package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)

	GiftSue := GiftSue()

	regex := regexp.MustCompile(`Sue ([[:digit:]]+): (?:([[:alpha:]]+): ([[:digit:]]+),?)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res := regex.FindStringSubmatch(line)

		n, err := strconv.Atoi(res[1])
		check(err)
		aSue := NewSue(n)
		for i := 2; i < len(res); i = i + 2 {
			attr := res[i]
			quantity, err := strconv.Atoi(res[i+1])
			check(err)
			aSue.SetProperty(attr, quantity)
		}
		//	allSues = append(allSues, aSue)
		if GiftSue.Compare(aSue) {
			fmt.Printf("%v != %v\n", GiftSue, aSue)
		}
	}
}

type Sue struct {
	id int
	// -1 means don't know; non-negative is known value
	children, cats, samoyeds, pomeranians, akitas, vizslas, goldfish, trees, cars, perfumes int
}

func GiftSue() Sue {
	s := NewSue(0)
	s.SetProperty("children", 3)
	s.SetProperty("cats", 7)
	s.SetProperty("samoyeds", 2)
	s.SetProperty("pomeranians", 3)
	s.SetProperty("akitas", 0)
	s.SetProperty("vizslas", 0)
	s.SetProperty("goldfish", 0)
	s.SetProperty("trees", 0)
	s.SetProperty("cars", 2)
	s.SetProperty("perfumes", 1)
	return s
}

func NewSue(id int) Sue {
	var s Sue = Sue{}
	s.id = id
	s.SetProperty("children", -1)
	s.SetProperty("cats", -1)
	s.SetProperty("samoyeds", -1)
	s.SetProperty("pomeranians", -1)
	s.SetProperty("akitas", -1)
	s.SetProperty("vizslas", -1)
	s.SetProperty("goldfish", -1)
	s.SetProperty("trees", -1)
	s.SetProperty("cars", -1)
	s.SetProperty("perfumes", -1)
	return s
}

func (aSue Sue) SetProperty(property string, value int) {
	ps := reflect.ValueOf(&aSue).Elem()
	p := ps.FieldByName(property)
	//if p.IsValid() && p.CanSet() {
	p.SetInt(int64(value))
	//	}
}

func (this Sue) Compare(that Sue) bool {
	if this.children != -1 && that.children != -1 && this.children != that.children {
		return false
	}
	if this.cats != -1 && that.cats != -1 && this.cats != that.cats {
		return false
	}
	if this.samoyeds != -1 && that.samoyeds != -1 && this.samoyeds != that.samoyeds {
		return false
	}
	if this.pomeranians != -1 && that.pomeranians != -1 && this.pomeranians != that.pomeranians {
		return false
	}
	if this.akitas != -1 && that.akitas != -1 && this.akitas != that.akitas {
		return false
	}
	if this.vizslas != -1 && that.vizslas != -1 && this.vizslas != that.vizslas {
		return false
	}
	if this.goldfish != -1 && that.goldfish != -1 && this.goldfish != that.goldfish {
		return false
	}
	if this.trees != -1 && that.trees != -1 && this.trees != that.trees {
		return false
	}
	if this.cars != -1 && that.cars != -1 && this.cars != that.cars {
		return false
	}
	if this.perfumes != -1 && that.perfumes != -1 && this.perfumes != that.perfumes {
		return false
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
