// Advent of Code 2023, Day 23
//
// Find the maximum number of steps (longest path) from top to bottom of a
// grid, ignorning blocks (hash marks), and not revisiting previous cells. In
// Part 1, pointer chars indicate that you must move in that direction
// (restriction removed in Part 2). Hard problem, used recursive depth-
// first search (brute force), but could probably simplify the graph to
// reduce the search space.
//
// AK, 23 Dec 2023

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
var rows [][]byte // each row of data
var nr, nc int    // number of rows and columns

// State information
var npaths, maxPath int
var part2 bool

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte{'\n'})

	// Need number of rows & cols to detect out of bounds
	nr = len(rows)
	nc = len(rows[0])

	// Find the start, explore from there
	var startX int
	for x := 0; x < nc; x++ {
		if rows[0][x] == '.' {
			startX = x
		}
	}

	// Part 1: Start exploring from here (94, 2430)
	initVis := map[Point]bool{}
	explore(Point{startX, 0}, initVis, 0) // take the first step
	fmt.Println("Part 1:", npaths, "paths found, longest =", maxPath)

	// Part 2: same, but ignore slopes (154, 6534)
	part2 = true
	maxPath = 0
	npaths = 0
	explore(Point{startX, 0}, initVis, 0)
	fmt.Println("Part 2:", npaths, "paths found, longest =", maxPath)

	// Attempted Djikstra (does not work for longest path)
	//ans := djikstra(Point{startX, 0}, Point{endX, nr - 1})
	//fmt.Println("Part 2:", ans) // 154,
}

// Recursively explore from this point, not stepping on any points visited
// already, oserving rules, set global maxPath to length of longest path
func explore(p Point, visited map[Point]bool, steps int) {

	// If we are on a period on the bottom row, we have arrived
	c := rows[p.y][p.x]
	if c == '.' && p.y == nr-1 {
		npaths++
		if steps > maxPath {
			maxPath = steps
		}
		fmt.Println("Path", npaths, "found, length =", steps, ", longest =", maxPath)
		return
	}

	// For Part 2, ignore slopes
	if part2 && in(c, []byte{'^', 'v', '<', '>'}) {
		c = '.'
	}

	// Determine the  possible points from here
	var allowed []Point
	if c == '^' { // only up
		allowed = []Point{{0, -1}}
	} else if c == 'v' { // only down
		allowed = []Point{{0, 1}}
	} else if c == '<' { // only left
		allowed = []Point{{-1, 0}}
	} else if c == '>' { // only right
		allowed = []Point{{1, 0}}
	} else if c == '.' { // down, right, left, up (last)
		allowed = []Point{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	} else {
		panic("Bad char")
	}

	// Explore in each allowed direction
	for _, d := range allowed {

		// Must be in range, not yet visited
		p1 := Point{p.x + d.x, p.y + d.y}
		if p1.x < 0 || p1.x >= nc || p1.y < 0 || p1.y >= nr || visited[p1] {
			continue
		}

		// Don't go into a wall
		if rows[p1.y][p1.x] == '#' {
			continue
		}

		// Make a copy of the visited map, adding this point
		vis1 := make(map[Point]bool, len(visited)+1)
		for k, v := range visited {
			vis1[k] = v
		}
		vis1[p] = true

		// Explore new point
		explore(p1, vis1, steps+1)
	}
}

// Attempted alternative solution to finding longest path for Part 2,
// using Djikstra modified for longest path instead of shortest
// Does not work, used brute force instead
func djikstra(start, end Point) int {

	// Distance initially zero for start, undefined (infinity)
	// everywhere else; no points visited yet
	dist := map[Point]int{start: 0}
	visited := map[Point]bool{}

	// Process each point, to find the distance from start
	p := start
	for {

		// Find and process the points adjacent to p, i.e., those
		// one step away, but within bounds and not a rock
		for _, d := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {

			// Get this point, ignore if visited, a wall, or out of bounds
			p1 := Point{p.x + d.x, p.y + d.y}
			if p1.x < 0 || p1.x >= nc || p1.y < 0 || p1.y >= nr {
				continue // out of bounds
			}
			if rows[p1.y][p1.x] == '#' || visited[p1] {
				continue // a wall or already visited
			}

			// Set dist if *higher* than current value
			d1 := dist[p] - 1  // the distance to get here
			if d1 < dist[p1] { // since map default is zero, no need to check if already in dist
				dist[p1] = d1 // Set the distance for this point if lower than previous
			}
		}

		// Mark p as visited
		visited[p] = true

		// Find the next point to process, any one
		// that has not been visited yet
		best := 99999999
		for p1, d1 := range dist {
			if d1 < best && !visited[p1] {
				p = p1
				best = d1
			}
		}

		// If no more points to process, we are done
		if best == 99999999 {
			break
		}
	}

	// Distance to end
	return dist[end] * -1
}

// Is element in a list?
func in[T int | float64 | byte | string](c T, s []T) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}
