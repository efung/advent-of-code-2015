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
	CHILDREN, CATS, SAMOYEDS, POMERANIANS, AKITAS, VIZSLAS, GOLDFISH, TREES, CARS, PERFUMES int
}

func GiftSue() Sue {
	var s Sue = Sue{}
	s.id = 0
	s.SetProperty("CHILDREN", 3)
	s.SetProperty("CATS", 7)
	s.SetProperty("SAMOYEDS", 2)
	s.SetProperty("POMERANIANS", 3)
	s.SetProperty("AKITAS", 0)
	s.SetProperty("VIZSLAS", 0)
	s.SetProperty("GOLDFISH", 5)
	s.SetProperty("TREES", 3)
	s.SetProperty("CARS", 2)
	s.SetProperty("PERFUMES", 1)
	return s
}

func NewSue(id int) Sue {
	var s Sue = Sue{}
	s.id = id
	s.SetProperty("CHILDREN", -1)
	s.SetProperty("CATS", -1)
	s.SetProperty("SAMOYEDS", -1)
	s.SetProperty("POMERANIANS", -1)
	s.SetProperty("AKITAS", -1)
	s.SetProperty("VIZSLAS", -1)
	s.SetProperty("GOLDFISH", -1)
	s.SetProperty("TREES", -1)
	s.SetProperty("CARS", -1)
	s.SetProperty("PERFUMES", -1)
	return s
}

func (aSue *Sue) SetProperty(property string, value int) {
	ps := reflect.ValueOf(aSue).Elem()
	p := ps.FieldByName(property)
	if p.IsValid() && p.CanSet() {
		p.SetInt(int64(value))
	}
}

func (this Sue) Compare(that Sue) bool {
	if this.CHILDREN != -1 && that.CHILDREN != -1 && this.CHILDREN != that.CHILDREN {
		fmt.Println("Children mismatch")
		return false
	}
	if this.CATS != -1 && that.CATS != -1 && this.CATS != that.CATS {
		fmt.Println("cats mismatch")
		return false
	}
	if this.SAMOYEDS != -1 && that.SAMOYEDS != -1 && this.SAMOYEDS != that.SAMOYEDS {
		fmt.Println("sam mismatch")
		return false
	}
	if this.POMERANIANS != -1 && that.POMERANIANS != -1 && this.POMERANIANS != that.POMERANIANS {
		fmt.Println("pom mismatch")
		return false
	}
	if this.AKITAS != -1 && that.AKITAS != -1 && this.AKITAS != that.AKITAS {
		fmt.Println("aki mismatch")
		return false
	}
	if this.VIZSLAS != -1 && that.VIZSLAS != -1 && this.VIZSLAS != that.VIZSLAS {
		fmt.Println("viz mismatch")
		return false
	}
	if this.GOLDFISH != -1 && that.GOLDFISH != -1 && this.GOLDFISH != that.GOLDFISH {
		fmt.Println("gf mismatch")
		return false
	}
	if this.TREES != -1 && that.TREES != -1 && this.TREES != that.TREES {
		fmt.Println("trees mismatch")
		return false
	}
	if this.CARS != -1 && that.CARS != -1 && this.CARS != that.CARS {
		fmt.Println("cars mismatch")
		return false
	}
	if this.PERFUMES != -1 && that.PERFUMES != -1 && this.PERFUMES != that.PERFUMES {
		fmt.Println("perf mismatch")
		return false
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
