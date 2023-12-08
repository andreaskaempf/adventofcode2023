// Advent of Code 2023, Day 08
//
// Read a list of left/right instructions, and the left/right adjacencies
// for a set of nodes. For  Part 1, calculate the number of steps required
// to get from node "AAA" to "ZZZ". For Part 2, calculate the first step
// where all routes from ??A lead to any ??Z. Part 2 solution finds and
// takes into account the repeating length of each ??A to ??Z, but uses
// brute force to find the time when they align, not very efficient.
//
// AK, 8 Dec 2023

package main

import (
	"fmt"
)

// Node, with next destinations for left and right branches
type Node struct {
	id, L, R string
}

// Map of nodes and path (global)
var nodes map[string]Node
var path string

func main() {

	// Read the input file into path and dict of nodes
	nodes = map[string]Node{} // initialize map
	fname := "sample3.txt"
	fname = "input.txt" // uncomment to use main input file
	for _, l := range readLines(fname) {
		if len(path) == 0 {
			path = l // First line is LRLR* path
		} else if len(l) > 0 { // e.g., QRX = (XNN, TCJ)
			n := Node{id: l[:3], L: l[7:10], R: l[12:15]}
			nodes[n.id] = n
		}
	}

	// Part 1: navigate path from AAA to ZZZ, count steps
	fmt.Println("Part 1:", nSteps("AAA", "ZZZ"))

	// Part 2: first find all starting nodes, i.e., those that end in A
	starts := []string{}
	for k := range nodes {
		if k[2] == 'A' {
			starts = append(starts, k)
		}
	}

	// Get path lengths from each start to any dest that ends in Z; these
	// seem to be repeating, so use them as the basis for finding an answer
	lengths := []int{}
	for _, s := range starts {
		n := nSteps(s, "any") // length of path to any "xxZ"
		lengths = append(lengths, n)
	}

	// Search multiples of the first path length to find the one that aligns
	// with all paths; multiplying them does not work, must be a common divisor
	// or something, so this is a bit slow
	fmt.Println("Lengths:", lengths)
	steps := lengths[0]
	for {
		// Check if the number of steps can be divided by all lengths,
		// stop if yes
		foundIt := true
		for i := 0; i < len(lengths); i++ {
			if steps%lengths[i] != 0 {
				foundIt = false
			}
		}
		if foundIt {
			break
		}

		// Otherwise, try the next increment
		steps += lengths[0]
	}
	fmt.Println("Part 2:", steps) // 22103062509257
}

// Find the number of steps to get from A to B, given starting path position.
// If B is "any", finds the first destination that ends in 'Z'
func nSteps(A, B string) int {
	var steps, pos int
	loc := A
	for loc != B { // continue until we reach destination
		n, _ := nodes[loc]    // get current node
		if path[pos] == 'L' { // get next location
			loc = n.L
		} else {
			loc = n.R
		}
		steps++
		pos++
		if pos >= len(path) {
			pos = 0
		}
		// If dest "any", stop at first destination that ends with 'Z'
		if B == "any" && loc[2] == 'Z' {
			fmt.Println(A, "->", loc, "after", steps, "steps")
			return steps // comment this to see the repeating intervals
		}
	}
	return steps
}
