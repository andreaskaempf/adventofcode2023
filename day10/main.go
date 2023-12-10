// Advent of Code 2023, Day 10
//
// Given a 2D terrain of symbols, start at 'S' and follow the chain of "pipe"
// shapes back to S. For Part 1, calculate the distance of the furthest-away
// shape, distance measured in either direction from start. For Part 2,
// count up how many positions are *not* inside the shape.
//
// AK, 10 Dec 2023

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Data is stored in a global list of lines
var lines [][]byte

// A point is defined as x,y position (necessary for map key)
type Point struct {
	x, y int
}

// For Part 1, the shortest distance to any point on path
var dist map[Point]int

// For Part 2, keep track of any position that has ever been visited
var everVisited map[Point]bool

func main() {

	// Read the input file (must have no empty lines at end)
	fname := "sample.txt"
	fname = "sample2c.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	lines = bytes.Split(data, []byte("\n"))

	// Find starting position
	var p, start Point
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			p = Point{x, y}
			if at(p) == 'S' {
				start = p
				break
			}
		}
	}
	fmt.Println("Start at", start)

	// Follow all paths from start, collect distances
	dist = map[Point]int{}         // Initialize shortest distance to any point on path
	visited := []Point{}           // Points already visited on this traversal
	everVisited = map[Point]bool{} // For part 2
	explore(start, 0, visited)     // Explore from start, get distances

	// Part 1: Find the maximum distance from start
	ans1 := 0
	for _, v := range dist {
		if v > ans1 {
			ans1 = v
		}
	}
	fmt.Println("Part 1:", ans1) // 6927

	// Part 2: Use ray tracing method on rows only, to count up points
	// that are enclosed by path. Involves traversing each row, and counting
	// up how many times we cross the outline of our path. Where the count
	// is odd, we are inside the shape, otherwise outside.
	// Source: https://www.quora.com/What-are-the-algorithms-for-determining-if-a-point-is-inside-an-arbitrary-closed-shape-or-not
	ans2 := 0
	for y := 0; y < len(lines); y++ { // each row
		bars := 0                            // number of borders encountered on this row
		for x := 0; x < len(lines[0]); x++ { // each character on row
			p := Point{x, y}
			if everVisited[p] && oneOf(at(p), "S|LJ") {
				bars++ // we just crossed another border
			} else if !everVisited[p] && bars%2 == 1 {
				ans2++ // odd number of borders crossed, inside shape
			}
		}
	}

	fmt.Println("Part 2 for", fname, ":", ans2) // 467
}

// Recursively explore all paths from this point, mark points as visited,
// and update global dictionary with shortest distances to each point
func explore(here Point, steps int, visited []Point) {

	// Mark this point as visited
	visited = append(visited, here) // local list
	everVisited[here] = true        // global, for Part 2

	// Update distance to this point if shorter
	if dist[here] == 0 || steps < dist[here] {
		dist[here] = steps
	}

	// Explore each point from here that has not been explored
	for _, next := range nextSteps(here) {
		if !in(next, visited) {
			explore(next, steps+1, visited)
		}
	}
}

// Return a list of the viable next steps from this position, depends on the
// current and next shape being compatible connectors.
// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
func nextSteps(p Point) []Point {

	res := []Point{} // list of possible moves
	c0 := at(p)      // character at this position

	// Above: bar, F, 7
	p1 := Point{p.x, p.y - 1}
	c := at(p1)
	if oneOf(c0, "S|LJ") && oneOf(c, "|F7") {
		res = append(res, p1)
	}

	// Below: vert bar, L, J
	p1 = Point{p.x, p.y + 1}
	c = at(p1)
	if oneOf(c0, "S|F7") && oneOf(c, "|LJ") {
		res = append(res, p1)
	}

	// Left: dash, F, L
	p1 = Point{p.x - 1, p.y}
	c = at(p1)
	if oneOf(c0, "S-J7") && oneOf(c, "-FL") {
		res = append(res, p1)
	}

	// Right: dash, 7, J
	p1 = Point{p.x + 1, p.y}
	c = at(p1)
	if oneOf(c0, "S-FL") && oneOf(c, "-7J") {
		res = append(res, p1)
	}

	return res
}

// Character at a point, 0 if out of bounds
func at(p Point) byte {
	if p.y < 0 || p.y >= len(lines) || p.x < 0 || p.x >= len(lines[0]) {
		return 0
	} else {
		return lines[p.y][p.x]
	}
}

// Is byte in string?
func oneOf(c byte, s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

// Is point in a list of points?
func in(p Point, pp []Point) bool {
	for _, x := range pp {
		if x == p {
			return true
		}
	}
	return false
}
