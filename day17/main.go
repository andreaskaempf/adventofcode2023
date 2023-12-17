// Advent of Code 2023, Day 17
//
//
//
// AK, 17 Dec 2023

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Data, number of rows & cols
var rows [][]byte
var nr, nc int

// A 2D point, required because we use x,y pair as map key
type Point struct {
	x, y int
}

// The four directions, in the directions that are mostly likely to take
// us bottom right
const (
	down = iota
	right
	left
	up
)

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample.txt"
	//fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte("\n"))
	nr = len(rows)
	nc = len(rows[0])
	// fmt.Println("Part 1 (Recursive):", attempt1())
	fmt.Println("Part 1 (Djikstra):", solve()) // 102,
	// 1641 too high, 1640 wrong, 724
}

const PENALTY int = 1e6

// Modified Djistra algorithm, transcribed from my Julia solution to
// AoC 2021, Day 15. Finds the lowest total cost path from top left to
// bottom right of a matrix. Modified for the problem constraint that you
// can't move more than 3 steps in the same direction.
func solve() int {

	// Mark all nodes as unvisited
	visited := make(map[Point]bool) // default false
	var path []Point

	// Assign to every node a tentative distance value: zero for
	// initial node and infinity for all other nodes
	dist := make(map[Point]int)
	for y := 0; y < nr; y++ {
		for x := 0; x < nc; x++ {
			p := Point{x, y}
			dist[p] = PENALTY
		}
	}

	// The current point, starts at top left, dist zero
	p := Point{0, 0}
	dist[p] = 0

	// Start algorithm iterations, stop when reach bottom right
	for !(p.y == nr-1 && p.x == nc-1) {

		// Add this point to the optimal path
		path = append(path, p)

		// Consider unvisited neighbors in each direction
		for dir := 0; dir < 4; dir++ {

			// Get the next location, based on direction
			p1 := p
			if dir == up {
				p1.y--
			} else if dir == down {
				p1.y++
			} else if dir == left {
				p1.x--
			} else if dir == right {
				p1.x++
			} else {
				panic("Invalid direction")
			}

			// Can't go here if already visited or out of bounds
			if visited[p1] || p1.x < 0 || p1.x >= nc || p1.y < 0 || p1.y >= nr {
				continue
			}

			// Can't go here if already 3 steps in this direction
			if sameSteps(dir, 3, path) {
				continue
			}

			// Update cheapest cost of getting to this location
			costHere := int(rows[p1.y][p1.x] - '0')
			c1 := costHere + dist[p]
			if c1 < dist[p1] {
				dist[p1] = c1
			}

		}

		// Mark this node as visited
		visited[p] = true

		// Next point will be the one that has minimum value in the matrix,
		// ignoring cells already visited
		var lowestDist int = PENALTY
		for x := 0; x < nc; x++ {
			for y := 0; y < nr; y++ {
				p1 := Point{x, y}
				d := dist[p1]
				if d < lowestDist && d < PENALTY && !visited[p1] {
					p = p1
				}
			}
		}

	}

	//Return solution, cost in the bottom right corner
	last := Point{nc - 1, nr - 1}
	return dist[last]

}

// Determine if n steps have been taken in the same direction
func sameSteps(checkDir, n int, path []Point) bool {

	// Obviously false if path too short
	if len(path) < n+1 {
		return false
	}

	// Only care about the last three steps
	if len(path) > n+1 {
		path = path[len(path)-n-1:]
		assert(len(path) == n+1, "Bad path slice")
	}

	// Check each direction, return false if not same as checking
	dirs := []int{}
	for i := 1; i < len(path); i++ {
		p0 := path[i-1]
		p1 := path[i]
		var dir1 int
		if p1.x > p0.x {
			dir1 = right
		} else if p1.x < p0.x {
			dir1 = left
		} else if p1.y > p0.y {
			dir1 = down
		} else if p1.y < p0.y {
			dir1 = up
		} else {
			fmt.Println(path)
			panic("Invalid movement")
		}
		dirs = append(dirs, dir1)
	}
	assert(len(dirs) == n, "Bad dirs len")
	//fmt.Println(path, dirs)

	// Check each direction
	for _, d := range dirs {
		if d != checkDir {
			return false
		}
	}

	// If we get to here, the last three directions are the same as the
	// one being checked
	return true
}

// Panic if a test condition is not true
func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
