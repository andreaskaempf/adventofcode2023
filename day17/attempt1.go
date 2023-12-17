// Initial recursive attempt, takes too long

package main

import (
	"fmt"
)

// To keep track of the best solution found so far
var bestSoFar int

// Memoization does not work. Tried current position, current direction,
// steps so far in this direction
/*type Config struct {
	pos      Point
	dir, ssd int
	vis      string
}

var Memo map[Config]int
*/

// First attempt: recursive depth-first search with pruning, but no memoization
func attempt1() int {

	//Memo = make(map[Config]int)
	bestSoFar = 1e6 // initial value of best solution found so far

	// Part 1: move from 0,0 to bottom right corner, only 90 degree turns,
	// max 3 steps in the same direction, to as to minimize total weight
	start := Point{} // starting position 0,0
	initCost := 0
	initDir := -1 // initially no direction
	initSSD := 0  //  have not yet travelled any number of steps in any direction
	initVis := make(map[Point]bool)
	return nav(start, initDir, initSSD, initVis, initCost)
}

// Recursively get the best from this position
func nav(p Point, curDir, stepsSameDir int, visited map[Point]bool, curCost int) int {

	// Add the cost of this square to the total cost, don't include cost of first cell
	if !(p.x == 0 && p.y == 0) {
		curCost += int(rows[p.y][p.x] - '0')
	}

	// Pruning: skip this branch if current cost is already higher than best
	// achieved so far
	if curCost >= bestSoFar {
		return 1e6
	}

	// If we have arrived at the bottom right, return the current cost
	if p.x == nc-1 && p.y == nr-1 {
		return curCost
	}

	// Try each direction, and find the lowest score
	var best int = 1e6
	for dir := 0; dir < 4; dir++ {

		// Can't go in this direction if already three steps in that direction
		if dir == curDir && stepsSameDir >= 3 {
			continue
		}

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

		// Create copy of visited map, adding this point
		vis := make(map[Point]bool, len(visited)+1)
		for k, v := range visited {
			vis[k] = v
		}
		vis[p] = true // this point visited, not the next one (yet)

		// Get steps in same direction
		ssd := 1 // steps in same direction
		if dir == curDir {
			ssd += stepsSameDir
		}

		// Memoization: does not work here because need to know entire history
		//config := Config{p1, dir, ssd, visString(vis)}
		//thisCost, already := Memo[config]
		//if true { //!already {
		//fmt.Println(p.x, p.y, "=>", p1.x, p1.y, curCost, vis)

		// Get the cost of the proposed move, update best if lower
		thisCost := nav(p1, dir, ssd, vis, curCost)

		//Memo[config] = thisCost
		//}

		// Update best for this round, if lower result found
		if thisCost < best {
			best = thisCost
		}
	}

	// Update global best
	if best < bestSoFar {
		fmt.Println("Best so far =", best)
		bestSoFar = best
	}

	// Return the best direction found
	return best
}
