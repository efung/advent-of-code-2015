package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)

	var wrap int
	var ribbon int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l, w, h int
		n, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &l, &w, &h)
		if n < 3 {
			panic(err)
		}
		wrap += SurfaceArea(l, w, h) + SmallestSideArea(l, w, h)
		ribbon += Volume(l, w, h) + SmallestPerimeter(l, w, h)
	}

	fmt.Printf("Total square feet: %d\n", wrap)
	fmt.Printf("Total ribbon length: %d\n", ribbon)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SurfaceArea(l int, w int, h int) int {
	return 2*l*w + 2*w*h + 2*h*l
}

func SmallestSideArea(l int, w int, h int) int {
	sorted := []int{l, w, h}
	sort.Ints(sorted)
	return sorted[0] * sorted[1]
}

func Volume(l int, w int, h int) int {
	return l * w * h
}

func SmallestPerimeter(l int, w int, h int) int {
	sorted := []int{l, w, h}
	sort.Ints(sorted)
	return 2*sorted[0] + 2*sorted[1]
}
