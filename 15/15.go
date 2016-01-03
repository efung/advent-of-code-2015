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
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)

	var ingredientsSlice []Ingredient = []Ingredient{}

	regex := regexp.MustCompilePOSIX(`([[:alpha:]]+): capacity ([0-9-]+), durability ([0-9-]+), flavor ([0-9-]+), texture ([0-9-]+), calories ([0-9-]+)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res := regex.FindStringSubmatch(line)

		name := res[1]
		cap, err := strconv.Atoi(res[2])
		check(err)
		dur, err := strconv.Atoi(res[3])
		check(err)
		fla, err := strconv.Atoi(res[4])
		check(err)
		tex, err := strconv.Atoi(res[5])
		check(err)
		cal, err := strconv.Atoi(res[6])
		check(err)

		ingredientsSlice = append(ingredientsSlice, Ingredient{name, cap, dur, fla, tex, cal})
	}

	comp := make([]int, len(ingredientsSlice))
	comp[len(comp)-1] = 100
	highScore := 0
	for {
		comp = NextCompositionInKParts(comp, 100)
		if comp == nil {
			break
		}
		score := Score(comp, ingredientsSlice)
		var exactCals int
		if len(os.Args) == 3 {
			exactCals, err = strconv.Atoi(os.Args[2])
			check(err)
		} else {
			exactCals = -1
		}
		if score > highScore && SatisfiesCalories(comp, ingredientsSlice, exactCals) {
			fmt.Printf("%v == %d\n", comp, score)
			highScore = score
		}
	}
	fmt.Printf("Best score: %d\n", highScore)
}

func SatisfiesCalories(composition []int, ingredients []Ingredient, calories int) bool {
	if calories == -1 {
		return true
	} else {
		return Calories(composition, ingredients) == calories
	}
}

func NextCompositionInKParts(composition []int, N int) []int {
	k := len(composition)
	out := make([]int, k)
	var leftSum int
	copy(out, composition)
	for i := k - 2; i >= 0; i-- {
		leftSum = SliceSum(out[0 : k-1])
		if leftSum < N {
			out[i]++
			break
		} else {
			out[i] = 0
		}
	}
	leftSum = SliceSum(out[0 : k-1])
	if leftSum == 0 {
		return nil
	} else {
		out[k-1] = N - leftSum
		return out
	}
}

func SliceSum(slice []int) int {
	var sum int
	for _, a := range slice {
		sum += a
	}
	return sum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Ingredient struct {
	name                    string
	cap, dur, fla, tex, cal int
}

func Score(partition []int, ingredients []Ingredient) int {
	cap := Capacity(partition, ingredients)
	if cap == 0 {
		return 0
	}
	dur := Durability(partition, ingredients)
	if dur == 0 {
		return 0
	}
	fla := Flavor(partition, ingredients)
	if fla == 0 {
		return 0
	}
	tex := Texture(partition, ingredients)
	if tex == 0 {
		return 0
	}

	return cap * dur * fla * tex
}

func Capacity(partition []int, ingredients []Ingredient) int {
	return CalculateProperty(partition, ingredients, "cap")
}
func Durability(partition []int, ingredients []Ingredient) int {
	return CalculateProperty(partition, ingredients, "dur")
}
func Flavor(partition []int, ingredients []Ingredient) int {
	return CalculateProperty(partition, ingredients, "fla")
}
func Texture(partition []int, ingredients []Ingredient) int {
	return CalculateProperty(partition, ingredients, "tex")
}
func Calories(partition []int, ingredients []Ingredient) int {
	return CalculateProperty(partition, ingredients, "cal")
}

func CalculateProperty(partition []int, ingredients []Ingredient, property string) int {
	sum := 0
	for i, ingredient := range ingredients {
		ps := reflect.ValueOf(&ingredient).Elem()
		p := ps.FieldByName(property)
		if p.IsValid() {
			sum += int(p.Int()) * partition[i]
		}
	}

	if sum < 0 {
		return 0
	} else {
		return sum
	}
}
