package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)

	grid := Make2DBoolArray(1000, 1000)
	brightgrid := Make2DIntArray(1000, 1000)
	regex := regexp.MustCompile(`([a-z ]+) (\d+),(\d+) through (\d+),(\d+)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res := regex.FindStringSubmatch(scanner.Text())
		cmd := res[1]
		var x1, y1, x2, y2 int
		x1, err = strconv.Atoi(res[2])
		y1, err = strconv.Atoi(res[3])
		x2, err = strconv.Atoi(res[4])
		y2, err = strconv.Atoi(res[5])
		switch cmd {
		case "turn on":
			Set(grid, x1, y1, x2, y2, true)
			Brightness(brightgrid, x1, y1, x2, y2, 1)
		case "toggle":
			Toggle(grid, x1, y1, x2, y2)
			Brightness(brightgrid, x1, y1, x2, y2, 2)
		case "turn off":
			Set(grid, x1, y1, x2, y2, false)
			Brightness(brightgrid, x1, y1, x2, y2, -1)
		}
	}

	lit := 0
	brightness := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				lit++
			}
			brightness += brightgrid[i][j]
		}
	}

	fmt.Printf("Lights lit: %d\n", lit)
	fmt.Printf("Brightness: %d\n", brightness)
}

func Set(grid [][]bool, x1 int, y1 int, x2 int, y2 int, value bool) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			grid[x][y] = value
		}
	}
}

func Toggle(grid [][]bool, x1 int, y1 int, x2 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			grid[x][y] = !grid[x][y]
		}
	}
}

func Brightness(grid [][]int, x1 int, y1 int, x2 int, y2 int, level int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			grid[x][y] += level
			if grid[x][y] < 0 {
				grid[x][y] = 0
			}
		}
	}
}

func Make2DBoolArray(rows int, cols int) [][]bool {
	a := make([][]bool, rows)
	for i := range a {
		a[i] = make([]bool, cols)
	}
	return a
}

func Make2DIntArray(rows int, cols int) [][]int {
	a := make([][]int, rows)
	for i := range a {
		a[i] = make([]int, cols)
	}
	return a
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
