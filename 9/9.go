package main

import (
	"bufio"
	"fmt"
	"os"
)

// Vertex in graph: holds name of city, and a map of the distances
// to each adjacent city
type Vertex struct {
	name  string
	edges map[*Vertex]int
}

// Maps city names to their graph vertex structure
var Graph map[string]*Vertex = make(map[string]*Vertex)

// Best visit so far
var best_tour int

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var src, dst string
		var distance int
		line := scanner.Text()
		_, err := fmt.Sscanf(line, "%s to %s = %d", &src, &dst, &distance)
		check(err)
		if Graph[src] == nil {
			Graph[src] = &Vertex{src, make(map[*Vertex]int)}
		}
		if Graph[dst] == nil {
			Graph[dst] = &Vertex{dst, make(map[*Vertex]int)}
		}
		Graph[src].edges[Graph[dst]] = distance
		Graph[dst].edges[Graph[src]] = distance
	}

	best_tour = GetWorstTourLength()
	allVertices := AllVertices(Graph)
	for _, city := range Graph {
		city.Visit("", VerticesMinusVertex(allVertices, city), 0)
	}

	fmt.Printf("Best tour of %d cities: %d\n", len(allVertices), best_tour)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Modify this function to change whether this
// program finds the shortest or longest tour
func GetWorstTourLength() int {
	// For when metric is shortest tour
	return 10000000
	// For when metric is longest tour
	//return 0
}

// Modify this function to change whether this
// program finds the shortest or longest tour
func IsBetter(current int, best int) bool {
	// For when metric is shortest tour
	return current < best
	// For when metric is longest tour
	//return current > best
}

func AllVertices(graph map[string]*Vertex) []*Vertex {
	all := make([]*Vertex, len(graph))
	i := 0
	for _, vertex := range graph {
		all[i] = vertex
		i++
	}
	return all
}

func VerticesMinusVertex(vertices []*Vertex, vertex *Vertex) []*Vertex {
	remaining := make([]*Vertex, len(vertices)-1)
	i := 0
	for _, v := range vertices {
		if v != vertex {
			remaining[i] = v
			i++
		}
	}
	return remaining
}

// After travelling `distanceTravelledToHere` units to arrive at city,
// we then visit each city in `unvisited`, keeping track of the best
// tour from the current city. We return the length of the best tour
// from this city, to the remaining cities in `unvisited`.
func (city *Vertex) Visit(trail string, unvisited []*Vertex, distanceTravelledToHere int) int {
	var distanceTravelledFromHere int
	var best = GetWorstTourLength()

	for _, dest := range unvisited {
		//trail := fmt.Sprintf("%s%s (%d) ", trail, city.name, city.edges[dest])
		next_segment := city.edges[dest]

		distanceTravelledFromHere = next_segment + dest.Visit(trail, VerticesMinusVertex(unvisited, dest),
			distanceTravelledToHere+next_segment)
		if IsBetter(distanceTravelledFromHere, best) {
			best = distanceTravelledFromHere
		}
	}
	if len(unvisited) == 0 {
		//fmt.Printf("%s%s == %d\n", trail, city.name, distanceTravelledToHere)
		if IsBetter(distanceTravelledToHere, best_tour) {
			best_tour = distanceTravelledToHere
		}
		return 0
	} else {
		return best
	}
}
