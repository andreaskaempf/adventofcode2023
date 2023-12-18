// Advent of Code 2023, Day 17
//
// Given a matrix of digits, find the sum of the digits along the shortest
// path from the top left to the bottom right, such that you never go in
// the same direction more than 3 steps. For Part 2, maximum 10 steps in
// the same direction, minimum 4 steps before changing direction. Solved
// by modifying the Djistra algorithm, transcribed from my Julia solution
// to AoC 2021, Day 15, to enforce constraints on direction and number of
// steps in the same direction.
//
// AK, 17-18 Dec 2023

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Data, number of rows & cols
var rows [][]byte
var nr, nc int

const INF int = 1e6

// State of a node, used as a key for distance & visited maps, defined by
// direction and number of steps already take in that direction, as well as
// position
type State struct {
	x, y      int // position
	direction int // direction we were in when we got here
	run       int // number of steps already in the same direction
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
	//fname = "sample2.txt" // second example for part 2
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte("\n"))
	nr = len(rows)
	nc = len(rows[0])

	// Recursive solution correctly give 102 for the sample input,
	// but takes 84 mins
	//fmt.Println("Part 1 (Recursive):", attempt1())

	//fmt.Println("Part 1:", solve(false)) // 102, 724
	fmt.Println("Part 2:", solve(true)) // 94,
}

// Modified Djistra algorithm, transcribed from my Julia solution to
// AoC 2021, Day 15. Finds the lowest total cost path from top left to
// bottom right of a matrix. Modified to keep track of path for state,
// i.e., direction of travel and number of steps in that direction, as
// well as position, and to allow for the problem constraint that you
// can't move backwards or more than 3 steps in the same direction.
func solve(part2 bool) int {

	// All states initially unvisited
	visited := make(map[State]bool) // initially all false

	// Distance for every state, assumed to be infinity if missing
	dist := make(map[State]int)

	// The initial state: top left corner, no initial direction or streak
	s := State{0, 0, -1, 0}
	dist[s] = 0 // only the start node has a distance, all others infinity

	// Start algorithm, stops when we reach the bottom, or no more states to explore
	var maxDist int
	for {

		// Report progress
		if dist[s] > maxDist {
			maxDist = dist[s]
			fmt.Println(s.x, s.y, maxDist)
		}

		// Consider all unvisited neighbors one step away from the current
		// state.  This is not just x,y proximity, but takes into account the
		// current direction and the number of movements already in that
		// direction.
		for dir := 0; dir < 4; dir++ {

			// Part 1: skip this route if same direction as last time,
			// and already 3 steps
			if !part2 && dir == s.direction && s.run == 3 {
				continue
			}

			// Part 2: maximum 10 steps in same direction, minimum 4 steps
			if part2 {
				if dir == s.direction && s.run >= 10 { // max 10 same dir
					continue
				} else if !(s.x == 0 && s.y == 0) && dir != s.direction && s.run < 4 { // min 4 blocks
					continue // before changing direction
				}
			}

			// Can't reverse
			if (dir == up && s.direction == down) || (dir == down && s.direction == up) || (dir == left && s.direction == right) || (dir == right && s.direction == left) {
				continue
			}

			// Get the next location one step away, based on direction
			x := s.x
			y := s.y
			if dir == up {
				y--
			} else if dir == down {
				y++
			} else if dir == left {
				x--
			} else { // right
				x++
			}

			// Can't go here if out of bounds
			if x < 0 || x >= nc || y < 0 || y >= nr {
				continue
			}

			// Create new state for proposed movement, skip if already visited
			run := 1                // 1 step if starting in new direction
			if dir == s.direction { // if going in same direction,
				run += s.run // add prev run
			}
			s1 := State{x, y, dir, run}

			// Skip if already visited, so we never backtrack
			if visited[s1] { // will be false if not in dictionary yet
				continue
			}

			// Get cost of getting to this location is cost of getting to
			// previous location, plus the cost of this one
			c1 := dist[s] + int(rows[y][x]-'0')

			// Update cost for the new node if lower than previous
			_, ok := dist[s1]
			if !ok || c1 < dist[s1] {
				dist[s1] = c1
			}
		}

		// Mark this state as visited, so we don't return to it
		visited[s] = true

		// Find the next state to explore, the one with the lowest cost in the
		// matrix, ignoring cells already visited; this emulates the
		// functionality of a priority queue
		var lowestDist int = INF
		for s1 := range dist {
			d, _ := dist[s1]
			if d < lowestDist && !visited[s1] {
				lowestDist = d
				s = s1 // the next state to be explored
			}
		}

		// If the selected state is the bottom, we are done, as it's the
		// first and cheapest route to the bottom
		if s.x == nc-1 && s.y == nr-1 {
			fmt.Println("Solution found")
			return dist[s]
		}

		// Stop when no more states (should not happen)
		if lowestDist == INF {
			fmt.Println("No solution found")
			break
		}
	}

	// No solution found
	return INF
}
