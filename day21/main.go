// Advent of Code 2023, Day 21
//
// AK, 22 Dec 2023

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// A point, used as key in maps
type Point struct {
	x, y int
}

// Global matrix of input data
var rows [][]byte        // each row of data
var nr, nc int           // number of rows and columns
var start Point          // starting point
var rocks map[Point]bool // where the rocks are

const INF int = 1000000 // infinity

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte{'\n'})

	// Remove last row if empty
	if len(rows[len(rows)-1]) == 0 {
		rows = rows[:len(rows)-1]
	}

	// Need number of rows & cols to detect out of bound
	nr = len(rows)
	nc = len(rows[0])
	gridToMap(rows) // sets rocks, start
	fmt.Println("Start at", start)

	// Part 1: number of squares that can be reached in n steps
	fmt.Println("Part 1:", part1(start, 64)) // sample: 6 -> 16
}

// Part 1: use Djikstra to get number of steps from start
// to any square, count up how many can be reached in n steps
func part1(start Point, n int) int {

	// Distance initially zero for start, undefined (infinity)
	// everywhere else; no points visited yet
	dist := map[Point]int{start: 0}
	visited := map[Point]bool{}

	// Process each point, to find the distance from start
	p := start // start at the current point
	for {

		// Find and process the points adjacent to p, i.e., those
		// one step away, but within bounds and not a rock
		for _, d := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {

			p1 := Point{p.x + d.x, p.y + d.y}
			if p1.x >= 0 && p1.x < nc && p1.y >= 0 && p1.y < nr && !rocks[p1] {

				// Set dist if lower than current value
				_, ok := dist[p1] // is there already a distance for this point?
				d1 := dist[p] + 1 // the distance to get here
				if d1 < dist[p1] || !ok {
					dist[p1] = d1 // Set the distance for this point if lower than previous
				}
			}
		}

		// Mark p as visited
		visited[p] = true

		// Find the next point to process, the one with the smallest distance
		// that has not been visited yet
		lowest := INF
		for q, d := range dist {
			if d < lowest && !visited[q] {
				p = q
				lowest = d
			}
		}

		// If no more points to process, we are done
		if lowest == INF {
			break
		}
	}

	// Count the number of distances that are <= n and even
	ans := 0
	dist[start] = 2 // so that start is included
	for _, d := range dist {
		if d > 0 && d <= n && d%2 == 0 {
			ans++
		}
	}

	return ans
}

// For debugging: print the grid
func printGrid() {
	for y := 0; y < nr; y++ {
		for x := 0; x < nc; x++ {
			c := rows[y][x]
			if c == '#' {
				fmt.Print("#")
			} else if x == start.x && y == start.y {
				fmt.Print("S")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Convert grid to map of where the rocks are,
// also returns the point that contains 'S'
func gridToMap(rows [][]byte) {
	rocks = map[Point]bool{} // remember rocks '#'
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[0]); x++ {
			c := rows[y][x]
			if c == 'S' {
				start = Point{x, y}
			} else if c == '#' {
				rocks[Point{x, y}] = true
			}
		}
	}
}
