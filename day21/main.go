// Advent of Code 2023, Day 21
//
// Given a map of points and rocks, find the number of points
// that can be reached in n steps, starting from a given point.
// Used Djikstra's algorithm (and later simple walk simulation)
// to find the number of points that can be reached in n steps.
// For Part 2, assume a much larger number of steps, infeasible
// using brute force, so count tiles reached in 1x1, 3x3, and 5x5
// area blocks, and extrapolate from these.
//
// AK, 22 and 29 Dec 2023

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
	fname := "sample.txt" // make sure no blank row at end
	fname = "input.txt"   // uncomment this line to use input file
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte{'\n'})
	nr = len(rows)
	nc = len(rows[0])

	// Get a map of just the rocks, and the start location
	gridToMap(rows) // sets rocks, start
	fmt.Println("Grid", nr, "rows x", nc, "cols, start at", start)

	// Part 1: number of squares that can be reached in 64 steps
	fmt.Println("Part 1:", walk(start, 64)) // 3578

	// Part 2: how many places could be visited in 26501365 steps?
	// The pattern is 131 wide, so each square "block" allows for 65 steps
	// from the centre. Calculate the number of tiles that can be visited in:
	// x0: 65 steps, i.e., within the inner square
	// x1: 65+one pattern width, i.e., 3x3 squares
	// x2: 65+two pattern widths, i.e, 5x5 squares
	x0 := int64(walk(start, 65))      // inner square = 3676
	x1 := int64(walk(start, 65+nc))   // 3x3 squares = 32808
	x2 := int64(walk(start, 65+nc*2)) // 5x5 squares = 90974

	// The whole area repeats itself, so calculate the width of
	// the 2.6M steps in terms of 131x131 squares
	var w int64 = 26501365 / int64(nc) // integer division, 202300
	//fmt.Println("x1/x2/x3 =", x0, x1, x2, ", w =", w)

	// Use the tiles reached across 1x1, 3x3 and 5x5 blocks
	ans := (x2 - 2*x1 + x0) / 2 // steps reached in one horizontal slice
	ans *= (w - 1)              // multiply by width in blocks
	ans += x1 - x0              // add the ring outside the middle
	ans *= w                    // multiply by height in blocks
	ans += x0                   // add the inner block
	fmt.Println("Part 2:", ans) // 594115391548176

	for i := 1; i <= 10; i++ {
		n := i * 65
		fmt.Println(n, walk(start, n))
	}
}

// Simple simulation of walking n steps from start, returns
// how many positions were visited in exactly n steps (i.e.,
// not any positions visited during the n steps, just the
// positions that are possible after exactly n steps).
// Uses breadth-first search with a queue.
func walk(start Point, n int) int {

	// Queue of points to explore, starting with start position
	assert(rows[start.y][start.x] == 'S', "Not starting on S")
	Q := map[Point]int{start: 1}

	// Explore up to the given number of steps, note that you don't need
	// to check for bounds in Part 2, even though you are going beyond the
	// edges of the original 131x131 area, since isRock() adjusts coordinates
	// to reflect that the pattern repeats itself indefinitely.
	var nextQ map[Point]int
	for s := 0; s < n; s++ {
		nextQ = map[Point]int{} // points to be explored in the next iteration
		for p, _ := range Q {
			for _, d := range []Point{Point{-1, 0}, Point{1, 0}, Point{0, -1}, Point{0, 1}} {
				p1 := Point{p.x + d.x, p.y + d.y} // Get adjacent point
				if !isRock(p1) {                  // Add to queue if not a rock
					nextQ[p1] = 1 // also acts as list of points visited in this iteration
				}
			}
		}
		Q = nextQ
	}

	// After the n steps have elapsed, the list of positions we would explore
	// next is effectively the number of tiles we have reached in exactly
	// n steps
	return len(nextQ)

}

// Is there a rock at this point? Accounts for 131x131 map being
// repeatedly repeated in any direction, so we don't need to check
// for bounds in Part 2 (TODO: use modulo)
func isRock(p Point) bool {
	for p.x < 0 {
		p.x += nc
	}
	for p.x >= nc {
		p.x -= nc
	}
	for p.y < 0 {
		p.y += nr
	}
	for p.y >= nr {
		p.y -= nr
	}
	assert(p.x >= 0 && p.x < nc && p.y >= 0 && p.y < nr, "Out of range")
	return rocks[p]
}

// Part 1: used Djikstra to get number of steps from start
// to any square, count up how many can be reached in exactly n steps.
// Newer walk() function is simpler, and can be used for Part 1
func djikstra(start Point, n int) int {

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
			//if /*p1.x >= 0 && p1.x < nc && p1.y >= 0 && p1.y < nr &&*/ !rocks[p1] {
			if !isRock(p1) {

				// Set dist if lower than current value
				d1 := dist[p] + 1 // the distance to get here
				if d1 <= n {      // don't bother if too many steps
					_, ok := dist[p1] // is there already a distance for this point?
					if d1 < dist[p1] || !ok {
						dist[p1] = d1 // Set the distance for this point if lower than previous
					}
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

// For debugging: print the grid, shows a diamond pattern with some
// unreachable positions after walk.
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
// also sets 'start' to the point that contains 'S'
func gridToMap(rows [][]byte) {
	rocks = map[Point]bool{} // remember rocks '#'
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[0]); x++ {
			c := rows[y][x]
			if c == 'S' {
				start = Point{x, y}
			} else if c == '#' {
				rocks[Point{x, y}] = true // global variable
			}
		}
	}
}

// Panic if a test condition is not true
func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
